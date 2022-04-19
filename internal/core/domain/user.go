package domain

// User represents users of the application
type User struct {
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Email        string `json:"email"`
	Password     string `json:"-"`
	HashPassword []byte `json:"hash_password"`
	Activated    bool   `json:"activated"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}
