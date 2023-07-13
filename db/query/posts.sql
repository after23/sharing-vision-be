-- name: createPost :execresult
INSERT INTO posts (
    title,
    content,
    category,
    status,
    created_date
) VALUES (
    ?, ?, ?, ?, ?
);

-- name: getPost :many
SELECT title, content, category, status FROM posts LIMIT ? OFFSET ?;

-- name: getPostById :one
SELECT title, content, category, status FROM posts where id=?;

-- name: updatePost :execresult
UPDATE posts SET title=?, content=?, category=?, status=?, updated_date=? WHERE id=?;

-- name: deletePost :execresult
DELETE FROM posts WHERE id=?;