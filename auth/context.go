package auth

import "github.com/google/uuid"

const (
	jwtTokenKey = "jwtToken"
	uidKey      = "uid"
)

type UID struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
}
