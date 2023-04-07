package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/test", func(context *gin.Context) {
		context.Redirect(http.StatusMovedPermanently, "http://www.google.com/")
	})
	r.POST("/test", func(context *gin.Context) {
		context.Redirect(http.StatusNotFound, "https://www.youandgentleness.cn/")
	})
	r.GET("/test2", func(context *gin.Context) {
		context.Request.URL.Path = "/test3"
		r.HandleContext(context)
	})
	r.GET("/test3", func(context *gin.Context) {
		context.JSON(200, gin.H{"hello": "world"})
	})

	r.Run(":8080")
}
