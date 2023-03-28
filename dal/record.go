package dal

import "time"

type Record struct {
	WarningId      int64     `json:"warning_id" xorm:"pk"`
	DeviceID       int64     `json:"device_id" xorm:"device_id"`
	DeviceAddress  string    `json:"device_address"`
	DevicePort     string    `json:"device_port"`
	WarningTime    time.Time `json:"warning_time"`
	WarningTimeEnd time.Time `json:"warning_time_end"`
	WarningCode    int64     `json:"warning_code"`
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
