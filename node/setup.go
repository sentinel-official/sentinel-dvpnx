package node

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sentinel-official/sentinel-go-sdk/libs/cron"
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
	log.Info("Setting up router")

	// Define middlewares to be used by the router.
	middlewares := []gin.HandlerFunc{
		cors.New(
			cors.Config{
				AllowAllOrigins: true,
				AllowMethods:    []string{http.MethodGet, http.MethodPost},
			},
		),
	}

	// Create a new Gin router and apply the middlewares.
	r := gin.New()
	r.Use(middlewares...)

	// Register API routes to the router.
	api.RegisterRoutes(n.Context, r)

	// Attach the configured router to the Node.
	n.WithRouter(r)
	return nil
}

// SetupScheduler sets up the cron scheduler with various workers.
func (n *Node) SetupScheduler(cfg *config.Config) error {
	log.Info("Setting up scheduler")

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
	if err := s.RegisterWorkers(items...); err != nil {
		return fmt.Errorf("failed to register workers: %w", err)
	}

	// Attach the configured scheduler to the Node.
	n.WithScheduler(s)
	return nil
}

// Setup sets up both the router and scheduler for the Node.
func (n *Node) Setup(cfg *config.Config) error {
	log.Info("Setting up node...")

	// Set up the HTTP router.
	if err := n.SetupRouter(cfg); err != nil {
		return fmt.Errorf("failed to setup router: %w", err)
	}

	// Set up the cron scheduler.
	if err := n.SetupScheduler(cfg); err != nil {
		return fmt.Errorf("failed to setup scheduler: %w", err)
	}

	return nil
}
