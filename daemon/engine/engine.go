package engine

import (
	"fmt"
	"os"
	"os/signal"
	"poker/daemon/internal/manager"
	"poker/daemon/internal/pkg/common"

	log "github.com/sirupsen/logrus"

	"poker/daemon/internal/pkg/logger"
	"poker/pkg/config"
)

func Run() {
	if err := logger.InitLogger(); err != nil {
		log.Fatalf("init logger fail, error: %s", err.Error())
	}

	if err := manager.InitManager(); err != nil {
		log.Fatalf("init manager fail, error: %s", err.Error())
	}

	killSignal := make(chan os.Signal, 1)
	signal.Notify(killSignal, os.Interrupt)

	socketFile := config.GetSocketFile()
	l, err := common.ListenOnUnix(socketFile)
	if err != nil {
		shutdown(fmt.Errorf("listen on unix socket fail, error: %s", err))
	}

	go func() {
		if err := serve(l); err != nil {
			shutdown(fmt.Errorf("serve fail, error: %s", err))
		}
	}()

	log.Infof("poker daemon is running now, poker version: %s", config.GetPokerVersion())
	<-killSignal
	shutdown(nil)
}

func shutdown(err error) {
	if err != nil {
		log.Errorf("poker unexpected exit, error: %s", err)
	} else {
		log.Info("poker daemon exit")
	}

	_ = os.Remove(config.GetSocketFile())

	os.Exit(1)
}
