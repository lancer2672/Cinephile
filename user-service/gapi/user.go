package gapi

import (
	"context"

	"github.com/vudinhan2525/TaskFlow/pb"
)


func (server *Server) CreateUser(ctx context.Context, req *pb.CreateUserReq) (*pb.CreateUserRes, error) {
	

	res := &pb.CreateUserRes{
		Data: &pb.User{
			UserId: 1,
		},
	}
	return res, nil
} 