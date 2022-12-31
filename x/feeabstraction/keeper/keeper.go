package keeper

import (
	"encoding/binary"
	"fmt"
	"time"

	"github.com/spf13/cast"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/nghuyenthevinh2000/fa-chain/x/feeabstraction/types"
	icqkeeper "github.com/nghuyenthevinh2000/fa-chain/x/interchainquery/keeper"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	ibctransferkeeper "github.com/cosmos/ibc-go/v3/modules/apps/transfer/keeper"
)

type (
	Keeper struct {
		cdc            codec.BinaryCodec
		storeKey       sdk.StoreKey
		memKey         sdk.StoreKey
		paramstore     paramtypes.Subspace
		icqKeeper      icqkeeper.Keeper
		transferKeeper ibctransferkeeper.Keeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey sdk.StoreKey,
	ps paramtypes.Subspace,
	icqKeeper icqkeeper.Keeper,
	transferKeeper ibctransferkeeper.Keeper,
) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{
		cdc:            cdc,
		storeKey:       storeKey,
		memKey:         memKey,
		paramstore:     ps,
		icqKeeper:      icqKeeper,
		transferKeeper: transferKeeper,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// fee of non - native compared to ujuno
// denom here is present on juno
func (k Keeper) SetFeeRate(ctx sdk.Context, denomJuno string, feeRate sdk.Dec) error {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.StoreFeeRate)
	data, err := feeRate.Marshal()
	if err != nil {
		return err
	}
	store.Set([]byte(denomJuno), data)

	return nil
}

func (k Keeper) GetFeeRate(ctx sdk.Context, denomJuno string) (sdk.Dec, error) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.StoreFeeRate)
	feeRate := sdk.Dec{}
	if err := feeRate.Unmarshal(store.Get([]byte(denomJuno))); err != nil {
		return sdk.Dec{}, err
	}

	return feeRate, nil
}

// record for coins on osmosis to juno
func (k Keeper) SetDenomTrack(ctx sdk.Context, denomOsmo, denomJuno string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.StoreDenomTrack)
	store.Set([]byte(denomOsmo), []byte(denomJuno))
}

func (k Keeper) HasDenomTrack(ctx sdk.Context, denomOsmo string) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.StoreDenomTrack)
	return store.Has([]byte(denomOsmo))
}

func (k Keeper) GetDenomTrack(ctx sdk.Context, denomOsmo string) string {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.StoreDenomTrack)
	denomJuno := store.Get([]byte(denomOsmo))
	return string(denomJuno)
}

func (k Keeper) IterateDenomTrack(ctx sdk.Context, f func(denomOsmo string, denomJuno string) bool) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.StoreDenomTrack)

	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		// key is appended with store prefix, this will remove initial prefix of store and return true value of denomOsmo
		if f(string(iterator.Key()[1:]), string(iterator.Value())) {
			break
		}
	}
}

// record for pool on Osmosis
func (k Keeper) SetPool(ctx sdk.Context, denomOsmo string, poolId uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.StorePool)
	data := make([]byte, 8)
	binary.LittleEndian.PutUint64(data, uint64(poolId))
	store.Set([]byte(denomOsmo), data)
}

func (k Keeper) HasPool(ctx sdk.Context, denomOsmo string) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.StorePool)
	return store.Has([]byte(denomOsmo))
}

func (k Keeper) IteratePool(ctx sdk.Context, f func(denomOsmo string, poolId uint64) bool) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.StorePool)

	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		// key is appended with store prefix, this will remove initial prefix of store and return true value of denomOsmo
		if f(string(iterator.Key()[1:]), uint64(binary.LittleEndian.Uint64(iterator.Value()))) {
			break
		}
	}
}

func (k Keeper) GetTtl(ctx sdk.Context) (uint64, error) {
	currentTime := ctx.BlockTime()

	// add 5 more mins to current time
	return cast.ToUint64E(currentTime.Add(time.Minute * 5).UnixNano())
}
