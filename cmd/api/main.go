package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"

	"github.com/kieranroneill/oarbit/internal/platform/application"
	"github.com/kieranroneill/oarbit/internal/platform/logging"
	"github.com/kieranroneill/oarbit/internal/platform/postgres"
	"github.com/kieranroneill/oarbit/internal/versions"
)

func main() {
	// create context for shutdown processes
	ctx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)
	defer stop()

	applicationConfig := application.NewConfig()
	databaseConfig := postgres.NewConfig()

	// config logger
	slog.SetLogLoggerLevel(logging.ParseLogLevel(applicationConfig.LogLevel))

	// get database connection
	database, err := postgres.NewConnection(ctx, databaseConfig)
	if err != nil {
		slog.Error("failed to connect to postgres", "error", err)

		return
	}
	defer func() {
		if _err := database.Close(); _err != nil {
			slog.Error("failed to close postgres connection", "error", err)
		}
	}()

	e := echo.New()

	// middlewares
	e.Use(middleware.RequestLogger())

	// routes
	// /versions
	e.GET(versions.Path, versions.NewGetVersionsRoute(applicationConfig))

	// set up server
	server := echo.StartConfig{
		Address:         fmt.Sprintf(":%s", applicationConfig.Port),
		GracefulTimeout: 10 * time.Second,
	}

	// start the server
	if err = server.Start(ctx, e); err != nil {
		slog.Error("failed to start server", "error", err)

		return
	}
}
