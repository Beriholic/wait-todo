package dao

import (
	"fmt"
	"wait-to-do/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var _db *gorm.DB

func InitDB() {
	host := config.DbHost
	port := config.DbPort
	user := config.DbUser
	password := config.DbPassword
	name := config.DbName
	Charset := config.Charset

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True", user, password, host, port, name, Charset)

	if err := DataBase(dsn); err != nil {
		fmt.Println("InitDb err\n dsn : %v\n  err:%v\n", dsn, err)
	}
}

func DataBase(connString string) error {
	var gormLogger logger.Interface
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                     connString,
		DefaultStringSize:       256,
		DontSupportRenameIndex:  true,
		DontSupportRenameColumn: true,
	}), &gorm.Config{
		Logger: gormLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err != nil {
		return err
	}

	_db = db

	if err := migration(); err != nil {
		fmt.Println("Database migration failed,err is :", err)
	}

	return nil
}
