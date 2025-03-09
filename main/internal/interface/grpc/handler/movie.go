package handler

import (
	"context"

	"github.com/lancer2672/cinephile/main/internal/app"
	"github.com/lancer2672/cinephile/main/internal/domain/model"
	"github.com/lancer2672/cinephile/main/internal/domain/repository"
	converter "github.com/lancer2672/cinephile/main/internal/interface/grpc/converter"
	"github.com/lancer2672/cinephile/main/internal/interface/grpc/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type MovieHandler struct {
	service app.MovieService
	pb.UnimplementedMovieServiceServer
}

func RegisterMovieHandler(grpcServer *grpc.Server, service repository.Store) {
	handler := &MovieHandler{service: service}
	pb.RegisterMovieServiceServer(grpcServer, handler)
	// return handler
}

func (m *MovieHandler) GetMovies(ctx context.Context, req *pb.GetMoviesRequest) (*pb.GetMoviesResponse, error) {
	filter := model.MovieFilter{}
	movies, _, err := m.service.GetMovies(ctx, filter)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to get recent movies")
	}

	result := make([]*pb.Movie, 0, len(movies))
	for _, movie := range movies {
		result = append(result, converter.ToProtoMovie(&movie))
	}

	return &pb.GetMoviesResponse{Movies: result}, nil
}

func (m *MovieHandler) GetMovieByID(ctx context.Context, req *pb.GetMovieRequest) (*pb.GetMovieResponse, error) {
	movie, err := m.service.GetMovieByID(ctx, req.MovieId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to get movie by id")
	}

	return &pb.GetMovieResponse{Movie: converter.ToProtoMovie(movie)}, nil
}
