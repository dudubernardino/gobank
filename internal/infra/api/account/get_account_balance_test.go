package account

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/dudubernardino/gobank/internal/domain/account/entities"
	"github.com/dudubernardino/gobank/internal/infra/repositories"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestHandleGetAccountBalance(t *testing.T) {
	mockAccountRepository := repositories.NewMockAccountRepository([]entities.Account{
		{
			ID:      uuid.New(),
			Name:    "Test Account",
			Balance: 1000,
		},
	})

	accountID := mockAccountRepository.Accounts[0].ID.String()

	fmt.Println("Testing with URL:", fmt.Sprintf("/api/v1/accounts/%s/balance", accountID))
	req := httptest.NewRequest("GET", "/api/v1/accounts/{account_id}/balance", nil)
	rec := httptest.NewRecorder()

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("account_id", accountID)

	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	handler := HandleGetAccountById(mockAccountRepository)
	handler.ServeHTTP(rec, req)

	assert.Equal(t, 200, rec.Code, "Expected status 200")

	var response entities.Account
	err := json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(t, err, "Error unmarshalling response JSON")

	assert.Equal(t, mockAccountRepository.Accounts[0].Balance, response.Balance, "Expected balance to be 1000")
}
