package conf

import (
	"fmt"
	"os"
	"time"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func InitLogger() *zap.Logger {
	core := zapcore.NewCore(getEncoder(), getWriteSyncer(), zap.DebugLevel)
	return zap.New(core)
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "datetime"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder //大写
	encoderConfig.EncodeTime = func(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(t.Local().Format(time.DateTime))
	}
	return zapcore.NewJSONEncoder(encoderConfig)
}

func getWriteSyncer() zapcore.WriteSyncer {

	t := time.Now().Format(time.DateOnly)

	pwd, _ := os.Getwd()
	logPath := fmt.Sprintf(pwd+"/log/%s.log", t)

	lumberJackLogger := &lumberjack.Logger{
		Filename:   logPath,
		MaxSize:    10,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   false,
	}
	return zapcore.AddSync(lumberJackLogger)
}
