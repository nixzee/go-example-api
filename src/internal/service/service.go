package service

import (
	"context"
	"net/http"
	"os"
	"strings"
	"time"

	chi "github.com/go-chi/chi/v5"
	middleware "github.com/go-chi/chi/v5/middleware"
	service "github.com/kardianos/service"
	storage "github.com/nixzee/go-example-api/internal/api/v1/storage"
	logging "github.com/nixzee/go-example-api/internal/logging"
)

var _ service.Interface = (*program)(nil)

type program struct {
	version string
	commit  string
	server  *http.Server
}

func NewProgram(version, commit string) service.Interface {
	return &program{
		version: version,
		commit:  commit,
	}
}

// Start the service
func (p *program) Start(s service.Service) (err error) {
	logging.Info.Println("Version: " + p.version)
	logging.Info.Println("Commit: " + p.commit)
	logging.Debug.Println("Starting Server")
	// Load Enviromental variables
	envs := loadEnvVars()
	// Create a router
	r := chi.NewRouter()
	// Add logging middle ware
	r.Use(middleware.Logger)
	// Add controllers
	storageService := storage.NewStorageService(envs["ACCOUNT_NAME"], envs["SAS_TOKEN"])  // TODO: This is not safe and should check first
	r.Mount("/v1/storage", storage.NewStorageController(storageService).RegisterRoutes()) // Dependency inject the service
	// Start the server
	serverAddr := ":8080"
	p.server = &http.Server{Addr: serverAddr, Handler: r}
	go func() {
		if err := p.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logging.Error.Println("Error starting server:", err)
		}
	}()
	logging.Info.Println("Server is running on", serverAddr)
	logging.Debug.Println("Started Server")

	return
}

// Stop the service
func (p *program) Stop(s service.Service) (err error) {
	logging.Debug.Println("Stopping Server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// Shutdown the server
	if err := p.server.Shutdown(ctx); err != nil {
		logging.Error.Println(err)
	}

	logging.Debug.Println("Stopped Server")
	return
}

// Load enviromental variables
func loadEnvVars() map[string]string {
	envs := make(map[string]string)
	for _, env := range os.Environ() {
		parts := strings.Split(env, "=")
		if len(parts) == 2 {
			envs[parts[0]] = parts[1]
		}
	}
	return envs
}
