package orchestrator

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
	comm      *communicator.Service
	db        sqlc.Querier
	tasks     map[string]*Task
	taskMutex sync.RWMutex
}

type Task struct {
	ID       string
	Steps    []Step
	Current  int
	Response *services.ChatResponse
}

type Step struct {
	Agent   string
	Subject string
	Request services.ChatRequest
}

func New(comm *communicator.Service, db sqlc.Querier) *Service {
	svc := &Service{
		comm:  comm,
		db:    db,
		tasks: make(map[string]*Task),
	}

	// Subscribe to orchestrator commands
	svc.comm.Subscribe("orchestrator.execute", svc.handleExecute, nats.DeliverNew())
	return svc
}

func (s *Service) handleExecute(msg *nats.Msg) {
	var req services.ChatRequest
	if err := json.Unmarshal(msg.Data, &req); err != nil {
		s.replyError(msg, fmt.Errorf("invalid request: %w", err))
		return
	}

	// Plan the steps
	task, err := s.planTask(req)
	if err != nil {
		s.replyError(msg, fmt.Errorf("failed to plan task: %w", err))
		return
	}

	// Execute steps
	resp, err := s.executeTask(task)
	if err != nil {
		s.replyError(msg, fmt.Errorf("failed to execute task: %w", err))
		return
	}

	s.replyJSON(msg, resp)
}

func (s *Service) planTask(req services.ChatRequest) (*Task, error) {
	task := &Task{
		ID:      fmt.Sprintf("task_%d", time.Now().UnixNano()),
		Steps:   []Step{},
		Current: 0,
	}

	// Add a single step for now
	task.Steps = append(task.Steps, Step{
		Agent:   "default",
		Subject: "agent.execute",
		Request: req,
	})

	return task, nil
}

func (s *Service) executeTask(task *Task) (*services.ChatResponse, error) {
	if len(task.Steps) == 0 {
		return nil, fmt.Errorf("task has no steps")
	}

	// Execute first step for now
	step := task.Steps[0]
	msg, err := json.Marshal(step.Request)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	resp, err := s.comm.Request(step.Subject, msg, 30*time.Second)
	if err != nil {
		return nil, fmt.Errorf("step execution failed: %w", err)
	}

	var response services.ChatResponse
	if err := json.Unmarshal(resp.Data, &response); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return &response, nil
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
	task, err := s.planTask(req)
	if err != nil {
		return nil, fmt.Errorf("failed to plan task: %w", err)
	}
	return s.executeTask(task)
}
