-- name: ListNews :many
SELECT n.*, nc.name AS category_name
FROM public."News" as n
INNER JOIN public."News_Category" as nc ON n.category_id = nc."ID"
ORDER BY n."created_at" DESC
LIMIT $1 OFFSET $2;