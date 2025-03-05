package jsonutils

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func EncodeJson[T any](w http.ResponseWriter, r *http.Request, statusCode int, data T) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		return fmt.Errorf("failed to encode json %w", err)
	}

	return nil
}

func DecodeJson[T any](r *http.Request) (T, []map[string]string, error) {
	var data T
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		return data, nil, fmt.Errorf("decode json failed: %w", err)
	}

	validate := validator.New()
	err := validate.Struct(data)

	if err != nil {
		errors := err.(validator.ValidationErrors)
		formattedErrors := make([]map[string]string, 0, len(errors))

		for _, e := range errors {
			formattedErrors = append(formattedErrors, map[string]string{
				"field":   e.Field(),
				"message": e.Tag(),
			})
		}

		return data, formattedErrors, nil
	}

	return data, nil, nil
}
