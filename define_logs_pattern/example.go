package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	r := gin.Default()
	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		log.Printf("endpoint %v_%v_%v_%v\n", httpMethod, absolutePath, handlerName, nuHandlers)
	}
	r.POST("/foo", func(context *gin.Context) {
		context.JSON(http.StatusOK, "foo")
	})
	r.GET("/bar", func(context *gin.Context) {
		context.JSON(http.StatusOK, "bar")
	})
	r.GET("/status", func(context *gin.Context) {
		context.JSON(http.StatusOK, "ok")
	})
	r.Run(":8080")
}
