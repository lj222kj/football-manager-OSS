package logger

import (
	"log/slog"
	"os"
)

func New() *slog.Logger {
	handler := slog.NewTextHandler(os.Stdout, nil)
	lgr := slog.New(handler)

	return lgr
}
