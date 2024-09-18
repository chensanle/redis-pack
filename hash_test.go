package redis_pack

import (
	"testing"
)

func TestHashRds_Hash(t *testing.T) {
	type Hash struct {
		Name string
		Age  int64
	}

	conn, _ := NewConnectionByPool(testPool)

	key := "hash"
	h := &Hash{"newgoo", 24}
	err := conn.Hash.HMSetFromStruct(key, h).error
	if err != nil {
		t.Error(err)
	}

	h2 := new(Hash)
	err = conn.Hash.HGetAll(key).ScanStruct(h2)
	if err != nil {
		t.Error(err)
	}
	if h2.Name != h.Name || h2.Age != h.Age {
		t.Error("收到值不正确")
	}
}

func TestHashRds_HGetAndSet(t *testing.T) {
	type HSet struct {
		Key   string
		Filed string
		Value int64
	}

	conn, _ := NewConnectionByPool(testPool)

	key := "HashSet"
	hs := []HSet{
		{Key: key, Filed: "filed1", Value: 1},
		{Key: key, Filed: "filed2", Value: 2},
		{Key: key, Filed: "filed3", Value: 3},
	}

	//测试对hash对象操作
	for _, hash := range hs {
		if err := conn.Hash.HSet(hash.Key, hash.Filed, hash.Value).error; err != nil {
			t.Error(err)
		}
	}

	for _, hash := range hs {
		v, err := conn.Hash.HGet(hash.Key, hash.Filed).Int64()
		if err != nil {
			t.Error(err)
		}
		if v != hash.Value {
			t.Error("获取到的值不正确")
		}
	}

	// 测试对字段操作
	value, err := conn.Hash.HMGet(key, []string{"filed1", "filed2"}).Int64s()
	if err != nil {
		t.Error(err)
	}

	if value[0] != 1 && value[1] != 2 {
		t.Error(err)
	}

	exist, err := conn.Hash.HExists(key, "filed1").Bool()
	if err != nil {
		t.Error(err)
	}
	if !exist {
		t.Error("判断出错")
	}

	// 获取hash所有字段
	fileds, err := conn.Hash.HKeys(key).Strings()
	if err != nil {
		t.Error(err)
	}
	if len(fileds) != len(hs) {
		t.Error("获取hash字段出错")
	}

	// 获取hash字段数量
	filedNum, err := conn.Hash.HLen(key).Int()
	if err != nil {
		t.Error(err)
	}
	if filedNum != len(hs) {
		t.Error("获取hash字段数量")
	}

	// 获取hash所有字段值
	values, err := conn.Hash.HVals(key).Int64s()
	if err != nil {
		t.Error(err)
	}
	if len(values) != len(hs) {
		t.Error("获取hash所有值失败")
	}

	// 测试删除字段
	err = conn.Hash.HDel(key, []string{"filed1", "filed2"}).error
	if err != nil {
		t.Error(err)
	}
	exist1, err := conn.Hash.HExists(key, "filed1").Bool()
	if err != nil {
		t.Error(err)
	}
	exist2, err := conn.Hash.HExists(key, "filed2").Bool()
	if err != nil {
		t.Error(err)
	}
	if exist1 || exist2 {
		t.Error("hdel,删除失败")
	}
}

func TestHashRds_HIncr(t *testing.T) {
	type HSet struct {
		Key   string
		Filed string
		Value int64
	}

	conn, _ := NewConnectionByPool(testPool)

	key := "HashSet"
	hs := []HSet{
		{Key: key, Filed: "filed1", Value: 1},
		{Key: key, Filed: "filed2", Value: 2},
		{Key: key, Filed: "filed3", Value: 3},
	}

	//测试对hash对象操作
	for _, hash := range hs {
		if err := conn.Hash.HSet(hash.Key, hash.Filed, hash.Value).error; err != nil {
			t.Error(err)
		}
	}

	for _, hash := range hs {
		err := conn.Hash.HIncrBy(hash.Key, hash.Filed, -1).error
		if err != nil {
			t.Error(err)
		}
	}
	for _, hash := range hs {
		v, err := conn.Hash.HGet(hash.Key, hash.Filed).Int64()
		if err != nil {
			t.Error(err)
		}
		if v != hash.Value-1 {
			t.Error("hincr,方法错误")
		}
	}
}

func TestHashRds_HIncrByFloat(t *testing.T) {
	type HSet struct {
		Key   string
		Filed string
		Value int64
	}

	conn, _ := NewConnectionByPool(testPool)

	key := "HashSet"
	hs := []HSet{
		{Key: key, Filed: "filed1", Value: 1},
		{Key: key, Filed: "filed2", Value: 2},
		{Key: key, Filed: "filed3", Value: 3},
	}
	//测试对hash对象操作
	for _, hash := range hs {
		if err := conn.Hash.HSet(hash.Key, hash.Filed, hash.Value).error; err != nil {
			t.Error(err)
		}
	}

	for _, hash := range hs {
		err := conn.Hash.HIncrByFloat(hash.Key, hash.Filed, 0.2).error
		if err != nil {
			t.Error(err)
		}
	}

	for _, hash := range hs {
		v, err := conn.Hash.HGet(hash.Key, hash.Filed).float64()
		if err != nil {
			t.Error(err)
		}
		if v != float64(hash.Value)+0.2 {
			t.Error("hincrbyfloat,方法错误")
		}
	}
}

func TestHashRds_HMSetFromMap(t *testing.T) {
	conn, _ := NewConnectionByPool(testPool)
	key := "maphash"
	mp := map[interface{}]interface{}{"filed1": 5, "filed2": 6}
	err := conn.Hash.HMSetFromMap(key, mp).error
	if err != nil {
		t.Error(err)
	}

	v, err := conn.Hash.HGet(key, "filed1").Int64()
	if err != nil {
		t.Error(err)
	}
	if v != 5 {
		t.Error("hmset map 失败")
	}

	v, err = conn.Hash.HGet(key, "filed2").Int64()
	if err != nil {
		t.Error(err)
	}
	if v != 6 {
		t.Error("hmset map 失败")
	}

}
