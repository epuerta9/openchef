// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: remote_clients.sql

package database

import (
	"context"
)

const createRemoteClient = `-- name: CreateRemoteClient :one
INSERT INTO remote_clients (id, name, status) 
VALUES (?, ?, ?) 
RETURNING id, name, status, created_at
`

type CreateRemoteClientParams struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
}

func (q *Queries) CreateRemoteClient(ctx context.Context, arg CreateRemoteClientParams) (RemoteClient, error) {
	row := q.queryRow(ctx, q.createRemoteClientStmt, createRemoteClient, arg.ID, arg.Name, arg.Status)
	var i RemoteClient
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Status,
		&i.CreatedAt,
	)
	return i, err
}

const deleteRemoteClient = `-- name: DeleteRemoteClient :exec
DELETE FROM remote_clients WHERE id = ?
`

func (q *Queries) DeleteRemoteClient(ctx context.Context, id string) error {
	_, err := q.exec(ctx, q.deleteRemoteClientStmt, deleteRemoteClient, id)
	return err
}

const getRemoteClient = `-- name: GetRemoteClient :one
SELECT id, name, status, created_at FROM remote_clients WHERE id = ?
`

func (q *Queries) GetRemoteClient(ctx context.Context, id string) (RemoteClient, error) {
	row := q.queryRow(ctx, q.getRemoteClientStmt, getRemoteClient, id)
	var i RemoteClient
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Status,
		&i.CreatedAt,
	)
	return i, err
}

const getRemoteClientByName = `-- name: GetRemoteClientByName :one
SELECT id, name, status, created_at FROM remote_clients
WHERE name = ? LIMIT 1
`

func (q *Queries) GetRemoteClientByName(ctx context.Context, name string) (RemoteClient, error) {
	row := q.queryRow(ctx, q.getRemoteClientByNameStmt, getRemoteClientByName, name)
	var i RemoteClient
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Status,
		&i.CreatedAt,
	)
	return i, err
}

const listRemoteClients = `-- name: ListRemoteClients :many
SELECT id, name, status, created_at FROM remote_clients
`

func (q *Queries) ListRemoteClients(ctx context.Context) ([]RemoteClient, error) {
	rows, err := q.query(ctx, q.listRemoteClientsStmt, listRemoteClients)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []RemoteClient
	for rows.Next() {
		var i RemoteClient
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Status,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateRemoteClientStatus = `-- name: UpdateRemoteClientStatus :one
UPDATE remote_clients 
SET status = ? 
WHERE id = ? 
RETURNING id, name, status, created_at
`

type UpdateRemoteClientStatusParams struct {
	Status string `json:"status"`
	ID     string `json:"id"`
}

func (q *Queries) UpdateRemoteClientStatus(ctx context.Context, arg UpdateRemoteClientStatusParams) (RemoteClient, error) {
	row := q.queryRow(ctx, q.updateRemoteClientStatusStmt, updateRemoteClientStatus, arg.Status, arg.ID)
	var i RemoteClient
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Status,
		&i.CreatedAt,
	)
	return i, err
}
