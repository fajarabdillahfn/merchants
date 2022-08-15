package postgre

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = 3306
	user     = "root"
	password = ""
	dbname   = "merchant"
)

// InitDB is
func OpenDB() *gorm.DB {
	var err error
	sqlInfo := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
	user, password, host, port, dbname)

	openConnection := mysql.Open(sqlInfo)

	DB, err := gorm.Open(openConnection, &gorm.Config{})
	DB = DB.Set("gorm:auto_preload", true)
	if err != nil {
		panic("failed to connect database")
	}

	return DB
}
