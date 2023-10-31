package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	Db  *gorm.DB
	err error
)

func InitMysql() {
	Db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
}
