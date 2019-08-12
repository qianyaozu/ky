package main

import (
	"database/sql"
)

func MakeSensor(){
	sql_table := `
    CREATE TABLE Sensor(
        ID INTEGER PRIMARY KEY ,
        Name VARCHAR(64) NULL,
        WorkPlace VARCHAR(64) NULL,
        SensorType int NULL,
        FirstValue int NULL,
        SecondValue int null,
        CreateTime DATE NULL
    );
    `
	db, err := sql.Open("sqlite3", "./kydatabase.db")
	if err != nil {
		checkErr(err)
	}
	defer db.Close()
	db.Exec(sql_table)
}

func MakeSensorHistory(){
	sql_table := `
    CREATE TABLE SensorHistory(
        ID INTEGER PRIMARY KEY ,
        FirstValue int NULL,
        SecondValue int null,
        CreateTime DATE NULL
    );
    `
	db, err := sql.Open("sqlite3", "./kydatabase.db")
	if err != nil {
		checkErr(err)
	}
	defer db.Close()
	db.Exec(sql_table)
}


func MakeMine() {
	sql_table := `
    CREATE TABLE SystemConfig(
        MineName VARCHAR(64) NULL,
        ComName VARCHAR(64) NULL,
        BaudRate int NULL,
        UserName VARCHAR(64) null,
        Password varchar(64) null
        
    );
    `
	db, err := sql.Open("sqlite3", "./kydatabase.db")
	if err != nil {
		checkErr(err)
	}
	defer db.Close()
	db.Exec(sql_table)

	rows, err := db.Query("SELECT * FROM SystemConfig")
	checkErr(err)

	exists := false
	for rows.Next() {
		exists = true

	}

	rows.Close()

	if !exists {
		res, err := db.Exec("INSERT INTO SystemConfig(MineName,ComName,BaudRate,UserName,Password) values('矿压系统','COM1',9600,'admin','admin')")
		checkErr(err)
		res.RowsAffected()


	}
}
