package db

import (
	"stone/cards/authorizer/internal/domain/entities"
)

type RiskRepository struct {
}

func (r *RiskRepository) InsertRisk(risk entities.Risk) {
	// ... implements here
}

func NewRiskRepository() *RiskRepository {
	return &RiskRepository{}
}
