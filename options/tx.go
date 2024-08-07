package options

import (
	sdkflags "github.com/sentinel-official/sentinel-go-sdk/cmd/flags"
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

// NewDefaultTx creates a new TxOptions instance with default values.
func NewDefaultTx() *TxOptions {
	return &TxOptions{
		ChainID:            sdkflags.DefaultTxChainID,
		Gas:                sdkflags.DefaultTxGas,
		GasAdjustment:      sdkflags.DefaultTxGasAdjustment,
		GasPrices:          sdkflags.DefaultTxGasPrices,
		SimulateAndExecute: sdkflags.DefaultTxSimulateAndExecute,
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
	sdkflags.SetFlagTxChainID(cmd)
	sdkflags.SetFlagTxFeeGranterAddr(cmd)
	sdkflags.SetFlagTxFromName(cmd)
	sdkflags.SetFlagTxGas(cmd)
	sdkflags.SetFlagTxGasAdjustment(cmd)
	sdkflags.SetFlagTxGasPrices(cmd)
	sdkflags.SetFlagTxSimulateAndExecute(cmd)
}

// NewTxOptionsFromCmd creates and returns TxOptions from the given cobra command's flags.
func NewTxOptionsFromCmd(cmd *cobra.Command) (*TxOptions, error) {
	// Retrieve the chain ID flag value from the command.
	chainID, err := sdkflags.GetTxChainIDFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieve the fee granter address flag value from the command.
	feeGranterAddr, err := sdkflags.GetTxFeeGranterAddrFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieve the from name flag value from the command.
	fromName, err := sdkflags.GetTxFromNameFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieve the gas flag value from the command.
	gas, err := sdkflags.GetTxGasFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieve the gas adjustment flag value from the command.
	gasAdjustment, err := sdkflags.GetTxGasAdjustmentFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieve the gas prices flag value from the command.
	gasPrices, err := sdkflags.GetTxGasPricesFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieve the simulate and execute flag value from the command.
	simulateAndExecute, err := sdkflags.GetTxSimulateAndExecuteFromCmd(cmd)
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
