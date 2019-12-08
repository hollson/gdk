package utilx

import (
	"reflect"
	"testing"
	"time"
)

func TestStrToTime(t *testing.T) {
	type args struct {
		timeStr  string
		template TimeTemplate
	}
	tests := []struct {
		name    string
		args    args
		want    time.Time
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StrToTime(tt.args.timeStr, tt.args.template)
			if (err != nil) != tt.wantErr {
				t.Errorf("StrToTime() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StrToTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRetroWeekN(t *testing.T) {
	type args struct {
		n time.Weekday
	}
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RetroWeekN(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RetroWeekN() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDuration(t *testing.T) {
	type args struct {
		hours int
	}
	tests := []struct {
		name string
		args args
		want time.Duration
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Duration(tt.args.hours); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Duration() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTimeOfYMD(t *testing.T) {
	type args struct {
		tm time.Time
	}
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TimeOfYMD(tt.args.tm); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TimeOfYMD() = %v, want %v", got, tt.want)
			}
		})
	}
}
