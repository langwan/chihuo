package main

import (
	"flag"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Options struct {
	WorkDir string
	Port    int
	Debug   bool
}

func main() {
	wd, _ := os.Getwd()

	opt := Options{WorkDir: wd, Port: 80, Debug: false}

	flag.BoolVar(&opt.Debug, "debug", false, "true is debug")
	flag.IntVar(&opt.Port, "port", 80, "web server port")
	flag.StringVar(&opt.WorkDir, "workdir", wd, "work directory")

	flag.Parse()

	if opt.Debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	} else {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}

	log.Info().Interface("opt", opt).Send()

	log.Debug().Msg("is debug.")

}
