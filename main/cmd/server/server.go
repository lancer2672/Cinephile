package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/lancer2672/cinephile/main/config"
	"github.com/lancer2672/cinephile/main/internal/interface/grpc/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	pb.UnimplementedMainServiceServer
	config      *config.Config
	grpcServer  *grpc.Server
	httpServer  *http.Server
	shutdownCtx context.Context
	cancel      context.CancelFunc
}

func New(cfg *config.Config) *Server {
	ctx, cancel := context.WithCancel(context.Background())

	return &Server{
		config:      cfg,
		grpcServer:  grpc.NewServer(),
		shutdownCtx: ctx,
		cancel:      cancel,
	}
}

// Khởi chạy gRPC server
func (s *Server) runGrpcServer(wg *sync.WaitGroup) {
	defer wg.Done()

	listener, err := net.Listen("tcp", ":15002")
	if err != nil {
		log.Fatal("Error when creating gRPC listener:", err)
	}
	log.Printf("🚀 gRPC server is running at %s", listener.Addr().String())

	// Đăng ký dịch vụ gRPC
	pb.RegisterMainServiceServer(s.grpcServer, s)
	reflection.Register(s.grpcServer)

	// Chạy gRPC server
	if err := s.grpcServer.Serve(listener); err != nil {
		log.Fatal("Error running gRPC server:", err)
	}
}

// Chạy cả gRPC và HTTP server đồng thời
func (s *Server) Run() {
	var wg sync.WaitGroup
	wg.Add(1)

	go s.runGrpcServer(&wg)

	// Lắng nghe tín hiệu shutdown
	s.handleShutdown()

	wg.Wait()
}

// Lắng nghe tín hiệu để shutdown server
func (s *Server) handleShutdown() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	<-sigChan
	log.Println("⚠️ Shutting down servers...")

	s.Shutdown()
}

// Shutdown cả gRPC và HTTP server
func (s *Server) Shutdown() {
	// Đóng gRPC server
	log.Println("🛑 Stopping gRPC server...")
	s.grpcServer.GracefulStop()

	// Đóng HTTP server với timeout
	log.Println("🛑 Stopping HTTP server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.httpServer.Shutdown(ctx); err != nil {
		log.Println("Error shutting down HTTP server:", err)
	}

	// Hủy context
	s.cancel()
	log.Println("✅ Server shutdown complete.")
}

func (s *Server) GetMovies(ctx context.Context, req *pb.GetMoviesRequest) (*pb.GetMoviesResponse, error) {
	return &pb.GetMoviesResponse{MovieId: req.MovieId}, nil
}
