package applog

import (
	"fmt"
	"os"

	"github.com/rs/zerolog"
)

var loggers = make(map[string]zerolog.Logger)

func Init(names ...string) {
	for _, name := range names {
		openFile, err := os.OpenFile(fmt.Sprintf("./logs/%s.log", name), os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModeAppend|os.ModePerm)
		if err != nil {
			fmt.Printf("open log file err: %v\n", err)
			return
		}
		loggers[name] = zerolog.New(openFile)
	}
}

func Logger(name string) *zerolog.Logger {
	instance := loggers[name]
	return &instance
}
