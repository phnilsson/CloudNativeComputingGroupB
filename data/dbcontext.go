package data

import (
	"fmt"

	"github.com/glebarez/sqlite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func openMySql(server, database, username, password string, port int) *gorm.DB {
	var url string
	url = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		username, password, server, port, database)

	db, err := gorm.Open(mysql.Open(url), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}

func InitDatabase(file, server, database, username, password string, port int) {
	if len(file) == 0 {
		DB = openMySql(server, database, username, password, port)
	} else {
		DB, _ = gorm.Open(sqlite.Open(file), &gorm.Config{})
	}
	DB.AutoMigrate(&Employee{})
	var antal int64
	DB.Model(&Employee{}).Count(&antal) // Seed
	if antal == 0 {
		DB.Create(&Employee{Age: 50, Namn: "Stefan", City: "Test"})
		DB.Create(&Employee{Age: 14, Namn: "Oliver", City: "Test"})
		DB.Create(&Employee{Age: 20, Namn: "Josefine", City: "Test"})
		DB.Create(&Employee{Age: 52, Namn: "Dorota", City: "Stokholm"})
		DB.Create(&Employee{Age: 32, Namn: "Philip", City: "Linköping"})
	}
	DB.AutoMigrate(&Team{})
	var teamAntal int64
	DB.Model(&Team{}).Count(&teamAntal) // Seed
	if teamAntal == 0 {
		DB.Create(&Team{FoundedYear: 1900, Name: "Bayern München", City: "München"})
		DB.Create(&Team{FoundedYear: 1880, Name: "Manchester City", City: "Manchester"})
		DB.Create(&Team{FoundedYear: 1980, Name: "Napoli", City: "Napoli"})
		DB.Create(&Team{FoundedYear: 1899, Name: "Barcelona", City: "Barcelona "})

	}
}
