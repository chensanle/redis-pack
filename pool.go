package redis_pack

import (
	"log"
	"time"

	"github.com/gomodule/redigo/redis"
)

// var pool *redis.Pool
func initPool(addr, password string) RedigoPack {
	pack := RedigoPack{}
	pool := &redis.Pool{
		MaxIdle:     10,
		MaxActive:   200,
		IdleTimeout: 180 * time.Second,
		Wait:        true,
		Dial: func() (redis.Conn, error) {
			return setDialog(addr, password)
		},
	}

	pack.String.pool = pool
	pack.List.pool = pool
	pack.Hash.pool = pool
	pack.Key.pool = pool
	pack.Set.pool = pool
	pack.ZSet.pool = pool
	pack.Bit.pool = pool
	pack.Db.pool = pool

	return pack
}

func initPoolByOld(pool *redis.Pool) RedigoPack {
	pack := RedigoPack{}

	pack.String.pool = pool
	pack.List.pool = pool
	pack.Hash.pool = pool
	pack.Key.pool = pool
	pack.Set.pool = pool
	pack.ZSet.pool = pool
	pack.Bit.pool = pool
	pack.Db.pool = pool

	return pack
}

func setDialog(addr, password string) (redis.Conn, error) {
	conn, err := redis.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}
	if len(password) != 0 {
		if _, err := conn.Do("AUTH", password); err != nil {
			_ = conn.Close()
			return nil, err
		}
	}

	r, err := redis.String(conn.Do("PING"))
	if err != nil || r != "PONG" {
		log.Fatalf("failed to connect redis: %v", addr)
	}

	return conn, nil
}
