package node

import (
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
)

// Key retrieves a key record by name using the client's options.
func (c *Context) Key(name string) (*keyring.Record, error) {
	opts := c.ClientOptions()
	return c.Client().Key(name, opts)
}
