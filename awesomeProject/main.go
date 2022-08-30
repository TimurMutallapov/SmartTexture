package main

import (
	"awesomeProject/dataBase"
	"awesomeProject/routes"
)

func main() {
	dataBase.Setup()
	routes.Setup()
}
