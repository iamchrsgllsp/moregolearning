package main

import (
	"fmt"
	"main/app"
)

func Workflow() {
	//Quick check on imports
	data := app.PassingVars()
	app.Discord()
  var data2 = app.Addition(400,2)
	fmt.Println(data)
  fmt.Println(data2)
}

func main() {
	Workflow()
}
