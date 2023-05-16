package dal

import "time"

type Record struct {
	WarningId     int64     `json:"warning_id" xorm:"pk"`       // 日志ID
	WarningCode   int64     `json:"warning_code"`               // 告警代码
	DeviceID      int64     `json:"device_id" xorm:"device_id"` // 网元ID
	WarningLevel  int64     `json:"warning_level"`              // 告警等级
	DeviceAddress string    `json:"device_address"`             // 线路ID
	DevicePort    string    `json:"device_port"`                // 线路端口
	WarningTime   time.Time `json:"warning_time"`               // 告警时间
}
type RecordResp struct {
	WarningId      int64  `json:"warning_id"`
	DeviceID       int64  `json:"device_id"`
	DeviceAddress  string `json:"device_address"`
	DevicePort     string `json:"device_port"`
	WarningTime    string `json:"warning_time"`
	WarningTimeEnd string `json:"warning_time_end"`
	WarningCode    int64  `json:"warning_code"`
}

func (record *Record) TableName() string {
	return "tb_record"
}
