package main

import (
	"context"
	"flag"
	"fmt"
	"ginweb/src/conf"
	"ginweb/src/gin-server/router"
	"ginweb/src/util"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var conFile = flag.String("conf", "", "configuration file absulote path")
var gconf *conf.Config

func main() {
	flag.Parse()

	// 加载配置文件
	conf.InitConfig(*conFile)
	gconf = conf.GlobalConfig

	// 初始化日志
	conf.InitLogger(gconf.LogConfig)

	// 初始化数据库
	conf.InitDB(gconf)

	// 注册路由
	router.Router()

	// 启动服务
	StartServer(gconf.SConfig)

}

func StartServer(sc conf.ServerConfig) {
	srv := &http.Server{
		Handler:      router.RouterMux,
		Addr:         fmt.Sprintf(":%d", sc.HttpPort),
		ReadTimeout:  time.Duration(sc.WriteTimeout) * time.Millisecond,
		WriteTimeout: time.Duration(sc.ReadTimeout) * time.Millisecond,
	}

	go func() {
		err := srv.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			util.Warningf("listen: %s\\n", err)
		}
	}()

	// 等待中断信号以超时 5 秒正常关闭服务器
	quit := make(chan os.Signal)
	// kill 命令发送信号 syscall.SIGTERM
	// kill -2 命令发送信号 syscall.SIGINT
	// kill -9 命令发送信号 syscall.SIGKILL

	//将对应的信号通知 quit
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	util.Info("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		util.Infof("Server Shutdown:", err)
	}

	// 5 秒后捕获 ctx.Done() 信号
	select {
	case <-ctx.Done():
		util.Info("timeout of 5 seconds.")
	}
	util.Info("Server exiting")
}
