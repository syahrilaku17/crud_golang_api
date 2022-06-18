package database

import (
	"awesomeProject/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var Instance *gorm.DB
var err error

func Connect(connectionString string) {
	Instance, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
		panic("failed to connect database")
	}

}

func Migrate() {
	Instance.AutoMigrate(&model.Product{})
	log.Fatalln("Migration complete")
}
