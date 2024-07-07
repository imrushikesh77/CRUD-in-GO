package main

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Database *gorm.DB
var urlDSN = "goapp:2004@tcp(localhost:3306)/employee_db?charset=utf8mb4&parseTime=True&loc=Local"
var err error

func DataMigration() {
	Database, err = gorm.Open(mysql.Open(urlDSN), &gorm.Config{})
	if err != nil {
		log.Printf("DSN used: %s", urlDSN)
		log.Printf("Database connection error: %v", err)
		panic("Could not connect to the database")
	}

	err = Database.AutoMigrate(&Employee{})
	if err != nil {
		log.Printf("Database migration error: %v", err)
		panic("Could not migrate the database")
	}
}
