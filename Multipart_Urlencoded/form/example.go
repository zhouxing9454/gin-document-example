package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	router.POST("/form_post", func(context *gin.Context) {
		message := context.PostForm("message")
		nick := context.DefaultPostForm("nick", "anonymous")
		context.JSON(200, gin.H{
			"status":  "posted",
			"message": message,
			"nick":    nick,
		})
	})
	router.Run(":8080")
}

//curl 中 -d 参数用于发送 POST 请求的数据体
//使用-d参数以后，HTTP 请求会自动加上标头Content-Type : application/x-www-form-urlencoded。
//并且会自动将请求转为 POST 方法，因此可以省略-X POST。

//curl -XPOST 127.0.0.1:8080/form_post
//curl -XPOST 127.0.0.1:8080/form_post?message  -d "message=11"
//curl -XPOST 127.0.0.1:8080/form_post?message -d "message=this_is_message&nick=1"
