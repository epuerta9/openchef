package swarm

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"

	sqlc "github.com/epuerta9/openchef/internal/database/sqlc"
	"github.com/epuerta9/openchef/internal/services"
	"github.com/epuerta9/openchef/internal/services/communicator"
	"github.com/nats-io/nats.go"
)

type Service struct {
	comm        *communicator.Service
	db          sqlc.Querier
	claims      map[string]*Claim
	claimMutex  sync.RWMutex
	claimExpiry time.Duration
}

type Claim struct {
	AgentID   string
	Timestamp time.Time
	Response  *services.ChatResponse
}

func New(comm *communicator.Service, db sqlc.Querier) *Service {
	svc := &Service{
		comm:        comm,
		db:          db,
		claims:      make(map[string]*Claim),
		claimExpiry: 30 * time.Second,
	}

	// Subscribe to swarm commands
	svc.comm.Subscribe("swarm.execute", svc.handleExecute, nats.DeliverNew())
	return svc
}

func (s *Service) handleExecute(msg *nats.Msg) {
	var req services.ChatRequest
	if err := json.Unmarshal(msg.Data, &req); err != nil {
		s.replyError(msg, fmt.Errorf("invalid request: %w", err))
		return
	}

	// Create unique conversation ID
	convoID := fmt.Sprintf("swarm_%d", time.Now().UnixNano())

	// Broadcast request to all agents
	if err := s.broadcastRequest(convoID, req); err != nil {
		s.replyError(msg, fmt.Errorf("failed to broadcast request: %w", err))
		return
	}

	// Wait for responses
	resp, err := s.gatherResponses(convoID)
	if err != nil {
		s.replyError(msg, fmt.Errorf("failed to gather responses: %w", err))
		return
	}

	s.replyJSON(msg, resp)
}

func (s *Service) broadcastRequest(convoID string, req services.ChatRequest) error {
	msg, err := json.Marshal(map[string]interface{}{
		"convo_id": convoID,
		"request":  req,
	})
	if err != nil {
		return fmt.Errorf("failed to marshal broadcast: %w", err)
	}

	return s.comm.Publish("swarm.broadcast", msg)
}

func (s *Service) gatherResponses(convoID string) (*services.ChatResponse, error) {
	// Wait for responses and aggregate them
	// This could involve:
	// 1. Collecting multiple responses
	// 2. Voting or consensus mechanism
	// 3. Timeout handling

	// For now, implement basic response gathering
	msgChan := make(chan *nats.Msg)
	if err := s.comm.Subscribe(fmt.Sprintf("swarm.responses.%s", convoID), func(msg *nats.Msg) {
		msgChan <- msg
	}); err != nil {
		return nil, fmt.Errorf("failed to subscribe to responses: %w", err)
	}
	defer s.comm.Unsubscribe(fmt.Sprintf("swarm.responses.%s", convoID))

	// Wait for first response
	select {
	case msg := <-msgChan:
		var response services.ChatResponse
		if err := json.Unmarshal(msg.Data, &response); err != nil {
			return nil, fmt.Errorf("failed to unmarshal response: %w", err)
		}
		return &response, nil
	case <-time.After(10 * time.Second):
		return nil, fmt.Errorf("timeout waiting for responses")
	}
}

func (s *Service) replyJSON(msg *nats.Msg, data interface{}) {
	response, err := json.Marshal(data)
	if err != nil {
		s.replyError(msg, fmt.Errorf("failed to marshal response: %w", err))
		return
	}
	msg.Respond(response)
}

func (s *Service) replyError(msg *nats.Msg, err error) {
	response, _ := json.Marshal(map[string]string{"error": err.Error()})
	msg.Respond(response)
}

func (s *Service) HandleRequest(req services.ChatRequest) (*services.ChatResponse, error) {
	convoID := fmt.Sprintf("swarm_%d", time.Now().UnixNano())
	if err := s.broadcastRequest(convoID, req); err != nil {
		return nil, err
	}
	return s.gatherResponses(convoID)
}
