package config

import (
	"fmt"
	"os"
)

const (
	AppNameEnv = "ORDER_SERVICE_APP_NAME"
	DebugEnv   = "ORDER_SERVICE_DEBUG"
	PortEnv    = "ORDER_SERVICE_PUBLIC_PORT"
)

type Config struct {
	AppName    string
	Debug      bool
	PublicPort string
}

func New() (*Config, error) {
	appName := os.Getenv(AppNameEnv)
	if appName == "" {
		return nil, fmt.Errorf("%s is empty", AppNameEnv)
	}

	debug := os.Getenv(DebugEnv)
	if debug == "" {
		return nil, fmt.Errorf("%s is empty", DebugEnv)
	}

	port := os.Getenv(PortEnv)
	if port == "" {
		return nil, fmt.Errorf("%s is empty", PortEnv)
	}

	return &Config{
		AppName:    appName,
		Debug:      debug == "true",
		PublicPort: port,
	}, nil
}
