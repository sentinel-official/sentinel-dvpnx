package info

import (
	"net/http"

	"github.com/cosmos/cosmos-sdk/version"
	"github.com/gin-gonic/gin"
	"github.com/sentinel-official/sentinel-go-sdk/libs/geoip"
	"github.com/sentinel-official/sentinel-go-sdk/types"

	"github.com/sentinel-official/dvpn-node/context"
)

// HandlerGetInfo returns a handler function to retrieve node information.
func HandlerGetInfo(c *context.Context) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		downLink, upLink := c.SpeedtestResults()
		loc := c.Location()

		// Construct the result structure with node information.
		res := ResultGetInfo{
			Addr:         c.NodeAddr().String(),
			DownLink:     downLink.String(),
			HandshakeDNS: false,
			Location: &geoip.Location{
				City:      loc.City,
				Country:   loc.Country,
				Latitude:  loc.Latitude,
				Longitude: loc.Longitude,
			},
			Moniker: c.Moniker(),
			Peers:   c.Service().PeerCount(),
			Type:    c.Service().Type().String(),
			UpLink:  upLink.String(),
			Version: version.Version,
		}

		// Send the result as a JSON response with HTTP status 200.
		ctx.JSON(http.StatusOK, types.NewResponseResult(res))
	}
}
