# fa-chain
fee abstraction chain


Related work

* [F1 Fee distribution](https://github.com/cosmos/cosmos-sdk/raw/main/docs/spec/fee_distribution/f1_fee_distr.pdf)


## Original Design from Osmosis Grants

Fee Abstraction

Description 

This initiative will fund a team(s) to develop infrastructure that leverages Osmosis to allow users on other chains to pass gas fees in any token. The architecture will include 3 primary components
Allow chains and applications to use IBC queries to pull Osmosis TWAPs. The TWAPs serve as oracles to set the relative values of different fee tokens 
Allow chains to and swap their accumulated fees using Osmosis 
SDK Version: easiest to integrate but requires each chain to opt in to accept new fee tokens
CosmWasm Version: applications can integrate without chain upgrade
JavaScript library to facilitate front-end integrations 

Purpose

After outposts are launched, it’s possible to allow other chains to pay fees in any IBC token. This  infrastructure stands to provide a significant boost to UX on other Cosmos chains as well as driving trading volume to Osmosis. If widely adopted, it will further cement Osmosis’s positioning within the ecosystem. 

Outcome

Milestone 1
Deliver Osmosis TWAP oracle infrastructure. Leverage Interchain Queries to pull Osmosis TWAPs and calculate conversion from native gas token to other tokens. 

Milestone 2 (SDK Version)
Deliver modified gas fee SDK module that allows non-native token to be used for transaction fees. This module would perform the swapping of accumulated fees using Osmosis on a periodic basis (e.g. once per day).   


