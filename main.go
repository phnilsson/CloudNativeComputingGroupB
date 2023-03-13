package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"systementor.se/yagolangapi/data"
)

type PageView struct {
	Title  string
	Rubrik string
}

func start(c *gin.Context) {
	c.HTML(http.StatusOK, "home.html", &PageView{Title: "test", Rubrik: "rwef"})
}

// HTML
// JSON

func employeesJson(c *gin.Context) {
	var employees []data.Employee
	data.DB.Find(&employees)
	c.JSON(http.StatusOK, employees)
}

func main() {
	data.InitDatabase()
	router := gin.Default()
	router.LoadHTMLGlob("templates/**")
	router.GET("/", start)
	router.GET("/api/employees", employeesJson)
	router.Run("localhost:8080")

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
