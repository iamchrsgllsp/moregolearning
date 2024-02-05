package main

import (
	"fmt"
	"main/anotherpkg"
	"main/app"
	"main/dbFunc"
	"time"
)

func Workflow() {
	//Quick check on imports
	// := is a type assertion based on the type of the variable
	data := app.PassingVars()
	// Running the discord.go function Discord
	app.Discord()
	// Using var based on the maths Addtion function
	var data2 = app.Addition(400, 2)
	anotherpkg.DryRun()
	// fmt.Println prints to terminal
	fmt.Println(data)
	fmt.Println(data2)
}

func DbActions() {
	dbFunc.CreateDB()
	dbFunc.AddtoDB("Raidy")
	dbFunc.AddtoDB("Chris")
  dbFunc.AddtoDB("Goose")
  dbFunc.AddtoDB("Bun")
  dbFunc.AddtoDB("Dave")
  time.Sleep(10)
	dbFunc.ReadDB()
}

func main() {
	// main is entry point to programme - then we run our actions
	Workflow()
  DbActions()
	

}
