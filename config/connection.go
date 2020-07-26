package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func ConnectDB() *gorm.DB {
	var err error

	db, err := gorm.Open("mysql","root:@/learngin?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic("failed connect to db")
	}

	return db
}