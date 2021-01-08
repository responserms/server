package services

import (
	"context"
	"fmt"

	"github.com/responserms/server/internal/services/cluster"
	"github.com/responserms/server/pkg/config"
)

type clusterService struct{}

func (s clusterService) Name() string {
	return "cluster"
}

func (s clusterService) generateOptions(cfg config.Config) *cluster.Options {
	c := &cluster.Options{
		EnvType:           cfg.Cluster.Environment,
		BindAddress:       cfg.Cluster.BindAddress,
		BindPort:          cfg.Cluster.BindPort,
		MemberBindAddress: cfg.Cluster.BindMemberAddress,
		MemberBindPort:    cfg.Cluster.BindMemberPort,
	}

	if c.BindAddress == "" {
		c.BindAddress = "0.0.0.0"
	}

	if c.BindPort == 0 {
		c.BindPort = 3320
	}

	if c.MemberBindAddress == "" {
		c.MemberBindAddress = "0.0.0.0"
	}

	if c.MemberBindPort == 0 {
		c.MemberBindPort = 3322
	}

	if cfg.Cluster.Join != nil {
		c.AutoJoin = &cluster.AutoJoinOptions{
			Type:    cfg.Cluster.Join.Type,
			URL:     cfg.Cluster.Join.URL,
			Subject: cfg.Cluster.Join.Subject,
		}

		// nats autojoin defaults
		if c.AutoJoin.Type == "nats" {
			if c.AutoJoin.URL == "" {
				c.AutoJoin.URL = "nats://localhost:4222"
			}

			if c.AutoJoin.Subject == "" {
				c.AutoJoin.Subject = "response-cluster"
			}
		}
	}

	return c
}

func (s clusterService) Startup(ctx context.Context, svcs *Services, errChan chan error) {
	cfg := svcs.cfg
	if cfg == nil {
		cfg = &config.Config{}
	}

	c, err := cluster.New(svcs.log, s.generateOptions(*cfg))
	if err != nil {
		errChan <- err

		return
	}

	svcs.cluster = c

	go func(c cluster.Cluster) {
		err := c.Start()

		if err != nil {
			errChan <- err
		}
	}(svcs.cluster)
}

func (s clusterService) Shutdown(ctx context.Context, svcs *Services) error {
	if svcs.cluster == nil {
		return nil
	}

	err := svcs.cluster.Shutdown(ctx)
	if err != nil {
		return fmt.Errorf("clusterService.Shutdown: %w", err)
	}

	svcs.cluster = nil

	return nil
}
