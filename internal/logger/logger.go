package logger

import (
	"log/slog"
	"os"
	"salon/internal/config"
	"time"
)

func New(env config.Env) *slog.Logger {
	level := slog.LevelDebug
	if env.IsProduction() {
		level = slog.LevelError
	}

	return slog.New(
		slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			AddSource:   true,
			Level:       level,
			ReplaceAttr: formatTimestamp,
		}),
	)
}

func formatTimestamp(groups []string, attr slog.Attr) slog.Attr {
	if attr.Key == slog.TimeKey {
		t := attr.Value.Time().Format(time.DateTime)
		attr.Value = slog.StringValue(t)
	}
	return attr
}
