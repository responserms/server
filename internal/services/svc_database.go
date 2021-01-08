package services

import (
	"context"
	"fmt"

	"github.com/responserms/server/internal/log"
	"github.com/responserms/server/internal/services/database"
	"github.com/responserms/server/pkg/config"
)

type databaseService struct{}

func (s databaseService) Name() string {
	return "database"
}

func (s databaseService) generateOptions(logger log.ComponentLogger, cfg config.Config) *database.Options {
	if cfg.Database == nil {
		cfg.Database = &config.DatabaseConfig{}
	}

	c := &database.Options{
		Debug:    cfg.LogLevel <= log.Debug,
		Driver:   database.Driver(cfg.Database.Type),
		Path:     cfg.Database.Path,
		Host:     cfg.Database.Host,
		Port:     cfg.Database.Port,
		URL:      cfg.Database.URL,
		Name:     cfg.Database.Name,
		Username: cfg.Database.Username,
		Password: cfg.Database.Password,
		Options:  cfg.Database.Options,
		Logger:   logger,
	}

	if c.Driver == "" {
		c.Driver = database.EmbeddedDriver
	}

	// TODO: Finish setting defaults

	return c
}

func (s databaseService) Startup(ctx context.Context, svcs *Services, errChan chan error) {
	c, err := database.Configure(ctx, s.generateOptions(svcs.log, *svcs.cfg))
	if err != nil {
		errChan <- err

		return
	}

	svcs.database = c
}

func (s databaseService) Shutdown(ctx context.Context, svcs *Services) error {
	if svcs.database == nil {
		return nil
	}

	err := svcs.database.Close()
	if err != nil {
		return fmt.Errorf("databaseService.Shutdown: %w", err)
	}

	svcs.database = nil

	return nil
}
