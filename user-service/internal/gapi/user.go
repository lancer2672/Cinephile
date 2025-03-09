package gapi

import (
	"context"

	"github.com/vudinhan2525/TaskFlow/pb"
)

func (server *Server) CreateUser(ctx context.Context, req *pb.CreateUserReq) (*pb.CreateUserRes, error) {

	// TODO: validator req & hashed password
	user, err := server.UserUseCase.CreateUser(req.Email, req.Password, req.FullName, req.Role)

	if err != nil {
		return nil, err
	}

	res := &pb.CreateUserRes{
		Data: &pb.User{
			UserId:   user.ID,
			Email:    user.Email,
			Role:     user.Role,
			FullName: user.FullName,
		},
	}
	return res, nil
}
