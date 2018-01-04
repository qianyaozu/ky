package main

import "github.com/qianyaozu/qconf"

var (
	DBServer = "127.0.0.1:27017"
	Port     = "30003" //服务端口号
	UserName = "admin"
	Password = "admin"
	MineName = "义乌小商品批发市场"
	MongoDBPath="" //
)

func init() {
	LoadConf()
}
func main() {
	handleServer()
}

func LoadConf() {
	conf, err := qconf.LoadConfiguration("conf.ini")
	if err != nil {
		panic(err)
	}
	Port = conf.GetString("ServerPort")
	DBServer = conf.GetString("DBServer")
	UserName = conf.GetString("UserName")
	Password = conf.GetString("Password")
	MineName = conf.GetString("MineName")
	MongoDBPath = conf.GetString("MongoDBPath")
	if err != nil {
		panic(err)
	}
}
