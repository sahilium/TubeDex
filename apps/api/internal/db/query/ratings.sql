-- name: GetRating :one
SELECT * FROM ratings WHERE user_id = $1 AND channel_id = $2;

-- name: UpsertRating :one
INSERT INTO ratings (user_id, channel_id, rating)
VALUES ($1, $2, $3)
ON CONFLICT (user_id, channel_id) DO UPDATE SET
    rating = EXCLUDED.rating,
    updated_at = NOW()
RETURNING *;

-- name: DeleteRating :exec
DELETE FROM ratings WHERE user_id = $1 AND channel_id = $2;

-- name: GetAverageRatingForChannel :one
SELECT COALESCE(AVG(rating::float), 0) as avg_rating, COUNT(*) as rating_count
FROM ratings WHERE channel_id = $1;
