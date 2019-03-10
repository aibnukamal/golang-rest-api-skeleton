package users

import (
	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to the gin Engine
func ApplyRoutes(r *gin.RouterGroup) {
	posts := r.Group("/users")
	{
		posts.GET("/", list)
	}
}

func list(c *gin.Context) {

	c.JSON(200, gin.H{
		"message": "pong",
	})

	return
}
