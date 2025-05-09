// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: product.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const countProducts = `-- name: CountProducts :one
SELECT COUNT(*) FROM products
`

// Counts the total number of products
func (q *Queries) CountProducts(ctx context.Context) (int64, error) {
	row := q.db.QueryRow(ctx, countProducts)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const countProductsByCategory = `-- name: CountProductsByCategory :many
SELECT 
  c.category_id,
  c.category_name,
  COUNT(*) as product_count
FROM products p
JOIN categories c ON p.category_id = c.category_id
GROUP BY c.category_id, c.category_name
ORDER BY COUNT(*) DESC
`

type CountProductsByCategoryRow struct {
	CategoryID   int16  `json:"category_id"`
	CategoryName string `json:"category_name"`
	ProductCount int64  `json:"product_count"`
}

// Counts products grouped by category
func (q *Queries) CountProductsByCategory(ctx context.Context) ([]CountProductsByCategoryRow, error) {
	rows, err := q.db.Query(ctx, countProductsByCategory)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []CountProductsByCategoryRow{}
	for rows.Next() {
		var i CountProductsByCategoryRow
		if err := rows.Scan(&i.CategoryID, &i.CategoryName, &i.ProductCount); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const countProductsBySupplier = `-- name: CountProductsBySupplier :many
SELECT 
  s.supplier_id,
  s.company_name,
  COUNT(*) as product_count
FROM products p
JOIN suppliers s ON p.supplier_id = s.supplier_id
GROUP BY s.supplier_id, s.company_name
ORDER BY COUNT(*) DESC
`

type CountProductsBySupplierRow struct {
	SupplierID   int16  `json:"supplier_id"`
	CompanyName  string `json:"company_name"`
	ProductCount int64  `json:"product_count"`
}

// Counts products grouped by supplier
func (q *Queries) CountProductsBySupplier(ctx context.Context) ([]CountProductsBySupplierRow, error) {
	rows, err := q.db.Query(ctx, countProductsBySupplier)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []CountProductsBySupplierRow{}
	for rows.Next() {
		var i CountProductsBySupplierRow
		if err := rows.Scan(&i.SupplierID, &i.CompanyName, &i.ProductCount); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const createProduct = `-- name: CreateProduct :one
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
RETURNING product_id, product_name, supplier_id, category_id, quantity_per_unit, unit_price, units_in_stock, units_on_order, reorder_level, discontinued
`

type CreateProductParams struct {
	ProductName     string        `json:"product_name"`
	SupplierID      pgtype.Int2   `json:"supplier_id"`
	CategoryID      pgtype.Int2   `json:"category_id"`
	QuantityPerUnit pgtype.Text   `json:"quantity_per_unit"`
	UnitPrice       pgtype.Float4 `json:"unit_price"`
	UnitsInStock    pgtype.Int2   `json:"units_in_stock"`
	UnitsOnOrder    pgtype.Int2   `json:"units_on_order"`
	ReorderLevel    pgtype.Int2   `json:"reorder_level"`
	Discontinued    int32         `json:"discontinued"`
}

// Creates a new product and returns it
func (q *Queries) CreateProduct(ctx context.Context, arg CreateProductParams) (Product, error) {
	row := q.db.QueryRow(ctx, createProduct,
		arg.ProductName,
		arg.SupplierID,
		arg.CategoryID,
		arg.QuantityPerUnit,
		arg.UnitPrice,
		arg.UnitsInStock,
		arg.UnitsOnOrder,
		arg.ReorderLevel,
		arg.Discontinued,
	)
	var i Product
	err := row.Scan(
		&i.ProductID,
		&i.ProductName,
		&i.SupplierID,
		&i.CategoryID,
		&i.QuantityPerUnit,
		&i.UnitPrice,
		&i.UnitsInStock,
		&i.UnitsOnOrder,
		&i.ReorderLevel,
		&i.Discontinued,
	)
	return i, err
}

const deleteProduct = `-- name: DeleteProduct :exec
DELETE FROM products
WHERE product_id = $1
`

// Deletes a product by ID
func (q *Queries) DeleteProduct(ctx context.Context, productID int16) error {
	_, err := q.db.Exec(ctx, deleteProduct, productID)
	return err
}

const discontinueProduct = `-- name: DiscontinueProduct :one
UPDATE products
SET
  discontinued = 1
WHERE product_id = $1
RETURNING product_id, product_name, supplier_id, category_id, quantity_per_unit, unit_price, units_in_stock, units_on_order, reorder_level, discontinued
`

// Marks a product as discontinued
func (q *Queries) DiscontinueProduct(ctx context.Context, productID int16) (Product, error) {
	row := q.db.QueryRow(ctx, discontinueProduct, productID)
	var i Product
	err := row.Scan(
		&i.ProductID,
		&i.ProductName,
		&i.SupplierID,
		&i.CategoryID,
		&i.QuantityPerUnit,
		&i.UnitPrice,
		&i.UnitsInStock,
		&i.UnitsOnOrder,
		&i.ReorderLevel,
		&i.Discontinued,
	)
	return i, err
}

const getProduct = `-- name: GetProduct :one
SELECT product_id, product_name, supplier_id, category_id, quantity_per_unit, unit_price, units_in_stock, units_on_order, reorder_level, discontinued
FROM products
WHERE product_id = $1
`

// Gets a product by ID
func (q *Queries) GetProduct(ctx context.Context, productID int16) (Product, error) {
	row := q.db.QueryRow(ctx, getProduct, productID)
	var i Product
	err := row.Scan(
		&i.ProductID,
		&i.ProductName,
		&i.SupplierID,
		&i.CategoryID,
		&i.QuantityPerUnit,
		&i.UnitPrice,
		&i.UnitsInStock,
		&i.UnitsOnOrder,
		&i.ReorderLevel,
		&i.Discontinued,
	)
	return i, err
}

const getProductValueByCategory = `-- name: GetProductValueByCategory :many
SELECT 
  c.category_id,
  c.category_name,
  SUM(p.unit_price * p.units_in_stock) as inventory_value
FROM products p
JOIN categories c ON p.category_id = c.category_id
GROUP BY c.category_id, c.category_name
ORDER BY SUM(p.unit_price * p.units_in_stock) DESC
`

type GetProductValueByCategoryRow struct {
	CategoryID     int16  `json:"category_id"`
	CategoryName   string `json:"category_name"`
	InventoryValue int64  `json:"inventory_value"`
}

// Gets the total inventory value by category
func (q *Queries) GetProductValueByCategory(ctx context.Context) ([]GetProductValueByCategoryRow, error) {
	rows, err := q.db.Query(ctx, getProductValueByCategory)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetProductValueByCategoryRow{}
	for rows.Next() {
		var i GetProductValueByCategoryRow
		if err := rows.Scan(&i.CategoryID, &i.CategoryName, &i.InventoryValue); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getProductWithDetails = `-- name: GetProductWithDetails :one
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
WHERE p.product_id = $1
`

type GetProductWithDetailsRow struct {
	ProductID           int16         `json:"product_id"`
	ProductName         string        `json:"product_name"`
	SupplierID          pgtype.Int2   `json:"supplier_id"`
	CategoryID          pgtype.Int2   `json:"category_id"`
	QuantityPerUnit     pgtype.Text   `json:"quantity_per_unit"`
	UnitPrice           pgtype.Float4 `json:"unit_price"`
	UnitsInStock        pgtype.Int2   `json:"units_in_stock"`
	UnitsOnOrder        pgtype.Int2   `json:"units_on_order"`
	ReorderLevel        pgtype.Int2   `json:"reorder_level"`
	Discontinued        int32         `json:"discontinued"`
	CategoryName        pgtype.Text   `json:"category_name"`
	CategoryDescription pgtype.Text   `json:"category_description"`
	SupplierName        pgtype.Text   `json:"supplier_name"`
	SupplierContact     pgtype.Text   `json:"supplier_contact"`
	SupplierCountry     pgtype.Text   `json:"supplier_country"`
}

// Gets a product by ID with category and supplier details
func (q *Queries) GetProductWithDetails(ctx context.Context, productID int16) (GetProductWithDetailsRow, error) {
	row := q.db.QueryRow(ctx, getProductWithDetails, productID)
	var i GetProductWithDetailsRow
	err := row.Scan(
		&i.ProductID,
		&i.ProductName,
		&i.SupplierID,
		&i.CategoryID,
		&i.QuantityPerUnit,
		&i.UnitPrice,
		&i.UnitsInStock,
		&i.UnitsOnOrder,
		&i.ReorderLevel,
		&i.Discontinued,
		&i.CategoryName,
		&i.CategoryDescription,
		&i.SupplierName,
		&i.SupplierContact,
		&i.SupplierCountry,
	)
	return i, err
}

const listDiscontinuedProducts = `-- name: ListDiscontinuedProducts :many
SELECT product_id, product_name, supplier_id, category_id, quantity_per_unit, unit_price, units_in_stock, units_on_order, reorder_level, discontinued
FROM products
WHERE discontinued = 1
ORDER BY product_name
`

// Lists all discontinued products
func (q *Queries) ListDiscontinuedProducts(ctx context.Context) ([]Product, error) {
	rows, err := q.db.Query(ctx, listDiscontinuedProducts)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Product{}
	for rows.Next() {
		var i Product
		if err := rows.Scan(
			&i.ProductID,
			&i.ProductName,
			&i.SupplierID,
			&i.CategoryID,
			&i.QuantityPerUnit,
			&i.UnitPrice,
			&i.UnitsInStock,
			&i.UnitsOnOrder,
			&i.ReorderLevel,
			&i.Discontinued,
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

const listProducts = `-- name: ListProducts :many
SELECT product_id, product_name, supplier_id, category_id, quantity_per_unit, unit_price, units_in_stock, units_on_order, reorder_level, discontinued
FROM products
ORDER BY product_name
`

// Lists all products
func (q *Queries) ListProducts(ctx context.Context) ([]Product, error) {
	rows, err := q.db.Query(ctx, listProducts)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Product{}
	for rows.Next() {
		var i Product
		if err := rows.Scan(
			&i.ProductID,
			&i.ProductName,
			&i.SupplierID,
			&i.CategoryID,
			&i.QuantityPerUnit,
			&i.UnitPrice,
			&i.UnitsInStock,
			&i.UnitsOnOrder,
			&i.ReorderLevel,
			&i.Discontinued,
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

const listProductsByCategory = `-- name: ListProductsByCategory :many
SELECT product_id, product_name, supplier_id, category_id, quantity_per_unit, unit_price, units_in_stock, units_on_order, reorder_level, discontinued
FROM products
WHERE category_id = $1
ORDER BY product_name
`

// Lists all products in a specific category
func (q *Queries) ListProductsByCategory(ctx context.Context, categoryID pgtype.Int2) ([]Product, error) {
	rows, err := q.db.Query(ctx, listProductsByCategory, categoryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Product{}
	for rows.Next() {
		var i Product
		if err := rows.Scan(
			&i.ProductID,
			&i.ProductName,
			&i.SupplierID,
			&i.CategoryID,
			&i.QuantityPerUnit,
			&i.UnitPrice,
			&i.UnitsInStock,
			&i.UnitsOnOrder,
			&i.ReorderLevel,
			&i.Discontinued,
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

const listProductsBySupplier = `-- name: ListProductsBySupplier :many
SELECT product_id, product_name, supplier_id, category_id, quantity_per_unit, unit_price, units_in_stock, units_on_order, reorder_level, discontinued
FROM products
WHERE supplier_id = $1
ORDER BY product_name
`

// Lists all products from a specific supplier
func (q *Queries) ListProductsBySupplier(ctx context.Context, supplierID pgtype.Int2) ([]Product, error) {
	rows, err := q.db.Query(ctx, listProductsBySupplier, supplierID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Product{}
	for rows.Next() {
		var i Product
		if err := rows.Scan(
			&i.ProductID,
			&i.ProductName,
			&i.SupplierID,
			&i.CategoryID,
			&i.QuantityPerUnit,
			&i.UnitPrice,
			&i.UnitsInStock,
			&i.UnitsOnOrder,
			&i.ReorderLevel,
			&i.Discontinued,
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

const listProductsNeedingReorder = `-- name: ListProductsNeedingReorder :many
SELECT product_id, product_name, supplier_id, category_id, quantity_per_unit, unit_price, units_in_stock, units_on_order, reorder_level, discontinued
FROM products
WHERE units_in_stock <= reorder_level AND discontinued = 0
ORDER BY product_name
`

// Lists all products that need to be reordered (stock below reorder level)
func (q *Queries) ListProductsNeedingReorder(ctx context.Context) ([]Product, error) {
	rows, err := q.db.Query(ctx, listProductsNeedingReorder)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Product{}
	for rows.Next() {
		var i Product
		if err := rows.Scan(
			&i.ProductID,
			&i.ProductName,
			&i.SupplierID,
			&i.CategoryID,
			&i.QuantityPerUnit,
			&i.UnitPrice,
			&i.UnitsInStock,
			&i.UnitsOnOrder,
			&i.ReorderLevel,
			&i.Discontinued,
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

const listProductsWithDetails = `-- name: ListProductsWithDetails :many
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
ORDER BY p.product_name
`

type ListProductsWithDetailsRow struct {
	ProductID       int16         `json:"product_id"`
	ProductName     string        `json:"product_name"`
	SupplierID      pgtype.Int2   `json:"supplier_id"`
	CategoryID      pgtype.Int2   `json:"category_id"`
	QuantityPerUnit pgtype.Text   `json:"quantity_per_unit"`
	UnitPrice       pgtype.Float4 `json:"unit_price"`
	UnitsInStock    pgtype.Int2   `json:"units_in_stock"`
	UnitsOnOrder    pgtype.Int2   `json:"units_on_order"`
	ReorderLevel    pgtype.Int2   `json:"reorder_level"`
	Discontinued    int32         `json:"discontinued"`
	CategoryName    pgtype.Text   `json:"category_name"`
	SupplierName    pgtype.Text   `json:"supplier_name"`
}

// Lists all products with category and supplier details
func (q *Queries) ListProductsWithDetails(ctx context.Context) ([]ListProductsWithDetailsRow, error) {
	rows, err := q.db.Query(ctx, listProductsWithDetails)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListProductsWithDetailsRow{}
	for rows.Next() {
		var i ListProductsWithDetailsRow
		if err := rows.Scan(
			&i.ProductID,
			&i.ProductName,
			&i.SupplierID,
			&i.CategoryID,
			&i.QuantityPerUnit,
			&i.UnitPrice,
			&i.UnitsInStock,
			&i.UnitsOnOrder,
			&i.ReorderLevel,
			&i.Discontinued,
			&i.CategoryName,
			&i.SupplierName,
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

const searchProductsByName = `-- name: SearchProductsByName :many
SELECT product_id, product_name, supplier_id, category_id, quantity_per_unit, unit_price, units_in_stock, units_on_order, reorder_level, discontinued
FROM products
WHERE product_name ILIKE '%' || $1 || '%'
ORDER BY product_name
`

// Searches products by name (case insensitive)
func (q *Queries) SearchProductsByName(ctx context.Context, dollar_1 pgtype.Text) ([]Product, error) {
	rows, err := q.db.Query(ctx, searchProductsByName, dollar_1)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Product{}
	for rows.Next() {
		var i Product
		if err := rows.Scan(
			&i.ProductID,
			&i.ProductName,
			&i.SupplierID,
			&i.CategoryID,
			&i.QuantityPerUnit,
			&i.UnitPrice,
			&i.UnitsInStock,
			&i.UnitsOnOrder,
			&i.ReorderLevel,
			&i.Discontinued,
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

const updateProduct = `-- name: UpdateProduct :one
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
RETURNING product_id, product_name, supplier_id, category_id, quantity_per_unit, unit_price, units_in_stock, units_on_order, reorder_level, discontinued
`

type UpdateProductParams struct {
	ProductID       int16         `json:"product_id"`
	ProductName     string        `json:"product_name"`
	SupplierID      pgtype.Int2   `json:"supplier_id"`
	CategoryID      pgtype.Int2   `json:"category_id"`
	QuantityPerUnit pgtype.Text   `json:"quantity_per_unit"`
	UnitPrice       pgtype.Float4 `json:"unit_price"`
	UnitsInStock    pgtype.Int2   `json:"units_in_stock"`
	UnitsOnOrder    pgtype.Int2   `json:"units_on_order"`
	ReorderLevel    pgtype.Int2   `json:"reorder_level"`
	Discontinued    int32         `json:"discontinued"`
}

// Updates a product by ID
func (q *Queries) UpdateProduct(ctx context.Context, arg UpdateProductParams) (Product, error) {
	row := q.db.QueryRow(ctx, updateProduct,
		arg.ProductID,
		arg.ProductName,
		arg.SupplierID,
		arg.CategoryID,
		arg.QuantityPerUnit,
		arg.UnitPrice,
		arg.UnitsInStock,
		arg.UnitsOnOrder,
		arg.ReorderLevel,
		arg.Discontinued,
	)
	var i Product
	err := row.Scan(
		&i.ProductID,
		&i.ProductName,
		&i.SupplierID,
		&i.CategoryID,
		&i.QuantityPerUnit,
		&i.UnitPrice,
		&i.UnitsInStock,
		&i.UnitsOnOrder,
		&i.ReorderLevel,
		&i.Discontinued,
	)
	return i, err
}

const updateProductPrice = `-- name: UpdateProductPrice :one
UPDATE products
SET
  unit_price = $2
WHERE product_id = $1
RETURNING product_id, product_name, supplier_id, category_id, quantity_per_unit, unit_price, units_in_stock, units_on_order, reorder_level, discontinued
`

type UpdateProductPriceParams struct {
	ProductID int16         `json:"product_id"`
	UnitPrice pgtype.Float4 `json:"unit_price"`
}

// Updates a product's price
func (q *Queries) UpdateProductPrice(ctx context.Context, arg UpdateProductPriceParams) (Product, error) {
	row := q.db.QueryRow(ctx, updateProductPrice, arg.ProductID, arg.UnitPrice)
	var i Product
	err := row.Scan(
		&i.ProductID,
		&i.ProductName,
		&i.SupplierID,
		&i.CategoryID,
		&i.QuantityPerUnit,
		&i.UnitPrice,
		&i.UnitsInStock,
		&i.UnitsOnOrder,
		&i.ReorderLevel,
		&i.Discontinued,
	)
	return i, err
}

const updateProductStock = `-- name: UpdateProductStock :one
UPDATE products
SET
  units_in_stock = $2,
  units_on_order = $3
WHERE product_id = $1
RETURNING product_id, product_name, supplier_id, category_id, quantity_per_unit, unit_price, units_in_stock, units_on_order, reorder_level, discontinued
`

type UpdateProductStockParams struct {
	ProductID    int16       `json:"product_id"`
	UnitsInStock pgtype.Int2 `json:"units_in_stock"`
	UnitsOnOrder pgtype.Int2 `json:"units_on_order"`
}

// Updates a product's stock levels
func (q *Queries) UpdateProductStock(ctx context.Context, arg UpdateProductStockParams) (Product, error) {
	row := q.db.QueryRow(ctx, updateProductStock, arg.ProductID, arg.UnitsInStock, arg.UnitsOnOrder)
	var i Product
	err := row.Scan(
		&i.ProductID,
		&i.ProductName,
		&i.SupplierID,
		&i.CategoryID,
		&i.QuantityPerUnit,
		&i.UnitPrice,
		&i.UnitsInStock,
		&i.UnitsOnOrder,
		&i.ReorderLevel,
		&i.Discontinued,
	)
	return i, err
}
