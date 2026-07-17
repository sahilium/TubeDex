-- name: GlobalSearch :many
SELECT
    'channel' as result_type,
    c.id as result_id,
    c.name as result_name,
    c.avatar as result_avatar,
    c.handle as result_handle,
    c.description as result_description,
    c.subscriber_count as result_subscriber_count
FROM channels c
JOIN subscriptions s ON s.channel_id = c.id
WHERE s.user_id = $1
  AND (c.name ILIKE '%' || $2 || '%'
    OR c.handle ILIKE '%' || $2 || '%'
    OR c.description ILIKE '%' || $2 || '%')
UNION ALL
SELECT
    'collection' as result_type,
    cl.id as result_id,
    cl.name as result_name,
    '' as result_avatar,
    '' as result_handle,
    '' as result_description,
    0 as result_subscriber_count
FROM collections cl
WHERE cl.user_id = $1 AND cl.name ILIKE '%' || $2 || '%'
UNION ALL
SELECT
    'note' as result_type,
    n.channel_id as result_id,
    c.name as result_name,
    c.avatar as result_avatar,
    c.handle as result_handle,
    n.body as result_description,
    c.subscriber_count as result_subscriber_count
FROM notes n
JOIN channels c ON c.id = n.channel_id
WHERE n.user_id = $1 AND n.body ILIKE '%' || $2 || '%'
ORDER BY result_name ASC
LIMIT $3;
