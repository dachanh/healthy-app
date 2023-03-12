// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.0
// source: ExceriseHistory.sql

package dbmodels

import (
	"context"
	"time"
)

const getExercseHistoryByDate = `-- name: GetExercseHistoryByDate :many
SELECT eh.date, e.name, eh.calories_burned, eh.timer 
FROM public."Exercise_History" eh 
INNER JOIN public."Exercise" e ON eh.exercise_id = e."ID" 
WHERE eh.user_id =  $1 AND eh.date BETWEEN  $2 AND $3
ORDER BY eh.date DESC
LIMIT  $4 OFFSET  $5
`

type GetExercseHistoryByDateParams struct {
	UserID int32     `json:"user_id"`
	Date   time.Time `json:"date"`
	Date_2 time.Time `json:"date_2"`
	Limit  int32     `json:"limit"`
	Offset int32     `json:"offset"`
}

type GetExercseHistoryByDateRow struct {
	Date           time.Time `json:"date"`
	Name           string    `json:"name"`
	CaloriesBurned int32     `json:"calories_burned"`
	Timer          int32     `json:"timer"`
}

func (q *Queries) GetExercseHistoryByDate(ctx context.Context, arg GetExercseHistoryByDateParams) ([]GetExercseHistoryByDateRow, error) {
	rows, err := q.db.QueryContext(ctx, getExercseHistoryByDate,
		arg.UserID,
		arg.Date,
		arg.Date_2,
		arg.Limit,
		arg.Offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetExercseHistoryByDateRow{}
	for rows.Next() {
		var i GetExercseHistoryByDateRow
		if err := rows.Scan(
			&i.Date,
			&i.Name,
			&i.CaloriesBurned,
			&i.Timer,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getExercseHistoryOfUser = `-- name: GetExercseHistoryOfUser :many
SELECT eh.date, e.name, eh.calories_burned, eh.timer 
FROM public."Exercise_History" as eh 
INNER JOIN public."Exercise" as e ON eh.exercise_id = e."ID" 
WHERE eh.user_id = $1
ORDER BY eh.date
LIMIT $2 OFFSET $3
`

type GetExercseHistoryOfUserParams struct {
	UserID int32 `json:"user_id"`
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

type GetExercseHistoryOfUserRow struct {
	Date           time.Time `json:"date"`
	Name           string    `json:"name"`
	CaloriesBurned int32     `json:"calories_burned"`
	Timer          int32     `json:"timer"`
}

func (q *Queries) GetExercseHistoryOfUser(ctx context.Context, arg GetExercseHistoryOfUserParams) ([]GetExercseHistoryOfUserRow, error) {
	rows, err := q.db.QueryContext(ctx, getExercseHistoryOfUser, arg.UserID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetExercseHistoryOfUserRow{}
	for rows.Next() {
		var i GetExercseHistoryOfUserRow
		if err := rows.Scan(
			&i.Date,
			&i.Name,
			&i.CaloriesBurned,
			&i.Timer,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
