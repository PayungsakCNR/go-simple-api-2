package main

import (
	"os"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	port := os.Getenv("APP_PORT")
	r := setupRouter()
	r.Run(":" + port)
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/healthz", GetHealthCheck)

	return r
}

func GetHealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": "container-2"})
}

