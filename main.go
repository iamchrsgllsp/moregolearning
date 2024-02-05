package main

import (
	"fmt"
	"main/app"
)

func main() {
	data := app.PassingVars()
	app.Discord()
	fmt.Println(data)
}
