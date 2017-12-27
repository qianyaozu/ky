package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-mgo/mgo"
	"github.com/go-mgo/mgo/bson"
	"github.com/qianyaozu/qcommon"
	"net/http"
	"reflect"
	"time"
)

func handleServer() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(gin.Recovery())
	router.POST("/api/common", common)
	router.POST("/api/upload", upload)
	router.Run(":" + Port)
}

//region 通用增删改查
func common(c *gin.Context) {
	var body qcommon.PostData
	err := c.Bind(&body)
	if err != nil {
		c.JSON(http.StatusOK, qcommon.ResponseJson(nil, err))
	}
	switch body.Method {
	case "get":
		c.JSON(http.StatusOK, qcommon.ResponseJson(get(body)))
		return
	case "update":
		c.JSON(http.StatusOK, qcommon.ResponseJson(update(body)))
		return
	case "count":
		c.JSON(http.StatusOK, qcommon.ResponseJson(count(body)))
		return
	case "add":
		c.JSON(http.StatusOK, qcommon.ResponseJson(add(body)))
		return
	case "delete":
		c.JSON(http.StatusOK, qcommon.ResponseJson(delete(body)))
		return
	}
	c.JSON(http.StatusOK, qcommon.ResponseJson(nil, errors.New("bad request:unsupported method")))
	return
}

//根据查询条件查询数据
func get(body qcommon.PostData) (interface{}, error) {
	data, ok := body.Data.(map[string]interface{})
	if !ok {
		return nil, errors.New("bad request:data type is not map[string]interface{} but a " + reflect.TypeOf(body.Data).Name())
	}
	session, err := qcommon.InitMongo(DBServer)
	if err != nil {
		return nil, err
	}
	defer session.Close()
	collection := session.DB(body.DBName).C(body.Table)
	var query *mgo.Query
	query.SetMaxTime(5 * time.Minute) //5分钟
	if body.ObjectID!=""{
		collection.Find(bson.M{"_id": bson.ObjectIdHex(body.ObjectID)})
	}else{
		query = collection.Find(data)
	}

	if body.OrderBy != "" {
		query = query.Sort(body.OrderBy)
	}
	if body.Skip != 0 {
		query = query.Skip(body.Skip)
	}
	if body.Limit != 0 {
		query = query.Limit(body.Limit)
	}
	if body.Select != nil {
		query = query.Select(body.Select)
	}
	if body.Distinct != "" {
		var res []string
		err = query.Distinct(body.Distinct, &res)
		if err == nil {
			return res, nil
		} else {
			return nil, err
		}
	} else {
		var res []interface{}
		err = query.All(&res)
		if err != nil {
			return nil, err
		} else {
			return res, nil
		}
	}
}

func update(body qcommon.PostData) (interface{}, error) {
	data, ok := body.Data.(map[string]interface{})
	if !ok {
		return nil, errors.New("bad request:data type is not map[string]interface{} but a " + reflect.TypeOf(body.Data).Name())
	}
	if body.Condition == nil {
		return nil, errors.New("bad request:body.Condition can't be empty")
	}

	condition, ok := body.Condition.(map[string]interface{})

	if !ok {
		return nil, errors.New("bad request:condition type is not map[string]interface{}   but a " + reflect.TypeOf(body.Condition).Name())
	}
	session, err := qcommon.InitMongo(DBServer)
	if err != nil {
		return nil, err
	}
	defer session.Close()
	collection := session.DB(body.DBName).C(body.Table)
	data["update_at"] = time.Now().Unix()

	if body.UpdateAll {
		if body.ObjectID != "" {
			collection.UpdateAll(bson.M{"_id": bson.ObjectIdHex(body.ObjectID)}, bson.M{"$set": data})
		} else {
			_, err = collection.UpdateAll(condition, bson.M{"$set": data})
		}

	} else {
		if body.ObjectID != "" {
			err = collection.Update(bson.M{"_id": bson.ObjectIdHex(body.ObjectID)}, bson.M{"$set": data})
		} else {
			err = collection.Update(condition, bson.M{"$set": data})
		}
	}
	if err != nil {
		return nil, err
	} else {
		return true, nil
	}
}

func count(body qcommon.PostData) (interface{}, error) {
	data, ok := body.Data.(map[string]interface{})
	if !ok {
		return nil, errors.New("bad request:data type is not map[string]interface{} but a " + reflect.TypeOf(body.Data).Name())
	}

	session, err := qcommon.InitMongo(DBServer)
	if err != nil {
		return nil, err
	}
	defer session.Close()
	query := session.DB(body.DBName).C(body.Table).Find(data)
	query.SetMaxTime(2 * time.Minute) //5分钟
	return query.Count()
}
func add(body qcommon.PostData) (interface{}, error) {

	if data, ok := body.Data.(map[string]interface{}); ok {
		session, err := qcommon.InitMongo(DBServer)
		if err != nil {
			return nil, err
		}
		defer session.Close()
		err = addFunc(data, body, session)
		if err != nil {
			return nil, err
		} else {
			return true, nil
		}
	} else if dataList, ok := body.Data.([]interface{}); ok {
		session, err := qcommon.InitMongo(DBServer)
		if err != nil {
			return nil, err
		}
		defer session.Close()
		for _, data := range dataList {
			if d, ok := data.(map[string]interface{}); ok {
				err := addFunc(d, body, session)
				if err != nil {
					return nil, err
				}
			}
		}
		return true, nil
	} else {
		return nil, errors.New("bad request:data type is not map[string]interface{} or []interface{} but a " + reflect.TypeOf(body.Data).Name())
	}
}
func addFunc(data map[string]interface{}, body qcommon.PostData, session *mgo.Session) error {
	collection := session.DB(body.DBName).C(body.Table)
	if body.Condition != nil {
		if condition, ok := body.Condition.(map[string]interface{}); ok {
			_, err := collection.Upsert(condition, bson.M{"$set": data})
			return err
		} else {
			return errors.New("bad request:body.Condition is not map[string]interface{}  but a " + reflect.TypeOf(body.Data).Name())
		}
	} else {
		return collection.Insert(data) //插入数据
	}
}
func delete(body qcommon.PostData) (interface{}, error) {
	data, ok := body.Data.(map[string]interface{})
	if !ok {
		return nil, errors.New("bad request:data type is not map[string]interface{} but a " + reflect.TypeOf(body.Data).Name())
	}
	session, err := qcommon.InitMongo(DBServer)
	if err != nil {
		return nil, err
	}
	defer session.Close()
	_, err = session.DB(body.DBName).C(body.Table).RemoveAll(data)
	if err != nil {
		return false, err
	} else {
		return true, err
	}
}

//endregion

//region 上传文件
func upload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusOK, qcommon.ResponseJson(nil, err))
	}
	// Upload the file to specific dst.
	var path = "d:\\" + file.Filename
	err = c.SaveUploadedFile(file, path)
	if err != nil {
		c.JSON(http.StatusOK, qcommon.ResponseJson(nil, err))
	}
	c.JSON(http.StatusOK, qcommon.ResponseJson(true, nil))
}

//endregion
