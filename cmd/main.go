package main

import (
	"fmt"
	"os"

	"github.com/Eduardo-Brehm/mobile-api-go/internal/config"
	"github.com/Eduardo-Brehm/mobile-api-go/internal/controllers"
	"github.com/Eduardo-Brehm/mobile-api-go/internal/db"
	"github.com/Eduardo-Brehm/mobile-api-go/internal/routes"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	//load the environment variables
	godotenv.Load()

	//connect to the database
	database, err := config.ConnectDB()
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		os.Exit(1)
	}

	// Create repositories
	userRepo := db.NewUserRepository(database)

	// Create controllers
	authController := controllers.NewAuthController(userRepo)

	// Create echo instance
	e := echo.New()

	// Setup routes
	routes.SetupAuthRoutes(e, authController)

	defer database.Close()

	fmt.Println("Server started on port 3000")

	// Start the server on port 3000
	e.Start(":3000")
}
