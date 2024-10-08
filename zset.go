package redis_pack

import (
	"github.com/gomodule/redigo/redis"
)

type zSetRds struct {
	pool *redis.Pool
}

// map[score]member  添加元素
func (z *zSetRds) ZAdd(key string, mp map[interface{}]interface{}) *Reply {
	c := z.pool.Get()
	defer c.Close()
	return getReply(c.Do("zadd", redis.Args{}.Add(key).AddFlat(mp)...))
}

// 增加元素权重
func (z *zSetRds) ZUncrBy(key string, increment, member interface{}) *Reply {
	c := z.pool.Get()
	defer c.Close()
	return getReply(c.Do("zuncrby", key, increment, member))
}

// 增加元素权重
func (z *zSetRds) ZCard(key string) *Reply {
	c := z.pool.Get()
	defer c.Close()
	return getReply(c.Do("zcard", key))
}

// 返回指定元素的排名
func (z *zSetRds) ZEank(key string, member interface{}) *Reply {
	c := z.pool.Get()
	defer c.Close()
	return getReply(c.Do("zrank", key, member))
}

// 返回指定元素的权重
func (z *zSetRds) ZScore(key string, member interface{}) *Reply {
	c := z.pool.Get()
	defer c.Close()
	return getReply(c.Do("zscore", key, member))
}

// 返回集合两个权重间的元素数
func (z *zSetRds) ZCount(key string, min, max interface{}) *Reply {
	c := z.pool.Get()
	defer c.Close()
	return getReply(c.Do("zcount", key, min, max))
}

// 返回指定区间内的元素
func (z *zSetRds) ZRange(key string, start, stop interface{}, withScore ...bool) *Reply {
	c := z.pool.Get()
	defer c.Close()
	if len(withScore) > 0 && withScore[0] {
		return getReply(c.Do("zrange", key, start, stop, withScore))
	}
	return getReply(c.Do("zrange", key, start, stop))
}

func (z *zSetRds) Zrem(key string, fileds interface{}) *Reply {
	c := z.pool.Get()
	defer c.Close()

	return getReply(c.Do("zrem", redis.Args{}.Add(key).AddFlat(fileds)...))
}

func (z *zSetRds) ZRangeByScore(key string, start, stop interface{}, args ...interface{}) *Reply {
	// ZRANGEBYSCORE key min max [WITHSCORES] [LIMIT offset count]
	c := z.pool.Get()
	defer c.Close()

	if len(args) == 1 {
		return getReply(c.Do("zrangebyscore", key, start, stop, args[0]))
	}
	if len(args) == 2 {
		return getReply(c.Do("zrangebyscore", key, start, stop, args[0], args[1]))
	}
	return getReply(c.Do("zrangebyscore", key, start, stop))
}

// 倒序返回指定区间内的元素
func (z *zSetRds) ZRevrange(key string, start, stop interface{}, withScore ...bool) *Reply {
	c := z.pool.Get()
	defer c.Close()
	if len(withScore) > 0 && withScore[0] {
		return getReply(c.Do("zrevrange", key, start, stop, "withscores"))
	}
	return getReply(c.Do("zrevrange", key, start, stop))
}
