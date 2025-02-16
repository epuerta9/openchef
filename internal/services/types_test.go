package services

import (
	"encoding/json"
	"testing"
)

func TestChatRequestJSON(t *testing.T) {
	tests := []struct {
		name    string
		req     ChatRequest
		want    string
		wantErr bool
	}{
		{
			name: "basic request",
			req: ChatRequest{
				Model: "gpt-4",
				Messages: []Message{
					{Role: "user", Content: "Hello"},
				},
				Temperature: 0.7,
			},
			want:    `{"model":"gpt-4","messages":[{"role":"user","content":"Hello"}],"mode":"","temperature":0.7}`,
			wantErr: false,
		},
		// Add more test cases
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := json.Marshal(tt.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("json.Marshal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if string(got) != tt.want {
				t.Errorf("json.Marshal() = %v, want %v", string(got), tt.want)
			}
		})
	}
}
