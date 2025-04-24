-- name: GetCustomer :one
-- Gets a customer by ID
SELECT *
FROM customers
WHERE customer_id = $1;

-- name: ListCustomers :many
-- Lists all customers
SELECT *
FROM customers
ORDER BY company_name;

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
-- Updates a customer by ID
UPDATE customers
SET
  company_name = $2,
  contact_name = $3,
  contact_title = $4,
  address = $5,
  city = $6,
  region = $7,
  postal_code = $8,
  country = $9,
  phone = $10,
  fax = $11
WHERE customer_id = $1
RETURNING *;

-- name: DeleteCustomer :exec
-- Deletes a customer by ID
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

-- name: CountCustomers :one
-- Counts the total number of customers
SELECT COUNT(*) FROM customers;

-- name: CountCustomersByCountry :many
-- Counts customers grouped by country
SELECT country, COUNT(*) as customer_count
FROM customers
GROUP BY country
ORDER BY COUNT(*) DESC;