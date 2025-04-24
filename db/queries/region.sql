-- name: GetRegion :one
-- Gets a region by ID
SELECT *
FROM region
WHERE region_id = $1;

-- name: ListRegions :many
-- Lists all regions
SELECT *
FROM region
ORDER BY region_id;

-- name: CreateRegion :one
-- Creates a new region and returns it
INSERT INTO region (
  region_description
) VALUES (
  $1
)
RETURNING *;

-- name: UpdateRegion :one
-- Updates a region by ID
UPDATE region
SET
  region_description = $2
WHERE region_id = $1
RETURNING *;

-- name: DeleteRegion :exec
-- Deletes a region by ID
DELETE FROM region
WHERE region_id = $1;

-- name: SearchRegionsByDescription :many
-- Searches regions by description (case insensitive)
SELECT *
FROM region
WHERE region_description ILIKE '%' || $1 || '%'
ORDER BY region_id;

-- name: CountRegions :one
-- Counts the total number of regions
SELECT COUNT(*) FROM region;

-- name: GetRegionByDescription :one
-- Gets a region by exact description (useful for lookups)
SELECT *
FROM region
WHERE region_description = $1;