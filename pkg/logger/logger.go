package logger

import (
	"log"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger

// init initializes the logger.
func init() {
	conf := zap.NewProductionConfig()
	conf.OutputPaths = []string{"stdout"}
	conf.Level.SetLevel(zap.InfoLevel)
	// if c.Get("Environment") == "development" {
	// 	conf.Level.SetLevel(zap.DebugLevel)
	// }
	conf.Level.SetLevel(zap.DebugLevel)
	conf.EncoderConfig.TimeKey = "timestamp"
	conf.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	zapLogger, err := conf.Build()
	if err != nil {
		log.Fatalf("Error initializing logger: %s", err)
	}
	logger = zapLogger.With(zap.String("service", "socialize"))
}

func Logger() *zap.Logger {
	return logger
}
