package redis_pack

import (
	"github.com/gomodule/redigo/redis"
	"time"
)

var testPool = &redis.Pool{
	MaxIdle:     100,
	MaxActive:   500,
	IdleTimeout: 180 * time.Second,
	Dial:        func() (redis.Conn, error) { return setDialog("127.0.0.1:6379", "") },
}

var testConn, _ = NewConnectionByPool(testPool)

type dbRds struct {
	pool *redis.Pool
}

func (d *dbRds) SelectDb(db int64) *Reply {
	c := d.pool.Get()
	defer c.Close()
	return getReply(c.Do("select", db))
}

func (d *dbRds) Do(commend string, args ...interface{}) *Reply {
	c := d.pool.Get()
	defer c.Close()
	return getReply(c.Do(commend, args...))
}

func (d *dbRds) PipLine(commends []map[string][]any) ([]any, []error) {
	conn := d.pool.Get()
	defer conn.Close()

	for _, cmd := range commends {
		for k, v := range cmd {
			err := conn.Send(k, v...)
			if err != nil {
				return nil, []error{err}
			}
		}
	}
	conn.Flush()

	replies, errs := make([]interface{}, 0), make([]error, 0)
	for i := 0; i < len(commends); i++ {
		tmp, err := conn.Receive()
		replies = append(replies, tmp)
		errs = append(errs, err)
	}
	return replies, errs
}
