-- name: CreateWorkOrder :one
INSERT INTO work_orders (
    quotation_id,
    customer_id,
    vehicle_id,
    status,
    due_date
)
VALUES ($1, $2, $3, $4, $5)
RETURNING id, quotation_id, customer_id, vehicle_id, status, due_date, created_at;


-- name: ListWorkOrders :many
SELECT
    id,
    quotation_id,
    customer_id,
    vehicle_id,
    status,
    due_date,
    created_at
FROM work_orders
ORDER BY id;


-- name: GetWorkOrder :one
SELECT
    id,
    quotation_id,
    customer_id,
    vehicle_id,
    status,
    due_date,
    created_at
FROM work_orders
WHERE id = $1;


-- name: UpdateWorkOrderStatus :one
UPDATE work_orders
SET status = $1
WHERE id = $2
RETURNING id, quotation_id, customer_id, vehicle_id, status, due_date, created_at;


-- name: DeleteWorkOrder :exec
DELETE FROM work_orders
WHERE id = $1;


-- name: GetWorkOrdersByStatus :many
SELECT id, quotation_id, customer_id, vehicle_id, status, due_date, created_at
FROM work_orders
WHERE status = $1
ORDER BY id;

-- name: CountWorkOrdersByStatus :one
SELECT COUNT(*)
FROM work_orders
WHERE status = $1;

