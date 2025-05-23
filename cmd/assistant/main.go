package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello, World!")

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		hoge := "hoge"
		fmt.Println(hoge)
		c.JSON(200, gin.H{"message": "pong!"})
	})
	r.Run() // デフォルトで :8080 で起動
}
