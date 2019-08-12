package main

import (
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"github.com/qianyaozu/qcommon"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

var Port = "30003"

func main() {

	//var data SensorHistoryModel
	//data.ID = 2
	//data.CreateTime = time.Date(time.Now().Year(), 5, 7, 12, 0, 0, 0, time.Local)
	//data.FirstValue = 3
	//data.SecondValue = 4
	//insertSensorHistoryData(data)
	//return
	//fmt.Println(querySystemConfigData())
	//var query SensorHistoryDataQueryModel
	//query.ID=2
	//query.StartTime=time.Date(time.Now().Year(), 5, 6, 0, 0, 0, 0, time.Local)
	//query.EndTime=time.Date(time.Now().Year(), 5, 7, 0, 0, 0, 0, time.Local)
	//fmt.Println(querySensorHistoryData(query))
	//
	//fmt.Println(querySensorRealTimeData(query))
	//handleMessage()
	if !verify() {
		fmt.Scan()
		os.Exit(0)
	}
	go startSerial()
	go func() {
		time.Sleep(1 * time.Second)
		Open("http://localhost:30003/#/Login")
	}()
	handleServer()
}

func handleServer() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(gin.Recovery())
	router.Static("/", "./")
	router.POST("/api/login", login)

	router.POST("/api/getAllSensor", getAllSensor)
	router.POST("/api/querySensorHistory", querySensorHistory)
	router.POST("/api/querySensorRealTime", querySensorRealTime)
	router.POST("/api/querySensor", querySensor)
	router.POST("/api/updateSensor", updateSensor)
	router.POST("/api/deleteSensor", deleteSensor)
	router.POST("/api/insertSensor", insertSensor)

	router.POST("/api/querySystemConfig", querySystemConfig)
	router.POST("/api/queryMineName", queryMineName)
	router.POST("/api/updateSystemConfig", updateSystemConfig)
	//router.POST("/api/common", common)
	//router.POST("/api/upload", upload)
	//router.POST("/api/backup", backup)
	router.POST("/api/backup", backup)
	router.POST("/api/restart", restartSerial)
	//router.POST("/api/startcommunication", startcommunication)
	//router.POST("/api/checkcommunication", checkcommunication)

	router.Run(":" + Port)
}

//region登陆
func login(c *gin.Context) {
	type Login struct {
		UserName string
		Password string
	}
	type UserInfo struct {
		UserName string
		Token    int64
	}
	var user Login
	c.Bind(&user)

	config, err := querySystemConfigData()
	if err != nil {
		c.JSON(http.StatusOK, qcommon.ResponseJson(nil, errors.New("登录失败")))
		return
	}
	if user.UserName == config.UserName && user.Password == config.Password {
		var user UserInfo
		user.UserName = config.UserName
		user.Token = time.Now().Unix()
		c.JSON(http.StatusOK, qcommon.ResponseJson(user, nil))
	} else {
		c.JSON(http.StatusOK, qcommon.ResponseJson(nil, errors.New("登录失败")))
	}
}

//endregion

//region 系统
func restartSerial(c *gin.Context) {
	if restart() {
		c.JSON(http.StatusOK, qcommon.ResponseJson("重启成功", nil))
		go initCom()
	} else {
		c.JSON(http.StatusOK, qcommon.ResponseJson(nil, errors.New("重启失败")))
	}
}

func backup(c *gin.Context) {
	sourceFileStat, err := os.Stat("./kydatabase.db")
	if err != nil {
		c.JSON(http.StatusOK, qcommon.ResponseJson(nil, err))
		return
	}
	if !sourceFileStat.Mode().IsRegular() {
		c.JSON(http.StatusOK, qcommon.ResponseJson(nil, errors.New("s not a regular file")))
		return
	}

	source, err := os.Open("./kydatabase.db")
	if err != nil {
		c.JSON(http.StatusOK, qcommon.ResponseJson(nil, err))
		return
	}
	defer source.Close()

	var dst = "./kydatabase" + fmt.Sprint(time.Now().Unix()) + ".db"
	destination, err := os.Create(dst)
	if err != nil {
		c.JSON(http.StatusOK, qcommon.ResponseJson(nil, err))
		return
	}
	defer destination.Close()
	b, err := io.Copy(destination, source)
	if err != nil {
		c.JSON(http.StatusOK, qcommon.ResponseJson(nil, err))
		return
	}
	c.JSON(http.StatusOK, qcommon.ResponseJson(dst+fmt.Sprint(b), nil))
}

//endregion

//region 传感器列表

func getAllSensor(c *gin.Context) {
	var sensor SensorQueryModel
	err := c.Bind(&sensor)
	if err != nil {
		c.JSON(http.StatusOK, qcommon.ResponseJson(nil, err))
		return
	}

	sList, err := queryAllSensorData()
	if err != nil {

		fmt.Println(err)
		c.JSON(http.StatusOK, qcommon.ResponseJson(nil, errors.New("查询失败")))
	} else {
		c.JSON(http.StatusOK, qcommon.ResponseJson(sList, nil))
	}
}
func querySensor(c *gin.Context) {
	var sensor SensorQueryModel
	err := c.Bind(&sensor)
	if err != nil {
		c.JSON(http.StatusOK, qcommon.ResponseJson(nil, err))
		return
	}

	sList, err := querySensorData(sensor)
	if err != nil {

		fmt.Println(err)
		c.JSON(http.StatusOK, qcommon.ResponseJson(nil, errors.New("查询失败")))
	} else {
		c.JSON(http.StatusOK, qcommon.ResponseJson(sList, nil))
	}
}

//获取实时传感器数据
func querySensorRealTime(c *gin.Context) {
	var query SensorHistoryDataQueryModel
	err := c.Bind(&query)
	if err != nil {
		c.JSON(http.StatusOK, qcommon.ResponseJson(nil, err))
		return
	}

	query.StartTime = time.Date(query.StartTime.Year(), time.Month(query.StartTime.Month()), query.StartTime.Day(), 0, 0, 0, 0, time.Local)
	query.StartTime = time.Date(query.EndTime.Year(), time.Month(query.EndTime.Month()), query.EndTime.Day(), 23, 59, 59, 0, time.Local)

	sList, err := querySensorRealTimeData(query)
	if err != nil {

		fmt.Println(err)
		c.JSON(http.StatusOK, qcommon.ResponseJson(nil, errors.New("查询失败")))
	} else {
		c.JSON(http.StatusOK, qcommon.ResponseJson(sList, nil))
	}
}

//获取历史传感器数据
func querySensorHistory(c *gin.Context) {
	var query SensorHistoryDataQueryModel
	err := c.Bind(&query)
	if err != nil {
		c.JSON(http.StatusOK, qcommon.ResponseJson(nil, err))
		return
	}

	query.StartTime = time.Date(query.StartTime.Year(), time.Month(query.StartTime.Month()), query.StartTime.Day(), 0, 0, 0, 0, time.Local)
	query.EndTime = time.Date(query.EndTime.Year(), time.Month(query.EndTime.Month()), query.EndTime.Day(), 23, 59, 59, 0, time.Local)

	sList, err := querySensorHistoryData(query)
	if err != nil {

		fmt.Println(err)
		c.JSON(http.StatusOK, qcommon.ResponseJson(nil, errors.New("查询失败")))
	} else {
		c.JSON(http.StatusOK, qcommon.ResponseJson(sList, nil))
	}
}

//更新传感器信息
func updateSensor(c *gin.Context) {
	var sensor SensorModel
	err := c.Bind(&sensor)
	if err != nil {
		c.JSON(http.StatusOK, qcommon.ResponseJson(nil, err))
		return
	}
	if updateSensorData(sensor) {
		c.JSON(http.StatusOK, qcommon.ResponseJson(1, nil))
	} else {
		c.JSON(http.StatusOK, qcommon.ResponseJson(nil, errors.New("更新失败")))
	}

}

//新增传感器
func insertSensor(c *gin.Context) {
	var sensor SensorModel
	err := c.Bind(&sensor)
	if err != nil {
		c.JSON(http.StatusOK, qcommon.ResponseJson(nil, err))
		return
	}
	if getSensorById(sensor.ID) != nil {
		c.JSON(http.StatusOK, qcommon.ResponseJson(nil, errors.New("该ID已存在，新增失败")))
		return
	}
	if getSensorByName(sensor.Name) != nil {
		c.JSON(http.StatusOK, qcommon.ResponseJson(nil, errors.New("该名称已存在，新增失败")))
		return
	}
	if insertSensorData(sensor) > 0 {
		c.JSON(http.StatusOK, qcommon.ResponseJson(1, nil))
	} else {
		c.JSON(http.StatusOK, qcommon.ResponseJson(nil, errors.New("新增失败")))
	}
}

//删除传感器
func deleteSensor(c *gin.Context) {
	var model QueryIdsModel
	err := c.Bind(&model)
	if err != nil {
		c.JSON(http.StatusOK, qcommon.ResponseJson(nil, err))
		return
	}

	if deleteSensorData(model.ID) {
		c.JSON(http.StatusOK, qcommon.ResponseJson(1, nil))
	} else {
		c.JSON(http.StatusOK, qcommon.ResponseJson(nil, errors.New("删除失败")))
	}
}

//endregion

//region 煤矿信息

func updateSystemConfig(c *gin.Context) {
	var config SystemConfigModel
	err := c.Bind(&config)
	if err != nil {
		c.JSON(http.StatusOK, qcommon.ResponseJson(nil, err))
		return
	}
	if updateSystemConfigData(config) {
		c.JSON(http.StatusOK, qcommon.ResponseJson(1, nil))
	} else {
		c.JSON(http.StatusOK, qcommon.ResponseJson(nil, errors.New("更新失败")))
	}

}

func querySystemConfig(c *gin.Context) {
	config, err := querySystemConfigData()
	if err == nil {
		c.JSON(http.StatusOK, qcommon.ResponseJson(config, nil))
	} else {
		c.JSON(http.StatusOK, qcommon.ResponseJson(nil, err))
	}

}

func queryMineName(c *gin.Context) {
	config, err := querySystemConfigData()
	if err == nil {
		c.JSON(http.StatusOK, qcommon.ResponseJson(config.MineName, nil))
	} else {
		c.JSON(http.StatusOK, qcommon.ResponseJson(nil, err))
	}

}

//endregion

var commands = map[string]string{
	"windows": "cmd /c start",
	"darwin":  "open",
	"linux":   "xdg-open",
}

var Version = "0.1.0"

// Open calls the OS default program for uri
func Open(uri string) error {
	run, ok := commands[runtime.GOOS]
	if !ok {
		return fmt.Errorf("don't know how to open things on %s platform", runtime.GOOS)
	}

	cmd := exec.Command(run, uri)
	return cmd.Start()
}

//验证码识别
func verify() bool {
	var verifyCode = ""
	//读取注册码
	if _, err := os.Stat("verify.txt"); err == nil {
		contents, err := ioutil.ReadFile("verify.txt")
		if err == nil {
			//因为contents是[]byte类型，直接转换成string类型后会多一行空格,需要使用strings.Replace替换换行符
			verifyCode = strings.Replace(string(contents), "\n", "", 1)
		}
	}

	var list = make([]string, 0)
	intf, err := net.Interfaces()
	if err != nil || len(intf) == 0 {
		return false
	}
	for _, v := range intf {
		if v.HardwareAddr.String() != "" {
			item := base64.StdEncoding.EncodeToString([]byte(v.HardwareAddr.String()))
			if verifyCode == base64.StdEncoding.EncodeToString([]byte(base64.StdEncoding.EncodeToString([]byte(item)))) {
				fmt.Println("软件授权成功")
				return true
			}
			list = append(list, item)
			//fmt.Println(item)
		}
	}
	fmt.Println("*****************软件未授权，请联系销售获取软件注册码*******************")
	fmt.Println("设备码:", list[0])
	fmt.Println("**************************请输入注册码********************************")
	var code = ""
	fmt.Scan(&code)
	//fmt.Println(n, err, code)
	rightcode := base64.StdEncoding.EncodeToString([]byte(base64.StdEncoding.EncodeToString([]byte(list[0]))))
	//fmt.Println(rightcode)
	if strings.ToLower(code) == strings.ToLower(rightcode) {
		fmt.Println("注册码验证成功")
		ioutil.WriteFile("verify.txt", []byte(code), 0644)
		return true
	}
	return false
}
