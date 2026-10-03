-- name: CreatePayment :one
INSERT INTO payments (
    invoice_id,
    amount,
    payment_method
)
VALUES ($1, $2, $3)
RETURNING id, invoice_id, amount, payment_method, payment_date;


-- name: ListPayments :many
SELECT id, invoice_id, amount, payment_method, payment_date
FROM payments
ORDER BY id;


-- name: GetPayment :one
SELECT id, invoice_id, amount, payment_method, payment_date
FROM payments
WHERE id = $1;


-- name: GetPaymentsByInvoice :many
SELECT id, invoice_id, amount, payment_method, payment_date
FROM payments
WHERE invoice_id = $1
ORDER BY id;


-- name: DeletePayment :exec
DELETE FROM payments
WHERE id = $1;

-- name: GetTotalRevenue :one
SELECT COALESCE(SUM(amount), 0)
FROM payments;