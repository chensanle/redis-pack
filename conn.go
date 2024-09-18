package redis_pack

import (
	"github.com/gomodule/redigo/redis"
)

type RedigoPack struct {
	String stringRds
	List   listRds
	Hash   hashRds
	Key    keyRds
	Set    setRds
	ZSet   zSetRds
	Bit    bitRds
	Db     dbRds
}

//var RedigoConn = new(RedigoPack)

func NewConnectionWithFile(addr, password string) (RedigoPack, error) {
	conn := initPool(addr, password)
	return conn, nil
}

func NewConnectionByPool(pool2 *redis.Pool) (RedigoPack, error) {
	conn := initPoolByOld(pool2)
	return conn, nil
}
