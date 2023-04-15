package data

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/glebarez/sqlite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

const arrayLength = 20

var firstNames = [arrayLength]string{
	"Emma", "Olivia", "Amelia", "Harper", "Sophia", "Liam", "Noah", "Oliver", "Elijah", "Lucas",
	"Anaya", "Aarav", "Reyansh", "Advik", "Vihaan", "Sofía", "María", "Martina", "Juan", "José",
}

var lastNames = [arrayLength]string{
	"Johnson", "Smith", "Brown", "Davis", "Taylor", "Jackson", "Williams", "Thomas", "Martinez", "Clark",
	"Wilson", "Patel", "Kumar", "Gupta", "Sharma", "Rodríguez", "García", "Pérez", "López", "González",
}

var cities = [arrayLength]string{
	"New York City", "London", "Paris", "Tokyo", "Sydney", "Rio de Janeiro", "Moscow", "Rome", "Istanbul", "Toronto",
	"Mumbai", "Shanghai", "Cape Town", "Buenos Aires", "Cairo", "Berlin", "Bangkok", "Dubai", "Athens", "Mexico City"}

var teamNames = [arrayLength]string{
	"Manchester United", "FC Barcelona", "Paris Saint-Germain", "Bayern Munich", "Juventus", "Real Madrid",
	"Liverpool", "AC Milan", "Chelsea", "Borussia Dortmund", "Ajax", "Atletico Madrid", "Tottenham Hotspur",
	"AS Roma", "Inter Milan", "Benfica", "Olympique Marseille", "Boca Juniors", "Flamengo", "Club América",
}

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

func randomArrayItem(array *[arrayLength]string) string {
	return array[rand.Intn(arrayLength)]
}

func existsIn(value *int, list *[]int) bool {
	for _, element := range *list {
		if *value == element {
			return true
		}
	}
	return false
}

func jerseyNumber(teamId int) int {
	var numberSuggestion int
	var jerseyNumbers []int
	const maxAttempts = 1000
	DB.Table("players").Where("team_id = ?", teamId).Pluck("jersey_number", &jerseyNumbers)

	for i := 0; i < maxAttempts; i++ {
		numberSuggestion = rand.Intn(30)
		if !existsIn(&numberSuggestion, &jerseyNumbers) {
			return numberSuggestion
		}
	}
	return -1
}

func createPlayersInTeam(teamId int) {
	const numberOfPlayers = 24
	for i := 0; i < numberOfPlayers; i++ {
		DB.Create(&Player{
			Name:         randomArrayItem(&firstNames) + " " + randomArrayItem(&lastNames),
			BirthYear:    rand.Intn(100) + 1970,
			TeamId:       teamId,
			JerseyNumber: jerseyNumber(teamId),
		})
	}
}

func createTeams() {
	for i := 0; i < 9; i++ {
		DB.Create(&Team{
			Name:        randomArrayItem(&teamNames),
			FoundedYear: 1900 + rand.Intn(100),
			City:        randomArrayItem(&cities),
		})
	}
}

func seedData() {
	var ids []int
	rand.Seed(time.Now().UnixNano())
	createTeams()

	DB.Model(&Team{}).Pluck("id", &ids)
	for _, id := range ids {
		createPlayersInTeam(id)
	}
}

func InitDatabase(file, server, database, username, password string, port int) {
	if len(file) == 0 {
		DB = openMySql(server, database, username, password, port)
	} else {
		DB, _ = gorm.Open(sqlite.Open(file), &gorm.Config{})
	}

	DB.AutoMigrate(&Employee{})
	DB.AutoMigrate(&Team{})
	DB.AutoMigrate(&Player{})

	var antal int64
	DB.Model(&Employee{}).Count(&antal) // Seed
	if antal == 0 {
		DB.Create(&Employee{Age: 52, Namn: "Dorota", City: "Stokholm"})
		DB.Create(&Employee{Age: 32, Namn: "Philip", City: "Linköping"})
		DB.Create(&Employee{Age: 44, Namn: "Christian", City: "Flyinge"})
	}
	// Seed in case we don't have any teams
	var numberOfTeams int64
	DB.Model(&Team{}).Count(&numberOfTeams)
	if numberOfTeams == 0 {
		seedData()
	}
}
