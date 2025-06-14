package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterHealthRoutes(r *gin.Engine) {
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Looking for something?",
		})
	})

}
