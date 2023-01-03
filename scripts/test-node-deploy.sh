#!/bin/bash

KEY="test"
CHAINID="fachain-1"
KEYRING="test"
MONIKER="localtestnet"
KEYALGO="secp256k1"
LOGLEVEL="info"

# retrieve all args
WILL_RECOVER=0
WILL_INSTALL=0
WILL_CONTINUE=0
# $# is to check number of arguments
if [ $# -gt 0 ];
then
    # $@ is for getting list of arguments
    for arg in "$@"; do
        case $arg in
        --recover)
            WILL_RECOVER=1
            shift
            ;;
        --install)
            WILL_INSTALL=1
            shift
            ;;
        --continue)
            WILL_CONTINUE=1
            shift
            ;;
        *)
            printf >&2 "wrong argument somewhere"; exit 1;
            ;;
        esac
    done
fi

# continue running if everything is configured
if [ $WILL_CONTINUE -eq 1 ];
then
    # Start the node (remove the --pruning=nothing flag if historical queries are not needed)
    fachaind start --pruning=nothing --log_level $LOGLEVEL --minimum-gas-prices=0.0001ufac
    exit 1;
fi

# validate dependencies are installed
command -v jq > /dev/null 2>&1 || { echo >&2 "jq not installed. More info: https://stedolan.github.io/jq/download/"; exit 1; }
command -v toml > /dev/null 2>&1 || { echo >&2 "toml not installed. More info: https://github.com/mrijken/toml-cli"; exit 1; }

# install fachaind if not exist
if [ $WILL_INSTALL -eq 0 ];
then 
    command -v fachaind > /dev/null 2>&1 || { echo >&1 "installing fachaind"; make install; }
else
    echo >&1 "installing fachaind"
    rm -rf $HOME/.fachain*
    make install
fi

fachaind config keyring-backend $KEYRING
fachaind config chain-id $CHAINID

# determine if user wants to recorver or create new
if [ $WILL_RECOVER -eq 0 ];
then
    fachaind keys add $KEY --keyring-backend $KEYRING --algo $KEYALGO
else
    fachaind keys add $KEY --keyring-backend $KEYRING --algo $KEYALGO --recover
fi

echo >&1 "\n"

# init chain
fachaind init $MONIKER --chain-id $CHAINID

# Change parameter token denominations to ufa
cat $HOME/.fachain/config/genesis.json | jq '.app_state["staking"]["params"]["bond_denom"]="ufac"' > $HOME/.fachain/config/tmp_genesis.json && mv $HOME/.fachain/config/tmp_genesis.json $HOME/.fachain/config/genesis.json
cat $HOME/.fachain/config/genesis.json | jq '.app_state["crisis"]["constant_fee"]["denom"]="ufac"' > $HOME/.fachain/config/tmp_genesis.json && mv $HOME/.fachain/config/tmp_genesis.json $HOME/.fachain/config/genesis.json
cat $HOME/.fachain/config/genesis.json | jq '.app_state["gov"]["deposit_params"]["min_deposit"][0]["denom"]="ufac"' > $HOME/.fachain/config/tmp_genesis.json && mv $HOME/.fachain/config/tmp_genesis.json $HOME/.fachain/config/genesis.json
cat $HOME/.fachain/config/genesis.json | jq '.app_state["mint"]["params"]["mint_denom"]="ufac"' > $HOME/.fachain/config/tmp_genesis.json && mv $HOME/.fachain/config/tmp_genesis.json $HOME/.fachain/config/genesis.json
cat $HOME/.fachain/config/genesis.json | jq '.app_state["feeabstraction"]["base_denom"]="ufac"' > $HOME/.fachain/config/tmp_genesis.json && mv $HOME/.fachain/config/tmp_genesis.json $HOME/.fachain/config/genesis.json
# Set gas limit in genesis
# cat $HOME/.fachain/config/genesis.json | jq '.consensus_params["block"]["max_gas"]="10000000"' > $HOME/.fachain/config/tmp_genesis.json && mv $HOME/.fachain/config/tmp_genesis.json $HOME/.fachain/config/genesis.json

# enable rest server and swagger
toml set --toml-path $HOME/.fachain/config/app.toml api.swagger true
toml set --toml-path $HOME/.fachain/config/app.toml api.enable true
toml set --toml-path $HOME/.fachain/config/app.toml api.address tcp://0.0.0.0:1310

# Allocate genesis accounts (cosmos formatted addresses)
fachaind add-genesis-account $KEY 1000000000000ufac --keyring-backend $KEYRING

# Sign genesis transaction
fachaind gentx $KEY 1000000ufac --keyring-backend $KEYRING --chain-id $CHAINID

# Collect genesis tx
fachaind collect-gentxs

# Run this to ensure everything worked and that the genesis file is setup correctly
fachaind validate-genesis

# Start the node (remove the --pruning=nothing flag if historical queries are not needed)
fachaind start --pruning=nothing --log_level $LOGLEVEL --minimum-gas-prices=0.0001ufac --p2p.laddr tcp://0.0.0.0:2280 --rpc.laddr tcp://0.0.0.0:2281 --grpc.address 0.0.0.0:2282 --grpc-web.address 0.0.0.0:2283