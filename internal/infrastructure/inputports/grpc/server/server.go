package server

import (
	"fmt"
	"net"
	"restapi/internal"
	"restapi/internal/app/commands"
	"restapi/internal/infrastructure/inputports/grpc/handler"

	"github.com/phuslu/log"
	"google.golang.org/grpc"
)

type Server struct {
	grpcServer *grpc.Server
}

func (s *Server) Run(service *internal.Service, voteService *commands.VoteService) error {
	s.grpcServer = grpc.NewServer()

	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		fmt.Printf("net err:%v", err)
		return err
	}

	log.Info().Msg("Server running")

	s.Register(service, voteService)

	return s.grpcServer.Serve(lis)
}

func (s *Server) Register(service *internal.Service, voteService *commands.VoteService) {
	handler.RegisterAuth(s.grpcServer, service, voteService)
	handler.RegisterUser(s.grpcServer, service, voteService)
	handler.RegisterPost(s.grpcServer, service, voteService)
	handler.RegisterRating(s.grpcServer, service, voteService)
}
