-- name: GetRemoteClient :one
SELECT * FROM remote_clients WHERE id = ?;

-- name: GetRemoteClientByName :one
SELECT * FROM remote_clients
WHERE name = ? LIMIT 1;

-- name: ListRemoteClients :many
SELECT * FROM remote_clients;

-- name: CreateRemoteClient :one
INSERT INTO remote_clients (id, name, status) 
VALUES (?, ?, ?) 
RETURNING *;

-- name: UpdateRemoteClientStatus :one
UPDATE remote_clients 
SET status = ? 
WHERE id = ? 
RETURNING *;

-- name: DeleteRemoteClient :exec
DELETE FROM remote_clients WHERE id = ?; 