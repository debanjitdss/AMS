package models

type User struct {
	ID       int64  `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password,omitempty"` // Omit password from JSON response
	Username string `json:"username"`
	IsActive bool   `json:"isActive"`
	GoogleID string `json:"googleId,omitempty"`
}
