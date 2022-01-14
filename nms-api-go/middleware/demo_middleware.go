package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func PrintHello(s string) gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("hello-" + s) // before request

		c.Next() // execute next middleware handler

		fmt.Println("bye-" + s) // after request
	}
}
