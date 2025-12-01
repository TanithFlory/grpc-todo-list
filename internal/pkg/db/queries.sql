-- name: CreateTodo :exec
INSERT INTO
    todos (id, title, description)
VALUES
    (?, ?, ?);

-- name: GetTodo :one
SELECT
    id,
    title,
    description,
    completed,
    created_at,
    updated_at
FROM
    todos
WHERE
    id = ?;

-- name: ListTodos :many
SELECT
    id,
    title,
    description,
    completed,
    created_at,
    updated_at
FROM
    todos
ORDER BY
    created_at DESC
LIMIT
    ? OFFSET ?;

-- name: UpdateTodo :exec
UPDATE
    todos
SET
    title = ?,
    description = ?,
    completed = ?,
    updated_at = CURRENT_TIMESTAMP
WHERE
    id = ?;

-- name: DeleteTodo :exec
DELETE FROM
    todos
WHERE
    id = ?;