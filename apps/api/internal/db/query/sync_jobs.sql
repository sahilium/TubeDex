-- name: CreateSyncJob :one
INSERT INTO sync_jobs (user_id, status)
VALUES ($1, 'pending')
RETURNING *;

-- name: GetLatestSyncJob :one
SELECT * FROM sync_jobs
WHERE user_id = $1
ORDER BY started_at DESC
LIMIT 1;

-- name: UpdateSyncJobStatus :exec
UPDATE sync_jobs
SET status = $2, finished_at = CASE WHEN $2 IN ('completed', 'failed') THEN NOW() ELSE NULL END, error = $3
WHERE id = $1;

-- name: ListSyncJobsByUser :many
SELECT * FROM sync_jobs
WHERE user_id = $1
ORDER BY started_at DESC
LIMIT $2 OFFSET $3;
