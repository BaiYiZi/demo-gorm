package ping

import (
	"fmt"

	"github.com/BaiYiZi/learn-gorm/database"
)

func Ping(name string) (any, error) {
	db := database.GetDB()

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("%s", "get instance of sql.DB error")
	}

	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("%s", "dead connection to database")
	}

	return "hello " + name, nil
}
