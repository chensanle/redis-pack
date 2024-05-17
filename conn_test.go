package redis_pack

import "testing"

func TestNewConnectionWithFile(t *testing.T) {
	type args struct {
		addr     string
		password string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := NewConnectionWithFile(tt.args.addr, tt.args.password); (err != nil) != tt.wantErr {
				t.Errorf("NewConnectionWithFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
