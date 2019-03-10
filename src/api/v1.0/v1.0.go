package apiv1

import (
	"accounts/src/api/v1.0/users"
	"github.com/gin-gonic/gin"
)

func health(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "success",
	})
}

// ApplyRoutes applies router to the gin Engine
func ApplyRoutes(r *gin.RouterGroup) {
	v1 := r.Group("/v1.0")
	{
		v1.GET("/health-check", health)
		users.ApplyRoutes(v1)
	}
}
