package config

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	db *gorm.DB
)

func Connect() {

	d, err := gorm.Open("mysql", "root:111111@/mysql?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db = d
}
func GetDB() *gorm.DB {
	return db
}
