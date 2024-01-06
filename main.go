// main.go
package main

import (
	"fmt"
	"go-microservice/db"
	"go-microservice/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	err := db.InitMongoDB()
	if err != nil {
		log.Fatal(err)
	}

	router := gin.Default()

	// Register routes
	routes.RegisterPeopleRoutes(router)

	fmt.Println("Server is running on :8080...")
	router.Run(":8080")
}
