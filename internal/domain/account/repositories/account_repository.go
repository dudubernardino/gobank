package repositories

import (
	"github.com/dudubernardino/gobank/internal/domain/account/entities"
	"github.com/google/uuid"
)

type AccountRepository interface {
	FindById(id uuid.UUID) (entities.Account, error)
}
