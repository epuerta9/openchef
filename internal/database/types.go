package database

import "context"

type Agent struct {
	ID        string
	Name      string
	Status    string
	LastSeen  string
	CreatedAt string
}

type RemoteClient struct {
	ID        string
	Name      string
	Status    string
	CreatedAt string
}

type CreateAgentParams struct {
	ID     string
	Name   string
	Status string
}

type CreateRemoteClientParams struct {
	ID     string
	Name   string
	Status string
}

type Querier interface {
	GetAgent(ctx context.Context, id string) (Agent, error)
	ListAgents(ctx context.Context) ([]Agent, error)
	CreateAgent(ctx context.Context, arg CreateAgentParams) (Agent, error)
	CreateRemoteClient(ctx context.Context, arg CreateRemoteClientParams) (RemoteClient, error)
}
