-- name: GetChannelByID :one
SELECT * FROM channels WHERE id = $1;

-- name: GetChannelByYouTubeID :one
SELECT * FROM channels WHERE youtube_channel_id = $1;

-- name: ListChannelsByIDs :many
SELECT * FROM channels WHERE id = ANY($1::bigint[]);

-- name: UpsertChannel :one
INSERT INTO channels (youtube_channel_id, name, handle, description, avatar, banner, subscriber_count, uploads_playlist_id)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
ON CONFLICT (youtube_channel_id) DO UPDATE SET
    name = EXCLUDED.name,
    handle = EXCLUDED.handle,
    description = EXCLUDED.description,
    avatar = EXCLUDED.avatar,
    banner = EXCLUDED.banner,
    subscriber_count = EXCLUDED.subscriber_count,
    uploads_playlist_id = EXCLUDED.uploads_playlist_id,
    updated_at = NOW()
RETURNING *;

-- name: SearchChannels :many
SELECT * FROM channels
WHERE to_tsvector('simple', name || ' ' || COALESCE(handle, '')) @@ plainto_tsquery('simple', $1)
   OR name ILIKE '%' || $1 || '%'
   OR handle ILIKE '%' || $1 || '%'
ORDER BY subscriber_count DESC
LIMIT $2 OFFSET $3;

-- name: GetChannelIDsBySubscriptionIDs :many
SELECT channel_id FROM subscriptions WHERE id = ANY($1::bigint[]);
