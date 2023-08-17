-- name: CreateTransfer :one
INSERT INTO transfers (
    from_account_id,
    to_account_id,
    amount
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: GetTransfer :one
SELECT * FROM transfers
WHERE id = $1 LIMIT 1;

-- name: ListTransfers :many
SELECT * FROM transfers
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateTransfer :one
UPDATE transfers
SET amount = $2
WHERE id = $1
RETURNING *;

-- name: DeleteTransfer :exec
DELETE FROM transfers
WHERE id = $1;

-- name: ListTransfersByAccount :many
SELECT * FROM transfers
WHERE from_account_id = $1
ORDER BY id
LIMIT $2
OFFSET $3;

-- name: ListTransfersByAccountAndTime :many
SELECT * FROM transfers
WHERE from_account_id = $1
AND created_at >= $2
AND created_at <= $3
ORDER BY id
LIMIT $4
OFFSET $5;

-- name: ListTransfersToAccount :many
SELECT * FROM transfers
WHERE to_account_id = $1
ORDER BY id
LIMIT $2
OFFSET $3;


-- name: ListTransfersToAccountAndTime :many
SELECT * FROM transfers
WHERE to_account_id = $1
  AND created_at >= $2
  AND created_at <= $3
ORDER BY id
    LIMIT $4
OFFSET $5;

-- name: ListTransfersByAccountAndToAccount :many
SELECT * FROM transfers
WHERE from_account_id = $1
AND to_account_id = $2
ORDER BY id
LIMIT $3
OFFSET $4;

-- name: ListTransfersByAccountAndToAccountAndTime :many
SELECT * FROM transfers
WHERE from_account_id = $1
AND to_account_id = $2
AND created_at >= $3
AND created_at <= $4
ORDER BY id
LIMIT $5
OFFSET $6;