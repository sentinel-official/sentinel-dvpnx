package node

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sentinel-official/sentinel-go-sdk/libs/cron"
	"github.com/sentinel-official/sentinel-go-sdk/libs/log"

	"github.com/sentinel-official/dvpn-node/api"
	"github.com/sentinel-official/dvpn-node/config"
	"github.com/sentinel-official/dvpn-node/context"
	"github.com/sentinel-official/dvpn-node/workers"
)

// init sets the Gin mode to ReleaseMode.
func init() {
	gin.SetMode(gin.ReleaseMode)
}

// SetupRouter sets up the HTTP router with necessary middlewares and routes.
func (n *Node) SetupRouter(ctx *context.Context) error {
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
	api.RegisterRoutes(ctx, r)

	// Attach the configured router to the Node.
	n.WithRouter(r)
	return nil
}

// SetupScheduler sets up the cron scheduler with various workers.
func (n *Node) SetupScheduler(ctx *context.Context, cfg *config.Config) error {
	log.Info("Setting up scheduler")

	// Define the list of cron workers with their respective handlers and intervals.
	items := []cron.Worker{
		workers.NewBestRPCAddrWorker(ctx, cfg.Node.GetIntervalBestRPCAddr()),
		workers.NewGeoIPLocationWorker(ctx, cfg.Node.GetIntervalGeoIPLocation()),
		workers.NewSessionUsageSyncWithBlockchainWorker(ctx, cfg.Node.GetIntervalSessionUsageSyncWithBlockchain()),
		workers.NewSessionUsageSyncWithDatabaseWorker(ctx, cfg.Node.GetIntervalSessionUsageSyncWithDatabase()),
		workers.NewSessionUsageValidateWorker(ctx, cfg.Node.GetIntervalSessionUsageValidate()),
		workers.NewSessionValidateWorker(ctx, cfg.Node.GetIntervalSessionValidate()),
		workers.NewSpeedtestWorker(ctx, cfg.Node.GetIntervalSpeedtest()),
		workers.NewNodeStatusUpdateWorker(ctx, cfg.Node.GetIntervalStatusUpdate()),
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
func (n *Node) Setup(ctx *context.Context, cfg *config.Config) error {
	log.Info("Setting up node...")

	// Set the listen address for the Node.
	n.WithListenAddr(cfg.Node.APIListenAddr())

	// Set up the HTTP router.
	if err := n.SetupRouter(ctx); err != nil {
		return fmt.Errorf("failed to setup router: %w", err)
	}

	// Set up the cron scheduler.
	if err := n.SetupScheduler(ctx, cfg); err != nil {
		return fmt.Errorf("failed to setup scheduler: %w", err)
	}

	return nil
}
