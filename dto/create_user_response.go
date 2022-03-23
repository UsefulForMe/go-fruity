package dto

import "github.com/google/uuid"

type CreateUserResponse struct {
	ID uuid.UUID `json:"id"`
}
