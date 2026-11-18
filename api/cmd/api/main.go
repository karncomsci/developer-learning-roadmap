// Command api serves the Developer Roadmap HTTP API.
package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v3"

	"github.com/karncomsci/developer-learning-roadmap/api/internal/httpapi"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	app := httpapi.NewApp(httpapi.NewServer())

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	go func() {
		<-ctx.Done()
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := app.ShutdownWithContext(shutdownCtx); err != nil {
			slog.Error("shutdown failed", "err", err)
		}
	}()

	if err := app.Listen(":"+port, fiber.ListenConfig{DisableStartupMessage: true}); err != nil {
		slog.Error("server stopped", "err", err)
		os.Exit(1)
	}
}
