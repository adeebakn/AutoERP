-- name: CreateVehicle :one 
INSERT INTO vehicles(customer_id,reg_no,vehicle_type) VALUES ($1,$2,$3)
RETURNING id,customer_id,reg_no,vehicle_type,created_at;

-- name: ListVehicles :many
SELECT id, customer_id, reg_no, vehicle_type, created_at
FROM vehicles
ORDER BY id;

-- name: GetVehicle :one
SELECT id, customer_id, reg_no, vehicle_type, created_at
FROM vehicles
WHERE id = $1;

-- name: GetVehiclesByCustomerID :many
SELECT id, customer_id, reg_no, vehicle_type, created_at
FROM vehicles
WHERE customer_id = $1
ORDER BY id;

-- name: UpdateVehicle :one
UPDATE vehicles
SET
    customer_id = $1,
    reg_no = $2,
    vehicle_type = $3
WHERE id = $4
RETURNING id, customer_id, reg_no, vehicle_type, created_at;


-- name: DeleteVehicle :exec
DELETE FROM vehicles
WHERE id = $1;
