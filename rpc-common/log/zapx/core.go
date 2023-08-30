package zapx

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"path/filepath"
	"time"
)

const (
	logTmFmtWithMS = "2006-01-02 15:04:05.000"
)

type Log struct {
	logDir      string
	logFileName string
	logMaxSize  int
	logMaxAge   int
	localTime   bool
	logCompress bool
	stdout      bool
	traceId     string
	enableColor bool
	jsonFormat  bool
	logMinLevel zapcore.Level
	logger      *zap.Logger
}

func NewLog() (*Log, error) {
	r := &Log{
		logDir:      "./log",
		logFileName: "userApi.log",
		logMaxSize:  512,
		logMaxAge:   30,
		localTime:   true,
		logCompress: false,
		stdout:      false,
		traceId:     "111",
		enableColor: true,
		jsonFormat:  true,
		logMinLevel: zapcore.DebugLevel,
	}
	r.logger = zap.New(initCore(r), zap.AddCaller())

	sugar := r.logger.Sugar()
	sugar.Infow("failed to fetch URL",
		"url", "sougou.com",
		"attempt", 3,
		"backoff", time.Second,
	)
	return r, nil
}

func NewCore() (logx.Writer, error) {
	logger, err := NewLog()
	if err != nil {
		return nil, err
	}

	return logger, nil
}

func initCore(l *Log) zapcore.Core {
	opts := []zapcore.WriteSyncer{
		zapcore.AddSync(&lumberjack.Logger{
			Filename:  filepath.Join(l.logDir, l.logFileName), // ⽇志⽂件路径
			MaxSize:   l.logMaxSize,                           // 单位为MB,默认为512MB
			MaxAge:    l.logMaxAge,                            // 文件最多保存多少天
			LocalTime: l.localTime,                            // 采用本地时间
			Compress:  l.logCompress,                          // 是否压缩日志
		}),
	}

	if l.stdout {
		opts = append(opts, zapcore.AddSync(os.Stdout))
	}

	syncWriter := zapcore.NewMultiWriteSyncer(opts...)

	// 自定义时间输出格式
	customTimeEncoder := func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString("[" + t.Format(logTmFmtWithMS) + "]")
	}
	// 自定义日志级别显示
	customLevelEncoder := func(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString("[" + level.CapitalString() + "]")
	}

	// 自定义文件：行号输出项
	customCallerEncoder := func(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString("[" + l.traceId + "]")
		enc.AppendString("[" + caller.TrimmedPath() + "]")
	}

	encoderConf := zapcore.EncoderConfig{
		CallerKey:      "caller_line", // 打印文件名和行数
		LevelKey:       "level_name",
		MessageKey:     "msg",
		TimeKey:        "ts",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeTime:     customTimeEncoder,   // 自定义时间格式
		EncodeLevel:    customLevelEncoder,  // 小写编码器
		EncodeCaller:   customCallerEncoder, // 全路径编码器
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeName:     zapcore.FullNameEncoder,
	}

	// level大写染色编码器
	if l.enableColor {
		encoderConf.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}

	// json 格式化处理
	if l.jsonFormat {
		return zapcore.NewCore(zapcore.NewJSONEncoder(encoderConf),
			syncWriter, zap.NewAtomicLevelAt(l.logMinLevel))
	}

	return zapcore.NewCore(zapcore.NewConsoleEncoder(encoderConf),
		syncWriter, zap.NewAtomicLevelAt(l.logMinLevel))
}

func (w *Log) Alert(v interface{}) {
	w.logger.Error(fmt.Sprint(v))
}

func (w *Log) Close() error {
	return w.logger.Sync()
}

func (w *Log) Debug(v interface{}, fields ...logx.LogField) {
	w.logger.Debug(fmt.Sprint(v), toZapCoreFields(fields...)...)
}

func (w *Log) Error(v interface{}, fields ...logx.LogField) {
	w.logger.Error(fmt.Sprint(v), toZapCoreFields(fields...)...)
}

func (w *Log) Info(v interface{}, fields ...logx.LogField) {
	w.logger.Info(fmt.Sprint(v), toZapCoreFields(fields...)...)
}

func (w *Log) Severe(v interface{}) {
	w.logger.Fatal(fmt.Sprint(v))
}

func (w *Log) Slow(v interface{}, fields ...logx.LogField) {
	w.logger.Warn(fmt.Sprint(v), toZapCoreFields(fields...)...)
}

func (w *Log) Stack(v interface{}) {
	w.logger.Error(fmt.Sprint(v), zap.Stack("stack"))
}

func (w *Log) Stat(v interface{}, fields ...logx.LogField) {
	w.logger.Info(fmt.Sprint(v), toZapCoreFields(fields...)...)
}

func toZapCoreFields(fields ...logx.LogField) []zap.Field {
	zapFields := make([]zap.Field, 0, len(fields))
	for _, f := range fields {
		zapFields = append(zapFields, zap.Any(f.Key, f.Value))
	}
	return zapFields
}
