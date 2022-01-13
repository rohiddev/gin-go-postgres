package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"go-gin-postgres/config"
	"go-gin-postgres/routes"


)

func main() {
	// Connect DB
	config.Connect()

	// Init Router
	router := gin.Default()

	// Route Handlers / Endpoints
	routes.Routes(router)

	log.Fatal(router.Run(":4747"))
}
