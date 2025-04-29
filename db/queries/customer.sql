-- name: GetCustomer :one
-- Gets a customer by ID
SELECT *
FROM customers
WHERE customer_id = $1;

-- name: ListCustomers :many
-- Lists all customers
SELECT *
FROM customers
ORDER BY company_name
LIMIT $1 OFFSET $2;


-- name: CreateCustomer :one
-- Creates a new customer and returns it
INSERT INTO customers (
  customer_id,
  company_name,
  contact_name,
  contact_title,
  address,
  city,
  region,
  postal_code,
  country,
  phone,
  fax
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11
)
RETURNING *;

-- name: UpdateCustomer :one
-- Updates specified columns for a customer, leaves others unchanged
UPDATE customers
SET
  company_name = COALESCE($2, company_name),
  contact_name = COALESCE($3, contact_name),
  contact_title = COALESCE($4, contact_title),
  address = COALESCE($5, address),
  city = COALESCE($6, city),
  region = COALESCE($7, region),
  postal_code = COALESCE($8, postal_code),
  country = COALESCE($9, country),
  phone = COALESCE($10, phone),
  fax = COALESCE($11, fax)
WHERE customer_id = $1
RETURNING *;






-- name: DeleteCustomer :exec
-- OBS! Completely deletes a customer by customer_id
DELETE FROM customers
WHERE customer_id = $1;

-- name: SearchCustomersByCompanyName :many
-- Searches customers by company name (case insensitive)
SELECT *
FROM customers
WHERE company_name ILIKE '%' || $1 || '%'
ORDER BY company_name;

-- name: SearchCustomersByContactName :many
-- Searches customers by contact name (case insensitive)
SELECT *
FROM customers
WHERE contact_name ILIKE '%' || $1 || '%'
ORDER BY contact_name;

-- name: ListCustomersByCountry :many
-- Lists all customers from a specific country
SELECT *
FROM customers
WHERE country = $1
ORDER BY company_name;

-- name: ListCustomersByCity :many
-- Lists all customers from a specific city
SELECT *
FROM customers
WHERE city = $1
ORDER BY company_name;

-- name: CountAllCustomers :one
-- Counts the total number of customers
SELECT COUNT(*) FROM customers;

-- name: CountCustomersByCountry :many
-- Counts customers grouped by country
SELECT country, COUNT(*) as customer_count
FROM customers
GROUP BY country
ORDER BY COUNT(*) DESC;

-- name: ListDistinctCountries :many
-- Returns all distinct countries in the customers table
SELECT DISTINCT country
FROM customers
WHERE country IS NOT NULL
ORDER BY country;

-- name: ToggleCustomerActiveStatus :one
-- Toggles the active status of a customer by ID
UPDATE customers
SET active = NOT active
WHERE customer_id = $1
RETURNING *;