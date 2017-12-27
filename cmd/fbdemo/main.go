package main

import (
	"bytes"
	"fmt"

	"github.com/ethanfrey/flatbuffers/users"
	flatbuffers "github.com/google/flatbuffers/go"
)

func MakeData() []byte {
	builder := flatbuffers.NewBuilder(512)

	acct := users.CreateSimpleAccount(builder,
		"eth", 123,
		users.KeyTypeEd25519, []byte("my-crypto-identity"))
	builder.Finish(acct)
	return builder.FinishedBytes()
}

func PrintAccount(acct *users.Account) {
	pk := acct.Pubkey(nil)
	kind := users.EnumNamesKeyType[int(pk.Type())]
	fmt.Printf("Account for pubkey %X (%s):\n",
		pk.KeyBytes(), kind)

	coin := new(users.Coin)
	numCoins := acct.CoinsLength()
	for j := 0; j < numCoins; j++ {
		acct.Coins(coin, j)
		fmt.Printf(" %6d %s\n", coin.Amount(), coin.Denom())
	}
}

func main() {
	raw := MakeData()
	orig := make([]byte, len(raw))
	copy(orig, raw)
	fmt.Printf("Encoded data (%d bytes):\n%X\n\n", len(raw), raw)

	acct := users.GetRootAsAccount(raw, 0)
	PrintAccount(acct)

	// let's change the key type
	// note that modifies raw buffer in-place
	acct.Pubkey(nil).MutateType(int8(users.KeyTypeSecp256k1))

	fmt.Printf("\nEncoded data (%d bytes):\n%X\n\n", len(raw), raw)
	if bytes.Equal(orig, raw) {
		fmt.Println("nothing changed.")
	}

	acct2 := users.GetRootAsAccount(raw, 0)
	PrintAccount(acct2)
}
