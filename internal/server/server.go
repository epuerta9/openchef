package server

import (
	"context"
	"fmt"
	"log"

	"github.com/epuerta9/openchef/internal/api/ochefopenai"
	"github.com/epuerta9/openchef/internal/config"
	"github.com/epuerta9/openchef/internal/database"
	"github.com/epuerta9/openchef/web/handlers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/epuerta9/openchef/internal/services/agent"
	"github.com/epuerta9/openchef/internal/services/chat"
	"github.com/epuerta9/openchef/internal/services/communicator"
	"github.com/epuerta9/openchef/internal/services/orchestrator"
	"github.com/epuerta9/openchef/internal/services/swarm"
)

type Server struct {
	config *config.Config
	db     *database.DB
	openai *ochefopenai.Service
	echo   *echo.Echo
}

func New(cfg *config.Config, db *database.DB) *Server {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Initialize communicator
	comm, err := communicator.New("nats://localhost:4222")
	if err != nil {
		log.Fatalf("Failed to create communicator: %v", err)
	}

	// Initialize services with DB and communicator
	chatService := chat.New(comm, db.Queries)
	agentService := agent.New(comm, db.Queries)
	orchestratorService := orchestrator.New(comm, db.Queries)
	swarmService := swarm.New(comm, db.Queries)

	// Initialize OpenAI API layer
	openaiService := ochefopenai.NewService(
		chatService,
		agentService,
		orchestratorService,
		swarmService,
		cfg.OpenAIKey,
	)

	// Initialize handlers
	webHandler := handlers.New(db)
	openaiHandler := ochefopenai.NewHandler(openaiService)

	// Register routes
	e.GET("/", webHandler.HandleHome)
	e.Static("/static", "web/static")
	openaiHandler.RegisterRoutes(e)

	return &Server{
		config: cfg,
		db:     db,
		openai: openaiService,
		echo:   e,
	}
}

func (s *Server) Start() error {
	return s.echo.Start(fmt.Sprintf(":%d", s.config.Port))
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.echo.Shutdown(ctx)
}
