package api

import (
	"context"
	"database/sql"
	"net/http"
	"os"
	"os/signal"
	"time"
	"user-app/internal/config"
	"user-app/internal/utils"

	"github.com/Masterminds/squirrel"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

// Server represents the HTTP server configuration and lifecycle.
type Server struct {
	db     *sql.DB
	config config.Config
	sq     squirrel.StatementBuilderType
	logger logrus.FieldLogger
	server *http.Server
}

// NewServer creates a new instance of the HTTP server.
func NewServer() *Server {
	s := &Server{
		logger: logrus.New(),
	}

	// Load configuration
	conf := config.NewConfig()
	config, err := conf.Load()
	if err != nil {
		s.logger.Fatalf("Failed to fetch configuration: %v", err)
	}
	s.config = *config

	// Initialize SQL builder
	s.sq = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

	// Retry database connection for 3 times
	for i := 0; i < 3; i++ {
		// Connect to the database
		db, err := utils.ConnectDB(s.config.PGUser, s.config.PGPassword, s.config.PGHost, s.config.PGDB, s.config.PGPort)
		if err != nil {
			s.logger.Errorf("Failed to initialize database: %v", err)
			if i < 2 {
				// Wait for a moment before retrying
				time.Sleep(5 * time.Second)
				continue
			} else {
				s.logger.Fatalf("Maximum retries reached. Exiting.")
			}
		}

		// Database connection successfully established
		s.db = db
		s.logger.Info("Database connection successfully established")
		break
	}

	// Initialize routes
	routes := s.InitRouter()

	// Initialize HTTP server
	s.server = &http.Server{
		Addr:    ":8080",
		Handler: routes,
	}

	return s
}

// Start launches the HTTP server in a separate goroutine.
func (s *Server) Start() {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				s.logger.Errorf("Recovered from panic: %v", r)
			}
		}()

		s.logger.Info("Server is listening on port 8080...")
		if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			s.logger.Fatalf("Server startup error: %v", err)
		}
	}()
}

// Shutdown gracefully shuts down the HTTP server and closes the database connection.
func (s *Server) Shutdown() {
	// Create a channel to receive signals
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	// Block until a signal is received
	<-stop

	// Create a context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Shutdown the server
	if err := s.server.Shutdown(ctx); err != nil {
		s.logger.Errorf("Server shutdown failed: %v", err)
	} else {
		s.logger.Info("Server shutdown successfully")
	}

	// Close the database connection
	if err := s.db.Close(); err != nil {
		s.logger.Errorf("Database close failed: %v", err)
	} else {
		s.logger.Info("Database connection closed")
	}
}
