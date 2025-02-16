package mocks

import (
	"context"

	sqlc "github.com/epuerta9/openchef/internal/database/sqlc"
)

type MockQuerier struct {
	// Add functions that will be called
	GetAgentFn                 func(ctx context.Context, id string) (sqlc.Agent, error)
	ListAgentsFn               func(ctx context.Context) ([]sqlc.Agent, error)
	CreateAgentFn              func(ctx context.Context, arg sqlc.CreateAgentParams) (sqlc.Agent, error)
	CreateRemoteClientFn       func(ctx context.Context, arg sqlc.CreateRemoteClientParams) (sqlc.RemoteClient, error)
	DeleteAgentFn              func(ctx context.Context, id string) error
	DeleteRemoteClientFn       func(ctx context.Context, id string) error
	GetAgentByNameFn           func(ctx context.Context, name string) (sqlc.Agent, error)
	GetRemoteClientFn          func(ctx context.Context, id string) (sqlc.RemoteClient, error)
	GetRemoteClientByNameFn    func(ctx context.Context, name string) (sqlc.RemoteClient, error)
	ListRemoteClientsFn        func(ctx context.Context) ([]sqlc.RemoteClient, error)
	UpdateAgentStatusFn        func(ctx context.Context, arg sqlc.UpdateAgentStatusParams) (sqlc.Agent, error)
	UpdateRemoteClientStatusFn func(ctx context.Context, arg sqlc.UpdateRemoteClientStatusParams) (sqlc.RemoteClient, error)
}

// Implement database.Querier interface
func (m *MockQuerier) GetAgent(ctx context.Context, id string) (sqlc.Agent, error) {
	return m.GetAgentFn(ctx, id)
}

func (m *MockQuerier) ListAgents(ctx context.Context) ([]sqlc.Agent, error) {
	return m.ListAgentsFn(ctx)
}

func (m *MockQuerier) CreateAgent(ctx context.Context, arg sqlc.CreateAgentParams) (sqlc.Agent, error) {
	return m.CreateAgentFn(ctx, arg)
}

func (m *MockQuerier) CreateRemoteClient(ctx context.Context, arg sqlc.CreateRemoteClientParams) (sqlc.RemoteClient, error) {
	return m.CreateRemoteClientFn(ctx, arg)
}

func (m *MockQuerier) DeleteAgent(ctx context.Context, id string) error {
	return m.DeleteAgentFn(ctx, id)
}

func (m *MockQuerier) DeleteRemoteClient(ctx context.Context, id string) error {
	return m.DeleteRemoteClientFn(ctx, id)
}

func (m *MockQuerier) GetAgentByName(ctx context.Context, name string) (sqlc.Agent, error) {
	return m.GetAgentByNameFn(ctx, name)
}

func (m *MockQuerier) GetRemoteClient(ctx context.Context, id string) (sqlc.RemoteClient, error) {
	return m.GetRemoteClientFn(ctx, id)
}

func (m *MockQuerier) GetRemoteClientByName(ctx context.Context, name string) (sqlc.RemoteClient, error) {
	return m.GetRemoteClientByNameFn(ctx, name)
}

func (m *MockQuerier) ListRemoteClients(ctx context.Context) ([]sqlc.RemoteClient, error) {
	return m.ListRemoteClientsFn(ctx)
}

func (m *MockQuerier) UpdateAgentStatus(ctx context.Context, arg sqlc.UpdateAgentStatusParams) (sqlc.Agent, error) {
	return m.UpdateAgentStatusFn(ctx, arg)
}

func (m *MockQuerier) UpdateRemoteClientStatus(ctx context.Context, arg sqlc.UpdateRemoteClientStatusParams) (sqlc.RemoteClient, error) {
	return m.UpdateRemoteClientStatusFn(ctx, arg)
}

// ... implement other methods
