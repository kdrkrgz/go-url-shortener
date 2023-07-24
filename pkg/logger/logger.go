package logger

import (
	"log"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger

// init initializes the logger.
func init() {
	conf := zap.NewProductionConfig()
	conf.OutputPaths = []string{"stdout"}
	conf.Level.SetLevel(zap.InfoLevel)
	if os.Getenv("Environment") == "development" {
		conf.Level.SetLevel(zap.DebugLevel)
	}
	conf.Level.SetLevel(zap.DebugLevel)
	conf.EncoderConfig.TimeKey = "timestamp"
	conf.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	zapLogger, err := conf.Build()
	if err != nil {
		log.Fatalf("Error initializing logger: %s", err)
	}
	logger = zapLogger.With(zap.String("service", "go-url-shortener"))
}

func Logger() *zap.Logger {
	return logger
}
