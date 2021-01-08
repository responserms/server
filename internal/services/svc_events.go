package services

import (
	"context"

	"github.com/responserms/server/internal/log"
	"github.com/responserms/server/internal/services/events"
	"github.com/responserms/server/pkg/config"
)

type eventService struct{}

func (s eventService) Name() string {
	return "events"
}

func (s eventService) generateOptions(log log.ComponentLogger, cfg config.Config) *events.Options {
	if cfg.Events == nil {
		cfg.Events = &config.EventsConfig{}
	}

	c := &events.Options{
		Logger:  log,
		Weight:  10,
		Driver:  events.Driver(cfg.Events.Type),
		URL:     cfg.Events.URL,
		Subject: cfg.Events.Subject,
	}

	switch events.Driver(cfg.Events.Type) {
	case events.NATS:
		if c.URL == "" {
			c.URL = "nats://localhost:4222"
		}

		if c.Subject == "" {
			c.Subject = "response-events"
		}
	case events.Embedded:
		if c.URL == "" {
			c.URL = "mem://"
		}

		if c.Subject == "" {
			c.Subject = "response-events"
		}
	default:
		c.Driver = events.Embedded
		c.URL = "mem://"
		c.Subject = "response-events"
	}

	return c
}

func (s eventService) Startup(ctx context.Context, svcs *Services, errChan chan error) {
	cfg := svcs.cfg
	if cfg == nil {
		cfg = &config.Config{}
	}

	e, err := events.New(ctx, s.generateOptions(svcs.log, *cfg))

	if err != nil {
		errChan <- err
		return
	}

	svcs.events = e
}

func (s eventService) Shutdown(ctx context.Context, svcs *Services) error {
	if svcs.events == nil {
		return nil
	}

	svcs.events = nil

	return nil
}
