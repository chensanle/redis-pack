package redis_pack

import (
	"github.com/gomodule/redigo/redis"
	"testing"
)

func Test_zSetRds_ZAdd(t *testing.T) {
	NewConnectionByPool(&redis.Pool{})

	err := RedigoConn.ZSet.ZAdd("key", map[interface{}]interface{}{3: 3, 1: 4, 5: 5}).error
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
	_ = NewConnectionWithFile("127.0.0.1:6379", "")

	sruRes := make([]ZSetUnion, 0)
	err := RedigoConn.ZSet.ZRevrange("clz", 0, -1, true).ScanSlice(&sruRes)
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

func TestZset(t *testing.T) {
	t.Log(a(15).(int))
}

func a(arg interface{}) interface{} {
	return arg
}

func Test_zSetRds_zrem(t *testing.T) {
	//z.zrem(tt.args.key, tt.args.fileds)
}
