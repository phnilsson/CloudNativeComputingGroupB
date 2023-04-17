package main

import (
	"net/http"

	"math/rand"
	"time"

	"github.com/Pallinder/go-randomdata"
	"github.com/gin-gonic/gin"
	"systementor.se/yagolangapi/data"
)

type PageView struct {
	Title  string
	Rubrik string
}

var theRandom *rand.Rand

func start(c *gin.Context) {
	c.HTML(http.StatusOK, "home.html", &PageView{Title: "Grupp B", Rubrik: "Grupp B - nu kör vi  Golang"})
}

// HTML
// JSON

func employeesJson(c *gin.Context) {
	var employees []data.Employee
	data.DB.Find(&employees)

	c.JSON(http.StatusOK, employees)
}

func christianJson(c *gin.Context) {
	var christian data.Employee
	data.DB.Where("Namn = ?", "Christian").First(&christian)

	c.JSON(http.StatusOK, struct {
		Name string
		City string
	}{
		Name: christian.Namn,
		City: christian.City})
}

func dorotaJson(c *gin.Context) {
	var dorota data.Employee
	data.DB.Where("Namn = ?", "Dorota").First(&dorota)

	c.JSON(http.StatusOK, struct {
		Name string
		City string
	}{
		Name: dorota.Namn,
		City: dorota.City})
}

func allJson(c *gin.Context) {
	var employees []data.Employee
	data.DB.Find(&employees)
	// An City slice to hold data from returned employees.
	var cities []data.CityName
	// Loop through employees, using Scan to assign column data to struct fields.
	for i := 0; i < len(employees); i++ {

		e := data.CityName{
			City: employees[i].City,
			Namn: employees[i].Namn,
		}
		cities = append(cities, e)
	}

	c.JSON(http.StatusOK, cities)
}

func philipJson(c *gin.Context) {
	var philipEmployee data.Employee
	data.DB.Where("Namn = ?", "Philip").First(&philipEmployee)
	c.JSON(http.StatusOK, struct {
		Name string
		City string
	}{
		Name: philipEmployee.Namn,
		City: philipEmployee.City})
}

func teamsJson(c *gin.Context) {
	var teams []data.Team
	data.DB.Find(&teams)
	var names []data.Name
	// Loop through employees, using Scan to assign column data to struct fields.
	for i := 0; i < len(teams); i++ {

		e := data.Name{
			Name: teams[i].Name,
		}
		names = append(names, e)
	}

	c.JSON(http.StatusOK, names)
}

func addEmployee(c *gin.Context) {

	data.DB.Create(&data.Employee{Age: theRandom.Intn(50) + 18, Namn: randomdata.FirstName(randomdata.RandomGender), City: randomdata.City()})

}

func addManyEmployees(c *gin.Context) {
	//Here we create 10 Employees
	for i := 0; i < 10; i++ {
		data.DB.Create(&data.Employee{Age: theRandom.Intn(50) + 18, Namn: randomdata.FirstName(randomdata.RandomGender), City: randomdata.City()})
	}

}

func getTeamName(teamID int) string {
	var team data.Team
	data.DB.Find(&team, teamID)
	return team.Name
}

func getPlayerInformation(c *gin.Context) {
	// Get the player ID from the URL path
	playerID := c.Param("playerid")

	var player data.Player
	data.DB.Find(&player, playerID)

	c.JSON(http.StatusOK, struct {
		Name         string
		BirthYear    int
		Team         string
		JerseyNumber int
	}{
		Name:         player.Name,
		BirthYear:    player.BirthYear,
		JerseyNumber: player.JerseyNumber,
		Team:         getTeamName(player.TeamId)})
}

var config Config

func main() {
	theRandom = rand.New(rand.NewSource(time.Now().UnixNano()))
	readConfig(&config)

	data.InitDatabase(config.Database.File,
		config.Database.Server,
		config.Database.Database,
		config.Database.Username,
		config.Database.Password,
		config.Database.Port)

	router := gin.Default()
	router.LoadHTMLGlob("templates/**")
	router.GET("/", start)
	router.GET("/api/employees", employeesJson)
	router.GET("/api/addemployee", addEmployee)
	router.GET("/api/addmanyemployees", addManyEmployees)
	router.GET("/api/dorota", dorotaJson)
	router.GET("/api/christian", christianJson)
	router.GET("/api/philip", philipJson)
	router.GET("/api/all", allJson)
	router.GET("/api/team", teamsJson)
	//router.GET("/api/team/:teamid", getTeamInformation)
	router.GET("/api/player/:playerid", getPlayerInformation)
	router.Run(":8080")

	// e := data.Employee{
	// 	Age:  1,
	// 	City: "Strefabn",
	// 	Namn: "wddsa",
	// }

	// if e.IsCool() {
	// 	fmt.Printf("Namn is cool:%s\n", e.Namn)
	// } else {
	// 	fmt.Printf("Namn:%s\n", e.Namn)
	// }

	// fmt.Println("Hello")
	// t := tabby.New()
	// t.AddHeader("Namn", "Age", "City")
	// t.AddLine("Stefan", "50", "Stockholm")
	// t.AddLine("Oliver", "14", "Stockholm")
	// t.Print()
}
