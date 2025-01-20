package authorizer

import (
	"stone/cards/authorizer/internal/domain/entities"

	"github.com/google/uuid"
)

const (
	riskAmountLimit      = 10_000.0
	riskAuthorizersLimit = 5
)

type AuthorizerRepository interface {
	InsertAuthorizer(authorizer entities.Authorizer) (uuid.UUID, error)
}

type RiskRepository interface {
	InsertRisk(risks entities.Risk)
}

type AuthorizerUC struct {
	authorizerRepo AuthorizerRepository
	riskRepo       RiskRepository
}

func (a AuthorizerUC) Authorize(authorize entities.Authorizer) (uuid.UUID, error) {
	// ... implements here

	return uuid.New(), nil
}

func NewAuthorizerUC(
	authRepo AuthorizerRepository,
	riskRepo RiskRepository,
) AuthorizerUC {
	return AuthorizerUC{
		authorizerRepo: authRepo,
		riskRepo:       riskRepo,
	}
}
