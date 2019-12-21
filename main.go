package main

import (
	"fmt"
	"github.com/wuyueCreator/gin_blog/pkg/setting"
	"github.com/wuyueCreator/gin_blog/routers"
	"net/http"
)

func main() {
	router := routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	err := s.ListenAndServe()
	fmt.Println(err)

}
