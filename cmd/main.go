package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	routes "restful.api.e-commerce.golang/router"
)

func init() {
	switch mode := os.Getenv("GIN_MODE"); mode {
	case "release":
		gin.SetMode(gin.ReleaseMode)

	case "test":
		gin.SetMode(gin.TestMode)

	default:
		gin.SetMode(gin.DebugMode)
	}
}

func main() {
	r := gin.Default()
	r.StaticFS("i", http.Dir("image"))
	routes.Index(r.Group("eCommerce"))
	r.Run(":" + os.Getenv("PORT"))
}
