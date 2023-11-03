package model

type Boy struct {
	ID           uint   `gorm:"primarykey" json:"id"`
	Name         string `gorm:"column:name" json:"name"`
	Age          int    `gorm:"column:age" json:"age"`
	GirlFriendID uint   `gorm:"column:girl_friend_id" json:"girl_friend_id"`
}

func (Boy) TableName() string {
	return "table_boy"
}

type Girl struct {
	ID   uint   `gorm:"primarykey" json:"id"`
	Name string `gorm:"column:name" json:"name"`
	Age  int    `gorm:"column:age" json:"age"`
}

func (Girl) TableName() string {
	return "table_girl"
}
