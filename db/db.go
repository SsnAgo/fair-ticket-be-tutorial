package db

import (
	"fmt"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	DB_USER     = "root"
	DB_PASSWORD = "root"
	DB_HOST     = "localhost"
	DB_PORT     = "3306"
	DB_NAME     = "fair-ticket-tutorial"
)

var once sync.Once
var dbInstance *gorm.DB

func DB() *gorm.DB {
	Dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", DB_USER, DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME)
	once.Do(func() {
		db, err := gorm.Open(mysql.Open(Dsn))
		if err != nil {
			panic(err)
		}
		dbInstance = db
	})
	return dbInstance
}
