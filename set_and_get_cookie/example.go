package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/cookie", func(context *gin.Context) {
		cookie, err := context.Cookie("gin_cookie")
		if err != nil {
			cookie = "NotSet"
			//path是指cookie的范围，/是指全站可用，是一种很通用的方式
			//domain可以访问该Cookie的域名
			//MaxAge	int	被访问后的存活时间;这个时间是个相对值(比如:3600s)；MaxAge=0,未指定该属性；MaxAge<0时，删除cookie，相当于“Max Age:0”
			//Secure	bool	是否需要安全传输，为true时只有https才会传输该cookie
			//HttpOnly	bool	为true时，不能通过js读取该cookie的值
			context.SetCookie("gin_cookie", "test", 3600, "/", "localhost", false, true)
		}
		fmt.Printf("Cookie value:%s \n", cookie)
	})
	router.Run()
}
