package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"net/url"
	"time"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

func NewConnection(ctx context.Context, config Config) (*bun.DB, error) {
	dsn := (&url.URL{
		Scheme: "postgres",
		User:   url.UserPassword(config.User, config.Password),
		Host:   fmt.Sprintf("%s:%s", config.Host, config.Port),
		Path:   "/" + config.Name,
		RawQuery: url.Values{
			"sslmode": []string{config.SSLMode},
		}.Encode(),
	}).String()

	sqlDB := sql.OpenDB(
		pgdriver.NewConnector(
			pgdriver.WithDSN(dsn),
		),
	)

	// database/sql pool configuration.
	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetMaxIdleConns(2)
	sqlDB.SetConnMaxLifetime(time.Hour)
	sqlDB.SetConnMaxIdleTime(30 * time.Minute)

	pingCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if err := sqlDB.PingContext(pingCtx); err != nil {
		sqlDB.Close()
		return nil, fmt.Errorf("ping postgresql: %w", err)
	}

	db := bun.NewDB(sqlDB, pgdialect.New())

	return db, nil
}
