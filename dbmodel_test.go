package main

import (
	"github.com/qianyaozu/qcommon"
	"testing"
	"fmt"
	"github.com/go-mgo/mgo/bson"
)

func TestInsert(t *testing.T) {
	var frame =make(map[string]interface{})
	frame["ID"] = 10
	frame["InitPower"] = 10
	frame["InstallPosition"] = "1"
	frame["MaxResistence"] = 20
	var postData = qcommon.PostData{Method: "add", DBName: "ky", Table: "Frame", Data: frame}
	fmt.Println(add(postData))

}

func TestGet(t *testing.T){
	var postData = qcommon.PostData{Method: "get", DBName: "ky", Table: "Frame", Data: make(map[string]interface{})}
	fmt.Println(get(postData))

}
<<<<<<< HEAD

func TestDelete(t *testing.T){
	var frame =make(map[string]interface{})
	frame["ID"] = 10
	var postData = qcommon.PostData{Method: "delete", DBName: "ky", Table: "Frame", Data: frame}
	fmt.Println(delete(postData))
}

func TestCount(t *testing.T){
	var postData = qcommon.PostData{Method: "get", DBName: "ky", Table: "Frame", Data: make(map[string]interface{})}
	fmt.Println(count(postData))

}
func TestGet1(t *testing.T){
	//5a47880dda81e2080bc3035c
	session, _ := qcommon.InitMongo("127.0.0.1:27017")
	defer session.Close()
	collection := session.DB("ky").C("workplace")
	fmt.Println(collection.Find(bson.M{"ID": "1"}).Count())
=======
func TestGet1(t * testing.T) {
	session, _ := qcommon.InitMongo(DBServer)
	collection := session.DB("ky").C("Frame")
	//query.SetMaxTime(5 * time.Minute) //5分钟

	fmt.Println(collection.Find(bson.D{{"_id", bson.ObjectId("5a498a5ad03c3df26bd9f885")}}).Count())
	session.Close()
>>>>>>> 53dc6ec96803f37e6945e33b3bbac3c3dc3a4f19
}
