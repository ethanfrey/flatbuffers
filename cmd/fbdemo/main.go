package main

import (
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
	fmt.Printf("Encoded data (%d bytes):\n%X\n", len(raw), raw)

	acct := users.GetRootAsAccount(raw, 0)
	PrintAccount(acct)
}
