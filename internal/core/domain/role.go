package domain

import "time"

// Role determines the role a user plays in the application
type Role struct {
	ID          string    `json:"id"`
	Title       string    `json:"title" validate:"required"`
	Description string    `json:"description" validate:"required"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
