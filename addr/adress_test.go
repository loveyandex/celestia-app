package addr_test

import (
	"encoding/hex"
	"fmt"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"

	tmrand "github.com/tendermint/tendermint/libs/rand"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
 

)

// RandomAccounts returns a list of n random accounts
func RandomAccounts(n int) (accounts []string) {
	for i := 0; i < n; i++ {
		account := tmrand.Str(20)
		accounts = append(accounts, account)
	}
	return accounts
}

func TestRandomAccounts(t *testing.T) {

	aa := estAddress()

	fmt.Printf("aa: %v\n", aa.String())

	aa.MarshalJSON()

	got := RandomAccounts(2)
	assert.Len(t, got, 2)
	assert.NotEqual(t, got[0], got[1])
}

var (
	TestAccName = "test-account"
	TestAccAddr = "celestia1g39egf59z8tud3lcyjg5a83m20df4kccx32qkp"
)

func estAddress() sdk.AccAddress {
	bz, err := sdk.GetFromBech32(TestAccAddr, "celestia")
	if err != nil {
		panic(err)
	}

	return sdk.AccAddress(bz)

}

func TestCOsmosToCelestia(t *testing.T) {
	// Set the desired prefix for the account address
	newPrefix := "celestia"

	// Modify the bech32PrefixAccAddr variable in the config package
	sdk.GetConfig().SetBech32PrefixForAccount(newPrefix, newPrefix+"pub")

	privKey := secp256k1.GenPrivKey()
	pubKey := privKey.PubKey()

	fmt.Printf("pubKey.Address().String(): %v\n", pubKey.Address().String())

	address := sdk.AccAddress(pubKey.Address())

	fmt.Println("Custom account address:", address.String())
}

func TestCOsmosToCelestiaMnemonia(t *testing.T) {
	// Set the desired prefix for the account address
	newPrefix := "celestia"
	// // Modify the bech32PrefixAccAddr variable in the config package
	sdk.GetConfig().SetBech32PrefixForAccount(newPrefix, newPrefix+"pub")

	privKey := secp256k1.GenPrivKey()
	fmt.Printf("privKey.String(): %v\n", privKey.String())
	fmt.Printf("privKey.Key: %v\n", string(privKey.Key))
	// Convert the private key to a string  

	privateKeyString := hex.EncodeToString(privKey.Bytes())
	fmt.Printf("privateKeyString: %v\n", privateKeyString)
	privateKeyBytes, err := hex.DecodeString(privateKeyString)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}


	for i, v := range privateKeyBytes {
		if  v==privKey.Bytes()[i] {
			println("ok")
			
		}else {
			// fmt.Println("nokey")
		}
		
	}

	pubKey := privKey.PubKey()

	fmt.Printf("pubKey.Address().String(): %v\n", pubKey.Address().String())

	address := sdk.AccAddress(pubKey.Address())

	fmt.Println("Custom account address:", address.String())
}

func TestJJ(t *testing.T) {

	// Replace privateKeyBytes with your existing private key in bytes //celestia1fnxgqll56hkcqajx905s2lhjq5qh2l0uauc4al
	privateKeyBytes, err := hex.DecodeString("8e2c88682b364f8db1192a6c956120419abcd8b23d15e9fd981ab4352084b2b7")
	if err != nil {
		
		
	}

	// Set the desired prefix for the account address
	// Modify the bech32PrefixAccAddr variable in the config package
	newPrefix := "celestia"
	sdk.GetConfig().SetBech32PrefixForAccount(newPrefix, newPrefix+"pub")

	privKey := secp256k1.PrivKey{Key: privateKeyBytes}
	pubKey := privKey.PubKey()
	address := sdk.AccAddress(pubKey.Address())

	fmt.Println("Cosmos SDK account address:", address.String())

	assert.Equal(t, "celestia1j555hdgulrt9jwq8752hqhnkunlpdszl2wyr3f", address.String())

}



// func TestMnemonic(t *testing.T) {
// 	mnemonic := "your mnemonic words go here"

// 	seed, err := types.Mnem(mnemonic)
// 	if err != nil {
// 		fmt.Println("Error generating seed from mnemonic:", err)
// 		return
// 	}

// 	privKey, err := secp256k1.DerivePrivateKey(seed)
// 	if err != nil {
// 		fmt.Println("Error deriving private key:", err)
// 		return
// 	}

// 	pubKey := privKey.PubKey()

// 	accAddress := types.AccAddress(pubKey.Address())

// 	fmt.Println("Generated AccAddress:", accAddress.String())
	
// }