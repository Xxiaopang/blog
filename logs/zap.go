package logs

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"project/config"
	"time"
)

var (
	logger *zap.Logger
	sugar  *zap.SugaredLogger
	Logger *zap.Logger
)

func init() {

	// Encoder，WriteSyncer，LogLevel

	core := zapcore.NewCore(getEncoder(), getWriteSyncer(), getLogLevel())
	logger := zap.New(core)
	sugar = logger.Sugar()
}

func InitLogger() {
	defer sugar.Sync()

}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.LowercaseLevelEncoder
	encoderConfig.EncodeTime = func(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(t.Local().Format("2006-01-02 15:04:05"))
	}

	return zapcore.NewJSONEncoder(encoderConfig)
}

func getWriteSyncer() zapcore.WriteSyncer {

	// 创建文件对象
	file, _ := os.Create("./getLog.log") // 或者是用 OpenFile函数，在原来基础上追加。
	// file, _ := os.OpenFile("./getLog.log", os.O_APPEND | os.O_RDWR, 0744)
	// 生成 WriteSyncer
	wSy := zapcore.AddSync(file)
	return wSy

}

func getLogLevel() zapcore.Level {
	if config.App.Mode == "dev" {
		return zapcore.DebugLevel
	} else {
		return zapcore.InfoLevel
	}
}
