package models

import (
	"core-go/config"
	"github.com/jinzhu/gorm"
	"github.com/gonrails/gonrails/models/mysql"
)

var (
	db *gorm.DB
)

func init() {

	db = mysql.Open(
		config.MySQL.Host,
		config.MySQL.Port,
		config.MySQL.Username,
		config.MySQL.Password,
		config.MySQL.Name,
	)

	mysql.ConfigureMySQL(db, config.MySQL.Idles, config.MySQL.Connections)
}

// DB returns the package inner variable db
func DB() *gorm.DB {
	return db
}
