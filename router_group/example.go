package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	// 简单的路由组: v1
	v1 := router.Group("/v1")
	{
		v1.POST("/login", loginEndpoint)
		v1.POST("/submit", submitEndpoint)
		v1.POST("/read", readEndpoint)
	}

	// 简单的路由组: v2
	v2 := router.Group("/v2")
	{
		v2.POST("/login", loginEndpoint)
		v2.POST("/submit", submitEndpoint)
		v2.POST("/read", readEndpoint)
	}

	router.Run(":8080")
}
func loginEndpoint(c *gin.Context) {
	c.String(200, c.Request.URL.String())
}
func submitEndpoint(c *gin.Context) {
	c.String(200, c.Request.URL.String())
}
func readEndpoint(c *gin.Context) {
	c.String(200, c.Request.URL.String())
}
