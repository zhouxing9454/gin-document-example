package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	//http:localhost:8080/assets/file(文件名）  你可以访问本地assets文件夹下的所有文件
	router.Static("/assets", "./assets")

	router.StaticFS("/more_static", http.Dir("D:\\Golang_workspace\\src")) ////gin.Dir("/var/log", true)

	//StaticFile注册单个路由，以便为本地文件系统的单个文件提供服务。router.StaticFile（“favicon.ico”，“./resources/favicon.co”）
	router.StaticFile("/1.html", "./resource/1.html")

	router.Run(":8081")
}
