package database

import (
"fmt"
"github.com/jinzhu/gorm"
_ "github.com/jinzhu/gorm/dialects/mysql"
"github.com/micro/go-micro/config"
"github.com/micro/go-micro/util/log"
)
// fmt.Println(conf)

func CreateConnection() (*gorm.DB, error) {
	//加载配置项
	err := config.LoadFile("../common-config/config.json")
	if err != nil {
		log.Fatalf("Could not load config file: %s", err.Error())
	}
	conf := config.Map()["mysql"]
	// fmt.Println(conf.(map[string]interface{})["host"])


	host := conf.(map[string]interface{})["host"]
	port := conf.(map[string]interface{})["port"]
	user := conf.(map[string]interface{})["user"]
	dbName := conf.(map[string]interface{})["database"]
	password := conf.(map[string]interface{})["password"]
	return gorm.Open("mysql", fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user, password, host, port, dbName,
		),
	)
}


