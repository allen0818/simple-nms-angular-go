package main

import (
	"fmt"
	"io"
	"nms-api/middleware"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {

	gin.DisableConsoleColor()
	file, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(file, os.Stdout)

	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// router.GET("/hello/:name", func(c *gin.Context) {
	// 	name := c.Param("name")
	// 	c.String(200, "hello %s", name)
	// })

	demoApi := router.Group("/demo")

	demoApi.Use(
		middleware.PrintHello("1"),
		middleware.PrintHello("2"),
	)

	v1 := demoApi.Group("v1")

	v1.GET("/hello", func(c *gin.Context) {
		name := c.Query("name")
		age := c.DefaultQuery("age", "20")

		fmt.Printf("%s (%s)\n", name, age)
		c.String(200, "hello %s (%s)", name, age)
	})

	v1.POST("/form", func(c *gin.Context) {
		name := c.PostForm("name")
		age := c.DefaultPostForm("age", "0")

		c.JSON(200, gin.H{
			"name": name,
			"age":  age,
		})
	})

	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
