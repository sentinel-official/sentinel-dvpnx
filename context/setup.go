package context

import (
	"io"

	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	sdkclient "github.com/sentinel-official/sentinel-go-sdk/client"
	"github.com/sentinel-official/sentinel-go-sdk/libs/geoip"
	sdk "github.com/sentinel-official/sentinel-go-sdk/types"

	"github.com/sentinel-official/dvpn-node/database"
)

// SetupClient initializes and sets up the SDK client.
func (c *Context) SetupClient(input io.Reader) error {
	// Create a new codec for proto encoding/decoding.
	cdc := sdk.NewProtoCodec()

	// Initialize a keyring with the specified settings.
	kr, err := keyring.New(c.AppName(), c.KeyringBackend(), c.KeyringDir(), input, cdc)
	if err != nil {
		return err
	}

	// Create a new client instance using the codec and keyring.
	client := sdkclient.New(cdc).
		WithKeyring(kr)

	// Set the client in the context.
	c.WithClient(client)
	return nil
}

// SetupDatabase initializes and sets up the database connection.
func (c *Context) SetupDatabase() error {
	// Initialize a new database connection using the determined file path.
	db, err := database.NewDefault(c.DatabaseFilePath())
	if err != nil {
		return err
	}

	// Set the newly created database connection in the context.
	c.WithDatabase(db)
	return nil
}

// SetupGeoIPClient initializes and sets up the GeoIP client in the context.
func (c *Context) SetupGeoIPClient() error {
	// Create a new default GeoIP client.
	client := geoip.NewDefaultClient()

	// Set the new GeoIP client in the context.
	c.WithGeoIPClient(client)
	return nil
}

func (c *Context) SetupService() error {
	return nil
}

// SetupTxFromAddr retrieves and sets the address of the transaction sender based on the account name.
func (c *Context) SetupTxFromAddr() error {
	// Retrieve the key record by name.
	key, err := c.Key(c.TxFromName())
	if err != nil {
		return err
	}

	// Get the address from the key record.
	addr, err := key.GetAddress()
	if err != nil {
		return err
	}

	// Set the address in the context.
	c.WithTxFromAddr(addr)
	return nil
}

// Setup initializes all the necessary components in the context.
func (c *Context) Setup(input io.Reader) error {
	// Setup the client with the provided input reader.
	if err := c.SetupClient(input); err != nil {
		return err
	}

	// Setup the database connection.
	if err := c.SetupDatabase(); err != nil {
		return err
	}

	// Setup the GeoIP client for geolocation services.
	if err := c.SetupGeoIPClient(); err != nil {
		return err
	}

	// Setup the service layer components.
	if err := c.SetupService(); err != nil {
		return err
	}

	// Setup the transaction from address for outgoing transactions.
	if err := c.SetupTxFromAddr(); err != nil {
		return err
	}

	return nil
}
