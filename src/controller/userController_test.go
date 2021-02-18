package controller

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestRegister(t *testing.T) {
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w) // gin用於test的方法 做出ctx

	tests := []struct {
		msg    string // 測試說明
		params gin.H  // 輸入引數
		errMsg string // 錯誤資訊
		want   string // 期望結果
	}{
		{
			msg: "資料正確",
			params: gin.H{
				"Name":     "ricky001",
				"Password": "qwe123",
			},
			want: `{"code":"error","msg":"已註冊的帳號"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.msg, func(t *testing.T) {
			data, _ := json.Marshal(tt.params)                                    // 將參數轉為json格式(byte)
			ctx.Request = httptest.NewRequest("POST", "/", bytes.NewBuffer(data)) // 設定請求參數
			ctx.Request.Header.Set("Content-Type", gin.MIMEJSON)                  // 設定Content-type為json
			Register(ctx)
			if tt.want != w.Body.String() {
				t.Errorf("error!! want: %s , get: %s", tt.want, w.Body.String())
			}
		})
	}
}

func TestLogin(t *testing.T) {
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w) // gin用於test的方法 做出ctx

	tests := []struct {
		msg    string // 測試說明
		params gin.H  // 輸入引數
		errMsg string // 錯誤資訊
		want   string // 期望結果
	}{
		{
			msg: "測試1",
			params: gin.H{
				"Name":     "ricky001",
				"Password": "qwe123",
			},
			want: `{"code":200,"msg":"登入成功"}`,
		},
		{
			msg: "測試2",
			params: gin.H{
				"Name":     "ricky001",
				"Password": "qwe12",
			},
			want: `{"code":"error","msg":"密碼錯誤"}`,
		},
		{
			msg: "測試3",
			params: gin.H{
				"Name":     "",
				"Password": "",
			},
			want: `{"code":"error","msg":"資料格式不正確"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.msg, func(t *testing.T) {
			data, _ := json.Marshal(tt.params)                                    // 將參數轉為json格式(byte)
			ctx.Request = httptest.NewRequest("POST", "/", bytes.NewBuffer(data)) // 設定請求參數
			ctx.Request.Header.Set("Content-Type", gin.MIMEJSON)                  // 設定Content-type為json
			Login(ctx)
			if tt.want != w.Body.String() {
				t.Errorf("error!! want: %s , get: %s", tt.want, w.Body.String())
			}
			w.Body = new(bytes.Buffer)
		})
	}
}
