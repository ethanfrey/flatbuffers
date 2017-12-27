package users

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

// CreatePubKey encodes a new pubkey
func CreatePubKey(builder *flatbuffers.Builder, typ int, data []byte) flatbuffers.UOffsetT {
	key := builder.CreateByteString(data)
	PubKeyStart(builder)
	PubKeyAddType(builder, int8(typ))
	PubKeyAddKey(builder, key)
	return PubKeyEnd(builder)
}

// CreateCoin encodes a new coin
func CreateCoin(builder *flatbuffers.Builder, denom string, amount int64) flatbuffers.UOffsetT {
	key := builder.CreateString(denom)
	CoinStart(builder)
	CoinAddDenom(builder, key)
	CoinAddAmount(builder, amount)
	return CoinEnd(builder)
}

// CreateSimpleAccount starts an account with a pub key and
// one type of coin
func CreateSimpleAccount(builder *flatbuffers.Builder, denom string,
	amount int64, typ int, data []byte) flatbuffers.UOffsetT {

	pk := CreatePubKey(builder, typ, data)
	coin := CreateCoin(builder, denom, amount)

	// TODO: generalize with more coins
	AccountStartCoinsVector(builder, 1)
	builder.PrependUOffsetT(coin)
	coins := builder.EndVector(1)

	AccountStart(builder)
	AccountAddPubkey(builder, pk)
	AccountAddCoins(builder, coins)
	return AccountEnd(builder)
}
