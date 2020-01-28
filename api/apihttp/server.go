package apihttp

import (
	"context"
	"log"
	"net"
	"net/http"
	"time"
)

// Server runs a HTTP server.
type Server struct {
	Listen       func() (net.Listener, error)
	Handler      http.Handler
	Logger       *log.Logger
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

// Run runs the server until error is returned or context is cancelled.
func (s *Server) Run(ctx context.Context) error {
	l, err := s.Listen()
	if err != nil {
		s.Logger.Printf(
			"unable to listen for HTTP API requests: %v",
			err,
		)
		return err
	}
	defer l.Close()

	svr := &http.Server{
		Handler:      s.Handler,
		ReadTimeout:  s.ReadTimeout,
		WriteTimeout: s.WriteTimeout,
	}

	go func() {
		<-ctx.Done()
		svr.Close()
	}()

	s.Logger.Printf("HTTP API server is listening on %v", l.Addr())

	if err := svr.Serve(l); err != nil {
		s.Logger.Printf("HTTP API server stopped: %v", err)
		return err
	}

	s.Logger.Print("HTTP API server stopped")
	return nil
}
