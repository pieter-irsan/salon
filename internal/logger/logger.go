package logger

import (
	"log/slog"
	"os"
	"time"
)

func New() *slog.Logger {
	return slog.New(
		slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			AddSource:   true,
			Level:       slog.LevelDebug,
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
