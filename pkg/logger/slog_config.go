package logger

import (
	"docero/internal/config"
	"io"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/natefinch/lumberjack"
)

var Logger *slog.Logger

// InitSlogLogger initializes a global slog logger with file rotation.
func InitSlogLogger(cfg *config.Config) {
	logConfig := cfg.Log

	// Ensure log directory exists
	logDir := filepath.Dir(logConfig.Filename)
	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		err = os.MkdirAll(logDir, 0755)
		if err != nil {
			slog.Default().Error("Failed to create log directory", "path", logDir, "error", err)
			// Fallback to stderr if directory can't be created
			Logger = slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
				AddSource: true,
				Level:     parseSlogLevel(logConfig.Level),
			}))
			return
		}
	}

	// Configure lumberjack for log rotation
	lumberjackLogger := &lumberjack.Logger{
		Filename:   logConfig.Filename,
		MaxSize:    logConfig.MaxSize, // megabytes
		MaxBackups: logConfig.MaxBackups,
		MaxAge:     logConfig.MaxAge, // days
		Compress:   logConfig.Compress,
	}

	// Create a MultiWriter to write to both stdout and the rotating file
	var writers []io.Writer
	writers = append(writers, os.Stdout) // Always log to stdout for real-time monitoring
	writers = append(writers, lumberjackLogger)

	multiWriter := io.MultiWriter(writers...)

	// Configure slog handler options
	handlerOptions := &slog.HandlerOptions{
		AddSource: true, // Shows file and line number
		Level:     parseSlogLevel(logConfig.Level),
	}

	var handler slog.Handler
	switch logConfig.Format {
	case "json":
		handler = slog.NewJSONHandler(multiWriter, handlerOptions)
	case "text":
		handler = slog.NewTextHandler(multiWriter, handlerOptions)
	default:
		slog.Default().Warn("Unknown log format, defaulting to text", "format", logConfig.Format)
		handler = slog.NewTextHandler(multiWriter, handlerOptions)
	}

	Logger = slog.New(handler)
	slog.SetDefault(Logger) // Set the global default logger
}

// parseSlogLevel converts string level to slog.Level
func parseSlogLevel(levelStr string) slog.Level {
	switch levelStr {
	case "debug":
		return slog.LevelDebug
	case "info":
		return slog.LevelInfo
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}

// CloseLogger closes the lumberjack file. Call this on application shutdown.
func CloseLogger() {
	// lumberjack doesn't expose a direct Close method on the Logger struct
	// It handles closing internally. However, if you had other resources
	// in multiWriter that needed explicit closing, you would do it here.
	// For example, if you opened os.Stdout, you wouldn't close it.
	// No explicit close needed for lumberjack.
}
