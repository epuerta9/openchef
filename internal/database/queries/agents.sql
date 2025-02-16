-- name: GetAgent :one
SELECT * FROM agents WHERE id = ?;

-- name: GetAgentByName :one
SELECT * FROM agents 
WHERE name = ? LIMIT 1;

-- name: ListAgents :many
SELECT * FROM agents;

-- name: CreateAgent :one
INSERT INTO agents (id, name, status, remote_client_id) 
VALUES (?, ?, ?, ?) 
RETURNING *;

-- name: UpdateAgentStatus :one
UPDATE agents 
SET status = ?, 
    last_seen = CURRENT_TIMESTAMP 
WHERE id = ? 
RETURNING *;

-- name: DeleteAgent :exec
DELETE FROM agents WHERE id = ?;
