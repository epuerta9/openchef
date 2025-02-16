-- Create remote clients table
CREATE TABLE remote_clients (
    id TEXT PRIMARY KEY,
    name TEXT NOT NULL,
    status TEXT NOT NULL DEFAULT 'active',
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Create agents table
CREATE TABLE agents (
    id TEXT PRIMARY KEY,
    name TEXT NOT NULL,
    status TEXT NOT NULL DEFAULT 'active',
    last_seen DATETIME,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    remote_client_id TEXT REFERENCES remote_clients(id) ON DELETE CASCADE,
    UNIQUE(name)
);

-- Create indices
CREATE INDEX idx_remote_clients_name ON remote_clients(name);
CREATE INDEX idx_agents_name ON agents(name);
