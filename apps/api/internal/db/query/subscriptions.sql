-- name: GetSubscription :one
SELECT s.*, c.id as channel_id, c.name as channel_name, c.handle as channel_handle,
       c.avatar as channel_avatar, c.subscriber_count as channel_subscriber_count,
       c.description as channel_description, c.uploads_playlist_id,
       c.banner as channel_banner, c.youtube_channel_id
FROM subscriptions s
JOIN channels c ON c.id = s.channel_id
WHERE s.user_id = $1 AND s.channel_id = $2;

-- name: ListSubscriptionsByName :many
SELECT s.*, c.id as channel_id, c.name as channel_name, c.handle as channel_handle,
       c.avatar as channel_avatar, c.subscriber_count as channel_subscriber_count,
       c.description as channel_description, c.uploads_playlist_id,
       c.banner as channel_banner, c.youtube_channel_id
FROM subscriptions s
JOIN channels c ON c.id = s.channel_id
WHERE s.user_id = $1
ORDER BY c.name ASC
LIMIT $2 OFFSET $3;

-- name: ListSubscriptionsByRecent :many
SELECT s.*, c.id as channel_id, c.name as channel_name, c.handle as channel_handle,
       c.avatar as channel_avatar, c.subscriber_count as channel_subscriber_count,
       c.description as channel_description, c.uploads_playlist_id,
       c.banner as channel_banner, c.youtube_channel_id
FROM subscriptions s
JOIN channels c ON c.id = s.channel_id
WHERE s.user_id = $1
ORDER BY s.subscribed_at DESC
LIMIT $2 OFFSET $3;

-- name: ListSubscriptionChannelsByUser :many
SELECT c.* FROM channels c
JOIN subscriptions s ON s.channel_id = c.id
WHERE s.user_id = $1;

-- name: UpsertSubscription :one
INSERT INTO subscriptions (user_id, channel_id, subscribed_at)
VALUES ($1, $2, $3)
ON CONFLICT (user_id, channel_id) DO UPDATE SET
    synced_at = NOW()
RETURNING *;

-- name: DeleteSubscription :exec
DELETE FROM subscriptions WHERE user_id = $1 AND channel_id = $2;

-- name: CountSubscriptionsByUser :one
SELECT COUNT(*) FROM subscriptions WHERE user_id = $1;
