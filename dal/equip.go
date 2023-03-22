package dal

type Equip struct {
	Id   int64  `json:"id" xorm:"pk" form:"id"`
	Name string `json:"name" form:"name"`
}

func (equip *Equip) TableName() string {
	return "tb_equipment"
}
