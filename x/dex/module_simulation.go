package dex

import (
	"math/rand"

	"github.com/NicholasDotSol/duality/testutil/sample"
	dexsimulation "github.com/NicholasDotSol/duality/x/dex/simulation"
	"github.com/NicholasDotSol/duality/x/dex/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = dexsimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgAddLiquidity = "op_weight_msg_add_liquidity"
	// TODO: Determine the simulation weight value
	defaultWeightMsgAddLiquidity int = 100

	opWeightMsgRemoveLiquidity = "op_weight_msg_remove_liquidity"
	// TODO: Determine the simulation weight value
	defaultWeightMsgRemoveLiquidity int = 100

	opWeightMsgCreatePair = "op_weight_msg_create_pair"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreatePair int = 100

	opWeightMsgSwap = "op_weight_msg_swap"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSwap int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	dexGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&dexGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {

	return []simtypes.ParamChange{}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgAddLiquidity int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgAddLiquidity, &weightMsgAddLiquidity, nil,
		func(_ *rand.Rand) {
			weightMsgAddLiquidity = defaultWeightMsgAddLiquidity
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgAddLiquidity,
		dexsimulation.SimulateMsgAddLiquidity(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgRemoveLiquidity int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgRemoveLiquidity, &weightMsgRemoveLiquidity, nil,
		func(_ *rand.Rand) {
			weightMsgRemoveLiquidity = defaultWeightMsgRemoveLiquidity
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRemoveLiquidity,
		dexsimulation.SimulateMsgRemoveLiquidity(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreatePair int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreatePair, &weightMsgCreatePair, nil,
		func(_ *rand.Rand) {
			weightMsgCreatePair = defaultWeightMsgCreatePair
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreatePair,
		dexsimulation.SimulateMsgCreatePair(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgSwap int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgSwap, &weightMsgSwap, nil,
		func(_ *rand.Rand) {
			weightMsgSwap = defaultWeightMsgSwap
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSwap,
		dexsimulation.SimulateMsgSwap(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}