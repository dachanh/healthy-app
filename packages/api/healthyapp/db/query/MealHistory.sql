-- name: GetMealHistory :many
SELECT *
FROM public."Meal_History" as mh
WHERE mh."user_id" = $1
ORDER BY mh."date" DESC, mh."session" DESC
LIMIT $2 OFFSET $3;

-- name: GetMealHistoryBySession :many
SELECT * FROM public."Meal_History" as mh
WHERE mh."user_id" = $1 AND mh."session" = $2
ORDER BY mh."date" DESC
LIMIT $3 OFFSET $4;