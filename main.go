package main

import (
	"github.com/gotoeasy/glang/cmn"
	"go-ocr/onstart"
	"go-ocr/pkg/settings"
	"go-ocr/redisutil"
	"runtime"
)

func main() {
	cmn.SetGlcClient(cmn.NewGlcClient(&cmn.GlcOptions{Enable: false, LogLevel: "INFO"})) // 控制台INFO日志级别输出
	runtime.GOMAXPROCS(settings.CpuMaxProcess)
	// 显式调用触发数据库、redis等
	redisutil.Start()
	onstart.Run()
}
