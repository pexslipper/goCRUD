package main

import "github.com/gin-gonic/gin"

type Users struct {
	ID    string
	Fname string
	Lname string
	Score float64
}

func main() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.Run() // listen and serve on 0.0.0.0:8080
}
