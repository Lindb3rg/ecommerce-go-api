-- name: GetTerritory :one
-- Gets a territory by ID
SELECT *
FROM territories
WHERE territory_id = $1;

-- name: ListTerritories :many
-- Lists all territories
SELECT *
FROM territories
ORDER BY territory_id;

-- name: CreateTerritory :one
-- Creates a new territory and returns it
INSERT INTO territories (
  territory_id,
  territory_description,
  region_id
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: UpdateTerritory :one
-- Updates a territory by ID
UPDATE territories
SET
  territory_description = $2,
  region_id = $3
WHERE territory_id = $1
RETURNING *;

-- name: DeleteTerritory :exec
-- Deletes a territory by ID
DELETE FROM territories
WHERE territory_id = $1;

-- name: ListTerritoriesByRegion :many
-- Lists all territories in a specific region
SELECT *
FROM territories
WHERE region_id = $1
ORDER BY territory_id;

-- name: GetTerritoryWithRegion :one
-- Gets a territory by ID with its region details
SELECT 
  t.territory_id,
  t.territory_description,
  t.region_id,
  r.region_description
FROM territories t
JOIN region r ON t.region_id = r.region_id
WHERE t.territory_id = $1;

-- name: ListTerritoriesWithRegion :many
-- Lists all territories with their region details
SELECT 
  t.territory_id,
  t.territory_description,
  t.region_id,
  r.region_description
FROM territories t
JOIN region r ON t.region_id = r.region_id
ORDER BY t.territory_id;

-- name: SearchTerritoriesByDescription :many
-- Searches territories by description (case insensitive)
SELECT *
FROM territories
WHERE territory_description ILIKE '%' || $1 || '%'
ORDER BY territory_id;

-- name: CountTerritories :one
-- Counts the total number of territories
SELECT COUNT(*) FROM territories;

-- name: CountTerritoriesByRegion :many
-- Counts territories grouped by region
SELECT 
  r.region_id,
  r.region_description,
  COUNT(*) as territory_count
FROM territories t
JOIN region r ON t.region_id = r.region_id
GROUP BY r.region_id, r.region_description
ORDER BY COUNT(*) DESC;