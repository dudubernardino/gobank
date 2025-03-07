-- name: CreateAccount :one
INSERT INTO accounts (tax_id, name, monthly_income, annual_revenue, email, balance) 
VALUES ($1, $2, $3, $4, $5, $6) RETURNING id;

-- name: GetAccountById :one
SELECT * FROM accounts 
WHERE id = $1;

-- name: GetAccountBalanceById :one
SELECT balance FROM accounts 
WHERE id = $1;

-- name: AccountDeposit :one
UPDATE accounts SET balance = balance + $1 
WHERE id = $2 RETURNING balance;

-- name: AccountWithdraw :one
UPDATE accounts SET balance = balance - $1 
WHERE id = $2 RETURNING balance;

-- name: Delete :one
UPDATE accounts SET deleted_at = CURRENT_TIMESTAMP
WHERE id = $1 RETURNING id;