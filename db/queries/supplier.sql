-- name: GetSupplier :one
-- Gets a supplier by ID
SELECT *
FROM suppliers
WHERE supplier_id = $1;

-- name: ListSuppliers :many
-- Lists all suppliers
SELECT *
FROM suppliers
ORDER BY company_name;

-- name: CreateSupplier :one
-- Creates a new supplier and returns it
INSERT INTO suppliers (
  company_name,
  contact_name,
  contact_title,
  address,
  city,
  region,
  postal_code,
  country,
  phone,
  fax,
  homepage
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11
)
RETURNING *;

-- name: UpdateSupplier :one
-- Updates a supplier by ID
UPDATE suppliers
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
  fax = $11,
  homepage = $12
WHERE supplier_id = $1
RETURNING *;

-- name: DeleteSupplier :exec
-- Deletes a supplier by ID
DELETE FROM suppliers
WHERE supplier_id = $1;

-- name: SearchSuppliersByCompanyName :many
-- Searches suppliers by company name (case insensitive)
SELECT *
FROM suppliers
WHERE company_name ILIKE '%' || $1 || '%'
ORDER BY company_name;

-- name: SearchSuppliersByContactName :many
-- Searches suppliers by contact name (case insensitive)
SELECT *
FROM suppliers
WHERE contact_name ILIKE '%' || $1 || '%'
ORDER BY contact_name;

-- name: ListSuppliersByCountry :many
-- Lists all suppliers from a specific country
SELECT *
FROM suppliers
WHERE country = $1
ORDER BY company_name;

-- name: ListSuppliersByCity :many
-- Lists all suppliers from a specific city
SELECT *
FROM suppliers
WHERE city = $1
ORDER BY company_name;

-- name: CountSuppliers :one
-- Counts the total number of suppliers
SELECT COUNT(*) FROM suppliers;

-- name: CountSuppliersByCountry :many
-- Counts suppliers grouped by country
SELECT country, COUNT(*) as supplier_count
FROM suppliers
GROUP BY country
ORDER BY COUNT(*) DESC;

-- name: GetSupplierWithContactInfo :one
-- Gets a supplier by ID with formatted contact information
SELECT
  supplier_id,
  company_name,
  contact_name,
  contact_title,
  address,
  city,
  region,
  postal_code,
  country,
  phone,
  fax,
  homepage,
  contact_name || ' (' || contact_title || ')' as formatted_contact,
  address || ', ' || city || COALESCE(', ' || region, '') || ' ' || COALESCE(postal_code, '') || ', ' || country as full_address
FROM suppliers
WHERE supplier_id = $1;