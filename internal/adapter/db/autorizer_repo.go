package db

import (
	"stone/cards/authorizer/internal/domain/entities"

	"github.com/google/uuid"
)

type AuthorizerRepository struct {
}

func (r *AuthorizerRepository) InsertAuthorizer(authorizer entities.Authorizer) (uuid.UUID, error) {
	// ... implements here
	return uuid.New(), nil
}

func NewAuthorizerRepository() *AuthorizerRepository {
	return &AuthorizerRepository{}
}
