package server

import (
	"context"
	"mini-stats-server/internal/repository"
	"net"

	"github.com/rs/zerolog/log"

	pb "github.com/Loag/mini-stats-proto/gen/go"

	"google.golang.org/grpc"
)

type RPCServer struct {
	server *grpc.Server
}

type grpc_server struct {
	pb.UnimplementedIngestServiceServer
	Repo *repository.Repo
}

func New(repo *repository.Repo) RPCServer {
	s := grpc.NewServer()

	pb.RegisterIngestServiceServer(s, &grpc_server{Repo: repo})

	return RPCServer{
		server: s,
	}
}

func (s *RPCServer) Start() {
	listener, err := net.Listen("tcp", ":5000")
	if err != nil {
		log.Fatal().Err(err).Msg("failed to create listener")
	}

	log.Info().Msgf("server listening at %v", listener.Addr())

	if err := s.server.Serve(listener); err != nil {
		log.Fatal().Err(err).Msg("failed to start grpc server")
	}
}

func (s *grpc_server) IngestMetric(ctx context.Context, in *pb.IngestRequest) (*pb.IngestResponse, error) {
	log.Info().Msgf("Received metric: %v with type: %v and value: %v for time: %v", in.GetName(), in.GetMetricType(), in.GetValue(), in.GetTime())

	err := s.Repo.Set(in)
	if err != nil {
		log.Err(err).Msg("unable to add metric")
		return &pb.IngestResponse{Status: 503}, nil
	}
	log.Info().Msg("added metric")
	return &pb.IngestResponse{Status: 201}, nil
}
