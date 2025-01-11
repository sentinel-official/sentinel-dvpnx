package session

import (
	"encoding/base64"
	"errors"

	"github.com/cosmos/cosmos-sdk/types"
	"github.com/gin-gonic/gin"
)

type RequestAddSession struct {
	AccAddr   types.AccAddress
	Data      []byte
	Signature []byte

	URI struct {
		AccAddr string `uri:"acc_addr"`
		ID      uint64 `uri:"id" binding:"gt=0"`
	}
	Body struct {
		Data      string `json:"data"`
		Signature string `json:"signature"`
	}
}

func (r *RequestAddSession) Msg() []byte {
	return append(types.Uint64ToBigEndian(r.URI.ID), r.Data...)
}

func NewRequestAddSession(c *gin.Context) (req *RequestAddSession, err error) {
	req = &RequestAddSession{}
	if err = c.ShouldBindUri(&req.URI); err != nil {
		return nil, err
	}
	if err = c.ShouldBindJSON(&req.Body); err != nil {
		return nil, err
	}

	if req.Body.Data == "" {
		return nil, errors.New("data cannot be empty")
	}
	if req.Body.Signature == "" {
		return nil, errors.New("signature cannot be empty")
	}

	req.AccAddr, err = types.AccAddressFromBech32(req.URI.AccAddr)
	if err != nil {
		return nil, err
	}

	req.Data, err = base64.StdEncoding.DecodeString(req.Body.Data)
	if err != nil {
		return nil, err
	}

	req.Signature, err = base64.StdEncoding.DecodeString(req.Body.Signature)
	if err != nil {
		return nil, err
	}

	return req, nil
}
