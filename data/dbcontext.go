package data

import (
	"os"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() {
	//dsn := "root:my-secret-pw@tcp(127.0.0.1:3307)/emp?charset=utf8mb4&parseTime=True&loc=Local"
	//DB, _ = gorm.Open(mysql.Open(dsn, &gorm.Config{}))
	//
	//os.MkdirAll("database", 0700)
	s := os.Getenv("RUNENVIRONMENT")
	var filePath = ""
	if s == "Production" {
		filePath = "/database/gorm.sqlite"
	} else {
		filePath = "database/gorm.sqlite"
	}

	filePath = "database/gorm.sqlite"
	DB, _ = gorm.Open(sqlite.Open(filePath), &gorm.Config{})
	DB.AutoMigrate(&Employee{})
	var antal int64
	DB.Model(&Employee{}).Count(&antal) // Seed
	if antal == 0 {
		DB.Create(&Employee{Age: 50, Namn: "Stefan", City: "Test"})
		DB.Create(&Employee{Age: 14, Namn: "Oliver", City: "Test"})
		DB.Create(&Employee{Age: 20, Namn: "Josefine", City: "Test"})
	}
}
