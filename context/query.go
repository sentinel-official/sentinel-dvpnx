package context

import (
	"context"

	sessiontypes "github.com/sentinel-official/hub/v12/x/session/types/v3"
	"github.com/v2fly/v2ray-core/v5/common/retry"
)

// QuerySession retrieves a session by its ID from the client.
func (c *Context) QuerySession(ctx context.Context, id uint64) (sessiontypes.Session, error) {
	// Get client options for the query.
	opts := c.ClientOptions()

	var err error
	var session sessiontypes.Session

	// Define the function that performs the session query.
	querySessionFunc := func() error {
		session, err = c.Client().Session(ctx, id, opts)
		if err != nil {
			return err
		}

		return nil
	}

	// Retry the session query.
	if err := retry.Timed(5, 1000).On(querySessionFunc); err != nil {
		return nil, err
	}

	return session, nil
}
