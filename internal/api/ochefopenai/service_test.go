package ochefopenai

import (
	"context"
	"testing"
	"time"

	"github.com/epuerta9/openchef/internal/mocks"
	"github.com/epuerta9/openchef/internal/services"
)

func TestCreateChatCompletion(t *testing.T) {
	expectedID := "test-response-1"
	expectedCreated := time.Now().Unix()

	// Setup mocked services
	mockChat := &mocks.MockChatService{
		HandleChatRequestFn: func(req services.ChatRequest) (*services.ChatResponse, error) {
			return &services.ChatResponse{
				ID:      expectedID,
				Model:   req.Model,
				Created: time.Unix(expectedCreated, 0),
				Messages: []services.Message{
					{Role: "assistant", Content: "Hello! How can I help you?"},
				},
				Usage: services.Usage{
					PromptTokens:     10,
					CompletionTokens: 20,
					TotalTokens:      30,
				},
			}, nil
		},
	}
	mockAgent := &mocks.MockAgentService{}
	mockOrch := &mocks.MockOrchestratorService{}
	mockSwarm := &mocks.MockSwarmService{}

	svc := NewService(mockChat, mockAgent, mockOrch, mockSwarm, "test-key")

	// Test cases
	tests := []struct {
		name    string
		req     *ChatCompletionRequest
		want    *ChatCompletionResponse
		wantErr bool
	}{
		{
			name: "direct mode success",
			req: &ChatCompletionRequest{
				Model: "gpt-4",
				Messages: []Message{
					{Role: "user", Content: "Hello"},
				},
			},
			want: &ChatCompletionResponse{
				ID:      expectedID,
				Object:  "chat.completion",
				Created: expectedCreated,
				Model:   "gpt-4",
				Choices: []Choice{{
					Index: 0,
					Message: Message{
						Role:    "assistant",
						Content: "Hello! How can I help you?",
					},
					FinishReason: "stop",
				}},
				Usage: Usage{
					PromptTokens:     10,
					CompletionTokens: 20,
					TotalTokens:      30,
				},
			},
			wantErr: false,
		},
		// Add more test cases
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := svc.CreateChatCompletion(context.Background(), tt.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateChatCompletion() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !compareResponses(got, tt.want) {
				t.Errorf("CreateChatCompletion() = %v, want %v", got, tt.want)
			}
		})
	}
}

func compareResponses(got, want *ChatCompletionResponse) bool {
	if got == nil || want == nil {
		return got == want
	}
	return got.ID == want.ID &&
		got.Model == want.Model &&
		len(got.Choices) == len(want.Choices) &&
		got.Choices[0].Message.Content == want.Choices[0].Message.Content
}
