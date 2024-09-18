package redis_pack

import (
	"testing"
)

func TestBitRds_Bit(t *testing.T) {
	conn, _ := NewConnectionWithFile("127.0.0.1:6379", "")
	type bit struct {
		key    string
		offset int64
		value  int64
	}

	md := []bit{
		{key: "bit_test", offset: 2, value: 1},
		{key: "bit_test", offset: 4, value: 1},
		{key: "bit_test", offset: 25, value: 1},
	}

	for _, bt := range md {
		conn.Bit.SetBit(bt.key, bt.offset, bt.value)
	}

	for _, bt := range md {
		v, err := conn.Bit.GetBit(bt.key, bt.offset).Int64()
		if err != nil {
			t.Error(err)
		}
		if v != bt.value {
			t.Error(err)
		}
	}

	l, err := conn.Bit.BitCount(md[0].key).Int()
	if err != nil {
		t.Error(err)
	}
	if l != len(md) {
		t.Error("bitcout值不正确")
	}
}
