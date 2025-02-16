package communicator

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/nats-io/nats.go"
)

type Service struct {
	nc     *nats.Conn
	js     nats.JetStreamContext
	kv     nats.KeyValue
	objIn  nats.ObjectStore
	objOut nats.ObjectStore
}

func New(natsURL string) (*Service, error) {
	nc, err := nats.Connect(natsURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to NATS: %w", err)
	}

	js, err := nc.JetStream()
	if err != nil {
		return nil, fmt.Errorf("failed to create jetstream context: %w", err)
	}

	kv, err := createOrGetKVStore(js)
	if err != nil {
		return nil, err
	}

	objIn, err := createOrGetObjectStore(js, "openchef-in")
	if err != nil {
		return nil, err
	}

	objOut, err := createOrGetObjectStore(js, "openchef-out")
	if err != nil {
		return nil, err
	}

	return &Service{
		nc:     nc,
		js:     js,
		kv:     kv,
		objIn:  objIn,
		objOut: objOut,
	}, nil
}

func (s *Service) Subscribe(subject string, handler nats.MsgHandler, opts ...nats.SubOpt) error {
	_, err := s.js.Subscribe(subject, handler, opts...)
	return err
}

func (s *Service) StoreRequest(key string, data []byte) error {
	_, err := s.kv.Put(key, data)
	return err
}

func (s *Service) StoreResponse(key string, data []byte) error {
	_, err := s.objOut.PutBytes(key, data)
	return err
}

func (s *Service) StoreAgent(key string, agent interface{}) error {
	data, err := json.Marshal(agent)
	if err != nil {
		return fmt.Errorf("failed to marshal agent: %w", err)
	}
	return s.StoreRequest(key, data)
}

func (s *Service) GetAgent(key string) ([]byte, error) {
	entry, err := s.kv.Get(key)
	if err != nil {
		return nil, err
	}
	return entry.Value(), nil
}

func (s *Service) WatchAgents(prefix string) (nats.KeyWatcher, error) {
	return s.kv.Watch(prefix + ".*")
}

func (s *Service) Close() error {
	return s.nc.Drain()
}

func (s *Service) Request(subject string, data []byte, timeout time.Duration) (*nats.Msg, error) {
	return s.nc.Request(subject, data, timeout)
}

func (s *Service) Publish(subject string, data []byte) error {
	return s.nc.Publish(subject, data)
}

func (s *Service) Unsubscribe(subject string) error {
	sub, err := s.js.PullSubscribe(subject, "")
	if err != nil {
		return err
	}
	return sub.Unsubscribe()
}

// Helper functions
func createOrGetKVStore(js nats.JetStreamContext) (nats.KeyValue, error) {
	kv, err := js.CreateKeyValue(&nats.KeyValueConfig{
		Bucket:      "openchef",
		Description: "OpenChef KV Store",
		TTL:         24 * time.Hour,
	})
	if err != nil {
		if err != nats.ErrBucketNotFound {
			return nil, fmt.Errorf("failed to create kv store: %w", err)
		}
		kv, err = js.KeyValue("openchef")
		if err != nil {
			return nil, fmt.Errorf("failed to get existing kv store: %w", err)
		}
	}
	return kv, nil
}

func createOrGetObjectStore(js nats.JetStreamContext, bucket string) (nats.ObjectStore, error) {
	obj, err := js.CreateObjectStore(&nats.ObjectStoreConfig{
		Bucket:      bucket,
		Description: fmt.Sprintf("OpenChef %s Object Store", bucket),
		TTL:         24 * time.Hour,
		Storage:     nats.FileStorage,
		Replicas:    1,
	})
	if err != nil {
		if err != nats.ErrBucketNotFound {
			return nil, err
		}
		obj, err = js.ObjectStore(bucket)
		if err != nil {
			return nil, err
		}
	}
	return obj, nil
}
