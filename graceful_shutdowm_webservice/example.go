package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	router := gin.Default()
	router.GET("/", func(context *gin.Context) {
		time.Sleep(5 * time.Second)
		context.String(http.StatusOK, "Welcome Gin server")
	})

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		//服务连接
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listern:%s\n", err)
		}
	}()

	//等待中断信号以优雅地关闭服务器（设置五秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server shutdown:", err)
	}
	log.Println("Server exiting")
}

//ctrl+c
