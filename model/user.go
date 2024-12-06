package model

import (
	"travelnest/enums"

	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID    `json:"id"`
	Username string    `json:"username"`
	Password string    `json:"password"`
	Email    string    `json:"email"`
	Role 	 	 enums.UserRole `json:"role"`
}