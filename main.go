package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.New()
	r.Use(func(c *gin.Context) {
		func(r *http.Request, w http.ResponseWriter) {
			fmt.Println("name")
			w.Header().Add("name", "ssss")
		}(c.Request, c.Writer)
		c.Next()
	})
	r.GET("/", func(c *gin.Context) {
		c.JSONP(200, gin.H{"message": "ok"})
	})
	r.Run(":8080")
}
