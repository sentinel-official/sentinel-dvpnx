package node

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sentinel-official/sentinel-go-sdk/libs/cron"

	"github.com/sentinel-official/dvpn-node/api"
	"github.com/sentinel-official/dvpn-node/workers"
)

// init sets the Gin mode to ReleaseMode.
func init() {
	gin.SetMode(gin.ReleaseMode)
}

// SetupRouter sets up the HTTP router with necessary middlewares and routes.
func (n *Node) SetupRouter() error {
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
	api.RegisterRoutes(n.ctx, r)

	// Attach the configured router to the Node.
	n.WithRouter(r)
	return nil
}

// SetupScheduler sets up the cron scheduler with various workers.
func (n *Node) SetupScheduler() error {
	// Define the list of cron workers with their respective handlers and intervals.
	items := []cron.Worker{
		cron.NewBasicWorker().WithName("BestRPCAddress").
			WithHandler(workers.HandlerBestRPCAddr(n.ctx)).
			WithInterval(n.ctx.IntervalBestRPCAddr()),
		cron.NewBasicWorker().WithName("GeoIPLocation").
			WithHandler(workers.HandlerGeoIPLocation(n.ctx)).
			WithInterval(n.ctx.IntervalGeoIPLocation()),
		cron.NewBasicWorker().WithName("SessionUsageSyncWithBlockchain").
			WithHandler(workers.HandlerSessionUsageSyncWithBlockchain(n.ctx)).
			WithInterval(n.ctx.IntervalSessionUsageSyncWithBlockchain()),
		cron.NewBasicWorker().WithName("SessionUsageSyncWithDatabase").
			WithHandler(workers.HandlerSessionUsageSyncWithDatabase(n.ctx)).
			WithInterval(n.ctx.IntervalSessionUsageSyncWithDatabase()),
		cron.NewBasicWorker().WithName("SessionUsageValidate").
			WithHandler(workers.HandlerSessionUsageValidate(n.ctx)).
			WithInterval(n.ctx.IntervalSessionUsageValidate()),
		cron.NewBasicWorker().WithName("SessionValidate").
			WithHandler(workers.HandlerSessionValidate(n.ctx)).
			WithInterval(n.ctx.IntervalSessionValidate()),
		cron.NewBasicWorker().WithName("Speedtest").
			WithHandler(workers.HandlerSpeedtest(n.ctx)).
			WithInterval(n.ctx.IntervalSpeedtest()),
		cron.NewBasicWorker().WithName("NodeStatusUpdate").
			WithHandler(workers.HandlerNodeStatusUpdate(n.ctx)).
			WithInterval(n.ctx.IntervalStatusUpdate()),
	}

	// Create a new cron scheduler and register the workers.
	s := cron.NewScheduler()
	if err := s.RegisterWorkers(items...); err != nil {
		return err
	}

	// Attach the configured scheduler to the Node.
	n.WithScheduler(s)
	return nil
}

// Setup sets up both the router and scheduler for the Node.
func (n *Node) Setup() error {
	// Set up the HTTP router.
	if err := n.SetupRouter(); err != nil {
		return err
	}

	// Set up the cron scheduler.
	if err := n.SetupScheduler(); err != nil {
		return err
	}

	return nil
}
