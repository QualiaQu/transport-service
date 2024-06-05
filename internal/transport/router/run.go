package router

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"transport-service/internal/helpers"
)

type serverOpts struct {
	*http.Server
}

func newServer(address string, handler http.Handler) serverOpts {
	return serverOpts{
		Server: &http.Server{
			Addr:    address,
			Handler: handler,
		},
	}
}

func (r *Router) Run(addr string) {
	if r == nil || addr == "" {
		slog.Error("ошибка запуска приложения: для запуска сервера не указаны необходимые параметры")

		return
	}

	server := newServer(addr, r)

	go server.start()

	server.stop()
}

func (server serverOpts) start() {
	if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		slog.Error("Error start server: ", err)
	}
}

func (server serverOpts) stop() {
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	slog.Info("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), helpers.RegisterSleepTime)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		slog.Info("Server Shutdown:", err)
	}

	select {
	case <-ctx.Done():
		slog.Info("timeout of 5 seconds.")
	}
	slog.Info("Server exiting")
}
