// automatically generated by the FlatBuffers compiler, do not modify

package users

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Account struct {
	_tab flatbuffers.Table
}

func GetRootAsAccount(buf []byte, offset flatbuffers.UOffsetT) *Account {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Account{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *Account) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Account) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Account) Pubkey(obj *PubKey) *PubKey {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(PubKey)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *Account) Coins(obj *Coin, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *Account) CoinsLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func AccountStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func AccountAddPubkey(builder *flatbuffers.Builder, pubkey flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(pubkey), 0)
}
func AccountAddCoins(builder *flatbuffers.Builder, coins flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(coins), 0)
}
func AccountStartCoinsVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func AccountEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
