package apigrpc

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
)

// Server runs a gRPC server.
type Server struct {
	Listen   func() (net.Listener, error)
	Register func(*grpc.Server)
	Logger   *log.Logger
	Opts     []grpc.ServerOption
}

// Run accepts API requests until ctx is canceled or an error occurs.
func (s *Server) Run(ctx context.Context) error {
	l, err := s.Listen()
	if err != nil {
		s.Logger.Printf("unable to listen for gRPC API requests: %v", err)
		return err
	}
	defer l.Close()

	svr := grpc.NewServer(s.Opts...)
	s.Register(svr)

	go func() {
		<-ctx.Done()
		svr.Stop()
	}()

	s.Logger.Printf("gRPC API server is listening on %v", l.Addr())

	if err := svr.Serve(l); err != nil {
		s.Logger.Printf("gRPC API server stopped: %v", err)
		return err
	}

	s.Logger.Print("gRPC API server stopped")
	return nil
}
