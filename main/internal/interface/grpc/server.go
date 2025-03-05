package grpc

import (
	"net"

	"github.com/lancer2672/cinephile/main/config"
	"github.com/lancer2672/cinephile/main/internal/domain/repository"
	"github.com/lancer2672/cinephile/main/internal/interface/grpc/handler"
	"github.com/lancer2672/cinephile/main/log"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	config     *config.Config
	grpcServer *grpc.Server
	store      repository.Store
}

func New(cfg *config.Config, store repository.Store) *Server {
	return &Server{
		config:     cfg,
		store:      store,
		grpcServer: grpc.NewServer(),
	}
}

// Khởi chạy gRPC server
func (s *Server) runGrpcServer() {

	listener, err := net.Listen("tcp", ":15002")
	if err != nil {
		log.Logger.Fatal("Error when creating gRPC listener:", zap.Error(err))
	}
	log.Logger.Info("🚀 gRPC server is running at %s", zap.String("address", listener.Addr().String()))

	reflection.Register(s.grpcServer)

	handler.RegisterMovieHandler(s.grpcServer, s.store)

	if err := s.grpcServer.Serve(listener); err != nil {
		log.Logger.Fatal("Error running gRPC server:", zap.Error(err))

	}
}

// Chạy cả gRPC và HTTP server đồng thời
func (s *Server) Run() {
	s.runGrpcServer()
}

func (s *Server) Shutdown() {
	log.Logger.Info("🛑 Stopping gRPC server...")

	s.grpcServer.GracefulStop()

	log.Logger.Info("✅ Server shutdown complete.")
}
