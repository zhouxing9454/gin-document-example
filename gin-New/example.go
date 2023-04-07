package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.New() //不使用 Logger 和 Recovery 中间件
	r.GET("/hello", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "hello world!",
		})
	})
	r.Run(":8080")
}
