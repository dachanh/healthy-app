-- name: GetExercseHistoryOfUser :many
SELECT eh.date, e.name, eh.calories_burned, eh.timer 
FROM public."Exercise_History" as eh 
INNER JOIN public."Exercise" as e ON eh.exercise_id = e."ID" 
WHERE eh.user_id = $1
ORDER BY eh.date
LIMIT $2 OFFSET $3;

-- name: GetExercseHistoryByDate :many
SELECT eh.date, e.name, eh.calories_burned, eh.timer 
FROM public."Exercise_History" eh 
INNER JOIN public."Exercise" e ON eh.exercise_id = e."ID" 
WHERE eh.user_id =  $1 AND eh.date BETWEEN  $2 AND $3
ORDER BY eh.date DESC
LIMIT  $4 OFFSET  $5;

