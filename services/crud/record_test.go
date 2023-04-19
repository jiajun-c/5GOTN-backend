package crud

import (
	"otn/config"
	"otn/dal"
	"testing"
	"time"
)

func TestInsertRecord(t *testing.T) {
	config.LoadDB()
	ti, _ := time.Parse("2006-01-02 15:04:05", "2009-02-04 12:12:12")
	InsertRecord(&dal.Record{
		WarningId:     2,
		DeviceID:      2,
		DeviceAddress: "whu",
		DevicePort:    "120",
		WarningTime:   ti,
		WarningCode:   0,
	})
}

func TestGetAllRecords(t *testing.T) {
	config.LoadDB()
	records, _ := GetAllRecords()
	println(records[0].WarningTime.String())
}

func TestSelectRecords(t *testing.T) {
	config.LoadDB()
	start, _ := time.Parse("2006-01-02 15:04:05", "2002-02-04 12:12:12")
	end, _ := time.Parse("2006-01-02 15:04:05", "2022-02-04 12:12:12")
	s := start.Format("2006-01-02 15:04:05")
	println(s)
	println(end.String())
	records, _ := SelectRecords(start, end, 1, "")
	println(len(records))
}
