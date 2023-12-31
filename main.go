package main

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/crypto/secp256k1"
)

func main() {
	privKey := secp256k1.GenPrivKey()
	pubKey := privKey.PubKey()

	fmt.Printf("pubKey.Address().String(): %v\n", pubKey.Address().String())

	address := sdk.AccAddress(pubKey.Address())

	

	

	fmt.Println("Cosmos SDK account address:", address.String())
}

