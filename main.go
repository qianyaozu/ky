package main

import "github.com/qianyaozu/qconf"

var (
	DBServer = "127.0.0.1:27017"
	Port     = "30003" //服务端口号
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
	if err != nil {
		panic(err)
	}
}
