-- name: CreateVehicle :one 
INSERT INTO vehicles(customer_id,reg_no,vehicle_type) VALUES ($1,$2,$3)
RETURNING id,customer_id,reg_no,vehicle_type,created_at;

-- name: ListVehicles :many
SELECT
    v.id,
    v.customer_id,
    v.reg_no,
    v.vehicle_type,
    v.created_at,
    c.name AS customer_name
FROM vehicles v
JOIN customers c ON c.id = v.customer_id
WHERE v.deleted_at IS NULL
  AND c.deleted_at IS NULL
ORDER BY v.id;

-- name: GetVehicle :one
SELECT
    v.id,
    v.customer_id,
    v.reg_no,
    v.vehicle_type,
    v.created_at
FROM vehicles v
JOIN customers c ON c.id = v.customer_id
WHERE v.id = $1
  AND v.deleted_at IS NULL
  AND c.deleted_at IS NULL;

-- name: GetVehiclesByCustomerID :many
SELECT
    v.id,
    v.customer_id,
    v.reg_no,
    v.vehicle_type,
    v.created_at
FROM vehicles v
JOIN customers c ON c.id = v.customer_id
WHERE v.customer_id = $1
  AND v.deleted_at IS NULL
  AND c.deleted_at IS NULL
ORDER BY v.id;

-- name: UpdateVehicle :one
UPDATE vehicles
SET
    customer_id = $1,
    reg_no = $2,
    vehicle_type = $3
WHERE id = $4
  AND deleted_at IS NULL
RETURNING id, customer_id, reg_no, vehicle_type, created_at;


-- name: DeleteVehicle :exec
UPDATE vehicles
SET deleted_at = CURRENT_TIMESTAMP
WHERE id = $1
  AND deleted_at IS NULL;