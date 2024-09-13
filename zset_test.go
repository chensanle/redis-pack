package redis_pack

import (
	"testing"
)

func Test_zSetRds_ZAdd(t *testing.T) {
	conn, _ := NewConnectionByPool(testPool)

	err := conn.ZSet.ZAdd("key", map[interface{}]interface{}{3: 3, 1: 4, 5: 5}).error
	if err != nil {
		t.Error(err)
	}
}

type ZSetUnion struct {
	Member int // 结构体上下顺序不可变
	Score  float64
}

func TestZsetReveng(t *testing.T) {
	// redis-server --port 6379
	// go test -run TestZsetReveng -v
	conn, _ := NewConnectionByPool(testPool)

	sruRes := make([]ZSetUnion, 0)
	err := conn.ZSet.ZRevrange("clz", 0, -1, true).ScanSlice(&sruRes)
	t.Logf("err1: %+v", err)

	//err = redis.ScanSlice(res, &sruRes)
	//t.Logf("err2: %+v", err)
	//
	//for _, val := range res {
	//	t.Logf("%+v", val)
	//}

	for _, val := range sruRes {
		t.Logf("%+v", val)
	}

}
