package main

import (
	"fmt"
	"net/http"

	"github.com/simanchou/go-gin-example/models"
	"github.com/simanchou/go-gin-example/routers"

	"github.com/simanchou/go-gin-example/pkg/setting"
)

func init() {
	models.Setup()
}

// @title Golang Gin API
// @version 1.0
// @description An example of gin
// @termsOfService https://github.com/simanchou/go-gin-example
// @license.name MIT
// @license.url https://github.com/simanchou/go-gin-example
func main() {
	router := routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.ServerSetting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ServerSetting.ReadTimeout,
		WriteTimeout:   setting.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
