package viperUtils

import (
	"testing"
)

func TestGetLogPath(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{"test", "/log/bento.log"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetLogPath(); got != tt.want {
				t.Errorf("GetLogPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
