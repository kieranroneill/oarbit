package authentication

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

//==========================================
// session repository
//==========================================

type SessionRepository struct {
	redisClient *redis.Client
}

/**
 * constructor functions
 */

func NewSessionRepository(
	redisClient *redis.Client,
) *SessionRepository {
	return &SessionRepository{
		redisClient: redisClient,
	}
}

/**
 * private functions
 */

func (r *SessionRepository) key(key string) string {
	return fmt.Sprintf("session:%s", key)
}

/**
 * public functions
 */

func (r *SessionRepository) Create(
	ctx context.Context,
	userId uuid.UUID,
) (*Session, error) {
	now := time.Now()
	dto := Session{
		ExpiresAt: now.Add(SessionTTLInSeconds).UTC(),
		ID:        uuid.New(),
		IssuedAt:  now,
		UserID:    userId,
	}
	key := r.key(dto.ID.String())
	fields := map[string]interface{}{
		"expiresAt": dto.ExpiresAt.UTC().Format(time.RFC3339),
		"id":        dto.ID.String(),
		"issuedAt":  dto.IssuedAt.UTC().Format(time.RFC3339),
		"userId":    dto.UserID.String(),
	}

	if err := r.redisClient.HSet(ctx, key, fields).Err(); err != nil {
		return nil, fmt.Errorf("failed to create session: %w", err)
	}

	if err := r.redisClient.Expire(ctx, key, SessionTTLInSeconds).Err(); err != nil {
		// rollback if the expiration fails
		_ = r.redisClient.Del(ctx, key).Err()

		return nil, fmt.Errorf("failed to set session expiration: %w", err)
	}

	return &dto, nil
}

func (r *SessionRepository) DeleteByID(
	ctx context.Context,
	id uuid.UUID,
) error {
	if id == uuid.Nil {
		slog.Debug("failed to delete session, empty id")

		return nil
	}

	deleted, err := r.redisClient.Del(ctx, r.key(id.String())).Result()
	if err != nil {
		return fmt.Errorf("failed to delete session: %w", err)
	}

	if deleted == 0 {
		slog.Debug("session not found", "id", id)

		return nil
	}

	slog.Debug("deleted session", "id", id)

	return nil
}

func (r *SessionRepository) FindByID(ctx context.Context, id uuid.UUID) (*Session, error) {
	if id == uuid.Nil {
		return nil, nil
	}

	values, err := r.redisClient.HGetAll(ctx, r.key(id.String())).Result()

	if err != nil {
		return nil, fmt.Errorf("failed to find session: %w", err)
	}

	if len(values) == 0 {
		return nil, nil
	}

	expiresAt, err := time.Parse(time.RFC3339, values["expiresAt"])

	if err != nil {
		return nil, fmt.Errorf(
			`failed to parse session "expiresAt": %w`,
			err,
		)
	}

	issuedAt, err := time.Parse(time.RFC3339, values["issuedAt"])

	if err != nil {
		return nil, fmt.Errorf(
			`failed to parse session "issuedAt": %w`,
			err,
		)
	}

	userId, err := uuid.Parse(values["userId"])
	if err != nil {
		return nil, fmt.Errorf(
			`failed to parse session "userId": %w`,
			err,
		)
	}

	return &Session{
		ExpiresAt: expiresAt,
		ID:        id,
		IssuedAt:  issuedAt,
		UserID:    userId,
	}, nil
}

func (r *SessionRepository) SlideTimeoutByID(ctx context.Context, id uuid.UUID) (*Session, error) {
	if id == uuid.Nil {
		slog.Debug("failed to update session, empty id")

		return nil, nil
	}

	dto, err := r.FindByID(ctx, id)

	if err != nil {
		return nil, err
	}

	if dto == nil {
		slog.Debug("failed to find session", "id", id.String())

		return nil, nil
	}

	dto.ExpiresAt = time.Now().Add(SessionTTLInSeconds)
	key := r.key(dto.ID.String())
	fields := map[string]interface{}{
		"expiresAt": dto.ExpiresAt.UTC().Format(time.RFC3339),
		"id":        dto.ID.String(),
		"issuedAt":  dto.IssuedAt.UTC().Format(time.RFC3339),
		"userId":    dto.UserID.String(),
	}

	if err = r.redisClient.HSet(ctx, key, fields).Err(); err != nil {
		return nil, fmt.Errorf("failed to update the session: %w", err)
	}

	if err = r.redisClient.Expire(ctx, key, SessionTTLInSeconds).Err(); err != nil {
		// rollback if the expiration fails
		_ = r.redisClient.Del(ctx, key).Err()

		return nil, fmt.Errorf("failed to set session expiration: %w", err)
	}

	return dto, nil
}
