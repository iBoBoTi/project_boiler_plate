package domain

// User represents users of the application
type User struct {
	ID           string `json:"id"`
	RoleID       string `json:"role_id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Email        string `json:"email"`
	Password     string `json:"-"`
	HashPassword []byte `json:"hash_password"`
	IsActivated  bool   `json:"activated"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}
