package main

import (
	"database/sql"
	"net"

	_ "github.com/lib/pq"
	"github.com/vudinhan2525/TaskFlow/internal/gapi"
	"github.com/vudinhan2525/TaskFlow/internal/repository"
	"github.com/vudinhan2525/TaskFlow/internal/usecase"
	"github.com/vudinhan2525/TaskFlow/pb"
	"github.com/vudinhan2525/TaskFlow/pkg/log"
	"github.com/vudinhan2525/TaskFlow/util"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Logger.Fatal("Error when loading env!!", err)
	}

	db, err := sql.Open(config.DbDriver, config.DbSource)
	if err != nil {
		log.Logger.Fatal("Cannot connect to database:", err)
	}
	log.Logger.Info("🚀 Connect success to PostGreSQL")
	defer db.Close()

	userRepo := repository.NewUserRepository(db)
	userUseCase := usecase.NewUserUseCase(userRepo)

	runGrpcServer(config, userUseCase)
}
func runGrpcServer(config util.Config, userUseCase usecase.UserUseCase) {

	server, err := gapi.NewServer(config, userUseCase)
	if err != nil {
		log.Logger.Fatal("Error when creating server")
	}

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, server)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", config.APIEndpoint)
	if err != nil {
		log.Logger.Fatal("Error when creating listener")
	}
	log.Logger.Printf("start gRPC server at %s", listener.Addr().String())

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Logger.Fatal("Cannot creating grpc server")
	}
}
