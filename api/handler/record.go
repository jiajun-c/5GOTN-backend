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
		wt, wte := record.WarningTime.Format("2006/01/02 15:04"), record.WarningTimeEnd.Format("2006/01/02 15:04")
		resp := &dal.RecordResp{
			WarningId:      record.WarningId,
			DeviceID:       record.DeviceID,
			DeviceAddress:  record.DeviceAddress,
			DevicePort:     record.DevicePort,
			WarningTime:    wt,
			WarningTimeEnd: wte,
			WarningCode:    record.WarningCode,
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
	device_id := ctx.Query("device_id")
	did, _ := strconv.Atoi(device_id)
	device_address := ctx.Query("device_address")
	warning_time_begin := ctx.Query("warning_time_begin")
	warning_time_end := ctx.Query("warning_time_end")
	start, _ := time.Parse("2006/01/02 15:04", warning_time_begin)
	end, _ := time.Parse("2006/01/02 15:04", warning_time_end)
	records, err := crud.SelectRecords(start, end, int64(did), device_address)
	if err != nil {
		ctx.JSON(200, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, records)
}

func InsertRecord(ctx *gin.Context) {
	warning_id := ctx.PostForm("warning_id")
	device_id := ctx.PostForm("device_id")
	device_port := ctx.PostForm("device_port")
	device_address := ctx.PostForm("device_address")
	warning_time := ctx.PostForm("warning_time")
	warning_time_end := ctx.PostForm("warning_time_end")
	warning_code := ctx.PostForm("warning_code")
	wid, _ := strconv.Atoi(warning_id)
	did, _ := strconv.Atoi(device_id)
	wtime, _ := time.Parse("2006/01/02 15:04", warning_time)
	wtime_end, _ := time.Parse("2006/01/02 15:04", warning_time_end)
	wcode, _ := strconv.Atoi(warning_code)
	crud.InsertRecord(&dal.Record{
		WarningId:      int64(wid),
		DeviceID:       int64(did),
		DeviceAddress:  device_address,
		DevicePort:     device_port,
		WarningTime:    wtime,
		WarningTimeEnd: wtime_end,
		WarningCode:    int64(wcode),
	})
	ctx.JSON(200, gin.H{
		"message": "success",
	})
	return
}
