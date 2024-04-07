package main

import (
	"github.com/gotoeasy/glang/cmn"
	"go-ocr/onstart"
	"go-ocr/pkg/settings"
	"runtime"
)

func main() {
	cmn.SetGlcClient(cmn.NewGlcClient(&cmn.GlcOptions{Enable: false, LogLevel: "INFO"})) // 控制台INFO日志级别输出
	runtime.GOMAXPROCS(settings.CpuMaxProcess)                                           // 使用最大CPU数量
	onstart.Run()
}
