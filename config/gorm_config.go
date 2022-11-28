package config

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitGorm() *gorm.DB {
	//util.ProcessEnv()

	host := os.Getenv("DB_HOST")
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		username,
		password,
		host,
		port,
		dbname,
	)

	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		logrus.Error("Can't connect mysql database!")
		panic(err)
	}

	MigrateDB(db)

	return db
}
