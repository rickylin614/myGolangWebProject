package dao

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/jinzhu/gorm"
)

func TestQueryUser(t *testing.T) {
	tests := []struct {
		name      string
		data      map[string]interface{}
		wantUsers []User
	}{
		// TODO: Add test cases.
		{"test1", map[string]interface{}{
			"pageNo":   1,
			"pageSize": 20,
		}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotUsers, _ := QueryUser(tt.data)
			if len(gotUsers) == 0 {
				b, _ := json.Marshal(gotUsers)
				t.Errorf("QueryUser() = %v, want %v", string(b), tt.wantUsers)
			}
		})
	}
}

func TestQueryUserByName(t *testing.T) {
	tests := []struct {
		name      string
		loginName string
		wantU     User
	}{
		// TODO: Add test cases.
		{"test", "ricky001", User{Model: gorm.Model{ID: 2}, Pwd: "qwe123"}},
		{"test2", "ricky002", User{Model: gorm.Model{ID: 1}, Pwd: "qwe123"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotU := QueryUserByName(tt.loginName)
			if !reflect.DeepEqual(gotU.ID, tt.wantU.ID) {
				t.Errorf("QueryUserByName() ID = %v, want %v", gotU.ID, tt.wantU.ID)
			}
			if !reflect.DeepEqual(gotU.Pwd, tt.wantU.Pwd) {
				t.Errorf("QueryUserByName() PWD = %v, want %v", gotU.Pwd, tt.wantU.Pwd)
			}
		})
	}
}
