package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("/post", func(context *gin.Context) {
		id := context.Query("id")
		page := context.DefaultQuery("page", "0")
		name := context.PostForm("name")
		message := context.PostForm("message")

		fmt.Printf("id: %s; page：%s;name: %s;message: %s", id, page, name, message)
	})
	router.Run(":8080")
}

//通过Query方法可以获取url 中? 之后的请求参数，通过PostForm方法可以获取到Post 的数据
//postman: http://localhost:8080/post?id=1234&page=1
