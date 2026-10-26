package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"

	"github.com/kieranroneill/oarbit/internal/authentication"
	"github.com/kieranroneill/oarbit/internal/conversation"
	"github.com/kieranroneill/oarbit/internal/platform/agent"
	"github.com/kieranroneill/oarbit/internal/platform/application"
	_http "github.com/kieranroneill/oarbit/internal/platform/http"
	"github.com/kieranroneill/oarbit/internal/platform/logging"
	"github.com/kieranroneill/oarbit/internal/platform/postgres"
	"github.com/kieranroneill/oarbit/internal/platform/redis"
	"github.com/kieranroneill/oarbit/internal/versions"
	"github.com/kieranroneill/oarbit/internal/vessel"
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
	postgresConfig := postgres.NewConfig()
	redisConfig := redis.NewConfig()

	// config logger
	slog.SetLogLoggerLevel(logging.ParseLogLevel(applicationConfig.LogLevel))

	// get database connection
	postgresConnection, err := postgres.NewConnection(ctx, postgresConfig)

	if err != nil {
		slog.Error("failed to connect to postgres", "error", err)

		return
	}

	defer func() {
		if _err := postgresConnection.Close(); _err != nil {
			slog.Error("failed to close postgres connection", "error", err)
		}
	}()

	// create redis connection
	redisConnection, err := redis.NewConnection(ctx, redisConfig)

	if err != nil {
		slog.Error("failed to connect to redis", "error", err)

		return
	}

	defer func() {
		if _err := redisConnection.Close(); _err != nil {
			slog.Error("failed to close redis connection", "error", err)
		}
	}()

	// create agent connection
	chatAdapter, err := agent.NewChatAdapter(context.Background())

	if err != nil {
		slog.Error("failed to connect to agent", "error", err)

		return
	}

	e := echo.New()

	if err != nil {
		slog.Error(fmt.Sprintf("failed to load static assets: %v", err))
	}

	// middlewares
	e.Use(middleware.RequestLogger())

	// routes
	// /api/v1/
	apiRouter := e.Group(fmt.Sprintf("%s/v%s", _http.APIPath, applicationConfig.ParseMajorVersion()))

	// /api/v1/auth/
	authenticationRouter := apiRouter.Group(authentication.Path)
	// /api/v1/auth/status
	authenticationRouter.GET(authentication.StatusEndpoint, authentication.NewGetAuthenticationStatusRoute(postgresConnection, redisConnection))

	// /api/v1/conversations/
	conversationsRouter := apiRouter.Group(conversation.Path)
	// /api/v1/conversations/chat
	conversationsRouter.POST(fmt.Sprintf("%s%s", conversation.Path, conversation.ChatEndpoint), conversation.NewPostChatRoute(chatAdapter, postgresConnection))
	// /api/v1/conversations/:conversation_id
	// /api/v1/conversations/:conversation_id/chat
	conversationsRouter.POST(fmt.Sprintf("%s/:conversation_id%s", conversation.Path, conversation.ChatEndpoint), conversation.NewPostChatRoute(chatAdapter, postgresConnection))

	// /api/v1/versions
	apiRouter.GET(versions.Path, versions.NewGetVersionsRoute(applicationConfig))

	// /api/v1/vessels/
	vesselsRouter := apiRouter.Group(vessel.Path)
	// /api/v1/vessels/specifications
	vesselsRouter.GET(_http.SpecificationsPath, vessel.NewGetVesselSpecificationsRoute(postgresConnection))

	// /app/*
	if applicationConfig.Environment != "development" {
		staticFiles := os.DirFS("/app/web/dist")

		e.GET(application.AppPath, application.NewGetAppRoute(staticFiles))
		e.GET(fmt.Sprintf("%s/*", application.AppPath), echo.WrapHandler(
			http.StripPrefix(fmt.Sprintf("%s/", application.AppPath), http.FileServer(http.FS(staticFiles))),
		))
	}

	// /healthcheck
	e.GET(application.HealthcheckPath, application.NewGetHealthcheckRoute())

	// set up server
	server := echo.StartConfig{
		Address:         fmt.Sprintf(":%s", applicationConfig.Port),
		HideBanner:      true,
		GracefulTimeout: 10 * time.Second,
	}

	// display a welcome message
	fmt.Printf(_http.WelcomeMessage, applicationConfig.Version, applicationConfig.Environment)

	// start the server
	if err = server.Start(ctx, e); err != nil {
		slog.Error("failed to start server", "error", err)

		return
	}
}
