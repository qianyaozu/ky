package main

import (
	"database/sql"
	"fmt"
	"time"
)

//region  database
func init() {
	MakeMine()
	MakeSensor()
	MakeSensorHistory()
}

//region 传感器信息
func getSensorById(id int) *SensorModel {
	db, err := sql.Open("sqlite3", "./kydatabase.db")
	if err != nil {
		checkErr(err)
	}
	defer db.Close()
	rows, err := db.Query("SELECT  * FROM Sensor where ID=?", id)
	checkErr(err)
	var sensors = make([]SensorModel, 0)
	for rows.Next() {
		var sensor SensorModel
		err = rows.Scan(&sensor.ID, &sensor.Name, &sensor.WorkPlace, &sensor.SensorType, &sensor.FirstValue, &sensor.SecondValue, &sensor.CreateTime)
		checkErr(err)
		sensors = append(sensors, sensor)
	}

	rows.Close()
	if len(sensors) > 0 {
		return &sensors[0]
	}
	return nil
}

func updateSensorValue(id int, v1 int, v2 int) bool {
	db, err := sql.Open("sqlite3", "./kydatabase.db")
	if err != nil {
		checkErr(err)
	}
	defer db.Close()
	stmt, err := db.Prepare("update Sensor set FirstValue=?,SecondValue=? where ID=?")
	checkErr(err)

	res, err := stmt.Exec(v1, v2, id)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	return affect > 0
}

func getSensorByName(name string) *SensorModel {
	db, err := sql.Open("sqlite3", "./kydatabase.db")
	if err != nil {
		checkErr(err)
	}
	defer db.Close()
	rows, err := db.Query("SELECT   * FROM Sensor where Name=?", name)
	checkErr(err)
	var sensors = make([]SensorModel, 0)
	for rows.Next() {
		var sensor SensorModel
		err = rows.Scan(&sensor.ID, &sensor.Name, &sensor.WorkPlace, &sensor.SensorType, &sensor.FirstValue, &sensor.SecondValue, &sensor.CreateTime)
		checkErr(err)
		sensors = append(sensors, sensor)
	}

	rows.Close()
	if len(sensors) > 0 {
		return &sensors[0]
	}
	return nil
}

func queryAllSensorData() (pList []SensorModel, err error) {

	db, err := sql.Open("sqlite3", "./kydatabase.db")
	if err != nil {
		checkErr(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM Sensor ")
	checkErr(err)

	var sensors = make([]SensorModel, 0)
	for rows.Next() {
		var sensor SensorModel
		err = rows.Scan(&sensor.ID, &sensor.Name, &sensor.WorkPlace, &sensor.SensorType, &sensor.FirstValue, &sensor.SecondValue, &sensor.CreateTime)
		checkErr(err)
		sensors = append(sensors, sensor)
	}

	rows.Close()
	return sensors, nil
}

func querySensorData(sensor SensorQueryModel) (pList []SensorModel, err error) {

	db, err := sql.Open("sqlite3", "./kydatabase.db")
	if err != nil {
		checkErr(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM Sensor where name like '%" + sensor.Name + "%'")
	checkErr(err)

	var sensors = make([]SensorModel, 0)
	for rows.Next() {
		var sensor SensorModel
		err = rows.Scan(&sensor.ID, &sensor.Name, &sensor.WorkPlace, &sensor.SensorType, &sensor.FirstValue, &sensor.SecondValue, &sensor.CreateTime)
		checkErr(err)
		sensors = append(sensors, sensor)
	}

	rows.Close()
	return sensors, nil
}

func insertSensorHistoryData(data SensorHistoryModel) bool {
	db, err := sql.Open("sqlite3", "./kydatabase.db")
	if err != nil {
		checkErr(err)
	}
	defer db.Close()
	stmt, err := db.Prepare("INSERT INTO SensorHistory(ID, FirstValue,SecondValue,CreateTime" +
		") values(?,?,?,?)")
	checkErr(err)

	res, err := stmt.Exec(data.ID, data.FirstValue, data.SecondValue, data.CreateTime)
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	return id > 0
}

func querySensorRealTimeData(query SensorHistoryDataQueryModel) (pList []SensorHistoryModel, err error) {

	db, err := sql.Open("sqlite3", "./kydatabase.db")
	if err != nil {
		checkErr(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT ID,Name,FirstValue,SecondValue FROM Sensor where SensorType=? ", query.ID)
	checkErr(err)

	var sensors = make([]SensorHistoryModel, 0)
	for rows.Next() {
		var sensor SensorHistoryModel
		err = rows.Scan(&sensor.ID, &sensor.Name, &sensor.FirstValue, &sensor.SecondValue)
		checkErr(err)
		sensors = append(sensors, sensor)
	}

	rows.Close()
	return sensors, nil
}

func querySensorHistoryData(query SensorHistoryDataQueryModel) (pList []SensorHistoryModel, err error) {

	db, err := sql.Open("sqlite3", "./kydatabase.db")
	if err != nil {
		checkErr(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT ID,FirstValue,SecondValue,CreateTime FROM SensorHistory where ID=? and CreateTime>? and CreateTime<?", query.ID, query.StartTime, query.EndTime)
	checkErr(err)
	//fmt.Println("SELECT ID,FirstValue,SecondValue,CreateTime FROM SensorHistory where ID=? and CreateTime>? and CreateTime<?", query.ID, query.StartTime, query.EndTime)
	var sensors = make([]SensorHistoryModel, 0)
	for rows.Next() {
		var sensor SensorHistoryModel
		err = rows.Scan(&sensor.ID, &sensor.FirstValue, &sensor.SecondValue, &sensor.CreateTime)
		checkErr(err)
		sensors = append(sensors, sensor)
	}

	rows.Close()
	return sensors, nil
}

//更新传感器
func updateSensorData(sensor SensorModel) bool {
	db, err := sql.Open("sqlite3", "./kydatabase.db")
	if err != nil {
		checkErr(err)
	}
	defer db.Close()
	stmt, err := db.Prepare("update Sensor set Name=?,WorkPlace=?,SensorType=? where ID=?")
	checkErr(err)

	res, err := stmt.Exec(sensor.Name, sensor.WorkPlace, sensor.SensorType, sensor.ID)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	return affect > 0
}

//新增传感器
func insertSensorData(sensor SensorModel) int64 {
	db, err := sql.Open("sqlite3", "./kydatabase.db")
	if err != nil {
		checkErr(err)
	}
	defer db.Close()
	stmt, err := db.Prepare("INSERT INTO Sensor(ID,Name,WorkPlace, SensorType, FirstValue,SecondValue,CreateTime" +
		") values(?,?,?,?,?,?,?)")
	checkErr(err)

	res, err := stmt.Exec(sensor.ID, sensor.Name, sensor.WorkPlace, sensor.SensorType, 0, 0, time.Now())
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	return id
}

//删除传感器信息
func deleteSensorData(id []int) bool {
	db, err := sql.Open("sqlite3", "./kydatabase.db")
	if err != nil {
		checkErr(err)
	}
	defer db.Close()

	where := "(" + fmt.Sprint(id[0])
	for i := 1; i < len(id); i++ {
		where += "," + fmt.Sprint(id[i])
	}
	where += ")"
	res, err := db.Exec("delete from Sensor where ID  in " + where)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)
	return affect > 0
}

//endregion

//region 系统信息
//更新传感器
func updateSystemConfigData(config SystemConfigModel) bool {
	db, err := sql.Open("sqlite3", "./kydatabase.db")
	if err != nil {
		checkErr(err)
	}
	defer db.Close()
	stmt, err := db.Prepare("update SystemConfig set MineName=?,ComName=?,BaudRate=?,Password=?")
	checkErr(err)

	res, err := stmt.Exec(config.MineName, config.ComName, config.BaudRate, config.Password)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	return affect > 0
}

func querySystemConfigData() (config SystemConfigModel, err error) {

	db, err := sql.Open("sqlite3", "./kydatabase.db")
	if err != nil {
		checkErr(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM SystemConfig")
	checkErr(err)

	for rows.Next() {
		err = rows.Scan(&config.MineName, &config.ComName, &config.BaudRate, &config.UserName, &config.Password)
		checkErr(err)

	}

	rows.Close()
	return config, nil
}

//endregion

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}

//endregion
