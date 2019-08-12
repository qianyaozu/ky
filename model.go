package main

import (
	"time"
)

//region model
type SensorModel struct{
	ID int
	Name string
	WorkPlace string
	SensorType int
	FirstValue int
	SecondValue int
	CreateTime time.Time
}


type SensorHistoryModel struct{
	ID int
	Name string
	FirstValue int
	SecondValue int
	CreateTime time.Time
}

type SensorQueryModel struct{
	Name string
	WorkPlace string
	SensorType int
}


type SensorHistoryDataQueryModel struct{
	ID int
	StartTime time.Time
	EndTime time.Time
}

type QueryIdsModel struct{
	ID []int
}




type SystemConfigModel struct {
	MineName string
	ComName  string
	BaudRate int
	UserName string
	Password string
}
//endregion
