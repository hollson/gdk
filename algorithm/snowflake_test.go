package algorithm

import "testing"

func TestSetMachineId(t *testing.T) {
	type args struct {
		mid int64
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetMachineId(tt.args.mid)
		})
	}
}

func TestGetSnowflakeId(t *testing.T) {
	tests := []struct {
		name string
		want int64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetSnowflakeId(); got != tt.want {
				t.Errorf("GetSnowflakeId() = %v, want %v", got, tt.want)
			}
		})
	}
}
