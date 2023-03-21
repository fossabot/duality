package types

import (
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

// Parameter store keys.
var (
	_ paramtypes.ParamSet = &Params{}
)

// ParamTable for lockup module.
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

func NewParams() Params {
	return Params{}
}

// DefaultParams returns default lockup module parameters.
func DefaultParams() Params {
	return Params{}
}

// validate params.
func (p Params) Validate() error {
	return nil
}

// Implements params.ParamSet.
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{}
}
