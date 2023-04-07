package main

import "github.com/gin-gonic/gin"

// 通常，JSON 使用 unicode 替换特殊 HTML 字符，例如 < 变为 \ u003c。
// 如果要按字面对这些字符进行编码，则可以使用 PureJSON。Go 1.6 及更低版本无法使用此功能。
func main() {
	r := gin.Default()

	//提供unicode实体
	//{"html":"\u003cb\u003eHello,world!\u003c/b\u003e"}
	r.GET("/json", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"html": "<b>Hello,world!</b>",
		})
	})

	//提供字面字符
	//{"html":"<b>Hello,world!</b>"}
	r.GET("/purejson", func(context *gin.Context) {
		context.PureJSON(200, gin.H{
			"html": "<b>Hello,world!</b>",
		})
	})
	r.Run(":8080")
}
