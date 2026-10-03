-- name: CreateInvoice :one
INSERT INTO invoices (
    work_order_id,
    customer_id,
    total_amount,
    status
)
VALUES ($1, $2, $3, $4)
RETURNING id, work_order_id, customer_id, total_amount, status, created_at;


-- name: ListInvoices :many
SELECT id, work_order_id, customer_id, total_amount, status, created_at
FROM invoices
ORDER BY id;


-- name: GetInvoice :one
SELECT id, work_order_id, customer_id, total_amount, status, created_at
FROM invoices
WHERE id = $1;


-- name: UpdateInvoiceStatus :one
UPDATE invoices
SET status = $1
WHERE id = $2
RETURNING id, work_order_id, customer_id, total_amount, status, created_at;


-- name: DeleteInvoice :exec
DELETE FROM invoices
WHERE id = $1;


-- name: GetWorkOrderForInvoice :one
SELECT id, quotation_id, customer_id, vehicle_id, status, due_date, created_at
FROM work_orders
WHERE id = $1;


-- name: GetInvoicesByStatus :many
SELECT id, work_order_id, customer_id, total_amount, status, created_at
FROM invoices
WHERE status = $1
ORDER BY id;

-- name: CountInvoicesByStatus :one
SELECT COUNT(*)
FROM invoices
WHERE status = $1;