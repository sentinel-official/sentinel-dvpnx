package session

import (
	"encoding/base64"
	"fmt"

	"github.com/cosmos/cosmos-sdk/crypto/types"
	cosmossdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gin-gonic/gin"
	"github.com/sentinel-official/sentinel-go-sdk/utils"
)

type RequestAddSession struct {
	Data      []byte
	PubKey    types.PubKey
	Signature []byte

	Body struct {
		Data      string `json:"data" binding:"required,base64,gt=0"`
		ID        uint64 `json:"id" binding:"required,gt=0"`
		PubKey    string `json:"pub_key" binding:"required,gt=0"`
		Signature string `json:"signature" binding:"required,base64,gt=0"`
	}
}

func (r *RequestAddSession) Msg() []byte {
	return append(cosmossdk.Uint64ToBigEndian(r.Body.ID), r.Body.Data...)
}

func newRequestAddSession(c *gin.Context) (req *RequestAddSession, err error) {
	req = &RequestAddSession{}
	if err = c.ShouldBindJSON(&req.Body); err != nil {
		return nil, fmt.Errorf("failed to bind json body: %w", err)
	}

	req.Data, err = base64.StdEncoding.DecodeString(req.Body.Data)
	if err != nil {
		return nil, fmt.Errorf("failed to decode data: %s", err)
	}

	req.Signature, err = base64.StdEncoding.DecodeString(req.Body.Signature)
	if err != nil {
		return nil, fmt.Errorf("failed to decode signature: %s", err)
	}

	req.PubKey, err = utils.DecodePubKey(req.Body.PubKey)
	if err != nil {
		return nil, fmt.Errorf("failed to decode pub_key: %s", err)
	}

	return req, nil
}
