package types

const (
	// ModuleName defines the module name
	ModuleName = "feeabstraction"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_feeabstraction"

	TWAP_STORE_QUERY            = "/osmosis.twap.v1beta1.Query/ArithmeticTwapToNow"
	POOL_STORE_QUERY            = "/osmosis.gamm.v1beta1.Query/PoolsWithFilter"
	BANK_STORE_QUERY_WITH_PROOF = "store/bank/key"

	// NonNativeFeeCollectorName the module account name for the alt fee collector account address (used for auto-swapping non-base-denom tx fees).
	NonNativeFeeCollectorName = "non_native_fee_collector"
)

// IBC events
const (
	AttributeKeyAck = "acknowledgement"
)

var (
	BaseDenomKey = []byte{0x01}
	FeeICAKey    = []byte{0x02}
	TempFeeKey   = []byte{0x03}

	StoreFeeRate        = []byte{0x11}
	StoreDenomOsmoTrack = []byte{0x12}
	StoreDenomJunoTrack = []byte{0x13}
	StorePool           = []byte{0x14}
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
