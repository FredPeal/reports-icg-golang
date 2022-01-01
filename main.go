package main

import (
	 "github.com/gin-gonic/gin"
	 controllers "reportsicg/controllers"
	 "github.com/gin-contrib/cors"

	)

func main() {
	router := gin.Default()
	router.Use(cors.Default())
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Woot Hola",
		})
	})
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.GET("/reports/home", controllers.Resume)

	router.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}