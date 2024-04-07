package onstart

import (
	"fmt"
	"github.com/gotoeasy/glang/cmn"
	"go-ocr/pkg/settings"
	"go-ocr/redisutil"
	"go-ocr/router"
	"net/http"
)

func Run() {
	cmn.Info("Http Server Start")
	r := router.InitRouter()
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", settings.HTTPPort),
		Handler:        r,
		ReadTimeout:    settings.ReadTimeout,
		WriteTimeout:   settings.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
	// 显式调用触发数据库、redis等
	//model.Start()
	redisutil.Start()
}
