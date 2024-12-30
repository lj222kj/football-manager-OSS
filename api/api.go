package api

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"time"
)

type Api struct {
	playerSvc PlayerService
	le        *slog.Logger
	srv       *http.Server
	mux       *http.ServeMux
}

func (a *Api) middleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	}
}

func (a *Api) GetPlayers() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(a.playerSvc.Get(r.Context()))
	}
}

func (a *Api) GetMux() *http.ServeMux {
	return a.mux
}

func NewRestApi(le *slog.Logger, playerSvc PlayerService) *Api {
	mux := http.NewServeMux()

	api := &Api{
		playerSvc: playerSvc,
		le:        le,
		mux:       mux,
		srv: &http.Server{
			Addr:         "127.0.0.1:8080",
			Handler:      mux,
			ReadTimeout:  time.Second * 15,
			WriteTimeout: time.Second * 15,
			IdleTimeout:  time.Second * 30,
		},
	}

	mux.Handle("GET /players/", api.middleware(api.GetPlayers()))
	return api
}

func (a Api) Shutdown(ctx context.Context) error {
	return a.srv.Shutdown(ctx)
}

func (a Api) ListenAndServe() error {
	return a.srv.ListenAndServe()
}
