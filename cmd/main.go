package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/LuaSavage/double-gis-test-task/internal/app"
	"github.com/LuaSavage/double-gis-test-task/internal/config"
	"github.com/LuaSavage/double-gis-test-task/pkg/logger"
)

func main() {
	stopCh := make(chan os.Signal, 1)
	signal.Notify(stopCh, os.Interrupt, os.Kill, syscall.SIGTERM)

	conf, err := config.New()
	if err != nil {
		logger.Panicf("config init: %v", err)
	}

	if conf.Debug {
		logger.Infof("config: %v", conf)
	}

	logger.Infof("application initialization")
	a, err := app.New(conf)
	if err != nil {
		logger.Panicf("app init: %v", err)
	}

	logger.Infof("starting application")
	err = a.Run()
	if err != nil {
		logger.Panicf("app run: %v", err)
	}

	receivedSignal := <-stopCh
	logger.Infof("catch stop signal: %v", receivedSignal.String())

	logger.Infof("finishing application")
	err = a.Stop()
	if err != nil {
		logger.Panicf("app stop: %v", err)
	}
}
