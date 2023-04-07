package main

import "github.com/gin-gonic/gin"

type Person struct {
	ID   string `uri:"id" binding:"required,uuid"`
	Name string `uri:"name" binding:"required"`
}

// 在gin框架中绑定 Uri，即在Uri 中传入指定的参数。
func main() {
	router := gin.Default()
	router.GET("/:name/:id", func(context *gin.Context) {
		var person Person
		if err := context.ShouldBindUri(&person); err != nil {
			context.JSON(400, gin.H{"msg": err.Error()})
			return
		}
		context.JSON(200, gin.H{"name": person.Name, "uuid": person.ID})
	})
	router.Run(":8080")
}
