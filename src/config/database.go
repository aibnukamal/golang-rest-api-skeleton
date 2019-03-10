package database

import (
	"accounts/src/models"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	// Register mysql dialect
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Initialize initializes the database
func Initialize() (*gorm.DB, error) {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	name := os.Getenv("DB_NAME")

	connectionstring := fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", user, password, name)

	db, err := gorm.Open("mysql", connectionstring)

	// logs SQL
	db.LogMode(true)

	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to database")
	models.Migrate(db)

	return db, err
}

// Inject injects database to gin context
func Inject(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	}
}
