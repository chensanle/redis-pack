package redis_pack

import (
	"testing"
)

func TestKeyRds_Key(t *testing.T) {
	conn, _ := NewConnectionByPool(testPool)
	type T struct {
		Key   string
		Value float64
	}
	ts := []T{
		{Key: "key1", Value: 1.1},
		{Key: "key2", Value: 2.1},
		{Key: "key3", Value: 2.2},
	}
	for _, one := range ts {
		err := conn.String.Set(one.Key, one.Value).error
		if err != nil {
			t.Error(err)
		}
	}
	key, err := conn.Key.RandomKey().String()
	if err != nil {
		t.Error(err)
	}
	for index, one := range ts {
		if one.Key == key {
			break
		}
		if index == len(ts) {
			t.Error("randomkey 随机key出错")
		}
	}

	for _, one := range ts {
		err := conn.Key.Rename(one.Key, one.Key+"1").error
		if err != nil {
			t.Error(err)
		}
	}

	for _, one := range ts {
		exist, err := conn.Key.Exists(one.Key).Bool()
		if err != nil {
			t.Error(err)
		}
		if exist {
			t.Error("rename 出错")
		}
		exist, err = conn.Key.Exists(one.Key + "1").Bool()
		if err != nil {
			t.Error(err)
		}
		if !exist {
			t.Error("rename 出错")
		}
	}

	for _, one := range ts {
		err = conn.Key.Expire(one.Key+"1", 1000).error
		if err != nil {
			t.Error(err)
		}
		ttl, err := conn.Key.TTL(one.Key + "1").Int64()
		if err != nil {
			t.Error(err)
		}
		if ttl == -1 || ttl == -2 {
			t.Error("过期key 失败")
		}
	}

	for _, one := range ts {
		err = conn.Key.Move(one.Key+"1", 2).error
		if err != nil {
			t.Error(err)
		}
	}

	err = conn.Db.SelectDb(2).error
	if err != nil {
		t.Error(err)
	}
	for _, one := range ts {
		exist, err := conn.Key.Exists(one.Key + "1").Bool()
		if err != nil {
			t.Error(err)
		}
		if !exist {
			t.Error("move 出错")
		}
	}

}

func TestKeyRds_Del(t *testing.T) {
	conn, _ := NewConnectionByPool(testPool)
	key := "1"
	err := conn.String.Set(key, 2, 10).error
	if err != nil {
		t.Error(err)
	}
	exist, err := conn.Key.Exists(key).Bool()
	if err != nil {
		t.Error(err)
	}
	if !exist {
		t.Error(err)
	}

	err = conn.Key.Del(key).error
	if err != nil {
		t.Error(err)
	}
	exist, err = conn.Key.Exists(key).Bool()
	if err != nil {
		t.Error(err)
	}
	if exist {
		t.Error(err)
	}

}
