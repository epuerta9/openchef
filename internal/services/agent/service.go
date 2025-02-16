package agent

import (
	"fmt"
	"sync"
	"time"

	sqlc "github.com/epuerta9/openchef/internal/database/sqlc"
	"github.com/epuerta9/openchef/internal/services"
	"github.com/epuerta9/openchef/internal/services/communicator"
)

type Service struct {
	comm    communicator.Interface
	db      sqlc.Querier
	agents  map[string]services.AgentInfo
	mu      sync.RWMutex
	timeout time.Duration
}

func New(comm communicator.Interface, db sqlc.Querier) *Service {
	return &Service{
		comm:    comm,
		db:      db,
		agents:  make(map[string]services.AgentInfo),
		timeout: 5 * time.Minute,
	}
}

func (s *Service) RegisterAgent(info services.AgentInfo) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	info.LastSeen = time.Now()
	info.Status = "active"
	s.agents[info.ID] = info

	return s.comm.StoreAgent(fmt.Sprintf("agent:%s", info.ID), info)
}

func (s *Service) ListAgents() []services.AgentInfo {
	s.mu.RLock()
	defer s.mu.RUnlock()

	agents := make([]services.AgentInfo, 0, len(s.agents))
	for _, agent := range s.agents {
		agents = append(agents, agent)
	}
	return agents
}

func (s *Service) FindAgentByName(name string) (services.AgentInfo, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	for _, agent := range s.agents {
		if agent.Name == name {
			return agent, nil
		}
	}
	return services.AgentInfo{}, fmt.Errorf("agent not found: %s", name)
}

func (s *Service) UpdateAgentStatus(id string, status string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if agent, exists := s.agents[id]; exists {
		agent.Status = status
		agent.LastSeen = time.Now()
		s.agents[id] = agent
		return s.comm.StoreAgent(fmt.Sprintf("agent:%s", id), agent)
	}
	return fmt.Errorf("agent not found: %s", id)
}
