package server

import (
	"context"
	"fmt"

	"github.com/responserms/server/internal/http"
)

func (s *Server) startHTTPService(_ context.Context, errChan chan error) error {
	logger := s.log.Component("http")

	options := &http.Options{
		BindAddress:   s.cfg.HTTP.BindAddress,
		Port:          s.cfg.HTTP.Port,
		MaxUploadSize: s.cfg.HTTP.MaxUploadSize,

		ServeProfiler:          s.cfg.Developer.Profiling,
		ServeGraphQL:           true,
		ServeGraphQLPlayground: true,
	}

	if s.cfg.HTTP.TLS != nil {
		logger.Info("using TLS")

		options.TLS = &http.TLSOptions{
			Port:     s.cfg.HTTP.TLS.Port,
			CertPath: s.cfg.HTTP.TLS.CertPath,
			KeyPath:  s.cfg.HTTP.TLS.KeyPath,
		}

		if s.cfg.HTTP.TLS.Auto != nil {
			logger.Info("using automatic TLS certificates")

			options.AutoTLS = &http.AutoTLSOptions{
				Production: s.cfg.HTTP.TLS.Auto.Production,
				Email:      s.cfg.HTTP.TLS.Auto.Email,
				Domains:    s.cfg.HTTP.TLS.Auto.Domains,
			}

			if s.cfg.HTTP.TLS.Auto.DNS != nil {
				logger.Info("using DNS solver")

				options.AutoTLS.DNS = true
				options.AutoTLS.DNSProvider = s.cfg.HTTP.TLS.Auto.DNS.Provider
				options.AutoTLS.DNSAPIToken = s.cfg.HTTP.TLS.Auto.DNS.APIToken
			}
		}
	}

	svr, err := http.New(s, options)
	if err != nil {
		return fmt.Errorf("new http: %w", err)
	}

	s.http = svr

	go func() {
		if err := s.http.Start(); err != nil {
			errChan <- err
		}
	}()

	return nil
}
