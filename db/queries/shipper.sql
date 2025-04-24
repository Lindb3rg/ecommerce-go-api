-- name: GetShipper :one
-- Gets a shipper by ID
SELECT *
FROM shippers
WHERE shipper_id = $1;

-- name: ListShippers :many
-- Lists all shippers
SELECT *
FROM shippers
ORDER BY company_name;

-- name: CreateShipper :one
-- Creates a new shipper and returns it
INSERT INTO shippers (
  company_name,
  phone
) VALUES (
  $1, $2
)
RETURNING *;

-- name: UpdateShipper :one
-- Updates a shipper by ID
UPDATE shippers
SET
  company_name = $2,
  phone = $3
WHERE shipper_id = $1
RETURNING *;

-- name: DeleteShipper :exec
-- Deletes a shipper by ID
DELETE FROM shippers
WHERE shipper_id = $1;

-- name: SearchShippersByName :many
-- Searches shippers by company name (case insensitive)
SELECT *
FROM shippers
WHERE company_name ILIKE '%' || $1 || '%'
ORDER BY company_name;

-- name: SearchShippersByPhone :many
-- Searches shippers by phone number
SELECT *
FROM shippers
WHERE phone LIKE '%' || $1 || '%'
ORDER BY company_name;

-- name: CountShippers :one
-- Counts the total number of shippers
SELECT COUNT(*) FROM shippers;

-- name: GetShipperByExactName :one
-- Gets a shipper by exact company name (useful for duplicate checks)
SELECT *
FROM shippers
WHERE company_name = $1;