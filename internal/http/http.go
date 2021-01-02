package http

import (
	"net/http"
	"net/http/pprof"
)

type AutoTLSOptions struct {
	Domains    []string
	Email      string
	Production bool

	DNS         bool
	DNSProvider string
	DNSAPIToken string
}

type TLSOptions struct {
	Port     int
	CertPath string
	KeyPath  string
}

type Options struct {
	BindAddress   string
	Port          int
	MaxUploadSize int

	TLS           *TLSOptions
	AutoTLS       *AutoTLSOptions
	ServeProfiler bool
}

type Server struct {
	options *Options
	mux     *http.ServeMux
}

// New creates a new http.Server instance and an internal multiplexer to handler
// HTTP routing.
func New(options *Options) (*Server, error) {
	server := &Server{
		options: options,
		mux:     http.NewServeMux(),
	}

	if options.ServeProfiler {
		server.mux.Handle("/debug", pprof.Handler("server"))
	}

	return server, nil
}

func (s *Server) Start() error {
	return nil
}
