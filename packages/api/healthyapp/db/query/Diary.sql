-- name: CreateDiary :one
INSERT INTO public."Diary" (content, user_id) 
VALUES ($1, $2)
RETURNING *;

-- name: GetUserDiary :many
SELECT d."ID", d."content", d."created_at", d."updated_at"
FROM public."Diary" as d
WHERE user_id = $1
ORDER BY created_at DESC
LIMIT $2 OFFSET $3;