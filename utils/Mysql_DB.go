package utils

import (
	"log"
	"time"
	"os"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/driver/mysql"
)

var (
	DB_MySQL *gorm.DB
	//DB_Redis *redis.Client
)

func init() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)
	//log.Println("root:123456@tcp(127.0.0.1:3306)/LibrarySystem?charset=utf8&parseTime=True&loc=Local")

	db, err := gorm.Open(mysql.Open("root:123456@tcp(127.0.0.1:3306)/LibrarySystem?charset=utf8&parseTime=True&loc=Local"),
		&gorm.Config{Logger: newLogger})
	if err != nil {
		panic("failed to connect database")
	}
	DB_MySQL = db
}
