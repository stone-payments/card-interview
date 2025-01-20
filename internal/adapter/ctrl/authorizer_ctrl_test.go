package ctrl

import (
	"encoding/json"
	"stone/cards/authorizer/internal/adapter/ctrl/schema"
	"stone/cards/authorizer/internal/adapter/db"
	"stone/cards/authorizer/internal/domain/authorizer"
	"stone/cards/authorizer/internal/domain/entities"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAuthorizerCtrlEx1(t *testing.T) {
	tests := []struct {
		name     string
		payload  string
		init     func() AuthorizerCtrl
		assertFn func(t *testing.T, got schema.AuthorizerResponse)
	}{
		{
			name:    "Test 1",
			payload: `{"card_number":"4111111111111111","amount": 100.50,"currency": "USD","merchant":"Amazon","timestamp":"2025-01-17T10:00:00Z"}`,
			init: func() AuthorizerCtrl {
				return NewAuthorizerCtrl(
					authorizer.NewAuthorizerUC(
						db.NewAuthorizerRepository(),
						db.NewRiskRepository(),
					),
				)
			},
			assertFn: func(t *testing.T, got schema.AuthorizerResponse) {
				assert.Equal(t, "approved", got.Status)
				assert.Len(t, got.AuthorizeID, 36)
				// is uuid
				assert.Regexp(t, `^[\d\w]{8}-[\d\w]{4}-[\d\w]{4}-[\d\w]{4}-[\d\w]{12}$`, got.AuthorizeID)

				bt, err := json.Marshal(got)
				assert.NoError(t, err)

				assert.NotContains(t, string(bt), "errors")
			},
		},
		{
			name:    "Test 2",
			payload: `{"card_number":"4111111111111111","amount": 100.50,"currency": "USD","merchant":"Amazon","timestamp":"2026-01-17T10:00:00Z"}`,
			init: func() AuthorizerCtrl {
				return NewAuthorizerCtrl(
					authorizer.NewAuthorizerUC(
						db.NewAuthorizerRepository(),
						db.NewRiskRepository(),
					),
				)
			},
			assertFn: func(t *testing.T, got schema.AuthorizerResponse) {
				assert.Equal(t, "rejected", got.Status)
				assert.Equal(t, "timestamp on future", got.Error)
				bt, err := json.Marshal(got)
				assert.NoError(t, err)

				assert.NotContains(t, string(bt), "authorize_id")
			},
		},
		{
			name:    "Test 3",
			payload: `{"card_number":"4111111111111111","amount": 100.50,"currency": "USD","merchant":"Amazon","timestamp":"Mon, 02 Jan 2006 15:04:05 MST"}`,
			init: func() AuthorizerCtrl {
				return NewAuthorizerCtrl(
					authorizer.NewAuthorizerUC(
						db.NewAuthorizerRepository(),
						db.NewRiskRepository(),
					),
				)
			},
			assertFn: func(t *testing.T, got schema.AuthorizerResponse) {
				assert.Equal(t, "rejected", got.Status)
				assert.Equal(t, "timestamp not valid", got.Error)
				bt, err := json.Marshal(got)
				assert.NoError(t, err)

				assert.NotContains(t, string(bt), "authorize_id")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.assertFn(t, tt.init().Authorize([]byte(tt.payload)))
		})
	}

}

func TestAuthorizerCtrlEx2(t *testing.T) {
	tests := []struct {
		name     string
		payload  string
		init     func() AuthorizerCtrl
		assertFn func(t *testing.T, got schema.AuthorizerResponse)
	}{
		{
			name:    "Test 1",
			payload: `{"card_number":"4111111111111111","amount": 100.50,"currency": "USD","merchant":"Amazon","timestamp":"2025-01-17T10:00:00Z"}`,
			init: func() AuthorizerCtrl {
				return NewAuthorizerCtrl(
					authorizer.NewAuthorizerUC(
						db.NewAuthorizerRepository(),
						db.NewRiskRepository(),
					),
				)
			},
			assertFn: func(t *testing.T, got schema.AuthorizerResponse) {
				assert.Equal(t, "approved", got.Status)
				assert.Len(t, got.AuthorizeID, 36)
				// is uuid
				assert.Regexp(t, `^[\d\w]{8}-[\d\w]{4}-[\d\w]{4}-[\d\w]{4}-[\d\w]{12}$`, got.AuthorizeID)

				bt, err := json.Marshal(got)
				assert.NoError(t, err)

				assert.NotContains(t, string(bt), "errors")
				assert.NotContains(t, string(bt), "warning")
			},
		},
		{
			name:    "Test 2",
			payload: `{"card_number":"4111111111111111","amount": 100.50,"currency": "USD","merchant":"Amazon","timestamp":"2026-01-17T10:00:00Z"}`,
			init: func() AuthorizerCtrl {
				return NewAuthorizerCtrl(
					authorizer.NewAuthorizerUC(
						db.NewAuthorizerRepository(),
						db.NewRiskRepository(),
					),
				)
			},
			assertFn: func(t *testing.T, got schema.AuthorizerResponse) {
				assert.Equal(t, "rejected", got.Status)
				assert.Equal(t, "timestamp on future", got.Error)
				bt, err := json.Marshal(got)
				assert.NoError(t, err)

				assert.NotContains(t, string(bt), "authorize_id")
				assert.NotContains(t, string(bt), "warning")
			},
		},
		{
			name:    "Test 3",
			payload: `{"card_number":"4111111111111111","amount": 100.50,"currency": "USD","merchant":"Amazon","timestamp":"Mon, 02 Jan 2006 15:04:05 MST"}`,
			init: func() AuthorizerCtrl {
				return NewAuthorizerCtrl(
					authorizer.NewAuthorizerUC(
						db.NewAuthorizerRepository(),
						db.NewRiskRepository(),
					),
				)
			},
			assertFn: func(t *testing.T, got schema.AuthorizerResponse) {
				assert.Equal(t, "rejected", got.Status)
				assert.Equal(t, "timestamp not valid", got.Error)
				bt, err := json.Marshal(got)
				assert.NoError(t, err)

				assert.NotContains(t, string(bt), "authorize_id")
				assert.NotContains(t, string(bt), "warning")
			},
		},
		{
			name:    "Test 4",
			payload: `{"card_number":"4111111111111111","amount": 100.50,"currency": "USD","merchant":"Amazon","timestamp":"2025-01-17T10:00:00Z"}`,
			init: func() AuthorizerCtrl {
				authRepo := db.NewAuthorizerRepository()
				for i := 0; i < 10; i++ {
					authRepo.InsertAuthorizer(entities.Authorizer{
						CardNumber: "4111111111111111",
						Amount:     100.50,
						Currency:   "USD",
						Merchant:   "Amazon",
						Timestamp:  time.Now().Add(-30 * time.Second),
					})
				}

				return NewAuthorizerCtrl(
					authorizer.NewAuthorizerUC(
						authRepo,
						db.NewRiskRepository(),
					),
				)
			},
			assertFn: func(t *testing.T, got schema.AuthorizerResponse) {
				assert.Equal(t, "approved_with_warning", got.Status)
				assert.Equal(t, "transaction marked as suspicious: not standard", got.Warning)
				assert.Len(t, got.AuthorizeID, 36)
				// is uuid
				assert.Regexp(t, `^[\d\w]{8}-[\d\w]{4}-[\d\w]{4}-[\d\w]{4}-[\d\w]{12}$`, got.AuthorizeID)
				bt, err := json.Marshal(got)
				assert.NoError(t, err)

				assert.NotContains(t, string(bt), "errors")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.assertFn(t, tt.init().Authorize([]byte(tt.payload)))
		})
	}

}
