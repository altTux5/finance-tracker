package main

import (
	"log"

	"github.com/altTux5/expense-tracker/config"
	"github.com/altTux5/expense-tracker/internal/db"
	"github.com/altTux5/expense-tracker/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	config.LoadConfig()

	db.InitDB()

	r := gin.Default()
	routes.SetupRoutes(r)

	log.Println("Server running on http://localhost:8080")
	err := r.Run(":8080")
	if err != nil {
		log.Fatal("Failed to start server:", err)
	}

}
