-- name: GetProduct :one
-- Gets a product by ID
SELECT *
FROM products
WHERE product_id = $1;

-- name: ListProducts :many
-- Lists all products
SELECT *
FROM products
ORDER BY product_name;

-- name: CreateProduct :one
-- Creates a new product and returns it
INSERT INTO products (
  product_name,
  supplier_id,
  category_id,
  quantity_per_unit,
  unit_price,
  units_in_stock,
  units_on_order,
  reorder_level,
  discontinued
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9
)
RETURNING *;

-- name: UpdateProduct :one
-- Updates a product by ID
UPDATE products
SET
  product_name = $2,
  supplier_id = $3,
  category_id = $4,
  quantity_per_unit = $5,
  unit_price = $6,
  units_in_stock = $7,
  units_on_order = $8,
  reorder_level = $9,
  discontinued = $10
WHERE product_id = $1
RETURNING *;

-- name: DeleteProduct :exec
-- Deletes a product by ID
DELETE FROM products
WHERE product_id = $1;

-- name: SearchProductsByName :many
-- Searches products by name (case insensitive)
SELECT *
FROM products
WHERE product_name ILIKE '%' || $1 || '%'
ORDER BY product_name;

-- name: ListProductsByCategory :many
-- Lists all products in a specific category
SELECT *
FROM products
WHERE category_id = $1
ORDER BY product_name;

-- name: ListProductsBySupplier :many
-- Lists all products from a specific supplier
SELECT *
FROM products
WHERE supplier_id = $1
ORDER BY product_name;

-- name: GetProductWithDetails :one
-- Gets a product by ID with category and supplier details
SELECT 
  p.product_id,
  p.product_name,
  p.supplier_id,
  p.category_id,
  p.quantity_per_unit,
  p.unit_price,
  p.units_in_stock,
  p.units_on_order,
  p.reorder_level,
  p.discontinued,
  c.category_name,
  c.description as category_description,
  s.company_name as supplier_name,
  s.contact_name as supplier_contact,
  s.country as supplier_country
FROM products p
LEFT JOIN categories c ON p.category_id = c.category_id
LEFT JOIN suppliers s ON p.supplier_id = s.supplier_id
WHERE p.product_id = $1;

-- name: ListProductsWithDetails :many
-- Lists all products with category and supplier details
SELECT 
  p.product_id,
  p.product_name,
  p.supplier_id,
  p.category_id,
  p.quantity_per_unit,
  p.unit_price,
  p.units_in_stock,
  p.units_on_order,
  p.reorder_level,
  p.discontinued,
  c.category_name,
  s.company_name as supplier_name
FROM products p
LEFT JOIN categories c ON p.category_id = c.category_id
LEFT JOIN suppliers s ON p.supplier_id = s.supplier_id
ORDER BY p.product_name;

-- name: ListProductsNeedingReorder :many
-- Lists all products that need to be reordered (stock below reorder level)
SELECT *
FROM products
WHERE units_in_stock <= reorder_level AND discontinued = 0
ORDER BY product_name;

-- name: ListDiscontinuedProducts :many
-- Lists all discontinued products
SELECT *
FROM products
WHERE discontinued = 1
ORDER BY product_name;

-- name: CountProducts :one
-- Counts the total number of products
SELECT COUNT(*) FROM products;

-- name: CountProductsByCategory :many
-- Counts products grouped by category
SELECT 
  c.category_id,
  c.category_name,
  COUNT(*) as product_count
FROM products p
JOIN categories c ON p.category_id = c.category_id
GROUP BY c.category_id, c.category_name
ORDER BY COUNT(*) DESC;

-- name: CountProductsBySupplier :many
-- Counts products grouped by supplier
SELECT 
  s.supplier_id,
  s.company_name,
  COUNT(*) as product_count
FROM products p
JOIN suppliers s ON p.supplier_id = s.supplier_id
GROUP BY s.supplier_id, s.company_name
ORDER BY COUNT(*) DESC;

-- name: GetProductValueByCategory :many
-- Gets the total inventory value by category
SELECT 
  c.category_id,
  c.category_name,
  SUM(p.unit_price * p.units_in_stock) as inventory_value
FROM products p
JOIN categories c ON p.category_id = c.category_id
GROUP BY c.category_id, c.category_name
ORDER BY SUM(p.unit_price * p.units_in_stock) DESC;

-- name: UpdateProductStock :one
-- Updates a product's stock levels
UPDATE products
SET
  units_in_stock = $2,
  units_on_order = $3
WHERE product_id = $1
RETURNING *;

-- name: UpdateProductPrice :one
-- Updates a product's price
UPDATE products
SET
  unit_price = $2
WHERE product_id = $1
RETURNING *;

-- name: DiscontinueProduct :one
-- Marks a product as discontinued
UPDATE products
SET
  discontinued = 1
WHERE product_id = $1
RETURNING *;