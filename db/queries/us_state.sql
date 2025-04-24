-- name: GetState :one
-- Gets a state by its primary key (employee_id).
SELECT * FROM us_states
WHERE employee_id = $1;

-- name: ListStates :many
-- Lists all states, optionally filtered by region.
SELECT * FROM us_states
WHERE 
    CASE 
        WHEN $1::varchar IS NULL THEN true
        ELSE state_region = $1
    END
ORDER BY state_name;

-- name: CreateState :one
-- Creates a new state record.
INSERT INTO us_states (
    state_name, 
    state_abbr, 
    state_region
) VALUES (
    $1, $2, $3
)
RETURNING *;

-- name: UpdateState :one
-- Updates an existing state record.
UPDATE us_states
SET 
    state_name = COALESCE($2, state_name),
    state_abbr = COALESCE($3, state_abbr),
    state_region = COALESCE($4, state_region)
WHERE employee_id = $1
RETURNING *;

-- name: DeleteState :exec
-- Deletes a state by its primary key (employee_id).
DELETE FROM us_states
WHERE employee_id = $1;

-- name: GetStateByAbbreviation :one
-- Gets a state by its abbreviation.
SELECT * FROM us_states
WHERE state_abbr = $1;

-- name: GetStatesByRegion :many
-- Gets all states in a specific region.
SELECT * FROM us_states
WHERE state_region = $1
ORDER BY state_name;

-- name: CountStatesByRegion :one
-- Counts the number of states in each region.
SELECT state_region, COUNT(*) as state_count
FROM us_states
GROUP BY state_region
ORDER BY state_count DESC;