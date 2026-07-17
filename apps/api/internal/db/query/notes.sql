-- name: GetNote :one
SELECT * FROM notes WHERE user_id = $1 AND channel_id = $2;

-- name: UpsertNote :one
INSERT INTO notes (user_id, channel_id, body)
VALUES ($1, $2, $3)
ON CONFLICT (user_id, channel_id) DO UPDATE SET
    body = EXCLUDED.body,
    updated_at = NOW()
RETURNING *;

-- name: DeleteNote :exec
DELETE FROM notes WHERE user_id = $1 AND channel_id = $2;

-- name: SearchNotes :many
SELECT n.*, c.name as channel_name, c.youtube_channel_id, c.avatar as channel_avatar
FROM notes n
JOIN channels c ON c.id = n.channel_id
WHERE n.user_id = $1 AND n.body ILIKE '%' || $2 || '%'
ORDER BY n.updated_at DESC;
