package ctrl

import (
	"encoding/json"
	"stone/cards/authorizer/internal/adapter/ctrl/schema"
	"stone/cards/authorizer/internal/domain/entities"

	"github.com/google/uuid"
)

type AuthorizerUseCase interface {
	Authorize(authorizer entities.Authorizer) (uuid.UUID, error)
}

type AuthorizerCtrl struct {
	authorizerUC AuthorizerUseCase
}

func (a AuthorizerCtrl) Authorize(payload json.RawMessage) schema.AuthorizerResponse {
	// ... implements here
	return schema.AuthorizerResponse{}
}

func NewAuthorizerCtrl(authorizerUC AuthorizerUseCase) AuthorizerCtrl {
	return AuthorizerCtrl{
		authorizerUC: authorizerUC,
	}
}
