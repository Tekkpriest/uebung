package main

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/uebung/data"
	"github.com/uebung/handlers"
	"github.com/uebung/middleware"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	mux := http.NewServeMux()

	storeTasks := data.NewStoreTasks()
	taskHandler := handlers.NewToDoHandler(storeTasks)

	mux.HandleFunc("GET /api/todos", taskHandler.GetToDos)

	cors := middleware.CorsHandler("*")
	server := &http.Server{
		Addr:         ":3000", // later cfg.port via env
		Handler:      cors(mux),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	go func() {
		slog.Info("server is starting on", "port", server.Addr)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			slog.Error("server encountered an error", "error", err)
			os.Exit(1)
		}
	}()

	<-ctx.Done()

	slog.Info("server shutdown started")

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdownCancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		slog.Error("error during shutdown", "error", err)
		os.Exit(1)
	}

	slog.Info("server shutdown without errors")
}
