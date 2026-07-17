-- name: ListCollectionsByUser :many
SELECT * FROM collections WHERE user_id = $1 ORDER BY name ASC;

-- name: GetCollectionByID :one
SELECT * FROM collections WHERE id = $1 AND user_id = $2;

-- name: CreateCollection :one
INSERT INTO collections (user_id, name, icon, color)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: UpdateCollection :one
UPDATE collections
SET name = $3, icon = $4, color = $5, updated_at = NOW()
WHERE id = $1 AND user_id = $2
RETURNING *;

-- name: DeleteCollection :exec
DELETE FROM collections WHERE id = $1 AND user_id = $2;

-- name: ListCollectionChannels :many
SELECT c.* FROM channels c
JOIN collection_channels cc ON cc.channel_id = c.id
WHERE cc.collection_id = $1
ORDER BY c.name ASC;

-- name: AddChannelToCollection :exec
INSERT INTO collection_channels (collection_id, channel_id)
VALUES ($1, $2)
ON CONFLICT DO NOTHING;

-- name: RemoveChannelFromCollection :exec
DELETE FROM collection_channels WHERE collection_id = $1 AND channel_id = $2;

-- name: ListCollectionsByChannel :many
SELECT cl.* FROM collections cl
JOIN collection_channels cc ON cc.collection_id = cl.id
WHERE cc.channel_id = $1 AND cl.user_id = $2
ORDER BY cl.name ASC;

-- name: SearchCollectionsByUser :many
SELECT * FROM collections
WHERE user_id = $1 AND name ILIKE '%' || $2 || '%'
ORDER BY name ASC;
