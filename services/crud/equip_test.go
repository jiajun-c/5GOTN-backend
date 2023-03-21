package crud

import (
	"otn/config"
	"otn/dal"
	"testing"
)

func TestInsert(t *testing.T) {
	config.LoadDB()
	err := InsertEquip(&dal.Equip{
		Id:   0,
		Name: "test",
	})
	if err != nil {
		t.Error(err)
	}
}

func TestDelete(t *testing.T) {
	config.LoadDB()
	err := DeleteEquip(0)
	if err != nil {
		t.Error(err)
	}
}

func TestGetAll(t *testing.T) {
	config.LoadDB()
	equips, err := GetAllEquip()
	if err != nil {
		t.Error(err)
	}
	println(equips)
}

func TestGet(t *testing.T) {
	config.LoadDB()
	equip, err := GetOneEquip(2)
	if err != nil {
		t.Error(err)
	}
	println(equip)
}
