package session

import (
	"encoding/base64"
	"fmt"

	"github.com/cosmos/cosmos-sdk/crypto/types"
	cosmossdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gin-gonic/gin"
	"github.com/sentinel-official/sentinel-go-sdk/node"
	"github.com/sentinel-official/sentinel-go-sdk/utils"
)

// AddSessionRequest represents the request for adding a session.
type AddSessionRequest struct {
	Body *node.AddSessionRequestBody

	Data      []byte
	PubKey    types.PubKey
	Signature []byte
}

// Msg generates the message from the session request.
func (r *AddSessionRequest) Msg() []byte {
	return append(cosmossdk.Uint64ToBigEndian(r.Body.ID), r.Data...)
}

// AccAddr returns the account address derived from the public key.
func (r *AddSessionRequest) AccAddr() cosmossdk.AccAddress {
	return r.PubKey.Address().Bytes()
}

// newAddSessionRequest binds and decodes the incoming request.
func newAddSessionRequest(c *gin.Context) (req *AddSessionRequest, err error) {
	req = &AddSessionRequest{}
	if err = c.ShouldBindJSON(&req.Body); err != nil {
		return nil, fmt.Errorf("failed to bind json body: %w", err)
	}

	// Decode base64 encoded data
	req.Data, err = base64.StdEncoding.DecodeString(req.Body.Data)
	if err != nil {
		return nil, fmt.Errorf("failed to decode data: %w", err)
	}

	// Decode base64 encoded signature
	req.Signature, err = base64.StdEncoding.DecodeString(req.Body.Signature)
	if err != nil {
		return nil, fmt.Errorf("failed to decode signature: %w", err)
	}

	// Decode public key
	req.PubKey, err = utils.DecodePubKey(req.Body.PubKey)
	if err != nil {
		return nil, fmt.Errorf("failed to decode pub_key: %w", err)
	}

	return req, nil
}
