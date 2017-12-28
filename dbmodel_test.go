package main

import (
	"github.com/qianyaozu/qcommon"
	"testing"
	"fmt"
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
