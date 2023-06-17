package dal

import (
	_ "github.com/gocarina/gocsv"
	"time"
)

type Record struct {
	WarningId     int64     `csv:"logid" json:"warning_id" xorm:"pk"`      // 日志ID
	WarningCode   int64     `csv:"alarmcode" json:"warning_code"`          // 告警代码
	DeviceID      int64     `csv:"neid" json:"device_id" xorm:"device_id"` // 网元ID
	WarningLevel  int64     `csv:"alarmlevel" json:"warning_level"`        // 告警等级
	DeviceAddress string    `csv:"locationinfo" json:"device_address"`     // 线路ID
	DevicePort    string    `csv:"lineport" json:"device_port"`            // 线路端口
	WarningTime   time.Time `csv:"occurtime" json:"warning_time"`          // 告警时间
}
type RecordResp struct {
	WarningId     int64  `json:"warning_id"`
	WarningCode   int64  `json:"warning_code"` // 告警代码
	DeviceID      int64  `json:"device_id"`
	WarningLevel  int64  `json:"warning_level"`  // 告警等级
	DeviceAddress string `json:"device_address"` // 线路ID
	DevicePort    string `json:"device_port"`    // 线路端口
	WarningTime   string `json:"warning_time"`   // 告警时间
}

func (record *Record) TableName() string {
	return "tb_record"
}
