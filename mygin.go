package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)
func main() {

	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET"}
	r.Use(cors.New(config))
    // r.Use(cors.Default())
	r.GET("/api/v1/ping", func(c *gin.Context) {

            c.JSON(200, gin.H {

				"message": "pong",
			})
	})

    r.Run()

}