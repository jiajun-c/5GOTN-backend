package handler

import (
	"net/http"
	"os"
	"otn/dal"
	"otn/services/core"
	"otn/services/crud"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gocarina/gocsv"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Record_1 struct {
	WarningId     int64  `csv:"logid" json:"warning_id" xorm:"pk"`      // 日志ID
	WarningCode   int64  `csv:"alarmcode" json:"warning_code"`          // 告警代码
	DeviceID      int64  `csv:"neid" json:"device_id" xorm:"device_id"` // 网元ID
	WarningLevel  int64  `csv:"alarmlevel" json:"warning_level"`        // 告警等级
	DeviceAddress string `csv:"locationinfo" json:"device_address"`     // 线路ID
	DevicePort    string `csv:"lineport" json:"device_port"`            // 线路端口
	WarningTime   string `csv:"occurtime" json:"warning_time"`          // 告警时间
}

func Train(c *gin.Context) {
	addr := viper.GetString("rpc.addr")
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	rFile, err := c.FormFile("file")
	c.SaveUploadedFile(rFile, rFile.Filename)
	var datas []*Record_1
	file, _ := os.Open(rFile.Filename)
	if err := gocsv.UnmarshalFile(file, &datas); err != nil {
		c.JSON(200, gin.H{
			"error": err.Error(),
		})
		return
	}
	for _, value := range datas {
		layout := "2006/1/2 15:04"
		t, _ := time.Parse(layout, value.WarningTime)
		realData := dal.Record{
			WarningId:     value.WarningId,
			WarningCode:   value.WarningCode,
			DeviceID:      value.DeviceID,
			WarningLevel:  value.WarningLevel,
			DeviceAddress: string([]rune(value.DeviceAddress)),
			DevicePort:    value.DevicePort,
			WarningTime:   t,
		}
		crud.InsertRecord(&realData)
		err := crud.InsertEquip(&dal.Equip{
			Id:   value.DeviceID,
			Name: value.DeviceAddress,
		})
		if err != nil {
			log.Error(err)
		}
	}
	data, err := core.PutTrainData(conn)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": data})
}

func StartTrain(c *gin.Context) {
	addr := viper.GetString("rpc.addr")
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	resp, err := core.StartTrain(conn)
	if err != nil {
		log.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": resp})
}
