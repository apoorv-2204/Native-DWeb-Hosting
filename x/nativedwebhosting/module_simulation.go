package nativedwebhosting

import (
	"math/rand"

	"github.com/apoorv-2204/Native-DWeb-Hosting/testutil/sample"
	nativedwebhostingsimulation "github.com/apoorv-2204/Native-DWeb-Hosting/x/nativedwebhosting/simulation"
	"github.com/apoorv-2204/Native-DWeb-Hosting/x/nativedwebhosting/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = nativedwebhostingsimulation.FindAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
	_ = rand.Rand{}
)

const (
	opWeightMsgHostManifest = "op_weight_msg_host_manifest"
	// TODO: Determine the simulation weight value
	defaultWeightMsgHostManifest int = 100

	opWeightMsgStorage = "op_weight_msg_storage"
	// TODO: Determine the simulation weight value
	defaultWeightMsgStorage int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	nativedwebhostingGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&nativedwebhostingGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// ProposalContents doesn't return any content functions for governance proposals.
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgHostManifest int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgHostManifest, &weightMsgHostManifest, nil,
		func(_ *rand.Rand) {
			weightMsgHostManifest = defaultWeightMsgHostManifest
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgHostManifest,
		nativedwebhostingsimulation.SimulateMsgHostManifest(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgStorage int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgStorage, &weightMsgStorage, nil,
		func(_ *rand.Rand) {
			weightMsgStorage = defaultWeightMsgStorage
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgStorage,
		nativedwebhostingsimulation.SimulateMsgStorage(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgHostManifest,
			defaultWeightMsgHostManifest,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				nativedwebhostingsimulation.SimulateMsgHostManifest(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgStorage,
			defaultWeightMsgStorage,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				nativedwebhostingsimulation.SimulateMsgStorage(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
