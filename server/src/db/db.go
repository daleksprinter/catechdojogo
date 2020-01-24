package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	DB *gorm.DB
)

func NewDB(datasource string) {
	db, err := gorm.Open("mysql", datasource)
	if err != nil {
		panic(err.Error())
	}

	DB = db
}
