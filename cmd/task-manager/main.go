package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/invocoder/task-manager/internal/config"
	"github.com/invocoder/task-manager/internal/config/http/handlers/task"
)

func main() {
	//load config
	cfg := config.MustLoad()

	//dataabse setup
	//setup router
	router := http.NewServeMux()

	router.HandleFunc("POST /api/tasks", task.New())

	//setup server
	server := http.Server{
		Addr:    cfg.Addr,
		Handler: router,
	}

	slog.Info("server started %s", slog.String("address", cfg.Addr))
	fmt.Println("server started")

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal("failed to start server")
		}
	}()
	<-done

	slog.Info("shutting down the server")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		slog.Error("failed to shutdown server", slog.String("error", err.Error()))
	}
	slog.Info("Server shutdown succssfully")

}
