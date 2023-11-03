package model

import "time"

type User struct {
	ID      uint      `gorm:"primarykey" json:"id"`
	Name    string    `gorm:"column:name" json:"name"`
	Age     int       `gorm:"column:age" json:"age"`
	Address string    `gorm:"column:address" json:"address"`
	Born    time.Time `gorm:"column:born" json:"born"`
}

func (User) TableName() string {
	return "table_user"
}
