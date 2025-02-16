package services

import "time"

type AgentInfo struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	Subject      string    `json:"subject"`
	Capabilities []string  `json:"capabilities"`
	LastSeen     time.Time `json:"last_seen"`
	Status       string    `json:"status"`
}

type ChatRequest struct {
	Model       string    `json:"model"`
	Messages    []Message `json:"messages"`
	Mode        string    `json:"mode"` // "direct", "orchestrator", or "swarm"
	Temperature float64   `json:"temperature"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatResponse struct {
	ID        string    `json:"id"`
	Created   time.Time `json:"created"`
	Model     string    `json:"model"`
	Messages  []Message `json:"messages"`
	Usage     Usage     `json:"usage"`
	Completed bool      `json:"completed"`
}

type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}
