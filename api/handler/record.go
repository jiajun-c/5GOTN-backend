package handler

import (
	"github.com/gin-gonic/gin"
	"otn/dal"
	"otn/services/crud"
	"strconv"
	"time"
)

func GetAllRecords(ctx *gin.Context) {
	records, err := crud.GetAllRecords()
	if err != nil {
		ctx.JSON(200, gin.H{
			"message": err.Error(),
		})
		return
	}
	resps := make([]*dal.RecordResp, 0)
	for _, record := range records {
		wt := record.WarningTime.Format("2006/01/02 15:04")
		resp := &dal.RecordResp{
			WarningId:     record.WarningId,
			DeviceID:      record.DeviceID,
			DeviceAddress: record.DeviceAddress,
			DevicePort:    record.DevicePort,
			WarningTime:   wt,
			WarningCode:   record.WarningCode,
		}
		resps = append(resps, resp)
	}
	ctx.JSON(200, resps)
}

func DeleteRecord(ctx *gin.Context) {
	id := ctx.Query("warning_id")
	warning_id, _ := strconv.Atoi(id)
	err := crud.DeleteRecord(int64(warning_id))
	if err != nil {
		ctx.JSON(200, gin.H{
			"message": "Error deleting",
		})
		return
	}
	ctx.JSON(200, gin.H{
		"message": "success",
	})
}

func FilterRecord(ctx *gin.Context) {
	deviceId := ctx.Query("device_id")
	did, _ := strconv.Atoi(deviceId)
	deviceAddress := ctx.Query("device_address")
	warningTimeBegin := ctx.Query("warning_time_begin")
	warningTimeEnd := ctx.Query("warning_time_end")
	start, _ := time.Parse("2006/01/02 15:04", warningTimeBegin)
	end, _ := time.Parse("2006/01/02 15:04", warningTimeEnd)
	records, err := crud.SelectRecords(start, end, int64(did), deviceAddress)
	if err != nil {
		ctx.JSON(200, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, records)
}

func InsertRecord(ctx *gin.Context) {
	warningId := ctx.PostForm("warning_id")
	deviceId := ctx.PostForm("device_id")
	devicePort := ctx.PostForm("device_port")
	deviceAddress := ctx.PostForm("device_address")
	warningTime := ctx.PostForm("warning_time")
	warningLevel := ctx.PostForm("warning_level")
	warningCode := ctx.PostForm("warning_code")
	wid, _ := strconv.Atoi(warningId)
	did, _ := strconv.Atoi(deviceId)
	wtime, _ := time.Parse("2006/01/02 15:04", warningTime)
	wcode, _ := strconv.Atoi(warningCode)
	wlevel, _ := strconv.Atoi(warningLevel)
	crud.InsertRecord(&dal.Record{
		WarningId:     int64(wid),
		DeviceID:      int64(did),
		DeviceAddress: deviceAddress,
		DevicePort:    devicePort,
		WarningTime:   wtime,
		WarningCode:   int64(wcode),
		WarningLevel:  int64(wlevel),
	})
	ctx.JSON(200, gin.H{
		"message": "success",
	})
	return
}
