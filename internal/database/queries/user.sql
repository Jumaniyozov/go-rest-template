-- name: ListUsers :many
SELECT *
FROM user
OFFSET $1 LIMIT $2;
