package types

func GetFeeICAAccountOwner(chainId string) (result string) {
	return chainId + "." + "FEE"
}
