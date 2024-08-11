package options

import (
	sdkflags "github.com/sentinel-official/sentinel-go-sdk/cmd/flags"
	"github.com/spf13/cobra"
)

// Tx represents configuration options for transactions.
type Tx struct {
	ChainID            string  `json:"chain_id" toml:"chain_id"`                         // ChainID is the identifier of the blockchain network.
	FeeGranterAddr     string  `json:"fee_granter_addr" toml:"fee_granter_addr"`         // FeeGranterAddr is the address of the entity granting fees.
	FromName           string  `json:"from_name" toml:"from_name"`                       // FromName is the name of the sender's account.
	Gas                uint64  `json:"gas" toml:"gas"`                                   // Gas is the gas limit for the transaction.
	GasAdjustment      float64 `json:"gas_adjustment" toml:"gas_adjustment"`             // GasAdjustment is the adjustment factor for gas estimation.
	GasPrices          string  `json:"gas_prices" toml:"gas_prices"`                     // GasPrices is the price of gas for the transaction.
	SimulateAndExecute bool    `json:"simulate_and_execute" toml:"simulate_and_execute"` // SimulateAndExecute indicates whether to simulate the transaction before execution.
}

// NewTx creates a new Tx instance with default values.
func NewTx() *Tx {
	return &Tx{
		ChainID:            sdkflags.DefaultTxChainID,
		Gas:                sdkflags.DefaultTxGas,
		GasAdjustment:      sdkflags.DefaultTxGasAdjustment,
		GasPrices:          sdkflags.DefaultTxGasPrices,
		SimulateAndExecute: sdkflags.DefaultTxSimulateAndExecute,
	}
}

// WithChainID sets the ChainID field and returns the modified Tx instance.
func (t *Tx) WithChainID(v string) *Tx {
	t.ChainID = v
	return t
}

// WithFeeGranterAddr sets the FeeGranterAddr field and returns the modified Tx instance.
func (t *Tx) WithFeeGranterAddr(v string) *Tx {
	t.FeeGranterAddr = v
	return t
}

// WithFromName sets the FromName field and returns the modified Tx instance.
func (t *Tx) WithFromName(v string) *Tx {
	t.FromName = v
	return t
}

// WithGas sets the Gas field and returns the modified Tx instance.
func (t *Tx) WithGas(v uint64) *Tx {
	t.Gas = v
	return t
}

// WithGasAdjustment sets the GasAdjustment field and returns the modified Tx instance.
func (t *Tx) WithGasAdjustment(v float64) *Tx {
	t.GasAdjustment = v
	return t
}

// WithGasPrices sets the GasPrices field and returns the modified Tx instance.
func (t *Tx) WithGasPrices(v string) *Tx {
	t.GasPrices = v
	return t
}

// WithSimulateAndExecute sets the SimulateAndExecute field and returns the modified Tx instance.
func (t *Tx) WithSimulateAndExecute(v bool) *Tx {
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

// NewTxFromCmd creates and returns Tx from the given cobra command's flags.
func NewTxFromCmd(cmd *cobra.Command) (*Tx, error) {
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

	// Return a new Tx instance populated with the retrieved flag values.
	return &Tx{
		ChainID:            chainID,
		FeeGranterAddr:     feeGranterAddr,
		FromName:           fromName,
		Gas:                gas,
		GasAdjustment:      gasAdjustment,
		GasPrices:          gasPrices,
		SimulateAndExecute: simulateAndExecute,
	}, nil
}
