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
	"github.com/uptrace/bun"

	"github.com/kieranroneill/oarbit/internal/platform/application"
	"github.com/kieranroneill/oarbit/internal/platform/postgres"
)

type Scope struct {
	controller         *Controller
	postgresConnection *bun.DB
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

	scope = &Scope{
		controller:         NewController(application.NewConfig(), postgresConnection),
		postgresConnection: postgresConnection,
	}
	// run tests with context
	exitCode := m.Run()

	_ = postgresConnection.Close()

	// clean up
	cancel()

	os.Exit(exitCode)
}

func TestSetupRequiredNoUsers(t *testing.T) {
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
}
