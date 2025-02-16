package mocks

import (
	"time"

	"github.com/nats-io/nats.go"
)

type MockCommunicator struct {
	PublishFn     func(subject string, data []byte) error
	SubscribeFn   func(subject string, handler nats.MsgHandler, opts ...nats.SubOpt) error
	RequestFn     func(subject string, data []byte, timeout time.Duration) (*nats.Msg, error)
	UnsubscribeFn func(subject string) error
}

func (m *MockCommunicator) Publish(subject string, data []byte) error {
	return m.PublishFn(subject, data)
}

func (m *MockCommunicator) Subscribe(subject string, handler nats.MsgHandler, opts ...nats.SubOpt) error {
	return m.SubscribeFn(subject, handler, opts...)
}

func (m *MockCommunicator) Request(subject string, data []byte, timeout time.Duration) (*nats.Msg, error) {
	return m.RequestFn(subject, data, timeout)
}

func (m *MockCommunicator) Unsubscribe(subject string) error {
	return m.UnsubscribeFn(subject)
}

func (m *MockCommunicator) StoreAgent(key string, agent interface{}) error {
	return nil // Implement as needed for tests
}

// ... implement other methods
