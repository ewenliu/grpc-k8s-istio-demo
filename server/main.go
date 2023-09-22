package main

import (
	"context"
	"k8s_demo_server/pb"
	"log"
	"net"

	// Change this for your own project
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	*pb.UnimplementedGCDServiceServer
}

func (s *server) Compute(ctx context.Context, r *pb.GCDRequest) (*pb.GCDResponse, error) {
	a, b := r.A, r.B
	logrus.Info("recieve req")
	return &pb.GCDResponse{Result: a * b}, nil
}

func (s *server) mustEmbedUnimplementedGCDServiceServer() {}

func main() {
	lis, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGCDServiceServer(s, &server{})
	reflection.Register(s)
	logrus.Info("start grpc")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
