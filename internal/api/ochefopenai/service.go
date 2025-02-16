package ochefopenai

import (
	"context"
	"fmt"
	"mime/multipart"
	"net/http"

	"github.com/epuerta9/openchef/internal/services"
)

type ChatServicer interface {
	HandleChatRequest(req services.ChatRequest) (*services.ChatResponse, error)
}

type AgentServicer interface {
	RegisterAgent(info services.AgentInfo) error
}

type OrchestratorServicer interface {
	HandleRequest(req services.ChatRequest) (*services.ChatResponse, error)
}

type SwarmServicer interface {
	HandleRequest(req services.ChatRequest) (*services.ChatResponse, error)
}

type Service struct {
	chat         ChatServicer
	agent        AgentServicer
	orchestrator OrchestratorServicer
	swarm        SwarmServicer
	client       *http.Client
	apiKey       string
}

func NewService(chat ChatServicer, agent AgentServicer,
	orchestrator OrchestratorServicer, swarm SwarmServicer, apiKey string) *Service {
	return &Service{
		chat:         chat,
		agent:        agent,
		orchestrator: orchestrator,
		swarm:        swarm,
		client:       &http.Client{},
		apiKey:       apiKey,
	}
}

func (s *Service) CreateChatCompletion(ctx context.Context, req *ChatCompletionRequest) (*ChatCompletionResponse, error) {
	// Convert OpenAI request to internal service request
	serviceReq := services.ChatRequest{
		Model:       req.Model,
		Messages:    convertMessages(req.Messages),
		Temperature: req.Temperature,
	}

	// Route to appropriate service based on model or other parameters
	var resp *services.ChatResponse
	var err error

	switch {
	case isDirectMode(req.Model):
		resp, err = s.chat.HandleChatRequest(serviceReq)
	case isOrchestratorMode(req.Model):
		resp, err = s.orchestrator.HandleRequest(serviceReq)
	case isSwarmMode(req.Model):
		resp, err = s.swarm.HandleRequest(serviceReq)
	default:
		return nil, fmt.Errorf("unsupported model: %s", req.Model)
	}

	if err != nil {
		return nil, fmt.Errorf("chat completion failed: %w", err)
	}

	// Convert internal response to OpenAI format
	return convertToOpenAIResponse(resp), nil
}

func (s *Service) UploadFile(ctx context.Context, file *multipart.FileHeader) (*File, error) {
	// Implementation for file upload
	// This could store in a local filesystem or cloud storage
	return nil, fmt.Errorf("file upload not implemented")
}

func (s *Service) GetFile(ctx context.Context, fileID string) (*File, error) {
	// Implementation for file retrieval
	return nil, fmt.Errorf("file retrieval not implemented")
}

// Helper functions
func convertMessages(msgs []Message) []services.Message {
	converted := make([]services.Message, len(msgs))
	for i, msg := range msgs {
		converted[i] = services.Message{
			Role:    msg.Role,
			Content: msg.Content,
		}
	}
	return converted
}

func convertToOpenAIResponse(resp *services.ChatResponse) *ChatCompletionResponse {
	return &ChatCompletionResponse{
		ID:      resp.ID,
		Object:  "chat.completion",
		Created: resp.Created.Unix(),
		Model:   resp.Model,
		Choices: []Choice{{
			Index: 0,
			Message: Message{
				Role:    resp.Messages[len(resp.Messages)-1].Role,
				Content: resp.Messages[len(resp.Messages)-1].Content,
			},
			FinishReason: "stop",
		}},
		Usage: Usage{
			PromptTokens:     resp.Usage.PromptTokens,
			CompletionTokens: resp.Usage.CompletionTokens,
			TotalTokens:      resp.Usage.TotalTokens,
		},
	}
}

func isDirectMode(model string) bool {
	return model == "gpt-3.5-turbo" || model == "gpt-4"
}

func isOrchestratorMode(model string) bool {
	return model == "orchestrator"
}

func isSwarmMode(model string) bool {
	return model == "swarm"
}
