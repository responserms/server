package services

import (
	"context"

	"github.com/responserms/server/internal/services/http"
	"github.com/responserms/server/pkg/config"
)

type httpService struct{}

func (s httpService) Name() string {
	return "http"
}

func (s httpService) generateOptions(cfg config.Config) *http.Options {
	if cfg.HTTP == nil {
		cfg.HTTP = &config.HTTPConfig{}
	}

	if cfg.Developer == nil {
		cfg.Developer = &config.DeveloperConfig{}
	}

	c := &http.Options{
		BindAddress:            cfg.HTTP.BindAddress,
		Port:                   cfg.HTTP.Port,
		MaxUploadSize:          cfg.HTTP.MaxUploadSize,
		ServeProfiler:          cfg.Developer.Profiling,
		ServeGraphQL:           true,
		ServeGraphQLPlayground: true,
	}

	if c.BindAddress == "" {
		c.BindAddress = "0.0.0.0"
	}

	if c.Port == 0 {
		c.Port = 8080
	}

	if c.MaxUploadSize == 0 {
		c.MaxUploadSize = 10000
	}

	// TODO: Need to implement TLS and Automatic TLS options

	return c
}

func (s httpService) Startup(ctx context.Context, svcs *Services, errChan chan error) {
	cfg := svcs.cfg
	if cfg == nil {
		cfg = &config.Config{}
	}

	h, err := http.New(svcs.log, svcs, s.generateOptions(*cfg))
	if err != nil {
		errChan <- err

		return
	}

	svcs.http = h

	go func() {
		err := svcs.http.Start()
		if err != nil {
			errChan <- err
		}
	}()
}

func (s httpService) Shutdown(_ context.Context, svcs *Services) error {
	if svcs.http == nil {
		return nil
	}

	svcs.http = nil

	return nil
}
