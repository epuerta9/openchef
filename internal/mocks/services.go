package mocks

import (
	"github.com/epuerta9/openchef/internal/services"
)

type MockChatService struct {
	HandleChatRequestFn func(req services.ChatRequest) (*services.ChatResponse, error)
}

type MockAgentService struct {
	RegisterAgentFn func(info services.AgentInfo) error
}

type MockOrchestratorService struct {
	HandleRequestFn func(req services.ChatRequest) (*services.ChatResponse, error)
}

type MockSwarmService struct {
	HandleRequestFn func(req services.ChatRequest) (*services.ChatResponse, error)
}

func (m *MockChatService) HandleChatRequest(req services.ChatRequest) (*services.ChatResponse, error) {
	return m.HandleChatRequestFn(req)
}

func (m *MockAgentService) RegisterAgent(info services.AgentInfo) error {
	return m.RegisterAgentFn(info)
}

func (m *MockOrchestratorService) HandleRequest(req services.ChatRequest) (*services.ChatResponse, error) {
	return m.HandleRequestFn(req)
}

func (m *MockSwarmService) HandleRequest(req services.ChatRequest) (*services.ChatResponse, error) {
	return m.HandleRequestFn(req)
}

// ... implement other methods
