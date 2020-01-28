package run

import (
	"context"
	"os"
	"os/signal"
	"syscall"
)

// WaitForSignal blocks until a SIGINT or SIGTERM signal is received, or ctx is
// canceled.
func WaitForSignal(ctx context.Context) {
	signals := make(chan os.Signal)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)

	select {
	case <-ctx.Done():
	case <-signals:
	}
}
