package database

import (
	"fmt"

	"github.com/BaiYiZi/learn-gorm/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func GetDB() *gorm.DB {
	return db
}

var dbConfig = struct {
	username string
	password string
	ip       string
	port     string
	dbName   string
	setting  string
}{
	username: "root",
	password: "root",
	ip:       "localhost",
	port:     "3307",
	dbName:   "demo_gorm",
	setting:  "charset=utf8mb4&parseTime=True&loc=UTC",
}

func init() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s", dbConfig.username, dbConfig.password, dbConfig.ip, dbConfig.port, dbConfig.dbName, dbConfig.setting)

	_db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	_db.AutoMigrate(&model.User{})

	_db.AutoMigrate(&model.Boy{})
	_db.AutoMigrate(&model.Girl{})

	_db.AutoMigrate(&model.People{})
	_db.AutoMigrate(&model.Phone{})

	_db.AutoMigrate(&model.Student{})
	_db.AutoMigrate(&model.Course{})
	_db.AutoMigrate(&model.SelectCourse{})

	db = _db
}
