package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type form struct {
	User     string `form:"user" json:"user" xml:"user" binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

func main() {
	router := gin.Default()
	router.POST("/loginJSON", func(context *gin.Context) {
		var json form
		if err := context.ShouldBindJSON(&json); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if json.User != "manu" || json.Password != "123" {
			context.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}
		context.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	})
	router.POST("/loginXML", func(context *gin.Context) {
		var xml form
		if err := context.ShouldBindXML(&xml); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if xml.User != "manu" || xml.Password != "123" {
			context.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}
		context.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	})
	router.POST("/loginForm", func(context *gin.Context) {
		var Form form
		if err := context.ShouldBind(&Form); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if Form.User != "manu" || Form.Password != "123" {
			context.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}
		context.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	})
	router.Run(":8080")
}
