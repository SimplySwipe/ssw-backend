package models

import "time"

type RefreshToken struct {
	ID        string
	UserID    string
	Token     string
	ExpiresAt time.Time
	Revoked   bool
	Used      bool
	CreatedAt time.Time
}
