package main

import (
	"fmt"
	"log"

	"github.com/FallenCap/url-shortner/config"
	"github.com/FallenCap/url-shortner/models"
	"github.com/FallenCap/url-shortner/routes"
)

func main() {
	config.InitDB()

	err := config.DB.AutoMigrate(&models.Url{}); if err != nil {
		log.Fatal("❌ AutoMigrate failed:", err)
	}

	fmt.Println("Connected to DB:", config.DB.Migrator().CurrentDatabase())

	r := routes.SetupRouter()

	r.Run(":6060")
}
