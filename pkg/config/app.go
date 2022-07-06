package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	db *gorm.DB
)

func Connect() {
	dbUri := "host=localhost user=postgres dbname=gotest sslmode=disable password=password"
	d, err := gorm.Open("postgres",dbUri)
	if err !=nil{
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB{
	return db
}
