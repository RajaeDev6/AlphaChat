package main

import (
	"github.com/RajaeDev6/AlphaChat/server/database"
	"github.com/RajaeDev6/AlphaChat/server/routes"
)

func main() {

	database.Initialize()

	routes.Configure()
}
