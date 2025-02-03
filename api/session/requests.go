package session

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/types"
	"github.com/gin-gonic/gin"
	"github.com/sentinel-official/sentinel-go-sdk/node"
)

// AddSessionRequest represents the request for adding a session.
type AddSessionRequest struct {
	Body node.AddSessionRequestBody
}

// AccAddr returns the account address derived from the body.
func (r *AddSessionRequest) AccAddr() types.AccAddress {
	addr, err := r.Body.AccAddr()
	if err != nil {
		panic(err)
	}

	return addr
}

// newAddSessionRequest parses, binds, and verifies the incoming JSON request.
func newAddSessionRequest(c *gin.Context) (req *AddSessionRequest, err error) {
	req = &AddSessionRequest{}

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
