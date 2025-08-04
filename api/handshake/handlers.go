package handshake

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"cosmossdk.io/math"
	cosmossdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gin-gonic/gin"
	"github.com/sentinel-official/sentinel-go-sdk/node"
	"github.com/sentinel-official/sentinel-go-sdk/types"
	"github.com/sentinel-official/sentinelhub/v12/types/v1"

	"github.com/sentinel-official/sentinel-dvpnx/core"
	"github.com/sentinel-official/sentinel-dvpnx/database/models"
	"github.com/sentinel-official/sentinel-dvpnx/database/operations"
)

// handlerInitHandshake returns a handler function to process the request for performing a handshake.
func handlerInitHandshake(c *core.Context) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// TODO: validate current peer count

		// Parse and verify the request.
		req, err := newInitHandshakeRequest(ctx)
		if err != nil {
			err = fmt.Errorf("invalid request format: %w", err)
			ctx.JSON(http.StatusBadRequest, types.NewResponseError(2, err))
			return
		}

		// Check if a session already exists by ID.
		query := map[string]interface{}{
			"id": req.Body.ID,
		}

		record, err := operations.SessionFindOne(c.Database(), query)
		if err != nil {
			err = fmt.Errorf("failed to retrieve session from database: %w", err)
			ctx.JSON(http.StatusInternalServerError, types.NewResponseError(3, err))
			return
		}
		if record != nil {
			err = fmt.Errorf("session already exists for id %d", req.Body.ID)
			ctx.JSON(http.StatusConflict, types.NewResponseError(3, err))
			return
		}

		// Check if a session already exists by peer request data.
		query = map[string]interface{}{
			"peer_request": base64.StdEncoding.EncodeToString(req.Body.Data),
		}

		record, err = operations.SessionFindOne(c.Database(), query)
		if err != nil {
			err = fmt.Errorf("failed to retrieve session from database: %w", err)
			ctx.JSON(http.StatusInternalServerError, types.NewResponseError(4, err))
			return
		}
		if record != nil {
			err = errors.New("session already exists for peer request")
			ctx.JSON(http.StatusConflict, types.NewResponseError(4, err))
			return
		}

		// Fetch session details from blockchain.
		session, err := c.Client().Session(ctx, req.Body.ID)
		if err != nil {
			err = fmt.Errorf("failed to query session from blockchain: %w", err)
			ctx.JSON(http.StatusInternalServerError, types.NewResponseError(5, err))
			return
		}
		if session == nil {
			err = fmt.Errorf("session %d does not exist", req.Body.ID)
			ctx.JSON(http.StatusNotFound, types.NewResponseError(5, err))
			return
		}

		// Validate session status.
		if !session.GetStatus().Equal(v1.StatusActive) {
			err = fmt.Errorf("invalid session status; got %s, expected %s", session.GetStatus(), v1.StatusActive)
			ctx.JSON(http.StatusBadRequest, types.NewResponseError(5, err))
			return
		}

		// Validate node address.
		if session.GetNodeAddress() != c.NodeAddr().String() {
			err = fmt.Errorf("node address mismatch; got %s, expected %s", session.GetNodeAddress(), c.NodeAddr())
			ctx.JSON(http.StatusBadRequest, types.NewResponseError(6, err))
			return
		}

		// Validate account address.
		accAddr, err := cosmossdk.AccAddressFromBech32(session.GetAccAddress())
		if err != nil {
			err = fmt.Errorf("invalid account address format: %w", err)
			ctx.JSON(http.StatusInternalServerError, types.NewResponseError(6, err))
			return
		}
		if got := req.AccAddr(); !got.Equals(accAddr) {
			err = fmt.Errorf("account address mismatch; got %s, expected %s", got, accAddr)
			ctx.JSON(http.StatusUnauthorized, types.NewResponseError(6, err))
			return
		}

		// Add the peer to the active service.
		id, data, err := c.Service().AddPeer(ctx, req.Body.Data)
		if err != nil {
			err = fmt.Errorf("failed to add peer: %w", err)
			ctx.JSON(http.StatusInternalServerError, types.NewResponseError(7, err))
			return
		}

		// Encode and prepare the handshake response.
		res := &node.InitHandshakeResult{Addrs: c.RemoteAddrs()}
		if res.Data, err = json.Marshal(data); err != nil {
			err = fmt.Errorf("failed to encode response: %w", err)
			ctx.JSON(http.StatusInternalServerError, types.NewResponseError(8, err))
			return
		}

		// Insert the session record into the database.
		item := models.NewSession().
			WithAccAddr(accAddr).
			WithDownloadBytes(math.ZeroInt()).
			WithDuration(0).
			WithID(session.GetID()).
			WithMaxBytes(session.GetMaxBytes()).
			WithMaxDuration(session.GetMaxDuration()).
			WithNodeAddr(c.NodeAddr()).
			WithPeerID(id).
			WithPeerRequest(req.Body.Data).
			WithServiceType(c.Service().Type()).
			WithSignature(nil).
			WithUploadBytes(math.ZeroInt())

		if err = operations.SessionInsertOne(c.Database(), item); err != nil {
			err = fmt.Errorf("failed to insert session into database: %w", err)
			ctx.JSON(http.StatusInternalServerError, types.NewResponseError(9, err))
			return
		}

		// Return a successful response.
		ctx.JSON(http.StatusOK, types.NewResponseResult(res))
	}
}
