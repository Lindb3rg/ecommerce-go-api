-- name: GetCategory :one
-- Gets a category by ID
SELECT *
FROM categories
WHERE category_id = $1;

-- name: ListCategories :many
-- Lists all categories
SELECT *
FROM categories
ORDER BY category_name;

-- name: CreateCategory :one
-- Creates a new category and returns it
INSERT INTO categories (
  category_name,
  description,
  picture
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: UpdateCategory :one
-- Updates a category by ID
UPDATE categories
SET
  category_name = $2,
  description = $3,
  picture = $4
WHERE category_id = $1
RETURNING *;

-- name: DeleteCategory :exec
-- Deletes a category by ID
DELETE FROM categories
WHERE category_id = $1;

-- name: SearchCategoriesByName :many
-- Searches categories by name (case insensitive)
SELECT *
FROM categories
WHERE category_name ILIKE '%' || $1 || '%'
ORDER BY category_name;

-- name: CountCategories :one
-- Counts the total number of categories
SELECT COUNT(*) FROM categories;