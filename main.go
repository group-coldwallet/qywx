package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/group-coldwallet/qywx/global"
	"github.com/group-coldwallet/qywx/middleware"
	"github.com/group-coldwallet/qywx/router"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

//CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w"
func main() {

	global.InitConfig("")
	global.InitWxAPP(global.Cfg.Wechat.Corpid)
	startServer()

}

func startServer() {

	r := gin.Default()
	//跨域
	r.Use(middleware.GinCors())
	router.InitV1Router(r)
	log.Println(fmt.Sprintf("port:%s", global.Cfg.Http.Port))
	srv := &http.Server{
		Addr:    ":" + global.Cfg.Http.Port,
		Handler: r,
	}

	go func() {
		// 服务连接
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")

}
