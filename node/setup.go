package node

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sentinel-official/sentinel-go-sdk/libs/cmux"
	"github.com/sentinel-official/sentinel-go-sdk/libs/cron"
	"github.com/sentinel-official/sentinel-go-sdk/libs/gin/middlewares"
	"github.com/sentinel-official/sentinel-go-sdk/libs/log"

	"github.com/sentinel-official/sentinel-dvpnx/api"
	"github.com/sentinel-official/sentinel-dvpnx/config"
	"github.com/sentinel-official/sentinel-dvpnx/core"
	"github.com/sentinel-official/sentinel-dvpnx/workers"
)

// init sets the Gin mode to ReleaseMode.
func init() {
	gin.SetMode(gin.ReleaseMode)
}

// SetupScheduler sets up the cron scheduler with various workers.
func (n *Node) SetupScheduler(ctx context.Context, cfg *config.Config) error {
	// Define the list of cron workers with their respective handlers and intervals.
	items := []cron.Worker{
		workers.NewBestRPCAddrWorker(n.Context(), cfg.Node.GetIntervalBestRPCAddr()),
		workers.NewGeoIPLocationWorker(n.Context(), cfg.Node.GetIntervalGeoIPLocation()),
		workers.NewNodeStatusUpdateWorker(n.Context(), cfg.Node.GetIntervalStatusUpdate()),
		workers.NewSessionUsageSyncWithBlockchainWorker(n.Context(), cfg.Node.GetIntervalSessionUsageSyncWithBlockchain()),
		workers.NewSessionUsageSyncWithDatabaseWorker(n.Context(), cfg.Node.GetIntervalSessionUsageSyncWithDatabase()),
		workers.NewSessionUsageValidateWorker(n.Context(), cfg.Node.GetIntervalSessionUsageValidate()),
		workers.NewSessionValidateWorker(n.Context(), cfg.Node.GetIntervalSessionValidate()),
		workers.NewSpeedtestWorker(n.Context(), cfg.Node.GetIntervalSpeedtest()),
	}

	log.Info("Initializing scheduler")

	s := cron.NewScheduler(ctx, "scheduler")
	if err := s.Setup(); err != nil {
		return err
	}

	for _, item := range items {
		log.Info("Registering scheduler worker",
			"name", item.Name(), "interval", item.Interval(),
		)
		if err := s.Register(item); err != nil {
			return fmt.Errorf("registering scheduler worker %q: %w", item.Name(), err)
		}
	}

	// Attach the configured scheduler to the Node.
	n.WithScheduler(s)
	return nil
}

// SetupServer sets up the API server with necessary middlewares and API routes.
func (n *Node) SetupServer(ctx context.Context, _ *config.Config) error {
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
	router := gin.New()
	router.Use(items...)

	// Register API routes to the router.
	api.RegisterRoutes(n.Context(), router)

	log.Info("Initializing API server")

	s := cmux.NewServer(
		ctx,
		"API-server",
		n.Context().APIListenAddr(),
		n.Context().TLSCertFile(),
		n.Context().TLSKeyFile(),
		router,
	)
	if err := s.Setup(); err != nil {
		return err
	}

	// Attach the API server to the Node instance.
	n.WithServer(s)
	return nil
}

// SetupContext sets up the core context.
func (n *Node) SetupContext(ctx context.Context, homeDir string, input io.Reader, cfg *config.Config) error {
	log.Info("Initializing context")

	c := core.NewContext().
		WithHomeDir(homeDir).
		WithInput(input)
	if err := c.Setup(ctx, cfg); err != nil {
		return err
	}

	// Seal the context.
	c.Seal()

	// Attach the code context to the Node instance.
	n.WithContext(c)
	return nil
}

// Setup sets up the context, scheduler and API server for the Node.
func (n *Node) Setup(homeDir string, input io.Reader, cfg *config.Config) error {
	return n.Manager.Setup(func(ctx context.Context) error {
		log.Info("Setting up context")
		if err := n.SetupContext(ctx, homeDir, input, cfg); err != nil {
			return fmt.Errorf("setting up context: %w", err)
		}

		log.Info("Setting up scheduler")
		if err := n.SetupScheduler(ctx, cfg); err != nil {
			return fmt.Errorf("setting up scheduler: %w", err)
		}

		log.Info("Setting up API server")
		if err := n.SetupServer(ctx, cfg); err != nil {
			return fmt.Errorf("setting up API server: %w", err)
		}

		return nil
	})
}
