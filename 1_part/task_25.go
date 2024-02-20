// Реализовать собственную функцию sleep.
package main

import (
	"context"
	"time"
)

func main() {
	ctx := context.Background()
	err := Sleep(ctx, 5*time.Second)
	if err != nil {
	}

}

// Sleep pauses the current goroutine for the given duration.
func Sleep(ctx context.Context, d time.Duration) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-time.After(d):
		return nil
	}
}
