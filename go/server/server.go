package server

import (
	"context"
	"fmt"
	"log"
	"net/http"

	pipanel "github.com/BenJetson/pipanel/go"
)

// Server provides a webserver that is capable of receiving and handling
// the PiPanel events.
type Server struct {
	log      *log.Logger
	frontend *pipanel.Frontend
	httpd    *http.Server
}

// New creates a new Server instance, binding to the given port and frontend.
func New(l *log.Logger, port int, frontend *pipanel.Frontend) *Server {
	// Create a multiplexer for routing requests.
	mux := http.NewServeMux()

	// Create a server instance.
	s := Server{
		log: l,
		httpd: &http.Server{
			Addr:     fmt.Sprintf(":%d", port),
			ErrorLog: l,
			Handler:  mux,
		},
		frontend: frontend,
	}

	// Define routes.
	mux.HandleFunc("/alert", s.handleAlertEvent)
	mux.HandleFunc("/sound", s.handleSoundEvent)
	mux.HandleFunc("/power", s.handlePowerEvent)
	mux.HandleFunc("/brightness", s.handleBrightnessEvent)

	return &s
}

// ListenAndServe instructs the server to bind to the configured port and
// listen for requests to handle. Will block until the server terminates.
// Upon termination, this function will close the channel given by the
// parameter, allowing for this server to run in a separate goroutine.
func (s *Server) ListenAndServe(closeOnReturn chan<- struct{}) {
	defer close(closeOnReturn)

	s.log.Println("Server started.")

	err := s.httpd.ListenAndServe()

	if err != nil && err != http.ErrServerClosed {
		s.log.Println("Server died with error:", err)
		return
	}
	s.log.Println("Server has gracefully stopped.")
}

// Shutdown tears down this Server and releases its resources.
func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpd.Shutdown(ctx)
}
