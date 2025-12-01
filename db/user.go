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

func GetOrCreateUserByGoogleID(ctx context.Context, db *pgxpool.Pool, googleID, email, name, photoURL string) (*models.User, error) {
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
		INSERT INTO users (google_id, email, name, photo_url) 
		VALUES ($1, $2, $3, $4)
		RETURNING id, google_id, email, name, phone, photo_url, created_at
`
		errIn := db.QueryRow(ctx, queryIn, googleID, email, name, photoURL).Scan(
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

func UpdateUser(ctx context.Context, db *pgxpool.Pool, userID, name string, phone, photoURL *string) (*models.User, error) {
	const query = `
	UPDATE users
	SET name = $1, phone = $2, photo_url = $3
	WHERE id = $4
	RETURNING id, google_id, email, name, phone, photo_url, created_at
	`
	var user models.User
	err := db.QueryRow(ctx, query, name, phone, photoURL, userID).Scan(
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
