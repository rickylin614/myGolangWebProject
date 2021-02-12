package viperUtils

import (
	"github.com/spf13/viper"
)

const basepath = "./resource/properties/"
const commonFile = basepath + "common.properties"
const logPathParam = "logPath"

// go get github.com/spf13/viper
func GetLogPath() string {
	viper.SetConfigFile(commonFile)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	return viper.GetString(logPathParam)
}
