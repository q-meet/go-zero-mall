package zapx

import (
	"fmt"
	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	"github.com/zeromicro/go-zero/core/logx"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"path/filepath"
	"time"
)

var (
	//logger                         *Logger
	sp                             = string(filepath.Separator)
	errWS, warnWS, infoWS, debugWS zapcore.WriteSyncer // IO输出
	//debugConsoleWS                 = zapcore.Lock(os.Stdout) // 控制台标准输出
	//errorConsoleWS                 = zapcore.Lock(os.Stderr)
)

type ZapWriter struct {
	logger *zap.Logger
}

func Level() zapcore.Level {
	return zapcore.DebugLevel
}
func InitLogger() (logx.Writer, error) {
	getLogWriter()
	encoder := getEncoder()

	errPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl > zapcore.WarnLevel && zapcore.WarnLevel-Level() > -1
	})
	warnPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl == zapcore.WarnLevel && zapcore.WarnLevel-Level() > -1
	})
	infoPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl == zapcore.InfoLevel && zapcore.InfoLevel-Level() > -1
	})
	debugPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl == zapcore.DebugLevel && zapcore.DebugLevel-Level() > -1
	})

	// 最后创建具体的Logger
	core := zapcore.NewTee(
		zapcore.NewCore(encoder, errWS, errPriority),
		zapcore.NewCore(encoder, warnWS, warnPriority),
		zapcore.NewCore(encoder, infoWS, infoPriority),
		zapcore.NewCore(encoder, debugWS, debugPriority),
	)

	logger := zap.New(core, zap.AddCaller()) // 需要传入 zap.AddCaller() 才会显示打日志点的文件名和行数, 有点小坑

	return &ZapWriter{
		logger: logger,
	}, nil
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.StacktraceKey = "Stacktrace"
	//encoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	//	enc.AppendString(t.Format("2006-01-02 15:04:05"))
	//}
	return zapcore.NewJSONEncoder(encoderConfig)
	//return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter() {
	f := func(fN string) zapcore.WriteSyncer {
		logf, _ := rotatelogs.New(
			"logs"+sp+fN+".%Y_%m%d.log",
			//"logs"+sp+fN+".%Y_%m%d_%H.log",
			//rotatelogs.WithLinkName("logs"+sp+fN),
			rotatelogs.WithMaxAge(30*24*time.Hour),
			rotatelogs.WithRotationTime(time.Minute),
		)
		return zapcore.AddSync(logf)
	}
	var ErrorFileName = "Error"
	var WarnFileName = "Warn"
	var InfoFileName = "Info"
	var DebugFileName = "Debug"
	errWS = f(ErrorFileName)
	warnWS = f(WarnFileName)
	infoWS = f(InfoFileName)
	debugWS = f(DebugFileName)
	/*
		Filename: 日志文件的位置
		MaxSize：在进行切割之前，日志文件的最大大小（以MB为单位）
		MaxBackups：保留旧文件的最大个数
		MaxAges：保留旧文件的最大天数
		Compress：是否压缩/归档旧文件
	*/
	/*
		   lumberJackLogger := &lumberjack.Logger{
				   Filename:   "./test.log",
				   MaxSize:    10,
				   MaxBackups: 5,
				   MaxAge:     30,
				   Compress:   false,
				}
			  return zapcore.AddSync(lumberJackLogger)
	*/
}

func (w *ZapWriter) Alert(v interface{}) {
	w.logger.Error(fmt.Sprint(v))
}

func (w *ZapWriter) Close() error {
	return w.logger.Sync()
}

func (w *ZapWriter) Debug(v interface{}, fields ...logx.LogField) {
	w.logger.Debug(fmt.Sprint(v), toZapDeFields(fields...)...)
}

func (w *ZapWriter) Error(v interface{}, fields ...logx.LogField) {
	w.logger.Error(fmt.Sprint(v), toZapDeFields(fields...)...)
}

func (w *ZapWriter) Info(v interface{}, fields ...logx.LogField) {
	w.logger.Info(fmt.Sprint(v), toZapDeFields(fields...)...)
}

func (w *ZapWriter) Severe(v interface{}) {
	w.logger.Fatal(fmt.Sprint(v))
}

func (w *ZapWriter) Slow(v interface{}, fields ...logx.LogField) {
	w.logger.Warn(fmt.Sprint(v), toZapDeFields(fields...)...)
}

func (w *ZapWriter) Stack(v interface{}) {
	w.logger.Error(fmt.Sprint(v), zap.Stack("stack"))
}

func (w *ZapWriter) Stat(v interface{}, fields ...logx.LogField) {
	w.logger.Info(fmt.Sprint(v), toZapDeFields(fields...)...)
}

func toZapDeFields(fields ...logx.LogField) []zap.Field {
	zapFields := make([]zap.Field, 0, len(fields))
	for _, f := range fields {
		zapFields = append(zapFields, zap.Any(f.Key, f.Value))
	}
	return zapFields
}
