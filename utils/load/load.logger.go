package load

import (
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

func init() {
	// Initialize the logger

	var encodeCfg zapcore.EncoderConfig

	if Cfg.GIN_MODE == "release" {
		// Set the logger to production mode
		encodeCfg = zap.NewProductionEncoderConfig()
	} else {
		// Set the logger to development mode
		encodeCfg = zap.NewDevelopmentEncoderConfig()

	}
	encodeCfg.EncodeTime = zapcore.TimeEncoder(func(t time.Time, pae zapcore.PrimitiveArrayEncoder) {
		pae.AppendString(t.Format("2006-01-02 15:04:05"))
	})
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encodeCfg),
		zapcore.AddSync(os.Stdout),
		zapcore.DebugLevel,
	)

	Logger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))

	zap.ReplaceGlobals(Logger)

	if Logger == nil {
		panic("Logger is not initialized")
	} else {
		Logger.Info("Logger initialized", zap.String("mode", Cfg.GIN_MODE))
	}

}
