package config

import (
	"GoSosmed/entity"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func LoadDB() {
	connectionStr := fmt.Sprintf("%v:%v@tcp(%v)/%v?%v", ENV.DB_USERNAME, ENV.DB_PASSWORD, ENV.DB_URL, ENV.DB_DATABASE, "charset=utf8mb4&parseTime=True&loc=Asia%2FJakarta")
	db, err := gorm.Open(mysql.Open(connectionStr), &gorm.Config{})
	if err != nil {
		panic("gagal konek database")
	}

	DB = db

	err = DB.AutoMigrate(
		&entity.User{},
		&entity.Post{},
	)
	if err != nil {
		log.Fatal("tidak bisa migrate")
	}
}
