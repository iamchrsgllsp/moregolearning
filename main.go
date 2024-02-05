package main

import (
	"fmt"
	"main/anotherpkg"
	"main/app"
	"main/dbFunc"
)

func Workflow() {
	//Quick check on imports
	// := is a type assertion based on the type of the variable
	data := app.PassingVars()
	// Running the discord.go function Discord
	app.Discord()
	// Using var based on the maths Addtion function
	var data2 = app.Addition(400, 2)
	dbFunc.CreateDB()
	anotherpkg.DryRun()
	// fmt.Println prints to terminal
	fmt.Println(data)
	fmt.Println(data2)
}

func main() {
	// main is entry point to programme - then we run workflow
	Workflow()
}
