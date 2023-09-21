package dex

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/duality-labs/duality/testutil/sample"
	dexsimulation "github.com/duality-labs/duality/x/dex/simulation"
	"github.com/duality-labs/duality/x/dex/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = dexsimulation.FindAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgDeposit = "op_weight_msg_deposit"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeposit int = 100

	opWeightMsgWithdrawal = "op_weight_msg_withdrawal"
	// TODO: Determine the simulation weight value
	defaultWeightMsgWithdrawal int = 100

	opWeightMsgPlaceLimitOrder = "op_weight_msg_place_limit_order"
	// TODO: Determine the simulation weight value
	defaultWeightMsgPlaceLimitOrder int = 100

	opWeightMsgWithdrawFilledLimitOrder = "op_weight_msg_withdrawal_withdrawn_limit_order"
	// TODO: Determine the simulation weight value
	defaultWeightMsgWithdrawFilledLimitOrder int = 100

	opWeightMsgCancelLimitOrder = "op_weight_msg_cancel_limit_order"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCancelLimitOrder int = 100

	opWeightMsgMultiHopSwap = "op_weight_msg_multi_hop_swap"
	// TODO: Determine the simulation weight value
	defaultWeightMsgMultiHopSwap int = 100

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

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(
	simState module.SimulationState,
) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgDeposit int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeposit, &weightMsgDeposit, nil,
		func(_ *rand.Rand) {
			weightMsgDeposit = defaultWeightMsgDeposit
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeposit,
		dexsimulation.SimulateMsgDeposit(am.bankKeeper, am.keeper),
	))

	var weightMsgWithdrawal int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgWithdrawal, &weightMsgWithdrawal, nil,
		func(_ *rand.Rand) {
			weightMsgWithdrawal = defaultWeightMsgWithdrawal
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgWithdrawal,
		dexsimulation.SimulateMsgWithdrawal(am.bankKeeper, am.keeper),
	))

	var weightMsgPlaceLimitOrder int
	simState.AppParams.GetOrGenerate(
		simState.Cdc,
		opWeightMsgPlaceLimitOrder,
		&weightMsgPlaceLimitOrder,
		nil,
		func(_ *rand.Rand) {
			weightMsgPlaceLimitOrder = defaultWeightMsgPlaceLimitOrder
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgPlaceLimitOrder,
		dexsimulation.SimulateMsgPlaceLimitOrder(am.bankKeeper, am.keeper),
	))

	var weightMsgWithdrawFilledLimitOrder int
	simState.AppParams.GetOrGenerate(
		simState.Cdc,
		opWeightMsgWithdrawFilledLimitOrder,
		&weightMsgWithdrawFilledLimitOrder,
		nil,
		func(_ *rand.Rand) {
			weightMsgWithdrawFilledLimitOrder = defaultWeightMsgWithdrawFilledLimitOrder
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgWithdrawFilledLimitOrder,
		dexsimulation.SimulateMsgWithdrawFilledLimitOrder(
			am.bankKeeper,
			am.keeper,
		),
	))

	var weightMsgCancelLimitOrder int
	simState.AppParams.GetOrGenerate(
		simState.Cdc,
		opWeightMsgCancelLimitOrder,
		&weightMsgCancelLimitOrder,
		nil,
		func(_ *rand.Rand) {
			weightMsgCancelLimitOrder = defaultWeightMsgCancelLimitOrder
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCancelLimitOrder,
		dexsimulation.SimulateMsgCancelLimitOrder(am.bankKeeper, am.keeper),
	))

	var weightMsgMultiHopSwap int
	simState.AppParams.GetOrGenerate(
		simState.Cdc,
		opWeightMsgMultiHopSwap,
		&weightMsgMultiHopSwap,
		nil,
		func(_ *rand.Rand) {
			weightMsgMultiHopSwap = defaultWeightMsgMultiHopSwap
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgMultiHopSwap,
		dexsimulation.SimulateMsgMultiHopSwap(am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
