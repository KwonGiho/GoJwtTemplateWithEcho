package config

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/gommon/log"
	"../api/handlers"
)

func InitDB() {

}
func initDB() {
	db := initializeDatabase()
	initializeTable(db)
}
func initializeDatabase() gorm.DB {
	log.Debug("gorm.DB init")

	db, error := gorm.Open("mysql", "root:zxcv1234@/goApi?charset=utf8&parseTime=True&loc=Local")
	if error != nil {
		log.Fatal(error)
	}

	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	return *db
}

func initializeTable(db gorm.DB) {
	var user handlers.User

	log.Info("initialize tables")

	db.DropTableIfExists(user)
	db.CreateTable(user)
}