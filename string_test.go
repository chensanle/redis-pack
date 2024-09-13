package redis_pack

import (
	"testing"
)

func TestStringRds_Get(t *testing.T) {
	conn, _ := NewConnectionByPool(testPool)
	conn.String.Set("namssss", 12)
	_, err := conn.String.Get("namssss").Int64()
	if err != nil {
		t.Error(err)
	}
}
