-- name: GetCustomerDemoRelation :one
-- Gets a specific customer-demographic relation
SELECT *
FROM customer_customer_demo
WHERE customer_id = $1 AND customer_type_id = $2;

-- name: ListCustomerDemographicsByCustomer :many
-- Lists all demographics for a specific customer with demographic details
SELECT cd.*
FROM customer_customer_demo ccd
JOIN customer_demographics cd ON ccd.customer_type_id = cd.customer_type_id
WHERE ccd.customer_id = $1;

-- name: ListCustomersByDemographic :many
-- Lists all customers belonging to a specific demographic with customer details
SELECT c.*
FROM customer_customer_demo ccd
JOIN customers c ON ccd.customer_id = c.customer_id
WHERE ccd.customer_type_id = $1
ORDER BY c.company_name;

-- name: CreateCustomerDemoRelation :one
-- Creates a new customer-demographic relation
INSERT INTO customer_customer_demo (
  customer_id,
  customer_type_id
) VALUES (
  $1, $2
)
RETURNING *;

-- name: DeleteCustomerDemoRelation :exec
-- Deletes a specific customer-demographic relation
DELETE FROM customer_customer_demo
WHERE customer_id = $1 AND customer_type_id = $2;

-- name: DeleteAllCustomerDemoRelations :exec
-- Deletes all demographic relations for a specific customer
DELETE FROM customer_customer_demo
WHERE customer_id = $1;

-- name: CountCustomersByDemographic :one
-- Counts how many customers belong to a specific demographic
SELECT COUNT(*)
FROM customer_customer_demo
WHERE customer_type_id = $1;

-- name: CountDemographicsByCustomer :one
-- Counts how many demographics a specific customer belongs to
SELECT COUNT(*)
FROM customer_customer_demo
WHERE customer_id = $1;

-- name: ListAllCustomerDemographicsWithDetails :many
-- Lists all customer-demographic relations with full details from both tables
SELECT 
  c.customer_id,
  c.company_name,
  c.contact_name,
  cd.customer_type_id,
  cd.customer_desc
FROM customer_customer_demo ccd
JOIN customers c ON ccd.customer_id = c.customer_id
JOIN customer_demographics cd ON ccd.customer_type_id = cd.customer_type_id
ORDER BY c.company_name, cd.customer_type_id;