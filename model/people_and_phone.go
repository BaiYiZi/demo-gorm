package model

type People struct {
	ID   uint   `gorm:"primarykey"`
	Name string `gorm:"colume:name" json:"name"`
}

func (People) TableName() string {
	return "table_people"
}

type Phone struct {
	ID       uint   `gorm:"primarykey"`
	PeopleID uint   `gorm:"people_id" json:"people_id"`
	Code     string `gorm:"code" json:"code"`
}

func (Phone) TableName() string {
	return "table_phone"
}
