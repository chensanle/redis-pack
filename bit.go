package redis_pack

import (
	"github.com/gomodule/redigo/redis"
)

type bitRds struct {
	pool *redis.Pool
}

func (b *bitRds) SetBit(key string, offset, value int64) *Reply {
	c := b.pool.Get()
	defer c.Close()
	return getReply(c.Do("setbit", key, offset, value))
}

func (b *bitRds) GetBit(key string, offset int64) *Reply {
	c := b.pool.Get()
	defer c.Close()
	return getReply(c.Do("getbit", key, offset))
}

func (b *bitRds) BitCount(key string, interval ...int64) *Reply {
	c := b.pool.Get()
	defer c.Close()
	if len(interval) == 2 {
		return getReply(c.Do("bitcount", key, interval[0], interval[1]))
	}
	return getReply(c.Do("bitcount", key))
}

// opt 包含 and、or、xor、not
func (b *bitRds) BitTop(opt, destKey string, keys ...string) *Reply {
	c := b.pool.Get()
	defer c.Close()
	return getReply(c.Do("bitop", opt, redis.Args{}.Add(keys).AddFlat(keys)))
}
