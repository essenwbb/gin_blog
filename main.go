package main

import (
	"fmt"
	"github.com/fvbock/endless"
	"github.com/wuyueCreator/gin_blog/models"
	"github.com/wuyueCreator/gin_blog/pkg/gredis"
	"github.com/wuyueCreator/gin_blog/pkg/logging"
	"github.com/wuyueCreator/gin_blog/pkg/setting"
	"github.com/wuyueCreator/gin_blog/pkg/util"
	"github.com/wuyueCreator/gin_blog/routers"
	"log"
	"syscall"
)

func main() {
	setting.Setup()
	models.Setup()
	logging.Setup()
	_ = gredis.Setup()
	util.Setup()

	endless.DefaultReadTimeOut = setting.ServerSetting.ReadTimeout
	endless.DefaultWriteTimeOut = setting.ServerSetting.WriteTimeout
	endless.DefaultMaxHeaderBytes = 1 << 20
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)

	server := endless.NewServer(endPoint, routers.InitRouter())
	server.BeforeBegin = func(add string) {
		log.Printf("Actual pid is %d", syscall.Getpid())
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Printf("Server err: %v", err)
	}

}
