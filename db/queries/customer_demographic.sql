-- name: GetCustomerDemographic :one
-- Gets a customer demographic by type ID
SELECT *
FROM customer_demographics
WHERE customer_type_id = $1;

-- name: ListCustomerDemographics :many
-- Lists all customer demographics
SELECT *
FROM customer_demographics
ORDER BY customer_type_id;

-- name: CreateCustomerDemographic :one
-- Creates a new customer demographic and returns it
INSERT INTO customer_demographics (
  customer_type_id,
  customer_desc
) VALUES (
  $1, $2
)
RETURNING *;

-- name: UpdateCustomerDemographic :one
-- Updates a customer demographic by type ID
UPDATE customer_demographics
SET
  customer_desc = $2
WHERE customer_type_id = $1
RETURNING *;

-- name: DeleteCustomerDemographic :exec
-- Deletes a customer demographic by type ID
DELETE FROM customer_demographics
WHERE customer_type_id = $1;

-- name: SearchCustomerDemographicsByDesc :many
-- Searches customer demographics by description (case insensitive)
SELECT *
FROM customer_demographics
WHERE customer_desc ILIKE '%' || $1 || '%'
ORDER BY customer_type_id;

-- name: CountCustomerDemographics :one
-- Counts the total number of customer demographics
SELECT COUNT(*) FROM customer_demographics;