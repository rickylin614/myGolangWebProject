package viperUtils

import (
	"fmt"
	"path/filepath"
	"runtime"

	"github.com/spf13/viper"
)

// go get github.com/spf13/viper

const basepath = "../../../resource/properties/"
const commonFile = basepath + "common.properties"
const logPathParam = "logPath"

var currentBase string

func init() {
	_, currentFile, _, _ := runtime.Caller(0)
	currentBase = filepath.Dir(currentFile)
}

func GetLogPath() string {
	path := AbsolutePath(commonFile)
	fmt.Println(path)
	viper.SetConfigFile(path)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	return viper.GetString(logPathParam)
}

func GetCommonParams(param string) string {
	path := AbsolutePath(commonFile)
	fmt.Println(path)
	viper.SetConfigFile(path)
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("err", err)
		return ""
	}
	return viper.GetString(param)
}

func AbsolutePath(rel string) string {
	return filepath.Join(currentBase, rel)
}
