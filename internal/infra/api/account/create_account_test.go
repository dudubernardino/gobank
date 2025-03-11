package account

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dudubernardino/gobank/internal/domain/account/entities"
	"github.com/dudubernardino/gobank/internal/infra/repositories"
	"github.com/stretchr/testify/assert"
)

func TestHandleCreateAccount(t *testing.T) {
	mockAccountRepository := repositories.NewMockAccountRepository([]entities.Account{})

	accountData := map[string]any{
		"tax_id":         "74553007006",
		"name":           "John Doe",
		"email":          "email@email.com",
		"monthly_income": 1000000,
		"annual_revenue": 12000000,
		"balance":        1000000,
	}

	body, err := json.Marshal(accountData)
	assert.NoError(t, err)

	fmt.Println("Testing with URL: /api/v1/accounts")
	req := httptest.NewRequest("POST", "/api/v1/accounts", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	handler := HandleCreateAccount(mockAccountRepository)
	handler.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusCreated, rec.Code, "Expected status 201 Created")
}
