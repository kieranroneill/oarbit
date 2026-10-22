package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/kieranroneill/oarbit/internal/platform/http"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"

	"github.com/kieranroneill/oarbit/internal/boat"
	"github.com/kieranroneill/oarbit/internal/conversation"
	"github.com/kieranroneill/oarbit/internal/platform/agent"
	"github.com/kieranroneill/oarbit/internal/platform/application"
	"github.com/kieranroneill/oarbit/internal/platform/logging"
	"github.com/kieranroneill/oarbit/internal/platform/postgres"
	"github.com/kieranroneill/oarbit/internal/user"
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

	// create agent connection
	chatAdapter, err := agent.NewChatAdapter(context.Background())
	if err != nil {
		slog.Error("failed to connect to agent", "error", err)

		return
	}

	e := echo.New()

	// middlewares
	e.Use(middleware.RequestLogger())

	// routes
	// /api/
	apiRouter := e.Group(http.APIPath)
	// /api/boats/
	boatsRouter := apiRouter.Group(boat.Path)
	// /api/boats/specifications
	boatsRouter.GET(http.SpecificationsPath, boat.NewGetBoatSpecificationsRoute(database))
	// /api/users/
	usersRouter := apiRouter.Group(user.Path)
	// /api/users/:user_id
	// /api/users/:user_id/conversations
	// /api/users/:user_id/conversations/chat
	usersRouter.POST(fmt.Sprintf("/:user_id%s%s", conversation.Path, conversation.ChatEndpoint), conversation.NewPostChatRoute(chatAdapter, database))
	// /api/users/:user_id/conversations/:conversation_id
	// /api/users/:user_id/conversations/:conversation_id/chat
	usersRouter.POST(fmt.Sprintf("/:user_id%s/:conversation_id%s", conversation.Path, conversation.ChatEndpoint), conversation.NewPostChatRoute(chatAdapter, database))
	// /api/versions
	apiRouter.GET(versions.Path, versions.NewGetVersionsRoute(applicationConfig))

	// set up server
	server := echo.StartConfig{
		Address:         fmt.Sprintf(":%s", applicationConfig.Port),
		HideBanner:      true,
		GracefulTimeout: 10 * time.Second,
	}

	// display a welcome message
	fmt.Printf(http.WelcomeMessage, applicationConfig.Version, applicationConfig.Environment)

	// start the server
	if err = server.Start(ctx, e); err != nil {
		slog.Error("failed to start server", "error", err)

		return
	}
}
