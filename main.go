package main

import (
	"github.com/qianyaozu/qconf"
	"os"
	"strconv"
)

type SerialPort struct {

}

var (
	DBServer = "127.0.0.1:27017"
	Port     = "30003" //服务端口号

	UserName    = "admin"
	Password    = "admin"
	MineName    = "义乌小商品批发市场"
	MongoDBPath = "" //
	Env         = ""
	ComName string
	Baud    int
	Parity  int // None,Even,Odd,0,1,2
	Conf *qconf.Config
)

func init() {
	Env, _ = os.Getwd()
	LoadConf()
}
func main() {
	handleServer()
}

func LoadConf() {
	Conf, err := qconf.LoadConfiguration("conf.ini")
	if err != nil {
		panic(err)
	}
	Port = Conf.GetString("ServerPort")
	DBServer = Conf.GetString("DBServer")
	UserName = Conf.GetString("UserName")
	Password = Conf.GetString("Password")
	MineName = Conf.GetString("MineName")
	MongoDBPath = Conf.GetString("MongoDBPath")

	ComName = Conf.GetString("ComName")
	Parity, err = strconv.Atoi(Conf.GetString("Parity"))
	Baud, err = strconv.Atoi(Conf.GetString("Baud"))
	if err != nil {
		panic(err)
	}
}
func SaveConf() {

}
