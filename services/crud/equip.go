package crud

import (
	"errors"
	"otn/config"
	"otn/dal"
)

func InsertEquip(equip *dal.Equip) error {
	get, err := config.Db.Where("id = ? and name = ?", equip.Id, equip.Name).Get(&dal.Equip{})
	if get {
		return errors.New("the equip is already exists")
	}
	_, err = config.Db.Insert(equip)
	if err != nil {
		return err
	}
	return nil
}

func DeleteEquip(id int64) error {
	_, err := config.Db.ID(id).Delete(&dal.Equip{})
	if err != nil {
		return err
	}
	return nil
}

func GetAllEquip() ([]*dal.Equip, error) {
	equips := make([]*dal.Equip, 0)
	err := config.Db.Where("1=1").Find(&equips)
	if err != nil {
		return nil, err
	}
	return equips, nil
}

func GetOneEquip(id int64) (*dal.Equip, error) {
	var equip dal.Equip
	_, err := config.Db.Get(&equip)
	if err != nil {
		return nil, err
	}
	return &equip, nil
}

func UpdateEquip(equip *dal.Equip) error {
	_, err := config.Db.ID(equip.Id).Update(equip)
	if err != nil {
		return err
	}
	return nil
}
