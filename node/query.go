package node

import (
	"context"

	cosmossdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	sessiontypes "github.com/sentinel-official/hub/v12/x/session/types/v3"
	"github.com/v2fly/v2ray-core/v5/common/retry"
)

// QueryAccount retrieves a account by its ID from the client.
func (c *Context) QueryAccount(ctx context.Context, accAddr cosmossdk.AccAddress) (authtypes.AccountI, error) {
	// Get client options for the query.
	opts := c.ClientOptions()

	var err error
	var account authtypes.AccountI

	// Define the function that performs the account query with retry logic.
	fn := func() error {
		account, err = c.Client().Account(ctx, accAddr, opts)
		if err != nil {
			return err
		}

		return nil
	}

	// Retry the account query.
	if err := retry.Timed(5, 1000).On(fn); err != nil {
		return nil, err
	}

	return account, nil
}

// QuerySession retrieves a session by its ID from the client.
func (c *Context) QuerySession(ctx context.Context, id uint64) (sessiontypes.Session, error) {
	// Get client options for the query.
	opts := c.ClientOptions()

	var err error
	var session sessiontypes.Session

	// Define the function that performs the session query with retry logic.
	fn := func() error {
		session, err = c.Client().Session(ctx, id, opts)
		if err != nil {
			return err
		}

		return nil
	}

	// Retry the session query.
	if err := retry.Timed(5, 1000).On(fn); err != nil {
		return nil, err
	}

	return session, nil
}
