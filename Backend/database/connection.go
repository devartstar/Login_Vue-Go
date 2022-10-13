package database

import (
	"github.com/devartstar/login/models"
	"gorm.io/gorm"
  	"gorm.io/driver/mysql"
)

var DB *gorm.DB

func Connect() {
	connection, err := gorm.Open(mysql.Open("root:rootroot@/yt_go_auth"), &gorm.Config{})

	if err != nil {
		panic("could not open connection")
	}

	DB = connection

	connection.AutoMigrate(&models.User{})
}