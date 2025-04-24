-- name: GetEmployee :one
-- Gets an employee by ID
SELECT *
FROM employees
WHERE employee_id = $1;

-- name: ListEmployees :many
-- Lists all employees
SELECT *
FROM employees
ORDER BY last_name, first_name;

-- name: CreateEmployee :one
-- Creates a new employee and returns it
INSERT INTO employees (
  last_name,
  first_name,
  title,
  title_of_courtesy,
  birth_date,
  hire_date,
  address,
  city,
  region,
  postal_code,
  country,
  home_phone,
  extension,
  photo,
  notes,
  reports_to,
  photo_path
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17
)
RETURNING *;

-- name: UpdateEmployee :one
-- Updates an employee by ID
UPDATE employees
SET
  last_name = $2,
  first_name = $3,
  title = $4,
  title_of_courtesy = $5,
  birth_date = $6,
  hire_date = $7,
  address = $8,
  city = $9,
  region = $10,
  postal_code = $11,
  country = $12,
  home_phone = $13,
  extension = $14,
  photo = $15,
  notes = $16,
  reports_to = $17,
  photo_path = $18
WHERE employee_id = $1
RETURNING *;

-- name: DeleteEmployee :exec
-- Deletes an employee by ID
DELETE FROM employees
WHERE employee_id = $1;

-- name: SearchEmployeesByName :many
-- Searches employees by first or last name (case insensitive)
SELECT *
FROM employees
WHERE first_name ILIKE '%' || $1 || '%' OR last_name ILIKE '%' || $1 || '%'
ORDER BY last_name, first_name;

-- name: ListEmployeesByManager :many
-- Lists all employees that report to a specific manager
SELECT *
FROM employees
WHERE reports_to = $1
ORDER BY last_name, first_name;

-- name: GetEmployeeWithManager :one
-- Gets an employee by ID along with their manager's details
SELECT 
  e.employee_id,
  e.last_name,
  e.first_name,
  e.title,
  e.title_of_courtesy,
  e.birth_date,
  e.hire_date,
  e.address,
  e.city,
  e.region,
  e.postal_code,
  e.country,
  e.home_phone,
  e.extension,
  e.photo,
  e.notes,
  e.reports_to,
  e.photo_path,
  m.employee_id as manager_id,
  m.last_name as manager_last_name,
  m.first_name as manager_first_name,
  m.title as manager_title
FROM employees e
LEFT JOIN employees m ON e.reports_to = m.employee_id
WHERE e.employee_id = $1;

-- name: ListEmployeesByTitle :many
-- Lists all employees with a specific title
SELECT *
FROM employees
WHERE title = $1
ORDER BY last_name, first_name;

-- name: ListEmployeesByCountry :many
-- Lists all employees from a specific country
SELECT *
FROM employees
WHERE country = $1
ORDER BY last_name, first_name;

-- name: CountEmployees :one
-- Counts the total number of employees
SELECT COUNT(*) FROM employees;

-- name: CountEmployeesByManager :many
-- Counts direct reports for each manager
SELECT 
  m.employee_id,
  m.last_name,
  m.first_name,
  COUNT(e.employee_id) as direct_reports
FROM employees m
LEFT JOIN employees e ON m.employee_id = e.reports_to
GROUP BY m.employee_id, m.last_name, m.first_name
ORDER BY COUNT(e.employee_id) DESC;

-- name: GetEmployeeHierarchy :many
-- Gets the entire reporting hierarchy for an employee
WITH RECURSIVE emp_hierarchy AS (
  SELECT 
    e.employee_id, 
    e.last_name, 
    e.first_name, 
    e.title, 
    e.reports_to, 
    0 as level
  FROM employees e
  WHERE e.employee_id = $1
  
  UNION ALL
  
  SELECT 
    e.employee_id, 
    e.last_name, 
    e.first_name, 
    e.title, 
    e.reports_to, 
    eh.level + 1
  FROM employees e
  JOIN emp_hierarchy eh ON e.reports_to = eh.employee_id
)
SELECT 
  employee_id,
  last_name,
  first_name,
  title,
  reports_to,
  level
FROM emp_hierarchy
ORDER BY level;