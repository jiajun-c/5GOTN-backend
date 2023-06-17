package handler

import (
	"github.com/gocarina/gocsv"
	"os"
	"testing"
	"time"
)

func TestHandler(t *testing.T) {
	var datas []*Record_1
	file, _ := os.Open("temp.csv")
	if err := gocsv.UnmarshalFile(file, &datas); err != nil {
		t.Error(err)
	}
	for _, value := range datas {
		layout := "2006/1/2 15:04"
		t, _ := time.Parse(layout, value.WarningTime)
		println(t.String())
		//crud.InsertRecord(&realData)
	}
}
