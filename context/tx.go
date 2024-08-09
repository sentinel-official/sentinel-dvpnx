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
	c.Client().Lock()
	defer c.Client().Unlock()

	// Get options for broadcasting and querying transactions.
	opts := c.ClientOptions()

	var err error
	var resp *coretypes.ResultBroadcastTx
	var result *coretypes.ResultTx

	// Define a function for broadcasting the transaction with retry logic.
	broadcastTxFunc := func() error {
		// Broadcast the transaction
		resp, err = c.Client().BroadcastTx(ctx, messages, opts)
		if err != nil {
			// If the error is not related to the transaction being already in the cache, return the error.
			if !isTxInMempoolCacheError(err) {
				return err
			}
		}

		// Check if the transaction was successful.
		if resp.Code != abci.CodeTypeOK {
			return errors.New(resp.Log)
		}

		return nil
	}

	// Retry broadcasting the transaction.
	if err := retry.Timed(5, 6000).On(broadcastTxFunc); err != nil {
		return nil, err
	}

	// Define a function for fetching the transaction result with retry logic.
	txFunc := func() error {
		result, err = c.Client().Tx(ctx, resp.Hash, opts)
		if err != nil {
			return err
		}

		return nil
	}

	// Retry fetching the transaction result.
	if err := retry.Timed(60, 1000).On(txFunc); err != nil {
		return nil, err
	}

	return result, nil
}
