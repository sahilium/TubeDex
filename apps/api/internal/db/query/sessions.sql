-- name: GetSessionByID :one
SELECT * FROM sessions WHERE id = $1;

-- name: GetSessionByUserID :one
SELECT * FROM sessions WHERE user_id = $1 ORDER BY created_at DESC LIMIT 1;

-- name: CreateSession :one
INSERT INTO sessions (id, user_id, data, expires_at)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: DeleteSession :exec
DELETE FROM sessions WHERE id = $1;

-- name: DeleteExpiredSessions :exec
DELETE FROM sessions WHERE expires_at < NOW();
