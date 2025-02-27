package repositories

import (
	"context"

	"github.com/dudubernardino/gobank/internal/domain/account/entities"
	"github.com/dudubernardino/gobank/internal/domain/account/repositories"
	"github.com/dudubernardino/gobank/internal/infra/store/pgstore"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type AccountRepositoryPostgres struct {
	pool    *pgxpool.Pool
	queries *pgstore.Queries
}

func NewAccountsRepository(pool *pgxpool.Pool) repositories.AccountRepository {
	return &AccountRepositoryPostgres{
		pool:    pool,
		queries: pgstore.New(pool),
	}
}

func (repository *AccountRepositoryPostgres) FindById(id uuid.UUID) (entities.Account, error) {
	ctx := context.Background()
	account, err := repository.queries.GetAccountById(ctx, id)

	if err != nil {
		return entities.Account{}, err
	}

	return entities.Account{
		ID:        account.ID,
		Name:      account.Name,
		TaxID:     account.TaxID,
		Balance:   account.Balance,
		CreatedAt: account.CreatedAt.Time,
	}, nil
}
