package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/gitlord/REST_API_WITH_GORM/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {


	//setup router
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
