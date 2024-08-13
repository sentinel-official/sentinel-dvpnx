package workers

import (
	"context"
	"errors"

	v1base "github.com/sentinel-official/hub/v12/types/v1"
	nodetypes "github.com/sentinel-official/hub/v12/x/node/types/v3"

	"github.com/sentinel-official/dvpn-node/node"
)

// HandlerNodeStatusUpdate returns a function that broadcasts a transaction to update the node's status to active.
func HandlerNodeStatusUpdate(ctx *node.Context) func() error {
	return func() error {
		// Create a new message to update the node's status to active.
		msg := nodetypes.NewMsgUpdateNodeStatusRequest(
			ctx.TxFromAddr().Bytes(),
			v1base.StatusActive,
		)

		// Broadcast the transaction message to the blockchain.
		res, err := ctx.BroadcastTx(context.TODO(), msg)
		if err != nil {
			return err
		}
		if !res.TxResult.IsOK() {
			return errors.New(res.TxResult.GetLog())
		}

		return nil
	}
}
