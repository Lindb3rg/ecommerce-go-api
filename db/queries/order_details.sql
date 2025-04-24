-- name: GetOrderDetail :one
-- Gets a specific order detail by order ID and product ID
SELECT *
FROM order_details
WHERE order_id = $1 AND product_id = $2;

-- name: ListOrderDetailsByOrder :many
-- Lists all details for a specific order
SELECT *
FROM order_details
WHERE order_id = $1
ORDER BY product_id;

-- name: ListOrderDetailsByProduct :many
-- Lists all orders containing a specific product
SELECT *
FROM order_details
WHERE product_id = $1
ORDER BY order_id DESC;

-- name: CreateOrderDetail :one
-- Creates a new order detail and returns it
INSERT INTO order_details (
  order_id,
  product_id,
  unit_price,
  quantity,
  discount
) VALUES (
  $1, $2, $3, $4, $5
)
RETURNING *;

-- name: UpdateOrderDetail :one
-- Updates an order detail by order ID and product ID
UPDATE order_details
SET
  unit_price = $3,
  quantity = $4,
  discount = $5
WHERE order_id = $1 AND product_id = $2
RETURNING *;

-- name: DeleteOrderDetail :exec
-- Deletes a specific order detail
DELETE FROM order_details
WHERE order_id = $1 AND product_id = $2;

-- name: DeleteAllOrderDetails :exec
-- Deletes all details for a specific order
DELETE FROM order_details
WHERE order_id = $1;

-- name: GetOrderDetailWithProductInfo :many
-- Gets order details with product information for a specific order
SELECT 
  od.order_id,
  od.product_id,
  od.unit_price,
  od.quantity,
  od.discount,
  p.product_name,
  p.supplier_id,
  p.category_id,
  p.quantity_per_unit,
  p.discontinued,
  (od.unit_price * od.quantity * (1 - od.discount)) as subtotal
FROM order_details od
JOIN products p ON od.product_id = p.product_id
WHERE od.order_id = $1
ORDER BY p.product_name;

-- name: GetOrderTotal :one
-- Calculates the total amount for a specific order
SELECT 
  SUM(unit_price * quantity * (1 - discount)) as order_total
FROM order_details
WHERE order_id = $1;

-- name: CountProductsInOrder :one
-- Counts how many different products are in a specific order
SELECT COUNT(*) 
FROM order_details
WHERE order_id = $1;

-- name: GetTotalQuantityInOrder :one
-- Gets the total quantity of items in a specific order
SELECT SUM(quantity) as total_items
FROM order_details
WHERE order_id = $1;

-- name: GetMostPopularProducts :many
-- Gets the most popular products based on quantity ordered
SELECT 
  p.product_id,
  p.product_name,
  SUM(od.quantity) as total_ordered
FROM order_details od
JOIN products p ON od.product_id = p.product_id
GROUP BY p.product_id, p.product_name
ORDER BY SUM(od.quantity) DESC
LIMIT $1;

-- name: GetOrderDetailsWithOrderInfo :many
-- Gets order details with order and product information
SELECT 
  od.order_id,
  o.order_date,
  od.product_id,
  p.product_name,
  od.unit_price,
  od.quantity,
  od.discount,
  (od.unit_price * od.quantity * (1 - od.discount)) as subtotal
FROM order_details od
JOIN orders o ON od.order_id = o.order_id
JOIN products p ON od.product_id = p.product_id
WHERE od.order_id = $1
ORDER BY p.product_name;

-- name: GetProductSalesByDateRange :many
-- Gets product sales within a specific date range
SELECT 
  p.product_id,
  p.product_name,
  SUM(od.quantity) as total_quantity,
  SUM(od.unit_price * od.quantity * (1 - od.discount)) as total_sales
FROM order_details od
JOIN products p ON od.product_id = p.product_id
JOIN orders o ON od.order_id = o.order_id
WHERE o.order_date BETWEEN $1 AND $2
GROUP BY p.product_id, p.product_name
ORDER BY SUM(od.quantity) DESC;

-- name: GetProductSalesByCategory :many
-- Gets product sales grouped by category
SELECT 
  c.category_id,
  c.category_name,
  SUM(od.quantity) as total_quantity,
  SUM(od.unit_price * od.quantity * (1 - od.discount)) as total_sales
FROM order_details od
JOIN products p ON od.product_id = p.product_id
JOIN categories c ON p.category_id = c.category_id
GROUP BY c.category_id, c.category_name
ORDER BY SUM(od.unit_price * od.quantity * (1 - od.discount)) DESC;

-- name: GetAverageOrderValue :one
-- Gets the average order value
SELECT 
  AVG(order_total) as average_order_value
FROM (
  SELECT 
    order_id,
    SUM(unit_price * quantity * (1 - discount)) as order_total
  FROM order_details
  GROUP BY order_id
) as order_totals;

-- name: UpdateOrderDetailQuantity :one
-- Updates only the quantity of an order detail
UPDATE order_details
SET quantity = $3
WHERE order_id = $1 AND product_id = $2
RETURNING *;