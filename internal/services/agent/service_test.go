package agent

import (
	"testing"

	"github.com/epuerta9/openchef/internal/mocks"
	"github.com/epuerta9/openchef/internal/services"
)

func TestRegisterAgent(t *testing.T) {
	// Setup
	mockDB := &mocks.MockQuerier{}
	mockComm := &mocks.MockCommunicator{}

	svc := New(mockComm, mockDB)

	// Test cases
	tests := []struct {
		name    string
		agent   services.AgentInfo
		wantErr bool
	}{
		{
			name: "successful registration",
			agent: services.AgentInfo{
				ID:   "test-1",
				Name: "test-agent",
			},
			wantErr: false,
		},
		// Add more test cases
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := svc.RegisterAgent(tt.agent)
			if (err != nil) != tt.wantErr {
				t.Errorf("RegisterAgent() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
