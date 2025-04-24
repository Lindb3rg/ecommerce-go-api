-- name: GetOrder :one
-- Gets an order by ID
SELECT *
FROM orders
WHERE order_id = $1;

-- name: ListOrders :many
-- Lists all orders sorted by date (newest first)
SELECT *
FROM orders
ORDER BY order_date DESC;

-- name: CreateOrder :one
-- Creates a new order and returns it
INSERT INTO orders (
  customer_id,
  employee_id,
  order_date,
  required_date,
  shipped_date,
  ship_via,
  freight,
  ship_name,
  ship_address,
  ship_city,
  ship_region,
  ship_postal_code,
  ship_country
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13
)
RETURNING *;

-- name: UpdateOrder :one
-- Updates an order by ID
UPDATE orders
SET
  customer_id = $2,
  employee_id = $3,
  order_date = $4,
  required_date = $5,
  shipped_date = $6,
  ship_via = $7,
  freight = $8,
  ship_name = $9,
  ship_address = $10,
  ship_city = $11,
  ship_region = $12,
  ship_postal_code = $13,
  ship_country = $14
WHERE order_id = $1
RETURNING *;

-- name: DeleteOrder :exec
-- Deletes an order by ID
DELETE FROM orders
WHERE order_id = $1;

-- name: ListOrdersByCustomer :many
-- Lists all orders for a specific customer
SELECT *
FROM orders
WHERE customer_id = $1
ORDER BY order_date DESC;

-- name: ListOrdersByEmployee :many
-- Lists all orders handled by a specific employee
SELECT *
FROM orders
WHERE employee_id = $1
ORDER BY order_date DESC;

-- name: ListOrdersByShipper :many
-- Lists all orders shipped by a specific shipper
SELECT *
FROM orders
WHERE ship_via = $1
ORDER BY order_date DESC;

-- name: GetOrderWithDetails :one
-- Gets an order by ID with customer, employee, and shipper details
SELECT 
  o.*,
  c.company_name as customer_name,
  c.contact_name as customer_contact,
  CONCAT(e.first_name, ' ', e.last_name) as employee_name,
  s.company_name as shipper_name
FROM orders o
LEFT JOIN customers c ON o.customer_id = c.customer_id
LEFT JOIN employees e ON o.employee_id = e.employee_id
LEFT JOIN shippers s ON o.ship_via = s.shipper_id
WHERE o.order_id = $1;

-- name: ListRecentOrders :many
-- Lists orders placed within a specified number of days
SELECT *
FROM orders
WHERE order_date >= CURRENT_DATE - $1::interval
ORDER BY order_date DESC;

-- name: ListPendingShipments :many
-- Lists orders that have not been shipped yet
SELECT *
FROM orders
WHERE shipped_date IS NULL
ORDER BY required_date ASC;

-- name: ListDelayedOrders :many
-- Lists orders that were shipped after the required date
SELECT *
FROM orders
WHERE shipped_date > required_date
ORDER BY shipped_date DESC;

-- name: UpdateShippingInfo :one
-- Updates just the shipping info for an order
UPDATE orders
SET
  shipped_date = $2,
  ship_via = $3,
  freight = $4
WHERE order_id = $1
RETURNING *;

-- name: CountOrders :one
-- Counts the total number of orders
SELECT COUNT(*) FROM orders;

-- name: CountOrdersByStatus :many
-- Counts orders grouped by their status
SELECT
  CASE
    WHEN shipped_date IS NULL THEN 'Pending'
    WHEN shipped_date <= required_date THEN 'On Time'
    ELSE 'Delayed'
  END as status,
  COUNT(*) as order_count
FROM orders
GROUP BY status
ORDER BY status;

-- name: CountOrdersByCountry :many
-- Counts orders grouped by shipping country
SELECT 
  ship_country,
  COUNT(*) as order_count
FROM orders
GROUP BY ship_country
ORDER BY COUNT(*) DESC;

-- name: ListOrdersByDateRange :many
-- Lists orders within a specific date range
SELECT *
FROM orders
WHERE order_date BETWEEN $1 AND $2
ORDER BY order_date DESC;

-- name: GetMonthlyOrderCounts :many
-- Gets order counts by month for a given year
SELECT
  EXTRACT(MONTH FROM order_date) as month,
  COUNT(*) as order_count
FROM orders
WHERE EXTRACT(YEAR FROM order_date) = $1
GROUP BY EXTRACT(MONTH FROM order_date)
ORDER BY month;

-- name: GetTotalFreightByCustomer :many
-- Gets total freight costs grouped by customer
SELECT
  o.customer_id,
  c.company_name,
  SUM(o.freight) as total_freight
FROM orders o
JOIN customers c ON o.customer_id = c.customer_id
GROUP BY o.customer_id, c.company_name
ORDER BY SUM(o.freight) DESC
LIMIT $1;

-- name: GetAverageProcessingTime :one
-- Gets the average time between order and shipment
SELECT
  AVG(shipped_date - order_date) as avg_processing_days
FROM orders
WHERE shipped_date IS NOT NULL;