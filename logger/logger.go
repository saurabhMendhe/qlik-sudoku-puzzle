package logger

import (
	"fmt"
	"os"

	"github.com/saurabhMendhe/qlik-sudoku-puzzzle/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// NewLogger creates a new configured zap logger based on config
func NewLogger(cfg *config.Config) (*zap.Logger, error) {
	// Parse log level
	level, err := parseLevel(cfg.Log.Level)
	if err != nil {
		return nil, err
	}

	// Create encoder config
	encoderConfig := zap.NewProductionEncoderConfig()
	if cfg.IsDevelopment() {
		encoderConfig = zap.NewDevelopmentEncoderConfig()
	}
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder

	// Create encoder
	var encoder zapcore.Encoder
	if cfg.Log.Format == "json" {
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	} else {
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	}

	// Configure output
	var writeSyncer zapcore.WriteSyncer
	switch cfg.Log.OutputPath {
	case "stdout":
		writeSyncer = zapcore.AddSync(os.Stdout)
	case "stderr":
		writeSyncer = zapcore.AddSync(os.Stderr)
	default:
		// File output
		file, err := os.OpenFile(cfg.Log.OutputPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
		if err != nil {
			return nil, fmt.Errorf("failed to open log file: %w", err)
		}
		writeSyncer = zapcore.AddSync(file)
	}

	// Create core
	core := zapcore.NewCore(encoder, writeSyncer, level)

	// Build logger with options
	opts := []zap.Option{}
	if cfg.Log.EnableCaller {
		opts = append(opts, zap.AddCaller())
	}
	if cfg.Log.EnableStacktrace {
		opts = append(opts, zap.AddStacktrace(zapcore.ErrorLevel))
	}

	logger := zap.New(core, opts...)

	// Add default fields
	logger = logger.With(
		zap.String("app", cfg.App.Name),
		zap.String("version", cfg.App.Version),
		zap.String("environment", cfg.App.Environment),
	)

	return logger, nil
}

// parseLevel converts string level to zapcore.Level
func parseLevel(level string) (zapcore.Level, error) {
	switch level {
	case "debug":
		return zapcore.DebugLevel, nil
	case "info":
		return zapcore.InfoLevel, nil
	case "warn":
		return zapcore.WarnLevel, nil
	case "error":
		return zapcore.ErrorLevel, nil
	default:
		return zapcore.InfoLevel, fmt.Errorf("invalid log level: %s", level)
	}
}

// NewDevelopmentLogger creates a preconfigured development logger
func NewDevelopmentLogger() (*zap.Logger, error) {
	return zap.NewDevelopment()
}

// NewProductionLogger creates a preconfigured production logger
func NewProductionLogger() (*zap.Logger, error) {
	return zap.NewProduction()
}
