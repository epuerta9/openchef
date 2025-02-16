package communicator

import (
	"time"

	"github.com/nats-io/nats.go"
)

type Interface interface {
	Publish(subject string, data []byte) error
	Subscribe(subject string, handler nats.MsgHandler, opts ...nats.SubOpt) error
	Request(subject string, data []byte, timeout time.Duration) (*nats.Msg, error)
	Unsubscribe(subject string) error
	StoreAgent(key string, agent interface{}) error
}
