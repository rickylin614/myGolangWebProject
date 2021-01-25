package zapLog

import (
	"errors"
	"testing"
	"time"

	"go.uber.org/zap"
)

func ExampleWriteLogInfo() {
	WriteLogInfo("info data",
		zap.String("string value", time.Now().Format("2006-01-02 15:04:05")),
		zap.Bool("bool value", true),
		zap.Int64("int64 value", 123),
		zap.Duration("duration", time.Second*3),
		zap.Time("time value", time.Now()),
		zap.Errors("erros value", []error{errors.New("my error")}),
	)
	// Output:
}

func TestWriteLogInfo(t *testing.T) {
	tests := []struct {
		name   string
		msg    string
		fields []zap.Field
	}{
		{name: "test1", msg: "info data", fields: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			WriteLogInfo(tt.msg, tt.fields...)
		})
	}
}

func TestWriteLogError(t *testing.T) {
	tests := []struct {
		name   string
		msg    string
		fields []zap.Field
	}{
		{name: "test1", msg: "error msg!!", fields: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			WriteLogError(tt.msg, tt.fields...)
		})
	}
}
