package options

import (
	"errors"

	cosmossdk "github.com/cosmos/cosmos-sdk/types"
	sdkflags "github.com/sentinel-official/sentinel-go-sdk/flags"
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
		FeeGranterAddr:     sdkflags.DefaultTxFeeGranterAddr,
		FromName:           sdkflags.DefaultTxFromName,
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
func (t *Tx) WithFeeGranterAddr(v cosmossdk.AccAddress) *Tx {
	t.FeeGranterAddr = v.String()
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
func (t *Tx) WithGasPrices(v cosmossdk.DecCoins) *Tx {
	t.GasPrices = v.String()
	return t
}

// WithSimulateAndExecute sets the SimulateAndExecute field and returns the modified Tx instance.
func (t *Tx) WithSimulateAndExecute(v bool) *Tx {
	t.SimulateAndExecute = v
	return t
}

// GetChainID returns the ChainID field.
func (t *Tx) GetChainID() string {
	return t.ChainID
}

// GetFeeGranterAddr returns the FeeGranterAddr field.
func (t *Tx) GetFeeGranterAddr() cosmossdk.AccAddress {
	if t.FeeGranterAddr == "" {
		return nil
	}
	addr, err := cosmossdk.AccAddressFromBech32(t.FeeGranterAddr)
	if err != nil {
		panic(err)
	}

	return addr
}

// GetFromName returns the FromName field.
func (t *Tx) GetFromName() string {
	return t.FromName
}

// GetGas returns the Gas field.
func (t *Tx) GetGas() uint64 {
	return t.Gas
}

// GetGasAdjustment returns the GasAdjustment field.
func (t *Tx) GetGasAdjustment() float64 {
	return t.GasAdjustment
}

// GetGasPrices returns the GasPrices field.
func (t *Tx) GetGasPrices() cosmossdk.DecCoins {
	decCoins, err := cosmossdk.ParseDecCoins(t.GasPrices)
	if err != nil {
		panic(err)
	}

	return decCoins
}

// GetSimulateAndExecute returns the SimulateAndExecute field.
func (t *Tx) GetSimulateAndExecute() bool {
	return t.SimulateAndExecute
}

// ValidateTxChainID validates the ChainID field.
func ValidateTxChainID(chainID string) error {
	if chainID == "" {
		return errors.New("chain_id must not be empty")
	}

	return nil
}

// ValidateTxFeeGranterAddr validates the FeeGranterAddr field.
func ValidateTxFeeGranterAddr(addr string) error {
	if addr == "" {
		return nil
	}
	if _, err := cosmossdk.AccAddressFromBech32(addr); err != nil {
		return errors.New("fee_granter_addr must be a valid address")
	}

	return nil
}

// ValidateTxFromName validates the FromName field.
func ValidateTxFromName(fromName string) error {
	if fromName == "" {
		return errors.New("from_name must not be empty")
	}

	return nil
}

// ValidateTxGas validates the Gas field.
func ValidateTxGas(gas uint64) error {
	if gas == 0 {
		return errors.New("gas must be greater than zero")
	}

	return nil
}

// ValidateTxGasAdjustment validates the GasAdjustment field.
func ValidateTxGasAdjustment(adjustment float64) error {
	if adjustment <= 0 {
		return errors.New("gas_adjustment must be greater than zero")
	}

	return nil
}

// ValidateTxGasPrices validates the GasPrices field.
func ValidateTxGasPrices(gasPrices string) error {
	if _, err := cosmossdk.ParseDecCoins(gasPrices); err != nil {
		return errors.New("gas_prices must be a valid decimal coins format")
	}

	return nil
}

// Validate validates all the fields of the Tx struct.
func (t *Tx) Validate() error {
	if err := ValidateTxChainID(t.ChainID); err != nil {
		return err
	}
	if err := ValidateTxFeeGranterAddr(t.FeeGranterAddr); err != nil {
		return err
	}
	if err := ValidateTxFromName(t.FromName); err != nil {
		return err
	}
	if err := ValidateTxGas(t.Gas); err != nil {
		return err
	}
	if err := ValidateTxGasAdjustment(t.GasAdjustment); err != nil {
		return err
	}
	if err := ValidateTxGasPrices(t.GasPrices); err != nil {
		return err
	}

	return nil
}

// NewTxFromCmd creates and returns Tx from the given cobra command's flags.
func NewTxFromCmd(cmd *cobra.Command) (*Tx, error) {
	// Retrieve the chain ID flag value from the command.
	chainID, err := sdkflags.GetTxChainID(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieve the fee granter address flag value from the command.
	feeGranterAddr, err := sdkflags.GetTxFeeGranterAddr(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieve the from name flag value from the command.
	fromName, err := sdkflags.GetTxFromName(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieve the gas flag value from the command.
	gas, err := sdkflags.GetTxGas(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieve the gas adjustment flag value from the command.
	gasAdjustment, err := sdkflags.GetTxGasAdjustment(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieve the gas prices flag value from the command.
	gasPrices, err := sdkflags.GetTxGasPrices(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieve the simulate and execute flag value from the command.
	simulateAndExecute, err := sdkflags.GetTxSimulateAndExecute(cmd)
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
