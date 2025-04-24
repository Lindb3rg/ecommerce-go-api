-- name: GetEmployeeTerritoryRelation :one
-- Gets a specific employee-territory relation
SELECT *
FROM employee_territories
WHERE employee_id = $1 AND territory_id = $2;

-- name: ListTerritoriesByEmployee :many
-- Lists all territories assigned to a specific employee
SELECT t.*
FROM employee_territories et
JOIN territories t ON et.territory_id = t.territory_id
WHERE et.employee_id = $1
ORDER BY t.territory_id;

-- name: ListEmployeesByTerritory :many
-- Lists all employees assigned to a specific territory
SELECT e.*
FROM employee_territories et
JOIN employees e ON et.employee_id = e.employee_id
WHERE et.territory_id = $1
ORDER BY e.last_name, e.first_name;

-- name: CreateEmployeeTerritoryRelation :one
-- Assigns a territory to an employee
INSERT INTO employee_territories (
  employee_id,
  territory_id
) VALUES (
  $1, $2
)
RETURNING *;

-- name: DeleteEmployeeTerritoryRelation :exec
-- Removes a specific territory assignment from an employee
DELETE FROM employee_territories
WHERE employee_id = $1 AND territory_id = $2;

-- name: DeleteAllTerritoryAssignmentsForEmployee :exec
-- Removes all territory assignments for a specific employee
DELETE FROM employee_territories
WHERE employee_id = $1;

-- name: DeleteAllEmployeeAssignmentsForTerritory :exec
-- Removes all employee assignments for a specific territory
DELETE FROM employee_territories
WHERE territory_id = $1;

-- name: ListEmployeesWithTerritoriesAndRegions :many
-- Lists employees with their territories and regions
SELECT 
  e.employee_id,
  e.first_name,
  e.last_name,
  t.territory_id,
  t.territory_description,
  r.region_id,
  r.region_description
FROM employee_territories et
JOIN employees e ON et.employee_id = e.employee_id
JOIN territories t ON et.territory_id = t.territory_id
JOIN region r ON t.region_id = r.region_id
ORDER BY e.last_name, e.first_name, r.region_id, t.territory_id;

-- name: CountTerritoriesByEmployee :many
-- Counts territories grouped by employee
SELECT 
  e.employee_id,
  e.first_name,
  e.last_name,
  COUNT(*) as territory_count
FROM employee_territories et
JOIN employees e ON et.employee_id = e.employee_id
GROUP BY e.employee_id, e.first_name, e.last_name
ORDER BY COUNT(*) DESC;

-- name: CountEmployeesByTerritory :many
-- Counts employees grouped by territory
SELECT 
  t.territory_id,
  t.territory_description,
  COUNT(*) as employee_count
FROM employee_territories et
JOIN territories t ON et.territory_id = t.territory_id
GROUP BY t.territory_id, t.territory_description
ORDER BY COUNT(*) DESC;

-- name: ListEmployeesByRegion :many
-- Lists employees assigned to territories in a specific region
SELECT DISTINCT
  e.employee_id,
  e.first_name,
  e.last_name,
  e.title
FROM employee_territories et
JOIN territories t ON et.territory_id = t.territory_id
JOIN employees e ON et.employee_id = e.employee_id
WHERE t.region_id = $1
ORDER BY e.last_name, e.first_name;