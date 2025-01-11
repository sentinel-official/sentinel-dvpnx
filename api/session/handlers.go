package session

import (
	"encoding/json"
	"fmt"
	"net/http"

	"cosmossdk.io/math"
	"github.com/gin-gonic/gin"
	"github.com/sentinel-official/hub/v12/types/v1"
	"github.com/sentinel-official/sentinel-go-sdk/types"
	"github.com/sentinel-official/sentinel-go-sdk/v2ray"
	"github.com/sentinel-official/sentinel-go-sdk/wireguard"

	"github.com/sentinel-official/dvpn-node/context"
	"github.com/sentinel-official/dvpn-node/database/models"
	"github.com/sentinel-official/dvpn-node/database/operations"
)

func HandlerAddSession(c *context.Context) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// TODO: validate current peer count

		req, err := NewRequestAddSession(ctx)
		if err != nil {
			err := fmt.Errorf("invalid request: %w", err)
			ctx.JSON(http.StatusBadRequest, types.NewResponseError(2, err.Error()))
			return
		}

		query := map[string]interface{}{
			"id": req.URI.ID,
		}

		record, err := operations.SessionFindOne(c.Database(), query)
		if err != nil {
			err := fmt.Errorf("failed to retrieve session from database: %w", err)
			ctx.JSON(http.StatusInternalServerError, types.NewResponseError(3, err.Error()))
			return
		}
		if record != nil {
			err := fmt.Errorf("session already exists for id %d", req.URI.ID)
			ctx.JSON(http.StatusConflict, types.NewResponseError(4, err.Error()))
			return
		}

		var peerKey string
		if c.Service().Type() == types.ServiceTypeV2Ray {
			var r v2ray.AddPeerRequest
			if err := json.Unmarshal(req.Data, &r); err != nil {
				err := fmt.Errorf("failed to unmarshal add peer request: %w", err)
				ctx.JSON(http.StatusInternalServerError, types.NewResponseError(5, err.Error()))
				return
			}

			peerKey = r.Key()
		}
		if c.Service().Type() == types.ServiceTypeWireGuard {
			var r wireguard.AddPeerRequest
			if err := json.Unmarshal(req.Data, &r); err != nil {
				err := fmt.Errorf("failed to unmarshal add peer request: %w", err)
				ctx.JSON(http.StatusInternalServerError, types.NewResponseError(5, err.Error()))
				return
			}

			peerKey = r.Key()
		}

		query = map[string]interface{}{
			"peer_key": peerKey,
		}

		record, err = operations.SessionFindOne(c.Database(), query)
		if err != nil {
			err := fmt.Errorf("failed to retrieve session from database: %w", err)
			ctx.JSON(http.StatusInternalServerError, types.NewResponseError(6, err.Error()))
			return
		}
		if record != nil {
			err := fmt.Errorf("session already exists for peer key %s", peerKey)
			ctx.JSON(http.StatusConflict, types.NewResponseError(7, err.Error()))
			return
		}

		account, err := c.Client().Account(ctx, req.AccAddr)
		if err != nil {
			err := fmt.Errorf("failed to query account for addr %s: %w", req.AccAddr, err)
			ctx.JSON(http.StatusInternalServerError, types.NewResponseError(5, err.Error()))
			return
		}
		if account == nil {
			err := fmt.Errorf("account for addr %s does not exist", req.AccAddr)
			ctx.JSON(http.StatusNotFound, types.NewResponseError(6, err.Error()))
			return
		}
		if account.GetPubKey() == nil {
			err := fmt.Errorf("public key for addr %s does not exist", req.AccAddr)
			ctx.JSON(http.StatusNotFound, types.NewResponseError(7, err.Error()))
			return
		}

		if ok := account.GetPubKey().VerifySignature(req.Msg(), req.Signature); !ok {
			err := fmt.Errorf("invalid signature")
			ctx.JSON(http.StatusBadRequest, types.NewResponseError(8, err.Error()))
			return
		}

		session, err := c.Client().Session(ctx, req.URI.ID)
		if err != nil {
			err := fmt.Errorf("failed to query session %d: %w", req.URI.ID, err)
			ctx.JSON(http.StatusInternalServerError, types.NewResponseError(9, err.Error()))
			return
		}
		if session == nil {
			err := fmt.Errorf("session %d does not exist", req.URI.ID)
			ctx.JSON(http.StatusNotFound, types.NewResponseError(10, err.Error()))
			return
		}
		if !session.GetStatus().Equal(v1.StatusActive) {
			err := fmt.Errorf("invalid session status %s; expected %s", session.GetStatus(), v1.StatusActive)
			ctx.JSON(http.StatusBadRequest, types.NewResponseError(11, err.Error()))
			return
		}
		if session.GetAccAddress() != req.URI.AccAddr {
			err := fmt.Errorf("invalid account addr %s; expected %s", session.GetAccAddress(), req.URI.AccAddr)
			ctx.JSON(http.StatusBadRequest, types.NewResponseError(1, err.Error()))
			return
		}
		if session.GetNodeAddress() != c.NodeAddr().String() {
			err := fmt.Errorf("invalid node addr %s; expected %s", session.GetNodeAddress(), req.URI.AccAddr)
			ctx.JSON(http.StatusBadRequest, types.NewResponseError(1, err.Error()))
			return
		}

		var data interface{}
		if c.Service().Type() == types.ServiceTypeV2Ray {
			var r v2ray.AddPeerRequest
			if err := json.Unmarshal(req.Data, &r); err != nil {
				err := fmt.Errorf("failed to unmarshal add peer request: %w", err)
				ctx.JSON(http.StatusInternalServerError, types.NewResponseError(5, err.Error()))
				return
			}

			data, err = c.Service().AddPeer(ctx, &r)
			if err != nil {
				err := fmt.Errorf("failed to add peer request: %w", err)
				ctx.JSON(http.StatusInternalServerError, types.NewResponseError(6, err.Error()))
				return
			}
		}
		if c.Service().Type() == types.ServiceTypeWireGuard {
			var r wireguard.AddPeerRequest
			if err := json.Unmarshal(req.Data, &r); err != nil {
				err := fmt.Errorf("failed to unmarshal add peer request: %w", err)
				ctx.JSON(http.StatusInternalServerError, types.NewResponseError(6, err.Error()))
				return
			}

			data, err = c.Service().AddPeer(ctx, &r)
			if err != nil {
				err := fmt.Errorf("failed to add peer request: %w", err)
				ctx.JSON(http.StatusInternalServerError, types.NewResponseError(6, err.Error()))
				return
			}
		}

		item := models.NewSession().
			WithID(session.GetID()).
			WithAccAddr(req.AccAddr).
			WithNodeAddr(c.NodeAddr()).
			WithDownloadBytes(math.ZeroInt()).
			WithUploadBytes(math.ZeroInt()).
			WithMaxBytes(session.GetMaxBytes()).
			WithDuration(0).
			WithMaxDuration(session.GetMaxDuration()).
			WithSignature(nil).
			WithPeerKey(peerKey).
			WithServiceType(c.Service().Type())

		if err := operations.SessionInsertOne(c.Database(), item); err != nil {
			err := fmt.Errorf("failed to insert session into database: %w", err)
			ctx.JSON(http.StatusInternalServerError, types.NewResponseError(12, err.Error()))
			return
		}

		res := &ResultAddSession{
			Addrs: c.RemoteAddrs(),
			Data:  data,
		}

		ctx.JSON(http.StatusOK, types.NewResponseResult(res))
	}
}
