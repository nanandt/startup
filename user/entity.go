package user

import "time"

// User is ...
type User struct {
	ID             int
	Name           string
	Occupation     string
	Email          string
	PasswordHash   string
	AvatarFileName string
	Role           string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}