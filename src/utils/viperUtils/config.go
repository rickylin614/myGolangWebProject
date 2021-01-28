package viperUtils

import (
	"github.com/spf13/viper"
)

// go get github.com/spf13/viper
func GetLogPath() string {
	viper.SetConfigFile("./resource/properties/logPath.properties")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	return viper.GetString("logPath")
}
