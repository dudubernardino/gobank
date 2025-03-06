package repositories

import (
	"context"

	"github.com/dudubernardino/gobank/internal/domain/account/entities"
	"github.com/dudubernardino/gobank/internal/domain/account/repositories"
	"github.com/dudubernardino/gobank/internal/infra/store/pgstore"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
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
		ID:            account.ID,
		Name:          account.Name,
		Email:         account.Email,
		TaxID:         account.TaxID,
		Balance:       account.Balance,
		MonthlyIncome: account.MonthlyIncome.Int64,
		AnnualRevenue: account.AnnualRevenue.Int64,
		CreatedAt:     account.CreatedAt.Time,
		UpdatedAt:     account.UpdatedAt,
	}, nil
}

func (repository *AccountRepositoryPostgres) Create(payload entities.Account) (uuid.UUID, error) {
	ctx := context.Background()
	id, err := repository.queries.CreateAccount(ctx, pgstore.CreateAccountParams{
		Name:          payload.Name,
		Email:         payload.Email,
		TaxID:         payload.TaxID,
		Balance:       payload.Balance,
		MonthlyIncome: pgtype.Int8{Int64: payload.MonthlyIncome, Valid: true},
		AnnualRevenue: pgtype.Int8{Int64: payload.AnnualRevenue, Valid: true},
	})

	if err != nil {
		return uuid.UUID{}, err
	}

	return id, nil
}

func (repository *AccountRepositoryPostgres) FindAccountBalanceById(id uuid.UUID) (int64, error) {
	ctx := context.Background()
	balance, err := repository.queries.GetAccountBalanceById(ctx, id)

	if err != nil {
		return 0, err
	}

	return balance, nil
}

func (repository *AccountRepositoryPostgres) AccountDeposit(id uuid.UUID, amount int64) (int64, error) {
	ctx := context.Background()

	balance, err := repository.queries.AccountDeposit(ctx, pgstore.AccountDepositParams{
		Balance: amount,
		ID:      id,
	})

	if err != nil {
		return 0, err
	}

	return balance, nil
}
