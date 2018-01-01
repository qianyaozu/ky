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
	"fmt"
	"os/exec"
	"sync"
)

func handleServer() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(gin.Recovery())
	router.POST("/api/login", login)
	router.POST("/api/common", common)
	router.POST("/api/upload", upload)
	router.POST("/api/backup",backup)
	router.POST("/api/restore",restore)
	router.POST("/api/startcommunication",startcommunication)
	router.POST("/api/checkcommunication",checkcommunication)
	router.Run(":" + Port)
}

//region 通用增删改查
func common(c *gin.Context) {
	var body qcommon.PostData
	err := c.Bind(&body)
	if err != nil {
		c.JSON(http.StatusOK, qcommon.ResponseJson(nil, err))
		return
	}
	switch body.Method {
	case "page":
		c.JSON(http.StatusOK, qcommon.ResponseJson(page(body)))
		return
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
func page(body qcommon.PostData) (interface{}, error) {
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
	//query.SetMaxTime(5 * time.Minute) //5分钟
	query = collection.Find(data)
	count, err := collection.Find(data).Count()
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
	var res []interface{}
	if body.Distinct != "" {
		err = query.Distinct(body.Distinct, &res)
	} else {
		err = query.All(&res)
	}
	type Page struct {
		Items []interface{}
		Count int
	}
	var page Page
	page.Items = res
	page.Count = count
	return page, nil
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
	//query.SetMaxTime(5 * time.Minute) //5分钟
	query = collection.Find(data)

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
	d, ok := body.Data.(map[string]interface{})
	if !ok {
		return nil, errors.New("bad request:data type is not map[string]interface{} but a " + reflect.TypeOf(body.Data).Name())
	}
	var data=make(map[string]interface{})
	for k,v:=range d {
		if k != "_id" {
			data[k] = v
		}
	}
	var query = make(map[string]interface{})

		if body.Condition == nil {
			return nil, errors.New("bad request:body.Condition can't be empty")
		} else {
			if condition, ok := body.Condition.(map[string]interface{}); ok {
				query = condition
			} else {
				return nil, errors.New("bad request:body.Condition is not map[string]interface{}  but a " + reflect.TypeOf(body.Data).Name())
			}
		}
	fmt.Println(query,body)
	session, err := qcommon.InitMongo(DBServer)
	if err != nil {
		return nil, err
	}
	defer session.Close()
	collection := session.DB(body.DBName).C(body.Table)
	data["update_at"] = time.Now().Unix()

	if body.UpdateAll {
		_, err = collection.UpdateAll(query, bson.M{"$set": data})

	} else {
		err = collection.Update(query, bson.M{"$set": data})
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
	data["create_at"] = time.Now().Unix()
	collection := session.DB(body.DBName).C(body.Table)
	query := make(map[string]interface{})
	if body.Condition != nil {

		if condition, ok := body.Condition.(map[string]interface{}); ok {
			query = condition
		} else {
			return errors.New("bad request:body.Condition is not map[string]interface{}  but a " + reflect.TypeOf(body.Data).Name())
		}
	}
	count, err := collection.Find(query).Count()
	if err != nil {
		return err
	} else if count > 0 {
		return errors.New("该记录已经存在")
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


//region登陆
func login(c *gin.Context) {
	type Login struct {
		UserName string
		Password string
	}
	type UserInfo struct{
		UserName string
		MineName string
		Token int64
	}
	var user Login
	c.Bind(&user)
	if user.UserName == UserName && user.Password == Password {
		var user UserInfo
		user.UserName=UserName
		user.MineName=MineName
		user.Token=time.Now().Unix()

		c.JSON(http.StatusOK, qcommon.ResponseJson(user, nil))
	} else {
		c.JSON(http.StatusOK, qcommon.ResponseJson(nil, errors.New("login failed")))
	}
}
//endregion

//region 备份还原
func backup(c *gin.Context) {

	list:=make([]string,0)

	cmd := exec.Command(MongoDBPath+"\\mongodump.exe", list...)
	err := cmd.Run()

	if err != nil {
		c.JSON(http.StatusOK, qcommon.ResponseJson(false, err))
	} else {
		c.JSON(http.StatusOK, qcommon.ResponseJson(true, nil))
	}
}
func restore(c *gin.Context) {
	cmd := exec.Command(MongoDBPath+"\\mongorestore.exe", "-d", "ky", "dump/ky")
	err := cmd.Run()
	if err != nil {
		c.JSON(http.StatusOK, qcommon.ResponseJson(false, err))
	} else {
		c.JSON(http.StatusOK, qcommon.ResponseJson(true, nil))
	}
}
//endregion


//region 开始接收数据
var IsCommunication=false
var CommunicateLock sync.RWMutex
func startcommunication(c *gin.Context) {
	CommunicateLock.Lock()
	defer CommunicateLock.Unlock()
	if IsCommunication {
		c.JSON(http.StatusOK, qcommon.ResponseJson(false, errors.New("通讯正在进行中")))
		return
	}
	IsCommunication = true
	//开始通讯
	go docomcommunication()
	c.JSON(http.StatusOK, qcommon.ResponseJson(true, nil))
}

func checkcommunication(c *gin.Context) {
	CommunicateLock.Lock()
	defer CommunicateLock.Unlock()
	if IsCommunication {
		c.JSON(http.StatusOK, qcommon.ResponseJson(false, errors.New("通讯正在进行中")))
		return
	}
	c.JSON(http.StatusOK, qcommon.ResponseJson(true, nil))
}
//执行通讯
func docomcommunication(){
	defer func() { IsCommunication = false }()
}

func doadbcommunication(){
	defer func() { IsCommunication = false }()
}
//endregion
