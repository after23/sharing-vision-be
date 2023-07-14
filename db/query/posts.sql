-- name: CreatePost :execresult
INSERT INTO posts (
    title,
    content,
    category,
    status,
    created_date
) VALUES (
    ?, ?, ?, ?, ?
);

-- name: GetPost :many
SELECT title, content, category, status FROM posts LIMIT ? OFFSET ?;

-- name: GetPostById :one
SELECT title, content, category, status FROM posts where id=?;

-- name: UpdatePost :execresult
UPDATE posts SET title=?, content=?, category=?, status=?, updated_date=? WHERE id=?;

-- name: DeletePost :execresult
DELETE FROM posts WHERE id=?;