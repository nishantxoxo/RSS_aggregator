-- name: Createfeed :one
INSERT INTO feeds(id, created_at, updated_at, name, url, user_id)
VALUES ($1, $2, $3, $4, $5, $6
)
RETURNING *;

-- name: GetFeeds :many
SELECT * from feeds;




-- name: GetNextFeedsToFetch :many

SELECT * from feeds
ORDER BY last_fetched_at asc nulls first
limit $1;

-- name: MarkFeedAsFetched :one

update feeds
set last_fetched_at = NOW(), updated_at = NOW()
where id = $1
RETURNING *;