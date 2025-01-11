package info

import (
	"github.com/sentinel-official/sentinel-go-sdk/libs/geoip"
)

// ResultGetInfo represents metadata about a node.
type ResultGetInfo struct {
	Addr         string          `json:"addr"`
	DownLink     string          `json:"down_link"`
	HandshakeDNS bool            `json:"handshake_dns"`
	Location     *geoip.Location `json:"location"`
	Moniker      string          `json:"moniker"`
	Peers        int             `json:"peers"`
	Type         string          `json:"type"`
	UpLink       string          `json:"up_link"`
	Version      string          `json:"version"`
}
