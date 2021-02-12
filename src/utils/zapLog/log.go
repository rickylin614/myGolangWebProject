package zapLog

import (
	"os"
	"strings"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"orderbento/src/utils/viperUtils"
)

//go get -u go.uber.org/zap

var logger *zap.Logger

/* 設定檔設定 */
func init() {
	path := viperUtils.GetLogPath()

	productionEncoder := zap.NewProductionEncoderConfig()
	productionEncoder.EncodeTime = definedTimeFormat
	config := &zap.Config{
		Level:       zap.NewAtomicLevelAt(zap.DebugLevel),
		Development: false,
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		Encoding:         "console",
		EncoderConfig:    productionEncoder,
		OutputPaths:      []string{"stderr", path}, // strerr use in debug mode
		ErrorOutputPaths: []string{"stderr", path}, //
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

/* 自定義時間輸出格式 */
func definedTimeFormat(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	type appendTimeEncoder interface {
		AppendTimeEncoder(time.Time, string)
	}
	layout := "2006-01-02 15:04:05.999"
	if enc, ok := enc.(appendTimeEncoder); ok {
		enc.AppendTimeEncoder(t, layout)
	}
	enc.AppendString(t.Format(layout))
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
