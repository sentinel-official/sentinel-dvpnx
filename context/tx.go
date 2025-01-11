package context

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/avast/retry-go/v4"
	abci "github.com/cometbft/cometbft/abci/types"
	core "github.com/cometbft/cometbft/rpc/core/types"
	cosmossdk "github.com/cosmos/cosmos-sdk/types"
)

// isTxInMempoolCacheError checks if the error indicates that the transaction is already present in the mempool cache.
func isTxInMempoolCacheError(err error) bool {
	return strings.Contains(strings.ToLower(err.Error()), "tx already exists in cache")
}

// BroadcastTx broadcasts a transaction to the network and returns the transaction result.
func (c *Context) BroadcastTx(ctx context.Context, msgs ...cosmossdk.Msg) (*core.ResultTx, error) {
	var err error
	var resp *core.ResultBroadcastTx
	var result *core.ResultTx

	// Define a function for broadcasting the transaction with retry logic.
	broadcastTxFunc := func() error {
		// Broadcast the transaction
		resp, err = c.Client().BroadcastTx(ctx, msgs)
		if err != nil {
			// Return the error if the error is not TxInMempoolCacheError.
			if !isTxInMempoolCacheError(err) {
				return fmt.Errorf("failed to broadcast tx: %w", err)
			}
		}

		return nil
	}

	// TODO: broadcast tx retry only on certain errors

	// Retry broadcasting the transaction.
	if err := retry.Do(
		broadcastTxFunc,
		retry.Attempts(5),
		retry.Delay(3*time.Second),
		retry.DelayType(retry.FixedDelay),
		retry.LastErrorOnly(true),
	); err != nil {
		return nil, fmt.Errorf("tx broadcast failed after retries: %w", err)
	}

	// Check if the transaction was successful.
	if resp.Code != abci.CodeTypeOK {
		err := errors.New(resp.Log)
		return nil, fmt.Errorf("tx broadcast failed with code %d: %w", resp.Code, err)
	}

	// Define a function for fetching the transaction result with retry logic.
	txFunc := func() error {
		result, err = c.Client().Tx(ctx, resp.Hash)
		if err != nil {
			return fmt.Errorf("failed to query tx result: %w", err)
		}

		return nil
	}

	// Retry fetching the transaction result.
	if err := retry.Do(
		txFunc,
		retry.Attempts(30),
		retry.Delay(1*time.Second),
		retry.DelayType(retry.FixedDelay),
		retry.LastErrorOnly(true),
	); err != nil {
		return nil, fmt.Errorf("tx query failed after retries: %w", err)
	}

	return result, nil
}
