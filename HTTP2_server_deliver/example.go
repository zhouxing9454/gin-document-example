package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"log"
)

var html = template.Must(template.New("https").Parse(`
<html>
<head>
  <title>Https Test</title>
  <script src="/assets/app.js"></script>
</head>
<body>
  <h1 style="color:red;">Welcome, Ginner!</h1>
</body>
</html>
`))

func main() {
	r := gin.Default()
	r.Static("/assets", "./assets")
	r.SetHTMLTemplate(html)
	r.GET("/", func(context *gin.Context) {
		if pusher := context.Writer.Pusher(); pusher != nil {
			if err := pusher.Push("/asset/app.js", nil); err != nil {
				log.Printf("Failed to push:%v", err)
			}
		}
		context.HTML(200, "https", gin.H{
			"status": "success",
		})
	})
	r.RunTLS(":8080", "cert.pem", "key.pem") //ListenAndServeTLS函数差不多，后面两个参数是SSL证书，服务器私钥
}

//要访问https://localhost:8080/
