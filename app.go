package main

import (
	"accounts/src/api"
	database "accounts/src/config"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// load .env environment variables
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	// initializes database
	db, _ := database.Initialize()

	port := os.Getenv("PORT")

	// create gin app
	app := gin.Default()
	app.Use(database.Inject(db))

	// apply api router
	api.ApplyRoutes(app)

	// listen to given port
	app.Run(":" + port)
}
