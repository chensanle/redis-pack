package redis_pack

import (
	"testing"
)

func Test_dbRds_PipLine(t *testing.T) {
	//d := &dbRds{}
	//got, got1 := d.PipLine(nil)
	//t.Log(got, got1)
}

func Test_dbRds_PipLine1(t *testing.T) {
	conn, _ := NewConnectionWithFile("127.0.0.1:6379", "")

	cmds := make([]map[string][]interface{}, 0)
	cmds = append(cmds, map[string][]interface{}{"INCR": {"test_pip_line"}})
	cmds = append(cmds, map[string][]interface{}{"expire": {"test_pip_line", 100}})

	got, got1 := conn.Db.PipLine(cmds)
	for _, val := range got {
		t.Log(val)
	}
	for _, val := range got1 {
		t.Log(val)
	}
}
