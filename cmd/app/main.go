package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"mini-stats-server/config"

	pb "github.com/Loag/mini-stats-proto/gen/go"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedIngestServiceServer
}

func (s *server) IngestMetric(ctx context.Context, in *pb.IngestRequest) (*pb.IngestResponse, error) {
	log.Printf("Received metric type: %v and value: %v", in.GetMetricType(), in.GetValue())

	return &pb.IngestResponse{Status: 200}, nil
}

func main() {

	conf := config.New()

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", *&conf.Port))
	if err != nil {
		panic(fmt.Sprintf("failed to start listener"))
	}

	s := grpc.NewServer()
	pb.RegisterIngestServiceServer(s, &server{})

	log.Printf("server listening at %v", listener.Addr())

	if err := s.Serve(listener); err != nil {
		panic(fmt.Sprintf("failed to serve: %v", err))
	}
}
