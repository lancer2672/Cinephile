package gapi

import (
	"github.com/vudinhan2525/TaskFlow/internal/usecase"
	"github.com/vudinhan2525/TaskFlow/pb"
	"github.com/vudinhan2525/TaskFlow/util"
)

type Server struct {
	pb.UnimplementedUserServiceServer
	Config      util.Config
	UserUseCase usecase.UserUseCase
}

func NewServer(config util.Config, userUseCase usecase.UserUseCase) (*Server, error) {

	server := Server{Config: config, UserUseCase: userUseCase}

	return &server, nil
}
