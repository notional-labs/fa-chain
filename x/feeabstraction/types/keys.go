package types

const (
	// ModuleName defines the module name
	ModuleName = "fachain"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_fachain"

	TWAP_STORE_QUERY = "/osmosis.twap.v1beta1.Query/ArithmeticTwapToNow"
	POOL_STORE_QUERY = "/osmosis.gamm.v1beta1.Query/PoolsWithFilter"
)

var (
	StoreFeeRate    = []byte{0x11}
	StoreDenomTrack = []byte{0x12}
	StorePool       = []byte{0x13}
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
