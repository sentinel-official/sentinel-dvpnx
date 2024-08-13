package workers

import (
	"net/http"
	"sync"
	"time"

	"github.com/sentinel-official/dvpn-node/node"
)

// HandlerBestRPCAddr determines the best RPC address based on latency.
func HandlerBestRPCAddr(ctx *node.Context) func() error {
	client := &http.Client{
		Timeout: 15 * time.Second,
	}

	return func() error {
		addrs := ctx.QueryRPCAddrs()
		if len(addrs) == 0 {
			return nil
		}

		var (
			bestAddr   = addrs[0]         // Initialize with the first address
			minLatency = 15 * time.Second // Initial latency threshold
			mu         sync.Mutex         // Mutex to synchronize access to shared variables
			wg         sync.WaitGroup     // WaitGroup to wait for all goroutines to finish
		)

		for _, addr := range addrs {
			wg.Add(1)
			go func(addr string) {
				defer wg.Done()

				// Record the start time of the request
				start := time.Now()

				// Perform GET request to the address
				resp, err := client.Get(addr)
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

				// Update the best address if the current one has lower latency
				if latency < minLatency {
					bestAddr = addr
					minLatency = latency
				}
			}(addr)
		}

		// Wait for all goroutines to complete
		wg.Wait()

		// Set the best RPC address in the context
		ctx.SetRPCAddr(bestAddr)

		return nil
	}
}
