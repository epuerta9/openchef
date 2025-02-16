package nats

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/nats-io/nats-server/v2/server"
)

type Server struct {
	nats *server.Server
}

func New() (*Server, error) {
	var ns *server.Server
	var err error
	var opts *server.Options

	if configFile := os.Getenv("NATS_CONFIG"); configFile != "" {
		opts, err = server.ProcessConfigFile(configFile)
		if err != nil {
			return nil, fmt.Errorf("error processing config file: %w", err)
		}
	} else {
		// Default development config
		opts = &server.Options{
			Host:           "0.0.0.0",
			Port:           4222,
			HTTPPort:       8222,
			Debug:          true,
			Trace:          true,
			NoLog:          false,
			NoSigs:         true,
			MaxPayload:     1024 * 1024,
			PingInterval:   2 * time.Minute,
			MaxPingsOut:    2,
			WriteDeadline:  2 * time.Second,
			MaxControlLine: 4096,
			JetStream:      true,
			StoreDir:       "./tmp/openchef-nats",
		}
	}

	ns, err = server.NewServer(opts)
	if err != nil {
		return nil, fmt.Errorf("failed to create nats server: %w", err)
	}

	go ns.Start()

	// Increased timeout and better error handling
	timeout := 10 * time.Second
	if !ns.ReadyForConnections(timeout) {
		ns.Shutdown()
		return nil, fmt.Errorf("nats server failed to start within %v", timeout)
	}

	log.Printf("NATS server started on port %d", opts.Port)
	return &Server{nats: ns}, nil
}

func (s *Server) Shutdown(ctx context.Context) error {
	s.nats.Shutdown()
	s.nats.WaitForShutdown()
	return nil
}

func (s *Server) URL() string {
	if s.nats.ClientURL() == "" {
		return "nats://localhost:4222"
	}
	return s.nats.ClientURL()
}
