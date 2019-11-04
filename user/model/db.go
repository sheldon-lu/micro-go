package model

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func CreateConnection() (*gorm.DB, error) {
	host := "127.0.0.1"
	user := "root"
	dbName := "shopping"
	password := "lsnxy1994"
	return gorm.Open("mysql", fmt.Sprintf(
		"%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local",
		user, password, host, dbName,
	),
	)
}
