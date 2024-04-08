package onstart

import (
	"fmt"
	"github.com/gotoeasy/glang/cmn"
	"go-ocr/pkg/settings"
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

}
