package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

func main() {
	router := gin.Default()
	// 自定义html模板渲染器，要指定所有的html路径，不推荐
	html := template.Must(template.ParseFiles(
		"templates/login.html",
		"templates/users/index.html",
		"templates/center/index.html",
	))
	//应用这些模板
	router.SetHTMLTemplate(html)
	router.GET("/users/index", func(context *gin.Context) {
		context.HTML(http.StatusOK, "users/index.html", gin.H{
			"title": "users/index.html",
		})
	})
}
