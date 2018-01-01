package main

import (
	"github.com/qianyaozu/qcommon"
	"testing"
	"fmt"
	"github.com/go-mgo/mgo/bson"
)

func TestInsert(t *testing.T) {
	var frame =make(map[string]interface{})
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
func TestGet1(t * testing.T) {
	session, _ := qcommon.InitMongo(DBServer)
	collection := session.DB("ky").C("Frame")
	//query.SetMaxTime(5 * time.Minute) //5分钟

	fmt.Println(collection.Find(bson.D{{"_id", bson.ObjectId("5a498a5ad03c3df26bd9f885")}}).Count())
	session.Close()
}
