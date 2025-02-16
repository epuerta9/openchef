package chat

import (
	"encoding/json"
	"fmt"
	"time"

	sqlc "github.com/epuerta9/openchef/internal/database/sqlc"
	"github.com/epuerta9/openchef/internal/services"
	"github.com/epuerta9/openchef/internal/services/communicator"
)

type Service struct {
	comm *communicator.Service
	db   sqlc.Querier
}

func New(comm *communicator.Service, db sqlc.Querier) *Service {
	return &Service{
		comm: comm,
		db:   db,
	}
}

func (s *Service) HandleChatRequest(req services.ChatRequest) (*services.ChatResponse, error) {
	switch req.Mode {
	case "direct":
		return s.handleDirectMode(req)
	case "orchestrator":
		return s.handleOrchestratorMode(req)
	case "swarm":
		return s.handleSwarmMode(req)
	default:
		return nil, fmt.Errorf("unsupported mode: %s", req.Mode)
	}
}

func (s *Service) handleDirectMode(req services.ChatRequest) (*services.ChatResponse, error) {
	msg, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	// Direct publish to agent
	resp, err := s.comm.Request("agent.direct", msg, 30*time.Second)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}

	var response services.ChatResponse
	if err := json.Unmarshal(resp.Data, &response); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return &response, nil
}

func (s *Service) handleOrchestratorMode(req services.ChatRequest) (*services.ChatResponse, error) {
	// Delegate to orchestrator service
	msg, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	resp, err := s.comm.Request("orchestrator.execute", msg, 60*time.Second)
	if err != nil {
		return nil, fmt.Errorf("orchestrator request failed: %w", err)
	}

	var response services.ChatResponse
	if err := json.Unmarshal(resp.Data, &response); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return &response, nil
}

func (s *Service) handleSwarmMode(req services.ChatRequest) (*services.ChatResponse, error) {
	// Delegate to swarm service
	msg, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	resp, err := s.comm.Request("swarm.execute", msg, 60*time.Second)
	if err != nil {
		return nil, fmt.Errorf("swarm request failed: %w", err)
	}

	var response services.ChatResponse
	if err := json.Unmarshal(resp.Data, &response); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return &response, nil
}
