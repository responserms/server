package server

import (
	"context"
	"errors"
	"fmt"

	"github.com/responserms/server/internal/cluster"
)

func (s *Server) startClusterService(ctx context.Context, errChan chan error) error {
	if s.cfg.Cluster == nil {
		return errors.New("cluster configuration filed")
	}

	options := &cluster.Options{
		EnvType:           s.cfg.Cluster.Environment,
		BindAddress:       s.cfg.Cluster.BindAddress,
		BindPort:          s.cfg.Cluster.BindPort,
		MemberBindAddress: s.cfg.Cluster.BindMemberAddress,
		MemberBindPort:    s.cfg.Cluster.BindMemberPort,
	}

	if s.cfg.Cluster.Join != nil {
		options.AutoJoin = &cluster.AutoJoinOptions{
			Type:    s.cfg.Cluster.Join.Type,
			URL:     s.cfg.Cluster.Join.URL,
			Subject: s.cfg.Cluster.Join.Subject,
		}
	}

	c, err := cluster.New(s.log, options)
	if err != nil {
		return fmt.Errorf("start cluster service: %w", err)
	}

	s.cluster = c

	// start the cluster
	go func() {
		if err := s.cluster.Start(); err != nil {
			errChan <- err
		}
	}()

	return nil
}
