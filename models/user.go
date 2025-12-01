package models

import "time"

type User struct {
	ID        string    `json:"id"`
	GoogleID  string    `json:"google_id"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	Phone     *string   `json:"phone,omitempty"`
	PhotoURL  *string   `json:"photo_url,omitempty"`
	CreatedAt time.Time `json:"created_at"`
}

type UpdateUserRequest struct {
	Name     string  `json:"name" binding:"required,min=2,max=50"`
	Phone    *string `json:"phone" binding:"omitempty"`
	PhotoURL *string `json:"photo_url" binding:"omitempty,url"`
}
