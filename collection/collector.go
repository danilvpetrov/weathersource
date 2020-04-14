package collection

import (
	"context"
	"log"
	"time"

	backoff "github.com/cenkalti/backoff/v4"
	"github.com/danilvpetrov/weathersource/internal/format"
	"github.com/danilvpetrov/weathersource/storage"
)

// Collector is a service that collects meteorological data at the given
// interval.
type Collector struct {
	Client    *Client
	Retry     backoff.BackOff
	Persister storage.Persister
	Interval  time.Duration
	Logger    *log.Logger
}

// Run runs the service until error is produced or context is cancelled.
func (c *Collector) Run(ctx context.Context) error {
	c.Logger.Print("started weather data collection")
	defer c.Logger.Print("stopped weather data collection")

	for {
		if err := c.do(ctx); err != nil {
			return err
		}
		if err := c.wait(ctx); err != nil {
			return err
		}
	}
}

func (c *Collector) do(ctx context.Context) error {
	var (
		res *Response
		err error
	)

	return backoff.RetryNotify(
		func() error {
			// Don't make a call if the response is already available from
			// previous retry call.
			if res == nil {
				res, err = c.Client.Forecast(ctx)
				if err != nil {
					return err
				}
			}

			if err := c.Persister.Save(ctx, res.ToData()); err != nil {
				return err
			}

			c.Retry.Reset()
			return nil
		},
		backoff.WithContext(c.Retry, ctx),
		func(err error, d time.Duration) {
			c.Logger.Printf(
				"error collecting data: %v, retrying in %s",
				err,
				format.DurationToHuman(d),
			)
		},
	)
}

func (c *Collector) wait(ctx context.Context) error {
	t := time.NewTimer(c.Interval)
	defer t.Stop()

	c.Logger.Printf(
		"next weather data collection in %v",
		c.Interval,
	)

	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-t.C:
		return nil
	}
}
