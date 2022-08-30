package main

import (
	"app/applog"
	"fmt"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func One() {
	user := struct {
		Name  string
		Level int64
	}{Name: "chihuo", Level: 1}
	log.Info().Interface("user", user).Str("name", user.Name).Int64("level", user.Level).Msg("打印用户信息")
}

func Two() {
	consoleWriter := zerolog.ConsoleWriter{Out: os.Stdout}

	logFile, err := os.OpenFile("./logs/app.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModeAppend|os.ModePerm)

	if err != nil {
		fmt.Printf("open log file err: %v\n", err)
		return
	}

	multi := zerolog.MultiLevelWriter(consoleWriter, logFile)

	logger := zerolog.New(multi).With().Timestamp().Logger()

	logger.Info().Msg("Hello World!")
}

func Three() {
	applog.Init("app", "pay")
	applog.Logger("app").Info().Msg("user info.")
	applog.Logger("pay").Info().Msg("pay info.")
}
