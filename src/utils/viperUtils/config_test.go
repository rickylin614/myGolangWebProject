package viperUtils

import (
	"testing"
)

func TestGetLogPath(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{"test", "/logs/server.log"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetLogPath(); got != tt.want {
				t.Errorf("GetLogPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetCommonParams(t *testing.T) {
	tests := []struct {
		name  string
		param string
		want  string
	}{
		{"1", "logPath", "/logs/server.log"},
		{"2", "redisPath", "127.0.0.1:6379"},
		{"3", "sqlPath", "root:qwe123@tcp(127.0.0.1:3306)/bento?charset=utf8&loc=Local&parseTime=true"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetCommonParams(tt.param); got != tt.want {
				t.Errorf("GetCommonParams() = %v, want %v", got, tt.want)
			}
		})
	}
}
