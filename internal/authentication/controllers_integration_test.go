//go:build integration

package authentication

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/labstack/echo/v5"
	"github.com/redis/go-redis/v9"
	"github.com/uptrace/bun"

	"github.com/kieranroneill/oarbit/internal/platform/application"
	"github.com/kieranroneill/oarbit/internal/platform/postgres"
	_redis "github.com/kieranroneill/oarbit/internal/platform/redis"
)

type Scope struct {
	controller         *Controller
	postgresConnection *bun.DB
	redisClient        *redis.Client
}

var scope *Scope

func TestMain(m *testing.M) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)

	defer cancel()

	postgresConnection, err := postgres.NewConnection(ctx, postgres.NewConfig())

	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to connect to postgres: %v\n", err)

		os.Exit(1)
	}

	redisConnection, err := _redis.NewConnection(ctx, _redis.NewConfig())

	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to connect to redis: %v\n", err)

		os.Exit(1)
	}

	scope = &Scope{
		controller:         NewController(application.NewConfig(), postgresConnection),
		postgresConnection: postgresConnection,
		redisClient:        redisConnection,
	}
	// run tests with context
	exitCode := m.Run()

	_ = postgresConnection.Close()
	_ = redisConnection.Close()

	// clean up
	cancel()

	os.Exit(exitCode)
}

func TestAuthenticate(t *testing.T) {
	t.Run("no session cookie", func(t *testing.T) {
		request := httptest.NewRequestWithContext(
			t.Context(),
			http.MethodGet,
			"/",
			nil,
		)
		response := httptest.NewRecorder()
		c := echo.New().NewContext(request, response)

		user, err := scope.controller.Authenticate(c, scope.postgresConnection, scope.redisClient)

		if user != nil {
			t.Error("expected user to be nil")
		}

		if err != nil {
			t.Error("expected no error, received: ", err)
		}
	})
}

func TestSetupRequired(t *testing.T) {
	t.Run("no users", func(t *testing.T) {
		var body StatusResponse

		request := httptest.NewRequestWithContext(
			t.Context(),
			http.MethodGet,
			"/",
			nil,
		)
		response := httptest.NewRecorder()
		c := echo.New().NewContext(request, response)

		err := scope.controller.SetupRequired(c, nil)

		if err != nil {
			t.Fatal(err)
		}

		if response.Code != http.StatusOK {
			t.Fatalf("expected status %d, actual status %d", http.StatusOK, response.Code)
		}

		if err = json.Unmarshal(response.Body.Bytes(), &body); err != nil {
			t.Fatal(err)
		}

		if !body.SetupRequired {
			t.Error(`expected "setupRequired" to be true`)
		}

		if body.User != nil {
			t.Errorf("expected no user, got %+v", body.User)
		}
	})
}
