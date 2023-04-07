package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 使用 SecureJSON 防止 json 劫持。如果给定的结构是数组值，则默认预置 "while(1)," 到响应体。
func main() {
	r := gin.Default()
	// 你也可以使用自己的 SecureJSON 前缀
	//r.SecureJsonPrefix(")]}',\n")
	r.GET("/someJSON", func(context *gin.Context) {
		name := []string{"lena", "austin", "foo"}
		context.SecureJSON(http.StatusOK, name)
	})
	r.Run(":8080")
}
