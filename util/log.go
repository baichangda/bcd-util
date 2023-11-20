package util

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

const EnableLogFile = true

var Log = initLog()

func initLog() *zap.SugaredLogger {
	// 获取日志写入位置
	writeSyncer := getLogWriter(EnableLogFile)
	// 获取日志编码格式
	encoder := getEncoder()

	// 获取日志最低等级，即>=该等级，才会被写入。
	var l = new(zapcore.Level)
	err := l.Set(zap.InfoLevel.String())
	if err != nil {
		panic(err)
	}

	// 创建一个将日志写入 WriteSyncer 的核心。
	core := zapcore.NewCore(encoder, writeSyncer, l)
	logger := zap.New(core, zap.AddCaller())

	// 替换zap包中全局的logger实例，后续在其他包中只需使用zap.L()调用即可
	zap.ReplaceGlobals(logger)
	return logger.Sugar()
}

// 负责设置 encoding 的日志格式
func getEncoder() zapcore.Encoder {
	// 获取一个指定的的EncoderConfig，进行自定义
	encodeConfig := zap.NewProductionEncoderConfig()
	// 设置每个日志条目使用的键。如果有任何键为空，则省略该条目的部分。
	// 序列化时间。eg: 2022-09-01T19:11:35.921+0800
	encodeConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	// "time":"2022-09-01T19:11:35.921+0800"
	encodeConfig.TimeKey = "time"
	// 将Level序列化为全大写字符串。例如，将info level序列化为INFO。
	encodeConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	// 以 package/file:行 的格式 序列化调用程序，从完整路径中删除除最后一个目录外的所有目录。
	encodeConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewConsoleEncoder(encodeConfig)
}

// 负责日志写入的位置
func getLogWriter(logFile bool) zapcore.WriteSyncer {
	if logFile {
		lumberJackLogger := &lumberjack.Logger{
			Filename:   "./logs/data.log", // 文件位置
			MaxSize:    1024,              // 进行切割之前,日志文件的最大大小(MB为单位)
			MaxAge:     30,                // 保留旧文件的最大天数
			MaxBackups: 10,                // 保留旧文件的最大个数
			Compress:   true,              // 是否压缩/归档旧文件
		}
		// AddSync 将 io.Writer 转换为 WriteSyncer。
		// 它试图变得智能：如果 io.Writer 的具体类型实现了 WriteSyncer，我们将使用现有的 Sync 方法。
		// 如果没有，我们将添加一个无操作同步。
		fileWriter := zapcore.AddSync(lumberJackLogger)
		consoleWriter := zapcore.AddSync(os.Stdout)
		return zapcore.NewMultiWriteSyncer(consoleWriter, fileWriter)
	} else {
		return zapcore.AddSync(os.Stdout)
	}
}
