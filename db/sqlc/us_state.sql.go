// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: us_state.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const countStatesByRegion = `-- name: CountStatesByRegion :one
SELECT state_region, COUNT(*) as state_count
FROM us_states
GROUP BY state_region
ORDER BY state_count DESC
`

type CountStatesByRegionRow struct {
	StateRegion pgtype.Text `json:"state_region"`
	StateCount  int64       `json:"state_count"`
}

// Counts the number of states in each region.
func (q *Queries) CountStatesByRegion(ctx context.Context) (CountStatesByRegionRow, error) {
	row := q.db.QueryRow(ctx, countStatesByRegion)
	var i CountStatesByRegionRow
	err := row.Scan(&i.StateRegion, &i.StateCount)
	return i, err
}

const createState = `-- name: CreateState :one
INSERT INTO us_states (
    state_name, 
    state_abbr, 
    state_region
) VALUES (
    $1, $2, $3
)
RETURNING employee_id, state_name, state_abbr, state_region
`

type CreateStateParams struct {
	StateName   pgtype.Text `json:"state_name"`
	StateAbbr   pgtype.Text `json:"state_abbr"`
	StateRegion pgtype.Text `json:"state_region"`
}

// Creates a new state record.
func (q *Queries) CreateState(ctx context.Context, arg CreateStateParams) (UsState, error) {
	row := q.db.QueryRow(ctx, createState, arg.StateName, arg.StateAbbr, arg.StateRegion)
	var i UsState
	err := row.Scan(
		&i.EmployeeID,
		&i.StateName,
		&i.StateAbbr,
		&i.StateRegion,
	)
	return i, err
}

const deleteState = `-- name: DeleteState :exec
DELETE FROM us_states
WHERE employee_id = $1
`

// Deletes a state by its primary key (employee_id).
func (q *Queries) DeleteState(ctx context.Context, employeeID int16) error {
	_, err := q.db.Exec(ctx, deleteState, employeeID)
	return err
}

const getState = `-- name: GetState :one
SELECT employee_id, state_name, state_abbr, state_region FROM us_states
WHERE employee_id = $1
`

// Gets a state by its primary key (employee_id).
func (q *Queries) GetState(ctx context.Context, employeeID int16) (UsState, error) {
	row := q.db.QueryRow(ctx, getState, employeeID)
	var i UsState
	err := row.Scan(
		&i.EmployeeID,
		&i.StateName,
		&i.StateAbbr,
		&i.StateRegion,
	)
	return i, err
}

const getStateByAbbreviation = `-- name: GetStateByAbbreviation :one
SELECT employee_id, state_name, state_abbr, state_region FROM us_states
WHERE state_abbr = $1
`

// Gets a state by its abbreviation.
func (q *Queries) GetStateByAbbreviation(ctx context.Context, stateAbbr pgtype.Text) (UsState, error) {
	row := q.db.QueryRow(ctx, getStateByAbbreviation, stateAbbr)
	var i UsState
	err := row.Scan(
		&i.EmployeeID,
		&i.StateName,
		&i.StateAbbr,
		&i.StateRegion,
	)
	return i, err
}

const getStatesByRegion = `-- name: GetStatesByRegion :many
SELECT employee_id, state_name, state_abbr, state_region FROM us_states
WHERE state_region = $1
ORDER BY state_name
`

// Gets all states in a specific region.
func (q *Queries) GetStatesByRegion(ctx context.Context, stateRegion pgtype.Text) ([]UsState, error) {
	rows, err := q.db.Query(ctx, getStatesByRegion, stateRegion)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []UsState{}
	for rows.Next() {
		var i UsState
		if err := rows.Scan(
			&i.EmployeeID,
			&i.StateName,
			&i.StateAbbr,
			&i.StateRegion,
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

const listStates = `-- name: ListStates :many
SELECT employee_id, state_name, state_abbr, state_region FROM us_states
WHERE 
    CASE 
        WHEN $1::varchar IS NULL THEN true
        ELSE state_region = $1
    END
ORDER BY state_name
`

// Lists all states, optionally filtered by region.
func (q *Queries) ListStates(ctx context.Context, dollar_1 string) ([]UsState, error) {
	rows, err := q.db.Query(ctx, listStates, dollar_1)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []UsState{}
	for rows.Next() {
		var i UsState
		if err := rows.Scan(
			&i.EmployeeID,
			&i.StateName,
			&i.StateAbbr,
			&i.StateRegion,
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

const updateState = `-- name: UpdateState :one
UPDATE us_states
SET 
    state_name = COALESCE($2, state_name),
    state_abbr = COALESCE($3, state_abbr),
    state_region = COALESCE($4, state_region)
WHERE employee_id = $1
RETURNING employee_id, state_name, state_abbr, state_region
`

type UpdateStateParams struct {
	EmployeeID  int16       `json:"employee_id"`
	StateName   pgtype.Text `json:"state_name"`
	StateAbbr   pgtype.Text `json:"state_abbr"`
	StateRegion pgtype.Text `json:"state_region"`
}

// Updates an existing state record.
func (q *Queries) UpdateState(ctx context.Context, arg UpdateStateParams) (UsState, error) {
	row := q.db.QueryRow(ctx, updateState,
		arg.EmployeeID,
		arg.StateName,
		arg.StateAbbr,
		arg.StateRegion,
	)
	var i UsState
	err := row.Scan(
		&i.EmployeeID,
		&i.StateName,
		&i.StateAbbr,
		&i.StateRegion,
	)
	return i, err
}
