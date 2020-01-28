package run

import (
	"context"

	"golang.org/x/sync/errgroup"
)

// Service is a generic long-running service. It must terminate on context
// cancellation.
type Service interface {
	// Run runs the service.
	Run(ctx context.Context) error
}

// Services runs implementations of Service until context is cancelled or
// one of the implementations returns an error.
func Services(ctx context.Context, services ...Service) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	g, ctx := errgroup.WithContext(ctx)

	for _, s := range services {
		s := s // capture the variable
		g.Go(func() error {
			return s.Run(ctx)
		})
	}

	// Do not treat context cancellation as an error.
	if err := g.Wait(); err != context.Canceled {
		return err
	}

	return nil
}
