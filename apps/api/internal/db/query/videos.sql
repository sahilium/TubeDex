-- name: GetVideoByYouTubeID :one
SELECT * FROM videos WHERE youtube_video_id = $1;

-- name: ListVideosByChannelID :many
SELECT * FROM videos
WHERE channel_id = $1
ORDER BY published_at DESC
LIMIT $2 OFFSET $3;

-- name: UpsertVideo :one
INSERT INTO videos (youtube_video_id, channel_id, title, description, published_at, thumbnail)
VALUES ($1, $2, $3, $4, $5, $6)
ON CONFLICT (youtube_video_id) DO UPDATE SET
    title = EXCLUDED.title,
    description = EXCLUDED.description,
    thumbnail = EXCLUDED.thumbnail
RETURNING *;

-- name: GetLatestVideoByChannelID :one
SELECT * FROM videos
WHERE channel_id = $1
ORDER BY published_at DESC
LIMIT 1;
