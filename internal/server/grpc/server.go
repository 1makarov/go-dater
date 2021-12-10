package grpcServer

import (
	"fmt"
	"github.com/1makarov/go-dater/server/internal/proto"
	"google.golang.org/grpc"
	"net"
)

type Server struct {
	server  *grpc.Server
	handler proto.DaterServer
}

func New(handler proto.DaterServer) *Server {
	return &Server{
		server:  grpc.NewServer(),
		handler: handler,
	}
}

func (s *Server) Start(port string) error {
	addr := fmt.Sprintf(":%s", port)

	listen, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	proto.RegisterDaterServer(s.server, s.handler)

	if err = s.server.Serve(listen); err != nil {
		return err
	}

	return nil
}

func (s *Server) Stop() {
	s.server.GracefulStop()
}
