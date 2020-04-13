package utils

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func ConnectDB() (*gorm.DB, error) {
	db, err := gorm.Open("mysql", "root:wasd@/drawerapp?charset=utf8&parseTime=True&loc=Local")
	return db, err
}
