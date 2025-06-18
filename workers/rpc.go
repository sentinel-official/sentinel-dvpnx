package workers

import (
	"net/http"
	"net/url"
	"sort"
	"sync"
	"time"

	"github.com/sentinel-official/sentinel-go-sdk/libs/cron"
	logger "github.com/sentinel-official/sentinel-go-sdk/libs/log"

	"github.com/sentinel-official/sentinel-dvpnx/context"
)

const nameBestRPCAddr = "best_rpc_addr"

// NewBestRPCAddrWorker creates a worker that determines the best RPC address based on latency.
// This worker periodically measures the latency of available RPC addresses,
// sorts them in ascending order of latency, and updates the context.
func NewBestRPCAddrWorker(c *context.Context, interval time.Duration) cron.Worker {
	client := &http.Client{Timeout: 5 * time.Second}
	log := logger.With("name", nameBestRPCAddr)

	// Handler function that measures RPC address latencies and updates the context.
	handlerFunc := func() error {
		// Retrieve the list of RPC addresses from the context.
		addrs := c.RPCAddrs()
		if len(addrs) == 0 {
			return nil
		}

		latencies := make(map[string]time.Duration) // Maps each address to its latency.
		mu := &sync.Mutex{}                         // Synchronizes access to shared resources.
		wg := &sync.WaitGroup{}                     // Ensures all goroutines complete.

		// Measure latency for each address concurrently.
		for _, addr := range addrs {
			wg.Add(1)
			go func(addr string) {
				defer wg.Done()

				endpoint, err := url.JoinPath(addr, "/status")
				if err != nil {
					return
				}

				// Record start time and perform HTTP GET request.
				start := time.Now()

				resp, err := client.Get(endpoint)
				if err != nil {
					return
				}

				defer resp.Body.Close()

				// Skip if the response status is not HTTP 200 OK.
				if resp.StatusCode != http.StatusOK {
					return
				}

				// Calculate and record the latency.
				latency := time.Since(start)

				mu.Lock()
				latencies[addr] = latency
				mu.Unlock()
			}(addr)
		}

		// Wait for all goroutines to complete.
		wg.Wait()

		// Sort the addresses by latency.
		addrs = make([]string, 0, len(latencies))
		for addr := range latencies {
			addrs = append(addrs, addr)
		}
		sort.Slice(addrs, func(i, j int) bool {
			return latencies[addrs[i]] < latencies[addrs[j]]
		})

		// Return early if no RPC addresses are available.
		if len(addrs) == 0 {
			return nil
		}

		// Update the context with the sorted list of RPC addresses.
		c.SetRPCAddrs(addrs)

		return nil
	}

	// Error handling function to log failures.
	onErrorFunc := func(err error) bool {
		log.Error("Failed to run scheduler worker", "msg", err)
		return false
	}

	// Initialize and return the worker.
	return cron.NewBasicWorker().
		WithName(nameBestRPCAddr).
		WithHandler(handlerFunc).
		WithInterval(interval).
		WithOnError(onErrorFunc)
}
