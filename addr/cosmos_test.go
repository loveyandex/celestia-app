package addr_test

import (
	"encoding/hex"
	"fmt"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"

	secp256k1_tend "github.com/tendermint/tendermint/crypto/secp256k1"
)


func TestMaink(t *testing.T){

	GetCosmosAddressFromPrivateKeyHex("cb735fa1305ca77eac4129616c2ca79af8c21b7a2ebff795750782663a29ee9b")

	//cosmos1fvm5hl5ez09588u8mu5uhl5d6feerrwvdjhhya
	//cosmos1e4pn42zv6yd853yrc2sm04nv8hcwev7c7j6yd2
}

func GetCosmosAddressFromPrivateKeyHex(hexPrivateKey string) (string, error) {
	privateKeyBytes, err := hex.DecodeString(hexPrivateKey)
	if err != nil {
		return "", err
	}

	privKey := secp256k1_tend.PrivKey(privateKeyBytes)
	pubKey := privKey.PubKey()

	accAddress := sdk.AccAddress(pubKey.Address().Bytes())
	fmt.Printf("accAddress.String(): %v\n", accAddress.String())


	return accAddress.String(), nil
}

