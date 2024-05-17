package redis_pack

type dbRds struct {
}

func (d *dbRds) SelectDb(db int64) *Reply {
	c := pool.Get()
	defer c.Close()
	return getReply(c.Do("select", db))
}

func (d *dbRds) Do(commend string, args ...interface{}) *Reply {
	c := pool.Get()
	defer c.Close()
	return getReply(c.Do(commend, args...))
}

func (d *dbRds) PipLine(commends []map[string][]any) ([]any, []error) {
	conn := pool.Get()
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
