package main

import (
	"fmt"

	"github.com/cheynewallace/tabby"
	"systementor.se/yagolangapi/data"
)

func main() {
	e := data.Employee{
		Age:  1,
		City: "Strefabn",
		Namn: "wddsa",
	}

	if e.IsCool() {
		fmt.Printf("Namn is cool:%s\n", e.Namn)
	} else {
		fmt.Printf("Namn:%s\n", e.Namn)
	}

	fmt.Println("Hello")
	t := tabby.New()
	t.AddHeader("Namn", "Age", "City")
	t.AddLine("Stefan", "50", "Stockholm")
	t.AddLine("Oliver", "14", "Stockholm")
	t.Print()
}
