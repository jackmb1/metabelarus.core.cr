package app

import (
	"encoding/json"
	"io"
	"os"

	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	tmos "github.com/tendermint/tendermint/libs/os"
	dbm "github.com/tendermint/tm-db"

	bam "github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/simapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/cosmos/cosmos-sdk/x/staking"
	"github.com/cosmos/cosmos-sdk/x/supply"
	"github.com/markvandal/metabelaruscorecr/x/metabelaruscorecr"
	metabelaruscorecrkeeper "github.com/markvandal/metabelaruscorecr/x/metabelaruscorecr/keeper"
	metabelaruscorecrtypes "github.com/markvandal/metabelaruscorecr/x/metabelaruscorecr/types"
  // this line is used by starport scaffolding # 1
		"github.com/markvandal/metabelaruscorecr/x/mbpassport"
		mbpassportkeeper "github.com/markvandal/metabelaruscorecr/x/mbpassport/keeper"
		mbpassporttypes "github.com/markvandal/metabelaruscorecr/x/mbpassport/types"
		"github.com/markvandal/metabelaruscorecr/x/mbpasstrust"
		mbpasstrustkeeper "github.com/markvandal/metabelaruscorecr/x/mbpasstrust/keeper"
		mbpasstrusttypes "github.com/markvandal/metabelaruscorecr/x/mbpasstrust/types"
		"github.com/markvandal/metabelaruscorecr/x/mbgovperm"
		mbgovpermkeeper "github.com/markvandal/metabelaruscorecr/x/mbgovperm/keeper"
		mbgovpermtypes "github.com/markvandal/metabelaruscorecr/x/mbgovperm/types"
)

const appName = "metabelaruscorecr"

var (
	DefaultCLIHome = os.ExpandEnv("$HOME/.metabelaruscorecrcli")
	DefaultNodeHome = os.ExpandEnv("$HOME/.metabelaruscorecrd")
	ModuleBasics = module.NewBasicManager(
		genutil.AppModuleBasic{},
		auth.AppModuleBasic{},
		bank.AppModuleBasic{},
		staking.AppModuleBasic{},
		params.AppModuleBasic{},
		supply.AppModuleBasic{},
		metabelaruscorecr.AppModuleBasic{},
    // this line is used by starport scaffolding # 2
		mbpassport.AppModuleBasic{},
		mbpasstrust.AppModuleBasic{},
		mbgovperm.AppModuleBasic{},
	)

	maccPerms = map[string][]string{
		auth.FeeCollectorName:     nil,
		// this line is used by starport scaffolding # 2.1
		staking.BondedPoolName:    {supply.Burner, supply.Staking},
		staking.NotBondedPoolName: {supply.Burner, supply.Staking},
	}
)

func MakeCodec() *codec.Codec {
	var cdc = codec.New()

	ModuleBasics.RegisterCodec(cdc)
	sdk.RegisterCodec(cdc)
	codec.RegisterCrypto(cdc)

	return cdc.Seal()
}

type NewApp struct {
	*bam.BaseApp
	cdc *codec.Codec

	invCheckPeriod uint

	keys  map[string]*sdk.KVStoreKey
	tKeys map[string]*sdk.TransientStoreKey

	subspaces map[string]params.Subspace

	accountKeeper  auth.AccountKeeper
	bankKeeper     bank.Keeper
	stakingKeeper  staking.Keeper
	supplyKeeper   supply.Keeper
	paramsKeeper   params.Keeper
	metabelaruscorecrKeeper metabelaruscorecrkeeper.Keeper
  // this line is used by starport scaffolding # 3
		mbpassportKeeper mbpassportkeeper.Keeper
		mbpasstrustKeeper mbpasstrustkeeper.Keeper
		mbgovpermKeeper mbgovpermkeeper.Keeper
	mm *module.Manager

	sm *module.SimulationManager
}

var _ simapp.App = (*NewApp)(nil)

func NewInitApp(
	logger log.Logger, db dbm.DB, traceStore io.Writer, loadLatest bool,
	invCheckPeriod uint, baseAppOptions ...func(*bam.BaseApp),
) *NewApp {
	cdc := MakeCodec()

	bApp := bam.NewBaseApp(appName, logger, db, auth.DefaultTxDecoder(cdc), baseAppOptions...)
	bApp.SetCommitMultiStoreTracer(traceStore)
	bApp.SetAppVersion(version.Version)

	keys := sdk.NewKVStoreKeys(
    bam.MainStoreKey,
    auth.StoreKey,
    staking.StoreKey,
		supply.StoreKey,
    params.StoreKey,
    metabelaruscorecrtypes.StoreKey,
    // this line is used by starport scaffolding # 5
		mbpassporttypes.StoreKey,
		mbpasstrusttypes.StoreKey,
		mbgovpermtypes.StoreKey,
  )

	tKeys := sdk.NewTransientStoreKeys(staking.TStoreKey, params.TStoreKey)

	var app = &NewApp{
		BaseApp:        bApp,
		cdc:            cdc,
		invCheckPeriod: invCheckPeriod,
		keys:           keys,
		tKeys:          tKeys,
		subspaces:      make(map[string]params.Subspace),
	}

	app.paramsKeeper = params.NewKeeper(app.cdc, keys[params.StoreKey], tKeys[params.TStoreKey])
	app.subspaces[auth.ModuleName] = app.paramsKeeper.Subspace(auth.DefaultParamspace)
	app.subspaces[bank.ModuleName] = app.paramsKeeper.Subspace(bank.DefaultParamspace)
	app.subspaces[staking.ModuleName] = app.paramsKeeper.Subspace(staking.DefaultParamspace)
	// this line is used by starport scaffolding # 5.1
		app.subspaces[mbpassporttypes.ModuleName] = app.paramsKeeper.Subspace(mbpassporttypes.DefaultParamspace)
		app.subspaces[mbpasstrusttypes.ModuleName] = app.paramsKeeper.Subspace(mbpasstrusttypes.DefaultParamspace)
		app.subspaces[mbgovpermtypes.ModuleName] = app.paramsKeeper.Subspace(mbgovpermtypes.DefaultParamspace)

	app.accountKeeper = auth.NewAccountKeeper(
		app.cdc,
		keys[auth.StoreKey],
		app.subspaces[auth.ModuleName],
		auth.ProtoBaseAccount,
	)

	app.bankKeeper = bank.NewBaseKeeper(
		app.accountKeeper,
		app.subspaces[bank.ModuleName],
		app.ModuleAccountAddrs(),
	)

	app.supplyKeeper = supply.NewKeeper(
		app.cdc,
		keys[supply.StoreKey],
		app.accountKeeper,
		app.bankKeeper,
		maccPerms,
	)

	stakingKeeper := staking.NewKeeper(
		app.cdc,
		keys[staking.StoreKey],
		app.supplyKeeper,
		app.subspaces[staking.ModuleName],
	)

	// this line is used by starport scaffolding # 5.2
		app.mbpassportKeeper = mbpassportkeeper.NewKeeper(
			app.cdc,
			keys[mbpassporttypes.StoreKey],
			app.subspaces[mbpassporttypes.ModuleName],
		)
		app.mbpasstrustKeeper = mbpasstrustkeeper.NewKeeper(
			app.cdc,
			keys[mbpasstrusttypes.StoreKey],
			app.subspaces[mbpasstrusttypes.ModuleName],
		)
		app.mbgovpermKeeper = mbgovpermkeeper.NewKeeper(
			app.cdc,
			keys[mbgovpermtypes.StoreKey],
			app.subspaces[mbgovpermtypes.ModuleName],
		)

	app.stakingKeeper = *stakingKeeper.SetHooks(
		staking.NewMultiStakingHooks(
			// this line is used by starport scaffolding # 5.3
		),
	)

	app.metabelaruscorecrKeeper = metabelaruscorecrkeeper.NewKeeper(
		app.bankKeeper,
		app.cdc,
		keys[metabelaruscorecrtypes.StoreKey],
	)

  // this line is used by starport scaffolding # 4

	app.mm = module.NewManager(
		genutil.NewAppModule(app.accountKeeper, app.stakingKeeper, app.BaseApp.DeliverTx),
		auth.NewAppModule(app.accountKeeper),
		bank.NewAppModule(app.bankKeeper, app.accountKeeper),
		supply.NewAppModule(app.supplyKeeper, app.accountKeeper),
		metabelaruscorecr.NewAppModule(app.metabelaruscorecrKeeper, app.bankKeeper),
		staking.NewAppModule(app.stakingKeeper, app.accountKeeper, app.supplyKeeper),
    // this line is used by starport scaffolding # 6
		mbpassport.NewAppModule(app.mbpassportKeeper),
		mbpasstrust.NewAppModule(app.mbpasstrustKeeper),
		mbgovperm.NewAppModule(app.mbgovpermKeeper),
	)

	app.mm.SetOrderEndBlockers(
		staking.ModuleName,
		// this line is used by starport scaffolding # 6.1
	)

	app.mm.SetOrderInitGenesis(
		// this line is used by starport scaffolding # 6.2
		staking.ModuleName,
		auth.ModuleName,
		bank.ModuleName,
		metabelaruscorecrtypes.ModuleName,
		supply.ModuleName,
		genutil.ModuleName,
    // this line is used by starport scaffolding # 7
		mbpassporttypes.ModuleName,
		mbpasstrusttypes.ModuleName,
		mbgovpermtypes.ModuleName,
	)

	app.mm.RegisterRoutes(app.Router(), app.QueryRouter())

	app.SetInitChainer(app.InitChainer)
	app.SetBeginBlocker(app.BeginBlocker)
	app.SetEndBlocker(app.EndBlocker)

	app.SetAnteHandler(
		auth.NewAnteHandler(
			app.accountKeeper,
			app.supplyKeeper,
			auth.DefaultSigVerificationGasConsumer,
		),
	)

	app.MountKVStores(keys)
	app.MountTransientStores(tKeys)

	if loadLatest {
		err := app.LoadLatestVersion(app.keys[bam.MainStoreKey])
		if err != nil {
			tmos.Exit(err.Error())
		}
	}

	return app
}

type GenesisState map[string]json.RawMessage

func NewDefaultGenesisState() GenesisState {
	return ModuleBasics.DefaultGenesis()
}

func (app *NewApp) InitChainer(ctx sdk.Context, req abci.RequestInitChain) abci.ResponseInitChain {
	var genesisState simapp.GenesisState

	app.cdc.MustUnmarshalJSON(req.AppStateBytes, &genesisState)

	return app.mm.InitGenesis(ctx, genesisState)
}

func (app *NewApp) BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock) abci.ResponseBeginBlock {
	return app.mm.BeginBlock(ctx, req)
}

func (app *NewApp) EndBlocker(ctx sdk.Context, req abci.RequestEndBlock) abci.ResponseEndBlock {
	return app.mm.EndBlock(ctx, req)
}

func (app *NewApp) LoadHeight(height int64) error {
	return app.LoadVersion(height, app.keys[bam.MainStoreKey])
}

func (app *NewApp) ModuleAccountAddrs() map[string]bool {
	modAccAddrs := make(map[string]bool)
	for acc := range maccPerms {
		modAccAddrs[supply.NewModuleAddress(acc).String()] = true
	}

	return modAccAddrs
}

func (app *NewApp) Codec() *codec.Codec {
	return app.cdc
}

func (app *NewApp) SimulationManager() *module.SimulationManager {
	return app.sm
}

func GetMaccPerms() map[string][]string {
	modAccPerms := make(map[string][]string)
	for k, v := range maccPerms {
		modAccPerms[k] = v
	}
	return modAccPerms
}
