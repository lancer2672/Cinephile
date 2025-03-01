package main

import (
	"flag"

	"github.com/lancer2672/cinephile/main/config"
)

func main() {
	environment := flag.String("e", "dev", "")
	flag.Usage = func() {
		// log.Logger.Info("Usage: server -e {mode}")
	}
	flag.Parse()

	cfg, err := config.LoadConfig(".", *environment)
	if err != nil {
		// log.Logger.Fatal("load config failed: ", err)
	}

	srv := New(cfg)
	if err != nil {
		// log.Logger.Fatal("create server failed: ", err)
	}
	srv.Run()

}
