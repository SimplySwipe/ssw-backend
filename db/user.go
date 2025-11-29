package db

import (
	"SimplySwipe/models"
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func GetUserByID(ctx context.Context, db *pgxpool.Pool, userID string) (*models.User, error) {
	const query = `
    SELECT id, google_id, email, name, phone, photo_url, created_at
    FROM users
    WHERE id = $1
	`
	var user models.User
	err := db.QueryRow(ctx, query, userID).Scan(
		&user.ID,
		&user.GoogleID,
		&user.Email,
		&user.Name,
		&user.Phone,
		&user.PhotoURL,
		&user.CreatedAt,
	)
	if err == pgx.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func GetOrCreateUserByGoogleID(ctx context.Context, db *pgxpool.Pool, googleID, email, name string) (*models.User, error) {
	const query = `
	SELECT id, google_id, email, name, phone, photo_url, created_at
    FROM users
    WHERE google_id = $1
	`
	var user models.User
	err := db.QueryRow(ctx, query, googleID).Scan(
		&user.ID,
		&user.GoogleID,
		&user.Email,
		&user.Name,
		&user.Phone,
		&user.PhotoURL,
		&user.CreatedAt,
	)
	if err == pgx.ErrNoRows {
		const queryIn = `
		INSERT INTO users (google_id, email, name) 
		VALUES ($1, $2, $3)
		RETURNING id, google_id, email, name, phone, photo_url, created_at
`
		errIn := db.QueryRow(ctx, queryIn, googleID, email, name).Scan(
			&user.ID,
			&user.GoogleID,
			&user.Email,
			&user.Name,
			&user.Phone,
			&user.PhotoURL,
			&user.CreatedAt,
		)
		if errIn != nil {
			return nil, errIn
		}
		return &user, nil
	}
	if err != nil {
		return nil, err
	}
	return &user, nil
}
