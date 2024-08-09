package workers

import (
	"context"
	"errors"

	v1sentinelhub "github.com/sentinel-official/hub/v12/types/v1"
	nodetypes "github.com/sentinel-official/hub/v12/x/node/types/v3"

	nodecontext "github.com/sentinel-official/dvpn-node/context"
)

// handleUpdateNodeStatus returns a function that broadcasts a transaction to update the node's status to active.
func handleUpdateNodeStatus(ctx *nodecontext.Context) func() error {
	return func() error {
		// Create a new message to update the node's status to active.
		msg := nodetypes.NewMsgUpdateNodeStatusRequest(
			ctx.TxFromAddr().Bytes(),
			v1sentinelhub.StatusActive,
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
