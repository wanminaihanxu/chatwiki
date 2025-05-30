// Copyright © 2016- 2024 Sesame Network Technology all right reserved

package main

import (
	"chatwiki/internal/app/user_domain_service"
	"chatwiki/internal/app/user_domain_service/define"
	"os"
	"os/signal"
	"syscall"

	"github.com/zhimaAi/go_tools/logs"
)

func main() {
	logs.SetLogsDir(define.AppRoot + `logs`)
	user_domain_service.Run()
	sc := make(chan os.Signal)
	sl := []os.Signal{
		syscall.SIGHUP,  //hangup
		syscall.SIGINT,  //interrupt
		syscall.SIGTERM, //terminated
	}
	signal.Notify(sc, sl...)
	sig := <-sc
	logs.Info(`sign:` + sig.String())
	user_domain_service.Stop()
	logs.Info(`stop finish`)
}
