package main

import (
	"context"
	"fmt"
	"os"

	"github.com/danilvpetrov/weathersource/cmd/weather/deps"
	"github.com/danilvpetrov/weathersource/internal/run"
)

func main() {
	os.Exit(
		func() int {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			go func() {
				run.WaitForSignal(ctx)
				cancel()
			}()

			services, cleanup, err := deps.Deps()
			if err != nil {
				fmt.Println(err)
				return 1
			}
			defer cleanup()

			if err := run.Services(ctx, services...); err != nil {
				fmt.Println(err)
				return 1
			}

			return 0
		}(),
	)
}
