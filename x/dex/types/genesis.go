package types

import (
	"fmt"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		TickMapList:                           []TickMap{},
		PairMapList:                           []PairMap{},
		TokensList:                            []Tokens{},
		TokenMapList:                          []TokenMap{},
		SharesList:                            []Shares{},
		FeeListList:                           []FeeList{},
		LimitOrderPoolUserShareMapList:        []LimitOrderPoolUserShareMap{},
		LimitOrderPoolUserSharesWithdrawnList: []LimitOrderPoolUserSharesWithdrawn{},
		LimitOrderPoolTotalSharesMapList:      []LimitOrderPoolTotalSharesMap{},
		LimitOrderPoolReserveMapList:          []LimitOrderPoolReserveMap{},
		LimitOrderPoolFillMapList:             []LimitOrderPoolFillMap{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in tickMap
	tickMapIndexMap := make(map[string]struct{})

	for _, elem := range gs.TickMapList {
		index := string(TickMapKey(elem.PairId, elem.TickIndex))
		if _, ok := tickMapIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for tickMap")
		}
		tickMapIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in pairMap
	pairMapIndexMap := make(map[string]struct{})

	for _, elem := range gs.PairMapList {
		index := string(PairMapKey(elem.PairId))
		if _, ok := pairMapIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for pairMap")
		}
		pairMapIndexMap[index] = struct{}{}
	}
	// Check for duplicated ID in tokens
	tokensIdMap := make(map[uint64]bool)
	tokensCount := gs.GetTokensCount()
	for _, elem := range gs.TokensList {
		if _, ok := tokensIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for tokens")
		}
		if elem.Id >= tokensCount {
			return fmt.Errorf("tokens id should be lower or equal than the last id")
		}
		tokensIdMap[elem.Id] = true
	}
	// Check for duplicated index in tokenMap
	tokenMapIndexMap := make(map[string]struct{})

	for _, elem := range gs.TokenMapList {
		index := string(TokenMapKey(elem.Address))
		if _, ok := tokenMapIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for tokenMap")
		}
		tokenMapIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in shares
	sharesIndexMap := make(map[string]struct{})

	for _, elem := range gs.SharesList {
		index := string(SharesKey(elem.Address, elem.PairId, elem.TickIndex, elem.FeeIndex))
		if _, ok := sharesIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for shares")
		}
		sharesIndexMap[index] = struct{}{}
	}
	// Check for duplicated ID in feeList
	feeListIdMap := make(map[uint64]bool)
	feeListCount := gs.GetFeeListCount()
	for _, elem := range gs.FeeListList {
		if _, ok := feeListIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for feeList")
		}
		if elem.Id >= feeListCount {
			return fmt.Errorf("feeList id should be lower or equal than the last id")
		}
		feeListIdMap[elem.Id] = true
	}
	// Check for duplicated index in limitOrderPoolUserShareMap
	limitOrderPoolUserShareMapIndexMap := make(map[string]struct{})

	for _, elem := range gs.LimitOrderPoolUserShareMapList {
		index := string(LimitOrderPoolUserShareMapKey(elem.PairId, elem.TickIndex, elem.Token, elem.Count, elem.Address))
		if _, ok := limitOrderPoolUserShareMapIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for limitOrderPoolUserShareMap")
		}
		limitOrderPoolUserShareMapIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in limitOrderPoolUserSharesWithdrawn
	limitOrderPoolUserSharesWithdrawnIndexMap := make(map[string]struct{})

	for _, elem := range gs.LimitOrderPoolUserSharesWithdrawnList {
		index := string(LimitOrderPoolUserSharesWithdrawnKey(elem.PairId, elem.TickIndex, elem.Token, elem.Count, elem.Address))
		if _, ok := limitOrderPoolUserSharesWithdrawnIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for limitOrderPoolUserSharesWithdrawn")
		}
		limitOrderPoolUserSharesWithdrawnIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in limitOrderPoolTotalSharesMap
	limitOrderPoolTotalSharesMapIndexMap := make(map[string]struct{})

	for _, elem := range gs.LimitOrderPoolTotalSharesMapList {
		index := string(LimitOrderPoolTotalSharesMapKey(elem.PairId, elem.TickIndex, elem.Token, elem.Count))
		if _, ok := limitOrderPoolTotalSharesMapIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for limitOrderPoolTotalSharesMap")
		}
		limitOrderPoolTotalSharesMapIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in limitOrderPoolReserveMap
	limitOrderPoolReserveMapIndexMap := make(map[string]struct{})

	for _, elem := range gs.LimitOrderPoolReserveMapList {
		index := string(LimitOrderPoolReserveMapKey(elem.PairId, elem.TickIndex, elem.Token, elem.Count))
		if _, ok := limitOrderPoolReserveMapIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for limitOrderPoolReserveMap")
		}
		limitOrderPoolReserveMapIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in limitOrderPoolFillMap
	limitOrderPoolFillMapIndexMap := make(map[string]struct{})

	for _, elem := range gs.LimitOrderPoolFillMapList {
		index := string(LimitOrderPoolFillMapKey(elem.PairId, elem.TickIndex, elem.Token, elem.Count))
		if _, ok := limitOrderPoolFillMapIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for limitOrderPoolFillMap")
		}
		limitOrderPoolFillMapIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
