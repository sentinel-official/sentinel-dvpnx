package node

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sentinel-official/sentinel-go-sdk/libs/cron"

	"github.com/sentinel-official/dvpn-node/api"
	"github.com/sentinel-official/dvpn-node/utils"
	"github.com/sentinel-official/dvpn-node/workers"
)

type Node struct {
	ctx       *Context
	router    *gin.Engine
	scheduler *cron.Scheduler
}

func New(ctx *Context) *Node {
	return &Node{
		ctx: ctx,
	}
}

func (n *Node) WithRouter(v *gin.Engine) *Node {
	n.router = v
	return n
}

func (n *Node) WithScheduler(v *cron.Scheduler) *Node {
	n.scheduler = v
	return n
}

func (n *Node) SetupRouter() error {
	middlewares := []gin.HandlerFunc{
		cors.New(
			cors.Config{
				AllowAllOrigins: true,
				AllowMethods:    []string{http.MethodGet, http.MethodPost},
			},
		),
	}

	r := gin.New()
	r.Use(middlewares...)
	api.RegisterRoutes(n.ctx, r)

	n.WithRouter(r)
	return nil
}

func (n *Node) SetupScheduler() error {
	items := []cron.Worker{
		cron.NewBasicWorker().WithName("BestRPCAddress").
			WithHandler(workers.HandlerBestRPCAddr(n.ctx)).
			WithInterval(n.ctx.NodeIntervalBestRPCAddr()),
		cron.NewBasicWorker().WithName("GeoIPLocation").
			WithHandler(workers.HandlerGeoIPLocation(n.ctx)).
			WithInterval(n.ctx.NodeIntervalGeoIPLocation()),
		cron.NewBasicWorker().WithName("SessionUsageSyncWithBlockchain").
			WithHandler(workers.HandlerSessionUsageSyncWithBlockchain(n.ctx)).
			WithInterval(n.ctx.NodeIntervalSessionUsageSyncWithBlockchain()),
		cron.NewBasicWorker().WithName("SessionUsageSyncWithDatabase").
			WithHandler(workers.HandlerSessionUsageSyncWithDatabase(n.ctx)).
			WithInterval(n.ctx.NodeIntervalSessionUsageSyncWithDatabase()),
		cron.NewBasicWorker().WithName("SessionUsageValidate").
			WithHandler(workers.HandlerSessionUsageValidate(n.ctx)).
			WithInterval(n.ctx.NodeIntervalSessionUsageValidate()),
		cron.NewBasicWorker().WithName("SessionValidate").
			WithHandler(workers.HandlerSessionValidate(n.ctx)).
			WithInterval(n.ctx.NodeIntervalSessionValidate()),
		cron.NewBasicWorker().WithName("Speedtest").
			WithHandler(workers.HandlerSpeedtest(n.ctx)).
			WithInterval(n.ctx.NodeIntervalSpeedtest()),
		cron.NewBasicWorker().WithName("NodeStatusUpdate").
			WithHandler(workers.HandlerNodeStatusUpdate(n.ctx)).
			WithInterval(n.ctx.NodeIntervalStatusUpdate()),
	}

	s := cron.NewScheduler()
	if err := s.RegisterWorkers(items...); err != nil {
		return err
	}

	n.WithScheduler(s)
	return nil
}

func (n *Node) Start() error {
	if err := n.scheduler.Start(); err != nil {
		return err
	}

	return utils.ListenAndServeTLS(
		n.ctx.NodeListenOn(),
		n.ctx.NodeTLSCertPath(),
		n.ctx.NodeTLSKeyPath(),
		n.router,
	)
}
