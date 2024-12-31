package app

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/lj222kj/api"
	"github.com/lj222kj/services/player"
	"github.com/lj222kj/xpkg/config"
	"github.com/lj222kj/xpkg/logger"
)

func New() error {
	cfg, err := config.NewConfig()
	if err != nil {
		return err
	}

	lgr := logger.New()
	lgr.Info("Application starting",
		slog.String("app name", "Football manager OOS"))

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	playerSvc := player.New()
	api := api.NewRestApi(cfg, lgr, playerSvc)

	go func() {
		err := api.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			lgr.ErrorContext(ctx, "http application shutdown", err)
		}
	}()

	<-ctx.Done()
	stop()

	lgr.Info("graceful shutdown")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := api.Shutdown(ctx); err != nil {
		lgr.With(err).ErrorContext(ctx, "Server forced to shutdown: %e")
	}
	return nil
}
