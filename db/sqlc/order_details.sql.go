// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: order_details.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const countProductsInOrder = `-- name: CountProductsInOrder :one
SELECT COUNT(*) 
FROM order_details
WHERE order_id = $1
`

// Counts how many different products are in a specific order
func (q *Queries) CountProductsInOrder(ctx context.Context, orderID int16) (int64, error) {
	row := q.db.QueryRow(ctx, countProductsInOrder, orderID)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const createOrderDetail = `-- name: CreateOrderDetail :one
INSERT INTO order_details (
  order_id,
  product_id,
  unit_price,
  quantity,
  discount
) VALUES (
  $1, $2, $3, $4, $5
)
RETURNING order_id, product_id, unit_price, quantity, discount
`

type CreateOrderDetailParams struct {
	OrderID   int16   `json:"order_id"`
	ProductID int16   `json:"product_id"`
	UnitPrice float32 `json:"unit_price"`
	Quantity  int16   `json:"quantity"`
	Discount  float32 `json:"discount"`
}

// Creates a new order detail and returns it
func (q *Queries) CreateOrderDetail(ctx context.Context, arg CreateOrderDetailParams) (OrderDetail, error) {
	row := q.db.QueryRow(ctx, createOrderDetail,
		arg.OrderID,
		arg.ProductID,
		arg.UnitPrice,
		arg.Quantity,
		arg.Discount,
	)
	var i OrderDetail
	err := row.Scan(
		&i.OrderID,
		&i.ProductID,
		&i.UnitPrice,
		&i.Quantity,
		&i.Discount,
	)
	return i, err
}

const deleteAllOrderDetails = `-- name: DeleteAllOrderDetails :exec
DELETE FROM order_details
WHERE order_id = $1
`

// Deletes all details for a specific order
func (q *Queries) DeleteAllOrderDetails(ctx context.Context, orderID int16) error {
	_, err := q.db.Exec(ctx, deleteAllOrderDetails, orderID)
	return err
}

const deleteOrderDetail = `-- name: DeleteOrderDetail :exec
DELETE FROM order_details
WHERE order_id = $1 AND product_id = $2
`

type DeleteOrderDetailParams struct {
	OrderID   int16 `json:"order_id"`
	ProductID int16 `json:"product_id"`
}

// Deletes a specific order detail
func (q *Queries) DeleteOrderDetail(ctx context.Context, arg DeleteOrderDetailParams) error {
	_, err := q.db.Exec(ctx, deleteOrderDetail, arg.OrderID, arg.ProductID)
	return err
}

const getAverageOrderValue = `-- name: GetAverageOrderValue :one
SELECT 
  AVG(order_total) as average_order_value
FROM (
  SELECT 
    order_id,
    SUM(unit_price * quantity * (1 - discount)) as order_total
  FROM order_details
  GROUP BY order_id
) as order_totals
`

// Gets the average order value
func (q *Queries) GetAverageOrderValue(ctx context.Context) (float64, error) {
	row := q.db.QueryRow(ctx, getAverageOrderValue)
	var average_order_value float64
	err := row.Scan(&average_order_value)
	return average_order_value, err
}

const getMostPopularProducts = `-- name: GetMostPopularProducts :many
SELECT 
  p.product_id,
  p.product_name,
  SUM(od.quantity) as total_ordered
FROM order_details od
JOIN products p ON od.product_id = p.product_id
GROUP BY p.product_id, p.product_name
ORDER BY SUM(od.quantity) DESC
LIMIT $1
`

type GetMostPopularProductsRow struct {
	ProductID    int16  `json:"product_id"`
	ProductName  string `json:"product_name"`
	TotalOrdered int64  `json:"total_ordered"`
}

// Gets the most popular products based on quantity ordered
func (q *Queries) GetMostPopularProducts(ctx context.Context, limit int32) ([]GetMostPopularProductsRow, error) {
	rows, err := q.db.Query(ctx, getMostPopularProducts, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetMostPopularProductsRow{}
	for rows.Next() {
		var i GetMostPopularProductsRow
		if err := rows.Scan(&i.ProductID, &i.ProductName, &i.TotalOrdered); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getOrderDetail = `-- name: GetOrderDetail :one
SELECT order_id, product_id, unit_price, quantity, discount
FROM order_details
WHERE order_id = $1 AND product_id = $2
`

type GetOrderDetailParams struct {
	OrderID   int16 `json:"order_id"`
	ProductID int16 `json:"product_id"`
}

// Gets a specific order detail by order ID and product ID
func (q *Queries) GetOrderDetail(ctx context.Context, arg GetOrderDetailParams) (OrderDetail, error) {
	row := q.db.QueryRow(ctx, getOrderDetail, arg.OrderID, arg.ProductID)
	var i OrderDetail
	err := row.Scan(
		&i.OrderID,
		&i.ProductID,
		&i.UnitPrice,
		&i.Quantity,
		&i.Discount,
	)
	return i, err
}

const getOrderDetailWithProductInfo = `-- name: GetOrderDetailWithProductInfo :many
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
ORDER BY p.product_name
`

type GetOrderDetailWithProductInfoRow struct {
	OrderID         int16       `json:"order_id"`
	ProductID       int16       `json:"product_id"`
	UnitPrice       float32     `json:"unit_price"`
	Quantity        int16       `json:"quantity"`
	Discount        float32     `json:"discount"`
	ProductName     string      `json:"product_name"`
	SupplierID      pgtype.Int2 `json:"supplier_id"`
	CategoryID      pgtype.Int2 `json:"category_id"`
	QuantityPerUnit pgtype.Text `json:"quantity_per_unit"`
	Discontinued    int32       `json:"discontinued"`
	Subtotal        int32       `json:"subtotal"`
}

// Gets order details with product information for a specific order
func (q *Queries) GetOrderDetailWithProductInfo(ctx context.Context, orderID int16) ([]GetOrderDetailWithProductInfoRow, error) {
	rows, err := q.db.Query(ctx, getOrderDetailWithProductInfo, orderID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetOrderDetailWithProductInfoRow{}
	for rows.Next() {
		var i GetOrderDetailWithProductInfoRow
		if err := rows.Scan(
			&i.OrderID,
			&i.ProductID,
			&i.UnitPrice,
			&i.Quantity,
			&i.Discount,
			&i.ProductName,
			&i.SupplierID,
			&i.CategoryID,
			&i.QuantityPerUnit,
			&i.Discontinued,
			&i.Subtotal,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getOrderDetailsWithOrderInfo = `-- name: GetOrderDetailsWithOrderInfo :many
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
ORDER BY p.product_name
`

type GetOrderDetailsWithOrderInfoRow struct {
	OrderID     int16       `json:"order_id"`
	OrderDate   pgtype.Date `json:"order_date"`
	ProductID   int16       `json:"product_id"`
	ProductName string      `json:"product_name"`
	UnitPrice   float32     `json:"unit_price"`
	Quantity    int16       `json:"quantity"`
	Discount    float32     `json:"discount"`
	Subtotal    int32       `json:"subtotal"`
}

// Gets order details with order and product information
func (q *Queries) GetOrderDetailsWithOrderInfo(ctx context.Context, orderID int16) ([]GetOrderDetailsWithOrderInfoRow, error) {
	rows, err := q.db.Query(ctx, getOrderDetailsWithOrderInfo, orderID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetOrderDetailsWithOrderInfoRow{}
	for rows.Next() {
		var i GetOrderDetailsWithOrderInfoRow
		if err := rows.Scan(
			&i.OrderID,
			&i.OrderDate,
			&i.ProductID,
			&i.ProductName,
			&i.UnitPrice,
			&i.Quantity,
			&i.Discount,
			&i.Subtotal,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getOrderTotal = `-- name: GetOrderTotal :one
SELECT 
  SUM(unit_price * quantity * (1 - discount)) as order_total
FROM order_details
WHERE order_id = $1
`

// Calculates the total amount for a specific order
func (q *Queries) GetOrderTotal(ctx context.Context, orderID int16) (int64, error) {
	row := q.db.QueryRow(ctx, getOrderTotal, orderID)
	var order_total int64
	err := row.Scan(&order_total)
	return order_total, err
}

const getProductSalesByCategory = `-- name: GetProductSalesByCategory :many
SELECT 
  c.category_id,
  c.category_name,
  SUM(od.quantity) as total_quantity,
  SUM(od.unit_price * od.quantity * (1 - od.discount)) as total_sales
FROM order_details od
JOIN products p ON od.product_id = p.product_id
JOIN categories c ON p.category_id = c.category_id
GROUP BY c.category_id, c.category_name
ORDER BY SUM(od.unit_price * od.quantity * (1 - od.discount)) DESC
`

type GetProductSalesByCategoryRow struct {
	CategoryID    int16  `json:"category_id"`
	CategoryName  string `json:"category_name"`
	TotalQuantity int64  `json:"total_quantity"`
	TotalSales    int64  `json:"total_sales"`
}

// Gets product sales grouped by category
func (q *Queries) GetProductSalesByCategory(ctx context.Context) ([]GetProductSalesByCategoryRow, error) {
	rows, err := q.db.Query(ctx, getProductSalesByCategory)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetProductSalesByCategoryRow{}
	for rows.Next() {
		var i GetProductSalesByCategoryRow
		if err := rows.Scan(
			&i.CategoryID,
			&i.CategoryName,
			&i.TotalQuantity,
			&i.TotalSales,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getProductSalesByDateRange = `-- name: GetProductSalesByDateRange :many
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
ORDER BY SUM(od.quantity) DESC
`

type GetProductSalesByDateRangeParams struct {
	OrderDate   pgtype.Date `json:"order_date"`
	OrderDate_2 pgtype.Date `json:"order_date_2"`
}

type GetProductSalesByDateRangeRow struct {
	ProductID     int16  `json:"product_id"`
	ProductName   string `json:"product_name"`
	TotalQuantity int64  `json:"total_quantity"`
	TotalSales    int64  `json:"total_sales"`
}

// Gets product sales within a specific date range
func (q *Queries) GetProductSalesByDateRange(ctx context.Context, arg GetProductSalesByDateRangeParams) ([]GetProductSalesByDateRangeRow, error) {
	rows, err := q.db.Query(ctx, getProductSalesByDateRange, arg.OrderDate, arg.OrderDate_2)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetProductSalesByDateRangeRow{}
	for rows.Next() {
		var i GetProductSalesByDateRangeRow
		if err := rows.Scan(
			&i.ProductID,
			&i.ProductName,
			&i.TotalQuantity,
			&i.TotalSales,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getTotalQuantityInOrder = `-- name: GetTotalQuantityInOrder :one
SELECT SUM(quantity) as total_items
FROM order_details
WHERE order_id = $1
`

// Gets the total quantity of items in a specific order
func (q *Queries) GetTotalQuantityInOrder(ctx context.Context, orderID int16) (int64, error) {
	row := q.db.QueryRow(ctx, getTotalQuantityInOrder, orderID)
	var total_items int64
	err := row.Scan(&total_items)
	return total_items, err
}

const listOrderDetailsByOrder = `-- name: ListOrderDetailsByOrder :many
SELECT order_id, product_id, unit_price, quantity, discount
FROM order_details
WHERE order_id = $1
ORDER BY product_id
`

// Lists all details for a specific order
func (q *Queries) ListOrderDetailsByOrder(ctx context.Context, orderID int16) ([]OrderDetail, error) {
	rows, err := q.db.Query(ctx, listOrderDetailsByOrder, orderID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []OrderDetail{}
	for rows.Next() {
		var i OrderDetail
		if err := rows.Scan(
			&i.OrderID,
			&i.ProductID,
			&i.UnitPrice,
			&i.Quantity,
			&i.Discount,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listOrderDetailsByProduct = `-- name: ListOrderDetailsByProduct :many
SELECT order_id, product_id, unit_price, quantity, discount
FROM order_details
WHERE product_id = $1
ORDER BY order_id DESC
`

// Lists all orders containing a specific product
func (q *Queries) ListOrderDetailsByProduct(ctx context.Context, productID int16) ([]OrderDetail, error) {
	rows, err := q.db.Query(ctx, listOrderDetailsByProduct, productID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []OrderDetail{}
	for rows.Next() {
		var i OrderDetail
		if err := rows.Scan(
			&i.OrderID,
			&i.ProductID,
			&i.UnitPrice,
			&i.Quantity,
			&i.Discount,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateOrderDetail = `-- name: UpdateOrderDetail :one
UPDATE order_details
SET
  unit_price = $3,
  quantity = $4,
  discount = $5
WHERE order_id = $1 AND product_id = $2
RETURNING order_id, product_id, unit_price, quantity, discount
`

type UpdateOrderDetailParams struct {
	OrderID   int16   `json:"order_id"`
	ProductID int16   `json:"product_id"`
	UnitPrice float32 `json:"unit_price"`
	Quantity  int16   `json:"quantity"`
	Discount  float32 `json:"discount"`
}

// Updates an order detail by order ID and product ID
func (q *Queries) UpdateOrderDetail(ctx context.Context, arg UpdateOrderDetailParams) (OrderDetail, error) {
	row := q.db.QueryRow(ctx, updateOrderDetail,
		arg.OrderID,
		arg.ProductID,
		arg.UnitPrice,
		arg.Quantity,
		arg.Discount,
	)
	var i OrderDetail
	err := row.Scan(
		&i.OrderID,
		&i.ProductID,
		&i.UnitPrice,
		&i.Quantity,
		&i.Discount,
	)
	return i, err
}

const updateOrderDetailQuantity = `-- name: UpdateOrderDetailQuantity :one
UPDATE order_details
SET quantity = $3
WHERE order_id = $1 AND product_id = $2
RETURNING order_id, product_id, unit_price, quantity, discount
`

type UpdateOrderDetailQuantityParams struct {
	OrderID   int16 `json:"order_id"`
	ProductID int16 `json:"product_id"`
	Quantity  int16 `json:"quantity"`
}

// Updates only the quantity of an order detail
func (q *Queries) UpdateOrderDetailQuantity(ctx context.Context, arg UpdateOrderDetailQuantityParams) (OrderDetail, error) {
	row := q.db.QueryRow(ctx, updateOrderDetailQuantity, arg.OrderID, arg.ProductID, arg.Quantity)
	var i OrderDetail
	err := row.Scan(
		&i.OrderID,
		&i.ProductID,
		&i.UnitPrice,
		&i.Quantity,
		&i.Discount,
	)
	return i, err
}
