-- name: CreateQuotation :one
INSERT INTO quotations (
    customer_id,
    vehicle_id,
    status,
    total_amount
)
VALUES ($1, $2, $3, $4)
RETURNING id, customer_id, vehicle_id, status, total_amount, created_at;


-- name: ListQuotations :many
SELECT id, customer_id, vehicle_id, status, total_amount, created_at
FROM quotations
ORDER BY id;


-- name: GetQuotation :one
SELECT id, customer_id, vehicle_id, status, total_amount, created_at
FROM quotations
WHERE id = $1;


-- name: UpdateQuotationStatus :one
UPDATE quotations
SET status = $1
WHERE id = $2
RETURNING id, customer_id, vehicle_id, status, total_amount, created_at;


-- name: AddQuotationItem :one
INSERT INTO quotation_items (
    quotation_id,
    service_id,
    quantity,
    price,
    amount
)
VALUES ($1, $2, $3, $4, $5)
RETURNING id, quotation_id, service_id, quantity, price, amount;


-- name: GetQuotationItems :many
SELECT
    id,
    quotation_id,
    service_id,
    quantity,
    price,
    amount
FROM quotation_items
WHERE quotation_id = $1
ORDER BY id;


-- name: DeleteQuotation :exec
DELETE FROM quotations
WHERE id = $1;


-- name: GetQuotationsByStatus :many
SELECT id, customer_id, vehicle_id, status, total_amount, created_at
FROM quotations
WHERE status = $1
ORDER BY id;
