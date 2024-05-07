package logger

import (
	"fmt"
	"log"
)

const (
	msgFormat = "[%s] - Msg: %s\n"
)

var logger = log.Default()

func Infof(format string, v ...any) {
	msg := fmt.Sprintf(format, v...)

	logger.Printf(msgFormat, "Info", msg)
}

func Errorf(format string, v ...any) {
	msg := fmt.Sprintf(format, v...)

	logger.Printf(msgFormat, "Error", msg)
}

func Warningf(format string, v ...any) {
	msg := fmt.Sprintf(format, v...)

	logger.Printf(msgFormat, "Warning", msg)
}

func Panicf(format string, v ...any) {
	msg := fmt.Sprintf(format, v...)

	logger.Panicf(msgFormat, "Fatal", msg)
}
