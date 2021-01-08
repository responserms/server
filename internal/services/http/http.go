package http

import (
	"errors"
	"fmt"
	"net"
	"net/http"
	"net/http/pprof"
	"strconv"

	"github.com/responserms/server/internal/log"
	"github.com/responserms/server/internal/reqdata"
	"github.com/responserms/server/internal/services/cluster"
	"github.com/responserms/server/internal/services/database"
	"github.com/responserms/server/internal/services/events"
)

// AutoTLSOptions configures automatic TLS for Response Server's TCP server.
type AutoTLSOptions struct {
	Domains    []string
	Email      string
	Production bool

	DNS         bool
	DNSProvider string
	DNSAPIToken string
}

// TLSOptions configured TLS for Response Server's TCP server.
type TLSOptions struct {
	Port     int
	CertPath string
	KeyPath  string
}

// Options configures Server.
type Options struct {
	BindAddress   string
	Port          int
	MaxUploadSize int
	TLS           *TLSOptions
	AutoTLS       *AutoTLSOptions

	ServeProfiler          bool
	ServeGraphQL           bool
	ServeGraphQLPlayground bool
}

// Server is the TCP server that powers Response Server's HTTP(S) services.
type Server struct {
	options *Options
	mux     *http.ServeMux
	logger  log.Logger
}

// Backend is implemented by a service capable of acting as the backend for the HTTP server.
type Backend interface {
	Database() *database.Database
	Cluster() cluster.Cluster
	Events() events.PubSub
}

// New creates a new http.Server instance and an internal multiplexer to handler
// HTTP routing.
func New(logger log.Logger, backend Backend, options *Options) (*Server, error) {
	server := &Server{
		options: options,
		logger:  logger,
		mux:     http.NewServeMux(),
	}

	if options.ServeProfiler {
		server.mux.HandleFunc("/debug/pprof/", pprof.Index)
		server.mux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
		server.mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
		server.mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
		server.mux.HandleFunc("/debug/pprof/trace", pprof.Trace)

		server.mux.Handle("/debug/pprof/block", pprof.Handler("block"))
		server.mux.Handle("/debug/pprof/goroutine", pprof.Handler("goroutine"))
		server.mux.Handle("/debug/pprof/heap", pprof.Handler("heap"))
		server.mux.Handle("/debug/pprof/threadcreate", pprof.Handler("threadcreate"))
		server.mux.Handle("/debug/pprof/mutex", pprof.Handler("mutex"))
	}

	if options.ServeGraphQL {
		server.registerGraphQL(backend)
	}

	if options.ServeGraphQLPlayground {
		server.registerGraphQLPlayground()
	}

	return server, nil
}

// Start serves the HTTP(S) endpoints for Response Server.
func (s *Server) Start() error {
	if s.mux == nil || s.options == nil {
		return errors.New("use the New() consutructor instead of initializing the Server directly")
	}

	if s.options.TLS == nil && s.options.AutoTLS == nil {
		err := http.ListenAndServe(
			net.JoinHostPort(
				s.options.BindAddress,
				strconv.Itoa(s.options.Port),
			),

			http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				s.mux.ServeHTTP(w, r.WithContext(reqdata.ContextFromRequest(r)))
			}),
		)

		if err != nil {
			return fmt.Errorf("start http: %w", err)
		}
	}

	return fmt.Errorf("start http: tls not yet implemented")
}
