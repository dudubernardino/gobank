package account

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dudubernardino/gobank/internal/domain/account/entities"
	"github.com/dudubernardino/gobank/internal/infra/repositories"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestHandleAccountWithdraw(t *testing.T) {
	mockAccountRepository := repositories.NewMockAccountRepository([]entities.Account{
		{
			ID:      uuid.New(),
			Name:    "Test Account",
			Balance: 1000,
		},
	})

	accountID := mockAccountRepository.Accounts[0].ID.String()

	payload := map[string]any{
		"amount": 50,
	}

	body, err := json.Marshal(payload)
	assert.NoError(t, err)

	fmt.Println("Testing with URL:", fmt.Sprintf("/api/v1/accounts/%s/withdraw", accountID))
	req := httptest.NewRequest("POST", "/api/v1/accounts/{account_id}/withdraw", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("account_id", accountID)

	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	handler := HandleAccountWithdraw(mockAccountRepository)
	handler.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code, "Expected status 200 OK")

	var response map[string]any
	err = json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(t, err, "Error unmarshalling response JSON")

	expectedMessage := float64(950)
	assert.Equal(t, expectedMessage, response["balance"], "Expected response balance to be: 950")
}
