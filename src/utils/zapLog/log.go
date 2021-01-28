package zapLog

import (
	"os"
	"strings"

	"go.uber.org/zap"

	"orderbento/src/utils/viperUtils"
)

//go get -u go.uber.org/zap

/* 使用預設LOG */
/* func GetLog() (logger *log.Logger) {
	out := gin.DefaultErrorWriter // TODO: 先同步gin的錯誤日至，之後應以設定檔為主
	if out != nil {
		logger = log.New(out, "\n\n\x1b[31m", log.LstdFlags)
	}
	return logger
} */

var logger *zap.Logger

func init() {
	path := viperUtils.GetLogPath()
	config := &zap.Config{
		Level:       zap.NewAtomicLevelAt(zap.DebugLevel),
		Development: false,
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		Encoding:         "console",
		EncoderConfig:    zap.NewProductionEncoderConfig(),
		OutputPaths:      []string{"stderr", path}, // TODO: 之後應以設定檔為主
		ErrorOutputPaths: []string{"stderr", path}, // TODO: 之後應以設定檔為主
	}

	//確認輸出路徑有資料夾，沒有就創建
	dir := path[:strings.LastIndex(path, "/")]
	_, err := os.Stat(dir)
	if os.IsNotExist(err) {
		errDir := os.MkdirAll(dir, os.ModePerm)
		if errDir != nil {
			panic(errDir)
		}
	}
	logger, err = config.Build()
	if err != nil {
		panic(err)
	}
}

/* 使用預設 zap.logger.info */
func WriteLogInfo(msg string, fields ...zap.Field) {
	defer sync()
	logger.Info(msg, fields...)
}

/* write log with errorLevel */
func WriteLogError(msg string, fields ...zap.Field) {
	defer sync()
	logger.Error(msg, fields...)
}

/* quick write error, skip annotation zap.Error() */
func ErrorW(msg string, err error) {
	defer sync()
	logger.Error(msg, zap.Error(err))
}

/* quick panic error, skip annotation zap.Error() */
func PanicW(msg string, err error) {
	defer sync()
	logger.Panic(msg, zap.Error(err))
}

/* write log with debugLevel */
func WriteLogDebug(msg string, fields ...zap.Field) {
	defer sync()
	logger.Debug(msg, fields...)
}

/* write log with warnLevel */
func WriteLogWarn(msg string, fields ...zap.Field) {
	defer sync()
	logger.Warn(msg, fields...)
}

/* write log then panic */
func WriteLogPanic(msg string, fields ...zap.Field) {
	defer sync()
	logger.Panic(msg, fields...)
}

/* write log then calls os.Exit(1) */
func WriteLogFatal(msg string, fields ...zap.Field) {
	defer sync()
	logger.Fatal(msg, fields...)
}

/* flushing any buffered log */
func sync() {
	err := logger.Sync()
	if err != nil {
		panic(err)
	}
}
