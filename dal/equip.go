package dal

type Equip struct {
	Id   int64  `json:"id" xorm:"pk" xorm:"pk"`
	Name string `json:"name" form:"name"`
}

func (equip *Equip) TableName() string {
	return "tb_equipment"
}
