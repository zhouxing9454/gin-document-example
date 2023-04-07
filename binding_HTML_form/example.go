package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)
type myForm struct {
	Colors []string `form:"colors[]"`
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("binding_HTML_form/*")
	r.GET("/", func(context *gin.Context) {
		context.HTML(http.StatusOK, "form.html", nil)
	})
	r.POST("/", formHandler)
	r.Run(":8080")
}

func formHandler(c *gin.Context) {
	var fakeForm myForm
	c.ShouldBind(&fakeForm)
	c.JSON(http.StatusOK, gin.H{"color": fakeForm.Colors})
}
