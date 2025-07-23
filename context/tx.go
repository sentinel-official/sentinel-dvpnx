package context

import (
	"context"

	"github.com/cosmos/cosmos-sdk/types"
)

// BroadcastTx safely broadcasts a transaction with the provided messages.
// It locks the transaction mutex to ensure only one transaction is broadcast at a time.
func (c *Context) BroadcastTx(ctx context.Context, msgs ...types.Msg) error {
	c.txm.Lock()
	defer c.txm.Unlock()

	// No messages to broadcast, skipping.
	if len(msgs) == 0 {
		return nil
	}

	// Broadcast the transaction and wait for it to be included in a block.
	_, _, err := c.Client().BroadcastTxBlock(ctx, msgs...)
	if err != nil {
		return err
	}

	return nil
}
