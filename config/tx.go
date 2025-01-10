package config

import (
	"errors"
	"fmt"

	"github.com/cosmos/cosmos-sdk/types"
)

type TxConfig struct {
	ChainID            string  `mapstructure:"chain_id"`             // ChainID is the identifier of the blockchain network.
	FeeGranterAddr     string  `mapstructure:"fee_granter_addr"`     // FeeGranterAddr is the address of the entity granting fees.
	FromName           string  `mapstructure:"from_name"`            // FromName is the name of the sender's account.
	Gas                uint64  `mapstructure:"gas"`                  // Gas is the gas limit for the transaction.
	GasAdjustment      float64 `mapstructure:"gas_adjustment"`       // GasAdjustment is the adjustment factor for gas estimation.
	GasPrices          string  `mapstructure:"gas_prices"`           // GasPrices is the price of gas for the transaction.
	SimulateAndExecute bool    `mapstructure:"simulate_and_execute"` // SimulateAndExecute indicates whether to simulate the transaction before execution.
}

// GetChainID returns the ChainID field.
func (c *TxConfig) GetChainID() string {
	return c.ChainID
}

// GetFeeGranterAddr returns the FeeGranterAddr field.
func (c *TxConfig) GetFeeGranterAddr() types.AccAddress {
	if c.FeeGranterAddr == "" {
		return nil
	}

	addr, err := types.AccAddressFromBech32(c.FeeGranterAddr)
	if err != nil {
		panic(err)
	}

	return addr
}

// GetFromName returns the FromName field.
func (c *TxConfig) GetFromName() string {
	return c.FromName
}

// GetGas returns the Gas field.
func (c *TxConfig) GetGas() uint64 {
	return c.Gas
}

// GetGasAdjustment returns the GasAdjustment field.
func (c *TxConfig) GetGasAdjustment() float64 {
	return c.GasAdjustment
}

// GetGasPrices returns the GasPrices field.
func (c *TxConfig) GetGasPrices() types.DecCoins {
	coins, err := types.ParseDecCoins(c.GasPrices)
	if err != nil {
		panic(err)
	}

	return coins
}

// GetSimulateAndExecute returns the SimulateAndExecute field.
func (c *TxConfig) GetSimulateAndExecute() bool {
	return c.SimulateAndExecute
}

// Validate validates the Tx configuration.
func (c *TxConfig) Validate() error {
	if c.ChainID == "" {
		return errors.New("chain_id cannot be empty")
	}
	if c.FeeGranterAddr != "" {
		if _, err := types.AccAddressFromBech32(c.FeeGranterAddr); err != nil {
			return fmt.Errorf("invalid fee_granter_addr: %w", err)
		}
	}
	if c.FromName == "" {
		return errors.New("from_name cannot be empty")
	}
	if c.GasAdjustment < 0 {
		return errors.New("gas_adjustment cannot be negative")
	}
	if c.GasPrices != "" {
		if _, err := types.ParseDecCoins(c.GasPrices); err != nil {
			return fmt.Errorf("invalid gas_prices: %w", err)
		}
	}

	return nil
}

func DefaultTxConfig() TxConfig {
	return TxConfig{
		ChainID:            "sentinelhub-2",
		FeeGranterAddr:     "",
		FromName:           "default",
		Gas:                200_000,
		GasAdjustment:      1.0 + 1.0/6,
		GasPrices:          "0.1udvpn",
		SimulateAndExecute: true,
	}
}
