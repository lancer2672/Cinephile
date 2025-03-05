package main

import (
	"flag"
	"fmt"

	"github.com/lancer2672/cinephile/main/config"
	"github.com/lancer2672/cinephile/main/internal/infras/mongo"
	"github.com/lancer2672/cinephile/main/internal/interface/grpc"
	"github.com/lancer2672/cinephile/main/log"
	"go.uber.org/zap"
)

func main() {
	environment := flag.String("e", "dev", "")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
	}
	flag.Parse()

	cfg, err := config.LoadConfig(".", *environment)
	if err != nil {
		fmt.Println("load config failed: " + err.Error())
	}
	log.InitLogger(cfg.LogLevel)
	store, err := mongo.NewMongoStore(cfg)
	if err != nil {
		log.Logger.Fatal("create store failed: ", zap.Error(err))
	}
	srv := grpc.New(cfg, store)
	if err != nil {
		log.Logger.Fatal("create server failed: ", zap.Error(err))
	}
	srv.Run()

}
