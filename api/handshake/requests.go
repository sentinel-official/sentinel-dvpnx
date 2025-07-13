package handshake

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/types"
	"github.com/gin-gonic/gin"
	"github.com/sentinel-official/sentinel-go-sdk/node"
)

// InitHandshakeRequest represents the request for performing a handshake.
type InitHandshakeRequest struct {
	Body node.InitHandshakeRequestBody
}

// AccAddr returns the account address from the request body.
func (r *InitHandshakeRequest) AccAddr() types.AccAddress {
	addr, err := r.Body.AccAddr()
	if err != nil {
		panic(err)
	}

	return addr
}

// newInitHandshakeRequest parses, binds, and verifies the handshake request.
func newInitHandshakeRequest(c *gin.Context) (req *InitHandshakeRequest, err error) {
	req = &InitHandshakeRequest{}

	// Bind JSON request to the struct.
	if err := c.ShouldBindJSON(&req.Body); err != nil {
		return nil, fmt.Errorf("failed to bind json body: %w", err)
	}

	// Verify the request body.
	if err := req.Body.Verify(); err != nil {
		return nil, fmt.Errorf("failed to verify body: %w", err)
	}

	return req, nil
}
