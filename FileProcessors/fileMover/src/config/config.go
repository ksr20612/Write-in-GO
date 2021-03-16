package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
	"io/ioutil"
)
/*
DB Config file
 */
func GetDBConnection() (*gorm.DB, error) {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second,   // Slow SQL threshold
			LogLevel:      logger.Silent, // Log level
			Colorful:      false,         // Disable color
		},
	)

	text, err := ioutil.ReadFile("/home/mz/goTools/changeName/searchDB.txt")
    if err != nil {
    	log.Fatal(err)
    }
    dsn := string(text[:])

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: dsn,
		DefaultStringSize: 512, // default size for string fields
		DisableDatetimePrecision: true, // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex: true, // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn: true, // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false, // auto configure based on currently MySQL version
	}), &gorm.Config{
		Logger: newLogger,
	})
	//fmt.Println(err, gorm.ErrRecordNotFound, " ===== ", err == gorm.ErrRecordNotFound, errors.Is(err, gorm.ErrRecordNotFound))
	if err != nil {
		return nil, err
	}
	return db, nil
}