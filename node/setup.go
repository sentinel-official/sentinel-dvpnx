package node

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sentinel-official/sentinel-go-sdk/libs/cron"
	"github.com/sentinel-official/sentinel-go-sdk/libs/gin/middlewares"
	"github.com/sentinel-official/sentinel-go-sdk/libs/log"

	"github.com/sentinel-official/sentinel-dvpnx/api"
	"github.com/sentinel-official/sentinel-dvpnx/config"
	"github.com/sentinel-official/sentinel-dvpnx/workers"
)

// init sets the Gin mode to ReleaseMode.
func init() {
	gin.SetMode(gin.ReleaseMode)
}

// SetupRouter sets up the HTTP router with necessary middlewares and routes.
func (n *Node) SetupRouter(_ *config.Config) error {
	// Define middlewares to be used by the router.
	items := []gin.HandlerFunc{
		cors.New(
			cors.Config{
				AllowAllOrigins: true,
				AllowMethods:    []string{http.MethodGet, http.MethodPost},
			},
		),
		middlewares.RateLimiter(nil),
	}

	// Create a new Gin router and apply the middlewares.
	r := gin.New()
	r.Use(items...)

	// Register API routes to the router.
	api.RegisterRoutes(n.Context, r)

	// Attach the configured router to the Node.
	n.WithRouter(r)
	return nil
}

// SetupScheduler sets up the cron scheduler with various workers.
func (n *Node) SetupScheduler(cfg *config.Config) error {
	// Define the list of cron workers with their respective handlers and intervals.
	items := []cron.Worker{
		workers.NewBestRPCAddrWorker(n.Context, cfg.Node.GetIntervalBestRPCAddr()),
		workers.NewGeoIPLocationWorker(n.Context, cfg.Node.GetIntervalGeoIPLocation()),
		workers.NewNodeStatusUpdateWorker(n.Context, cfg.Node.GetIntervalStatusUpdate()),
		workers.NewSessionUsageSyncWithBlockchainWorker(n.Context, cfg.Node.GetIntervalSessionUsageSyncWithBlockchain()),
		workers.NewSessionUsageSyncWithDatabaseWorker(n.Context, cfg.Node.GetIntervalSessionUsageSyncWithDatabase()),
		workers.NewSessionUsageValidateWorker(n.Context, cfg.Node.GetIntervalSessionUsageValidate()),
		workers.NewSessionValidateWorker(n.Context, cfg.Node.GetIntervalSessionValidate()),
		workers.NewSpeedtestWorker(n.Context, cfg.Node.GetIntervalSpeedtest()),
	}

	// Create a new cron scheduler and register the workers.
	s := cron.NewScheduler()

	log.Info("Registering scheduler workers", "count", len(items))
	if err := s.RegisterWorkers(items...); err != nil {
		return fmt.Errorf("failed to register workers: %w", err)
	}

	// Attach the configured scheduler to the Node.
	n.WithScheduler(s)
	return nil
}

// Setup sets up both the router and scheduler for the Node.
func (n *Node) Setup(cfg *config.Config) error {
	log.Info("Setting up router")
	if err := n.SetupRouter(cfg); err != nil {
		return fmt.Errorf("failed to setup router: %w", err)
	}

	log.Info("Setting up scheduler")
	if err := n.SetupScheduler(cfg); err != nil {
		return fmt.Errorf("failed to setup scheduler: %w", err)
	}

	return nil
}
