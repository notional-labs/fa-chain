package app

import (
	"github.com/cosmos/cosmos-sdk/std"
	"github.com/notional-labs/fa-chain/app/params"
	"github.com/osmosis-labs/osmosis/v13/x/gamm"
)

var (
	encodingConfig params.EncodingConfig = MakeEncodingConfig()
)

func GetEncodingConfig() params.EncodingConfig {
	return encodingConfig
}

// MakeEncodingConfig creates an EncodingConfig.
func MakeEncodingConfig() params.EncodingConfig {
	encodingConfig := params.MakeEncodingConfig()
	std.RegisterLegacyAminoCodec(encodingConfig.Amino)
	std.RegisterInterfaces(encodingConfig.InterfaceRegistry)
	ModuleBasics.RegisterLegacyAminoCodec(encodingConfig.Amino)
	ModuleBasics.RegisterInterfaces(encodingConfig.InterfaceRegistry)
	// register gamm to unmarshall Pool interface
	gamm.AppModuleBasic{}.RegisterLegacyAminoCodec(encodingConfig.Amino)
	gamm.AppModuleBasic{}.RegisterInterfaces(encodingConfig.InterfaceRegistry)
	return encodingConfig
}
