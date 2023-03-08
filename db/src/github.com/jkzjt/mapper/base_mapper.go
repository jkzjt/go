package mapper

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

const DSN = "root:123456@tcp(127.0.0.1:3306)/go_test?charset=utf8mb4&parseTime=True&loc=Local"

func init() {
	db, err := gorm.Open(mysql.Open(DSN), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}
	DB = db
}
