package crud

import (
	"otn/config"
	"otn/dal"
	"time"
)

func GetAllRecords() ([]*dal.Record, error) {
	records := make([]*dal.Record, 0)
	err := config.Db.Where("1=1").Find(&records)
	if err != nil {
		return nil, err
	}
	return records, nil
}
func SelectRecords(start, end time.Time, device_id int64, device_address string) ([]*dal.Record, error) {
	records := make([]*dal.Record, 0)
	//s := start.Format("2006-01-02 15:04:05")
	//e := end.Format("2006-01-02 15:04:05")
	err := config.Db.Where("warning_time >= ? and warning_time <= ? and device_id = ? and device_address = ?", start, end, device_id, device_address).Table("tb_record").Find(&records)
	if err != nil {
		return nil, err
	}
	return records, nil
}

func InsertRecord(record *dal.Record) {
	_, err := config.Db.Insert(record)
	if err != nil {
		return
	}
}

func DeleteRecord(id int64) error {
	_, err := config.Db.Where("warning_id = ?", id).Delete(&dal.Record{})
	if err != nil {
		return err
	}
	return nil
}
