package config

import (
	"fmt"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"otn/dal"
	"xorm.io/xorm"
)

var Db *xorm.Engine

func CheckTable(table interface{}) {
	exist, _ := Db.IsTableExist(table)
	if !exist {
		err := Db.CreateTables(table)
		if err != nil {
			panic("failed to init the database")
		}
	}
}
func LoadDB() {
	var err error
	driver := viper.GetString("db.driver")
	dsn := viper.GetString("db.dsn")
	fmt.Println(dsn)
	Db, err = xorm.NewEngine(driver, dsn)
	if err != nil {
		panic("failed to initialize the database")
	}
	//Db.SetMapper(names.GonicMapper{})
	CheckTable(dal.Equip{})
}
