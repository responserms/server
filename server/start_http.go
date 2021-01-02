package server

import (
	"context"
	"fmt"

	"github.com/responserms/server/internal/http"
)

func (s *Server) startHTTPService(_ context.Context) error {
	logger := s.log.Component("http")

	options := &http.Options{
		BindAddress:   s.cfg.HTTP.BindAddress,
		Port:          s.cfg.HTTP.Port,
		MaxUploadSize: s.cfg.HTTP.MaxUploadSize,
	}

	if s.cfg.HTTP.TLS != nil {
		logger.Info("using TLS")

		options.TLS = &http.TLSOptions{
			Port:     s.cfg.HTTP.TLS.Port,
			CertPath: s.cfg.HTTP.TLS.CertPath,
			KeyPath:  s.cfg.HTTP.TLS.KeyPath,
		}
	}

	if s.cfg.HTTP.TLS.Auto != nil {
		logger.Info("automatically configuring certificate")

		options.AutoTLS = &http.AutoTLSOptions{
			Production: s.cfg.HTTP.TLS.Auto.Production,
			Email:      s.cfg.HTTP.TLS.Auto.Email,
			Domains:    s.cfg.HTTP.TLS.Auto.Domains,
		}

		if s.cfg.HTTP.TLS.Auto.DNS != nil {
			logger.Info("configuring DNS solver")

			options.AutoTLS.DNS = true
			options.AutoTLS.DNSProvider = s.cfg.HTTP.TLS.Auto.DNS.Provider
			options.AutoTLS.DNSAPIToken = s.cfg.HTTP.TLS.Auto.DNS.APIToken
		}
	}

	svr, err := http.New(options)
	if err != nil {
		return fmt.Errorf("new http: %w", err)
	}

	s.http = svr

	return nil
}
