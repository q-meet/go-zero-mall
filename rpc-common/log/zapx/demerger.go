package zapx

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"time"
)

type ZapDeWriter struct {
	logger *zap.Logger
}

func InitLogger() (logx.Writer, error) {
	writeSyncer := getLogWriter()
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)

	logger := zap.New(core, zap.AddCaller())

	sugar := logger.Sugar()
	sugar.Infow("failed to fetch URL",
		"url", "sougou.com",
		"attempt", 3,
		"backoff", time.Second,
	)

	return &ZapWriter{
		logger: logger,
	}, nil
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.StacktraceKey = "Stacktrace"
	return zapcore.NewJSONEncoder(encoderConfig)
	//return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter() zapcore.WriteSyncer {
	/*
		Filename: 日志文件的位置
		MaxSize：在进行切割之前，日志文件的最大大小（以MB为单位）
		MaxBackups：保留旧文件的最大个数
		MaxAges：保留旧文件的最大天数
		Compress：是否压缩/归档旧文件
	*/
	lumberJackLogger := &lumberjack.Logger{
		Filename:   "./test.log",
		MaxSize:    10,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   false,
	}
	return zapcore.AddSync(lumberJackLogger)
}

func (w *ZapDeWriter) Alert(v interface{}) {
	w.logger.Error(fmt.Sprint(v))
}

func (w *ZapDeWriter) Close() error {
	return w.logger.Sync()
}

func (w *ZapDeWriter) Debug(v interface{}, fields ...logx.LogField) {
	w.logger.Debug(fmt.Sprint(v), toZapDeFields(fields...)...)
}

func (w *ZapDeWriter) Error(v interface{}, fields ...logx.LogField) {
	w.logger.Error(fmt.Sprint(v), toZapDeFields(fields...)...)
}

func (w *ZapDeWriter) Info(v interface{}, fields ...logx.LogField) {
	w.logger.Info(fmt.Sprint(v), toZapDeFields(fields...)...)
}

func (w *ZapDeWriter) Severe(v interface{}) {
	w.logger.Fatal(fmt.Sprint(v))
}

func (w *ZapDeWriter) Slow(v interface{}, fields ...logx.LogField) {
	w.logger.Warn(fmt.Sprint(v), toZapDeFields(fields...)...)
}

func (w *ZapDeWriter) Stack(v interface{}) {
	w.logger.Error(fmt.Sprint(v), zap.Stack("stack"))
}

func (w *ZapDeWriter) Stat(v interface{}, fields ...logx.LogField) {
	w.logger.Info(fmt.Sprint(v), toZapDeFields(fields...)...)
}

func toZapDeFields(fields ...logx.LogField) []zap.Field {
	zapFields := make([]zap.Field, 0, len(fields))
	for _, f := range fields {
		zapFields = append(zapFields, zap.Any(f.Key, f.Value))
	}
	return zapFields
}
