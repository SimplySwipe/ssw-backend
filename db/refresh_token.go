package db

import (
	"SimplySwipe/models"
	"context"
	"crypto/rand"
	"encoding/base64"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func GenerateRefreshToken() (string, error) {
	slice := make([]byte, 32)
	_, err := rand.Read(slice)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(slice), nil

}

func InsertRefreshToken(ctx context.Context, db *pgxpool.Pool, userID, token string, ExpiresAt time.Time) (*models.RefreshToken, error) {
	const query = `
	INSERT INTO refresh_tokens (user_id, token, expires_at)
	VALUES ($1, $2, $3)
	RETURNING id, user_id, token, expires_at, revoked, used, created_at
	`
	var refreshToken models.RefreshToken
	err := db.QueryRow(ctx, query, userID, token, ExpiresAt).Scan(
		&refreshToken.ID,
		&refreshToken.UserID,
		&refreshToken.Token,
		&refreshToken.ExpiresAt,
		&refreshToken.Revoked,
		&refreshToken.Used,
		&refreshToken.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &refreshToken, nil

}

func GetRefreshToken(ctx context.Context, db *pgxpool.Pool, token string) (*models.RefreshToken, error) {
	const query = `
	SELECT * 
	FROM refresh_tokens
	WHERE token = $1
	`
	var refreshToken models.RefreshToken
	err := db.QueryRow(ctx, query, token).Scan(
		&refreshToken.ID,
		&refreshToken.UserID,
		&refreshToken.Token,
		&refreshToken.ExpiresAt,
		&refreshToken.Revoked,
		&refreshToken.Used,
		&refreshToken.CreatedAt,
	)
	if err == pgx.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &refreshToken, nil
}
