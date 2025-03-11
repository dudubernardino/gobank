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
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestHandleAccountAmountTransfer(t *testing.T) {
	mockAccountRepository := repositories.NewMockAccountRepository([]entities.Account{
		{
			ID:      uuid.New(),
			Name:    "Test Account 1",
			Balance: 1000,
		},
		{
			ID:      uuid.New(),
			Name:    "Test Account 2",
			Balance: 1000,
		},
	})

	accountOriginID := mockAccountRepository.Accounts[0].ID.String()
	accountDestinationID := mockAccountRepository.Accounts[1].ID.String()

	payload := map[string]any{
		"account_origin_id":      accountOriginID,
		"account_destination_id": accountDestinationID,
		"amount":                 50,
	}

	body, err := json.Marshal(payload)
	assert.NoError(t, err)

	fmt.Println("Testing with URL: /api/v1/accounts/transfer")
	req := httptest.NewRequest("POST", "/api/v1/accounts/transfer", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	handler := HandleAccountAmountTransfer(mockAccountRepository)
	handler.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code, "Expected status 200 OK")

	var response map[string]any
	err = json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(t, err, "Error unmarshalling response JSON")

	expectedMessage := "amount transfered successfully"
	assert.Equal(t, expectedMessage, response["message"], "Expected response message to be: "+expectedMessage)

	assert.Equal(t, mockAccountRepository.Accounts[0].Balance, int64(950), "Expected account origin balance to be 950")
	assert.Equal(t, mockAccountRepository.Accounts[1].Balance, int64(1050), "Expected account destination balance to be 1050")
}
