package jobs

import (
	"net/http"
	"sync"
	"time"

	"github.com/sentinel-official/dvpn-node/context"
)

// handleBestRPCEndpoint determines the best RPC endpoint based on latency.
func handleBestRPCEndpoint(ctx *context.Context, endpoints []string) func() error {
	client := &http.Client{
		Timeout: 15 * time.Second,
	}

	return func() error {
		if len(endpoints) == 0 {
			return nil
		}

		var (
			bestEndpoint = endpoints[0]     // Initialize with the first endpoint
			minLatency   = 15 * time.Second // Initial latency threshold
			mu           sync.Mutex         // Mutex to synchronize access to shared variables
			wg           sync.WaitGroup     // WaitGroup to wait for all goroutines to finish
		)

		for _, endpoint := range endpoints {
			wg.Add(1)
			go func(endpoint string) {
				defer wg.Done()

				// Record the start time of the request
				start := time.Now()

				// Perform GET request to the endpoint
				resp, err := client.Get(endpoint)
				if err != nil {
					return
				}
				defer resp.Body.Close()

				// Check if the response status code indicates success.
				if resp.StatusCode != http.StatusOK {
					return
				}

				// Calculate the latency of the request
				latency := time.Since(start)

				mu.Lock()
				defer mu.Unlock()

				// Update the best endpoint if the current one has lower latency
				if latency < minLatency {
					bestEndpoint = endpoint
					minLatency = latency
				}
			}(endpoint)
		}

		// Wait for all goroutines to complete
		wg.Wait()

		// Set the best RPC address in the context
		if err := ctx.SetRPCAddr(bestEndpoint); err != nil {
			return err
		}

		return nil
	}
}

// onErrorBestRPCEndpoint returns a function to handle errors.
func onErrorBestRPCEndpoint(_ *context.Context) func(error) bool {
	return func(err error) bool {
		// Do nothing on error
		return false
	}
}
