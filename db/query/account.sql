-- name: CreateAccount :one
INSERT INTO account (
  owner, balance, currency
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: ListAccount :many
SELECT * FROM account
ORDER BY owner;

-- name: ListAccountLimit :many
SELECT * FROM account
ORDER BY owner LIMIT $1;