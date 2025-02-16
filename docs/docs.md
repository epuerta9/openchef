docs/
├── database/
│   ├── migrations.md
│   ├── queries.md
│   └── schema.md
├── messaging/
│   ├── nats-config.md
│   └── nats-setup.md
├── server/
│   ├── api.md
│   └── handlers.md
└── development/
    ├── installation.md
    └── tooling.md

# Database Schema

## Remote Clients Table
This table stores information about remote clients connecting to the system.
sql
CREATE TABLE remote_clients (
id TEXT PRIMARY KEY,
name TEXT UNIQUE NOT NULL,
description TEXT NOT NULL DEFAULT '',
ack BOOLEAN NOT NULL DEFAULT FALSE,
message TEXT NOT NULL DEFAULT '',
last_seen DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
bento_box JSON NOT NULL DEFAULT '{}',
version TEXT NOT NULL DEFAULT '0.0.1',
created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);


root = "."
tmp_dir = "tmp"

[build]
  cmd = "go build -o ./tmp/main ./cmd/api"
  bin = "./tmp/main"
  delay = 1000
  exclude_dir = ["assets", "tmp", "vendor"]
  include_ext = ["go", "tpl", "tmpl", "templ"]
  exclude_regex = ["_test\\.go"]

[screen]
  clear_on_rebuild = true


Make Commands
General
make build: Build the application
make run: Run the application
make dev: Run with hot reload
make test: Run tests
make generate: Generate templ files
make clean: Clean build artifacts
Database
make db-setup: Install database tools
make db-status: Check database status
make db-migrate: Apply migrations
make db-new: Create new migration
make db-generate: Generate SQLC code
NATS
make nats-dev: Run NATS in development mode
make nats-prod: Run NATS in production mode


This documentation structure:
1. Organizes information by domain
2. Provides clear setup instructions
3. Includes configuration examples
4. Documents all available commands
5. Separates development and production concerns

You can access this documentation in the repository or generate it as HTML using a tool like MkDocs.

### Indexes
- `idx_remote_clients_name`: Index for faster lookups by name

# Database Queries

## SQLC Generated Queries

### Remote Clients

1. Get Remote Client
```sql
-- name: GetRemoteClient :one
SELECT * FROM remote_clients
WHERE id = ? LIMIT 1;
```

2. Get Remote Client by Name
```sql
-- name: GetRemoteClientByName :one
SELECT * FROM remote_clients
WHERE name = ? LIMIT 1;
```

3. List Remote Clients
```sql
-- name: ListRemoteClients :many
SELECT * FROM remote_clients
ORDER BY created_at DESC;
```

4. Create Remote Client
```sql
-- name: CreateRemoteClient :one
INSERT INTO remote_clients (
    id, name, description, ack, message, bento_box, version
) VALUES (?, ?, ?, ?, ?, ?, ?)
RETURNING *;
```

5. Update Remote Client
```sql
-- name: UpdateRemoteClient :one
UPDATE remote_clients
SET 
    description = ?,
    ack = ?,
    message = ?,
    bento_box = ?,
    version = ?,
    last_seen = CURRENT_TIMESTAMP,
    updated_at = CURRENT_TIMESTAMP
WHERE id = ?
RETURNING *;
```

6. Update Last Seen
```sql
-- name: UpdateLastSeen :exec
UPDATE remote_clients
SET last_seen = CURRENT_TIMESTAMP
WHERE id = ?;
```

7. Delete Remote Client
```sql
-- name: DeleteRemoteClient :exec
DELETE FROM remote_clients
WHERE id = ?;
```