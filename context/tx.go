package context

import (
	"context"
	"errors"
	"strings"

	abci "github.com/cometbft/cometbft/abci/types"
	coretypes "github.com/cometbft/cometbft/rpc/core/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/v2fly/v2ray-core/v5/common/retry"
)

// isTxInMempoolCacheError checks if the error indicates that the transaction is already present in the mempool cache.
func isTxInMempoolCacheError(err error) bool {
	return strings.Contains(strings.ToLower(err.Error()), "tx already exists in cache")
}

// BroadcastTx broadcasts a transaction to the network and returns the transaction result.
func (c *Context) BroadcastTx(ctx context.Context, messages ...sdk.Msg) (*coretypes.ResultTx, error) {
	// Get options for broadcasting and querying transactions.
	opts := c.ClientOptions()

	var err error
	var resp *coretypes.ResultBroadcastTx

	// Retry broadcasting the transaction
	err = retry.Timed(5, 6000).On(func() error {
		c.Client().Lock()
		defer c.Client().Unlock()

		// Broadcast the transaction
		resp, err = c.Client().BroadcastTx(ctx, messages, opts)
		if err != nil {
			if !isTxInMempoolCacheError(err) {
				return err
			}
		}

		// Check if the transaction was successful
		if resp.Code != abci.CodeTypeOK {
			return errors.New(resp.Log)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	var result *coretypes.ResultTx

	// Retry fetching the transaction result
	err = retry.Timed(60, 1000).On(func() error {
		result, err = c.Client().Tx(ctx, resp.Hash, opts)
		return err
	})

	if err != nil {
		return nil, err
	}

	return result, nil
}
