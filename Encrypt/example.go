package main

import (
	"log"

	"github.com/gin-gonic/autotls"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/acme/autocert"
)

// Let's Encrypt是一个数字证书认证机构，
// 旨在以自动化流程消除手动创建和安装证书的复杂流程，
// 并推广使万维网服务器的加密连接无所不在，为安全网站提供免费的SSL/TLS证书。
func main() {
	r := gin.Default()

	// Ping handler
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	//一行式LetsEncrypt证书, 处理https
	log.Fatal(autotls.Run(r, "example1.com", "example2.com"))
}

func main2() {
	r := gin.Default()

	// Ping handler
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	m := autocert.Manager{
		Prompt:     autocert.AcceptTOS,                                     //Prompt指定一个回调函数有条件的接受证书机构CA的TOS服务, 使用AcceptTOS总是接受服务条款
		HostPolicy: autocert.HostWhitelist("example1.com", "example2.com"), //HostPolicy用于控制指定哪些域名, 管理器将检索新证书
		Cache:      autocert.DirCache("/var/www/.cache"),                   //缓存证书和其他状态
	}

	log.Fatal(autotls.RunWithManager(r, &m))
}
