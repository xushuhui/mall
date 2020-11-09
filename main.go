package main

import (
	"mall_go/internal/router"
	"mall_go/pkg/core"
	"mall_go/pkg/msg"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	core.StartModule()

	router.HttpServerRun()
	msg.StartMsg()
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGKILL, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	router.HttpServerStop()

}
