package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	arr := []int{1, 2, 4, 5, 6}
	for i, item := range arr {
		fmt.Println(i, item)
	}
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "hllo",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
