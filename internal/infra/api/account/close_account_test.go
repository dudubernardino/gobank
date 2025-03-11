package account

import (
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

func TestHandleCloseAccount(t *testing.T) {
	mockAccountRepository := repositories.NewMockAccountRepository([]entities.Account{
		{
			ID:      uuid.New(),
			Name:    "Test Account",
			Balance: 1000,
		},
	})

	accountID := mockAccountRepository.Accounts[0].ID.String()

	fmt.Println("Testing with URL:", fmt.Sprintf("/api/v1/accounts/%s", accountID))
	req := httptest.NewRequest("DELETE", "/api/v1/accounts/{account_id}", nil)
	rec := httptest.NewRecorder()

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("account_id", accountID)

	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	handler := HandleCloseAccount(mockAccountRepository)
	handler.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code, "Expected status 200 OK")

	var response map[string]any
	err := json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(t, err, "Error unmarshalling response JSON")

	expectedMessage := "account closed successfully"
	assert.Equal(t, expectedMessage, response["message"], "Expected response message to be: "+expectedMessage)
}
