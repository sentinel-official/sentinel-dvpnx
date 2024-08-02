package options

import (
	"github.com/sentinel-official/sentinel-go-sdk/client/options"
	"github.com/spf13/cobra"
)

// TxOptions represents configuration options for transactions.
type TxOptions struct {
	ChainID            string  `json:"chain_id" toml:"chain_id"`                         // ChainID is the identifier of the blockchain network.
	FeeGranterAddr     string  `json:"fee_granter_addr" toml:"fee_granter_addr"`         // FeeGranterAddr is the address of the entity granting fees.
	FromName           string  `json:"from_name" toml:"from_name"`                       // FromName is the name of the sender's account.
	Gas                uint64  `json:"gas" toml:"gas"`                                   // Gas is the gas limit for the transaction.
	GasAdjustment      float64 `json:"gas_adjustment" toml:"gas_adjustment"`             // GasAdjustment is the adjustment factor for gas estimation.
	GasPrices          string  `json:"gas_prices" toml:"gas_prices"`                     // GasPrices is the price of gas for the transaction.
	SimulateAndExecute bool    `json:"simulate_and_execute" toml:"simulate_and_execute"` // SimulateAndExecute indicates whether to simulate the transaction before execution.
}

// NewDefaultTxOptions creates a new TxOptions instance with default values.
func NewDefaultTxOptions() *TxOptions {
	return &TxOptions{
		ChainID:            options.DefaultTxChainID,
		Gas:                options.DefaultTxGas,
		GasAdjustment:      options.DefaultTxGasAdjustment,
		GasPrices:          options.DefaultTxGasPrices,
		SimulateAndExecute: options.DefaultTxSimulateAndExecute,
	}
}

// WithChainID sets the ChainID field and returns the modified TxOptions instance.
func (t *TxOptions) WithChainID(v string) *TxOptions {
	t.ChainID = v
	return t
}

// WithFeeGranterAddr sets the FeeGranterAddr field and returns the modified TxOptions instance.
func (t *TxOptions) WithFeeGranterAddr(v string) *TxOptions {
	t.FeeGranterAddr = v
	return t
}

// WithFromName sets the FromName field and returns the modified TxOptions instance.
func (t *TxOptions) WithFromName(v string) *TxOptions {
	t.FromName = v
	return t
}

// WithGas sets the Gas field and returns the modified TxOptions instance.
func (t *TxOptions) WithGas(v uint64) *TxOptions {
	t.Gas = v
	return t
}

// WithGasAdjustment sets the GasAdjustment field and returns the modified TxOptions instance.
func (t *TxOptions) WithGasAdjustment(v float64) *TxOptions {
	t.GasAdjustment = v
	return t
}

// WithGasPrices sets the GasPrices field and returns the modified TxOptions instance.
func (t *TxOptions) WithGasPrices(v string) *TxOptions {
	t.GasPrices = v
	return t
}

// WithSimulateAndExecute sets the SimulateAndExecute field and returns the modified TxOptions instance.
func (t *TxOptions) WithSimulateAndExecute(v bool) *TxOptions {
	t.SimulateAndExecute = v
	return t
}

// AddTxFlagsToCmd configures all transaction-related flags for the given command.
func AddTxFlagsToCmd(cmd *cobra.Command) {
	options.SetFlagTxChainID(cmd)
	options.SetFlagTxFeeGranterAddr(cmd)
	options.SetFlagTxFromName(cmd)
	options.SetFlagTxGas(cmd)
	options.SetFlagTxGasAdjustment(cmd)
	options.SetFlagTxGasPrices(cmd)
	options.SetFlagTxSimulateAndExecute(cmd)
}

// NewTxOptionsFromCmd creates and returns TxOptions from the given cobra command's flags.
func NewTxOptionsFromCmd(cmd *cobra.Command) (*TxOptions, error) {
	chainID, err := options.GetTxChainIDFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	feeGranterAddr, err := options.GetTxFeeGranterAddrFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	fromName, err := options.GetTxFromNameFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	gas, err := options.GetTxGasFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	gasAdjustment, err := options.GetTxGasAdjustmentFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	gasPrices, err := options.GetTxGasPricesFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	simulateAndExecute, err := options.GetTxSimulateAndExecuteFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	// Return a new TxOptions instance populated with the retrieved flag values.
	return &TxOptions{
		ChainID:            chainID,
		FeeGranterAddr:     feeGranterAddr,
		FromName:           fromName,
		Gas:                gas,
		GasAdjustment:      gasAdjustment,
		GasPrices:          gasPrices,
		SimulateAndExecute: simulateAndExecute,
	}, nil
}
