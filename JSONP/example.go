package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/JSONP", func(context *gin.Context) {
		data := map[string]interface{}{
			"foo": "bar",
		}
		context.JSONP(http.StatusOK, data)
	})
	r.Run(":8080")
}

//使用 JSONP 向不同域的服务器请求数据。
//如果查询参数存在回调，则将回调添加到响应体中。
