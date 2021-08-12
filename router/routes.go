package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(r *gin.RouterGroup) {
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "server is alive",
		})
	})
}
