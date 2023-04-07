package main

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"time"
)

var (
	//实现一个goroutine等待一组goroutine干活结束，
	//更好的实现了任务同步，但是waitGroup却无法返回错误，
	//当一组Goroutine中的某个goroutine出错时，我们是无法感知到的，
	//所以errGroup对waitGroup进行了一层封装,会返回错误
	g errgroup.Group
)

func router01() http.Handler {
	e := gin.New()
	e.Use(gin.Recovery())
	e.GET("/", func(context *gin.Context) {
		context.JSON(
			http.StatusOK,
			gin.H{
				"code":  http.StatusOK,
				"error": "welcome sever 01",
			},
		)
	})
	return e
}

func router02() http.Handler {
	e := gin.New()
	e.Use(gin.Recovery())
	e.GET("/", func(context *gin.Context) {
		context.JSON(
			http.StatusOK,
			gin.H{
				"code":  http.StatusOK,
				"error": "welcome sever 02",
			},
		)
	})
	return e
}

func main() {
	server01 := &http.Server{
		Addr:         ":8080",
		Handler:      router01(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	server02 := &http.Server{
		Addr:         ":8081",
		Handler:      router02(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	//执行Add()方法增加一个计数器
	//开启一个协程，运行我们传入的函数f，使用waitGroup的Done()方法控制是否结束
	//如果有一个函数f运行出错了，我们把它保存起来，如果有cancel()方法，则执行cancel()取消其他goroutine
	g.Go(func() error {
		return server01.ListenAndServe()
	})

	g.Go(func() error {
		return server02.ListenAndServe()
	})

	//调用waitGroup的Wait()等待一组Goroutine的运行结束
	//这里为了保证代码的健壮性，如果前面赋值了cancel，要执行cancel()方法
	//返回错误信息，如果有goroutine出现了错误才会有值
	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}
