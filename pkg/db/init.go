package db

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/youshintop/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func openDB(username, password, addr, name string) *gorm.DB {
	config := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, addr, name)
	log.Info(config)
	database, err := gorm.Open(mysql.Open(config), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	setupDB(database)
	return database
}

func setupDB(db *gorm.DB) {
	db.Logger.LogMode(4)
}

func initDB() *gorm.DB {
	return openDB(viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.address"),
		viper.GetString("db.database"))
}

func Database() *gorm.DB {
	if db == nil {
		db = initDB()
	}
	return db
}
