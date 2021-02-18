package viperUtils

import (
	"fmt"
	"path/filepath"
	"runtime"

	"github.com/spf13/viper"
)

// go get github.com/spf13/viper

// 根據此檔的路徑 配置靜態文件的相對路徑
const basepath = "../../../resource/properties/"
const commonFile = basepath + "common.properties"
const logPathParam = "logPath"

// 存入絕對路徑
var currentBase string

func init() {
	_, currentFile, _, _ := runtime.Caller(0) // 取得此時本地的檔案路徑
	currentBase = filepath.Dir(currentFile)   // 取得此時本地資料夾路徑
}

/* 此檔絕對路徑 與 靜態文件相對路徑 組合 靜態文件絕對路徑 */
func AbsolutePath(rel string) string {
	return filepath.Join(currentBase, rel)
}

func GetLogPath() string {
	path := AbsolutePath(commonFile)
	viper.SetConfigFile(path)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	return viper.GetString(logPathParam)
}

func GetCommonParams(param string) string {
	path := AbsolutePath(commonFile)
	viper.SetConfigFile(path)
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("err", err)
		return ""
	}
	return viper.GetString(param)
}
