package keeper

import (
	transfertypes "github.com/cosmos/ibc-go/v3/modules/apps/transfer/types"
)

// params changeable through gov (future milestones)
const (
	host_zone_chain_id      = "test-osmo"
	juno_osmo_connection_id = "connection-0"
	osmo_juno_connection_id = "connection-0"
	juno_osmo_channel_id    = "channel-0"
	osmo_juno_channel_id    = "channel-0"
)

func GetIBCDenom(channelId, baseDenom string) transfertypes.DenomTrace {
	sourcePrefix := transfertypes.GetDenomPrefix("transfer", channelId)
	prefixedDenom := sourcePrefix + baseDenom

	return transfertypes.ParseDenomTrace(prefixedDenom)
}
