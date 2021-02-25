package errRecordDao

import "testing"

func TestInsert(t *testing.T) {
	tests := []ErrRecord{
		{Name: "123", UserId: 456},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			Insert(&tt)
		})
	}
}
