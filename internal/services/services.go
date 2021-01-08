package services

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/responserms/server/internal/log"
	"github.com/responserms/server/internal/services/cluster"
	"github.com/responserms/server/internal/services/database"
	"github.com/responserms/server/internal/services/events"
	"github.com/responserms/server/internal/services/http"
	"github.com/responserms/server/pkg/config"
)

// Services are the services available to Response Server.
type Services struct {
	once sync.Once
	log  log.Logger
	cfg  *config.Config

	http     *http.Server
	database *database.Database
	cluster  cluster.Cluster
	events   events.PubSub
}

var (
	serviceOrder = []service{
		&clusterService{},
		&databaseService{},
		&eventService{},
		&httpService{},
	}
)

// service is implemented bt each service that should be started and shutdown
type service interface {
	Name() string
	Startup(ctx context.Context, svcs *Services, errChan chan error)
	Shutdown(ctx context.Context, svcs *Services) error
}

func reverseServiceOrder(workingOrder []service) []service {
	for i := len(workingOrder)/2 - 1; i >= 0; i-- {
		opp := len(workingOrder) - 1 - i
		workingOrder[i], workingOrder[opp] = workingOrder[opp], workingOrder[i]
	}

	return workingOrder
}

// New initializes all of the services using the application config.
func New(ctx context.Context, logger log.Logger, cfg *config.Config) (*Services, error) {
	if cfg == nil {
		cfg = &config.Config{}
	}

	s := &Services{
		log: logger,
		cfg: cfg,
	}

	return s, nil
}

// Start starts services in the defined startup order.
func (s *Services) Start(ctx context.Context, errChan chan error) {
	s.log.Info("starting services")

	for order, svc := range serviceOrder {
		s.log.Info(fmt.Sprintf("starting %s", svc.Name()), log.Attributes{
			"order": order + 1,
		})

		svc.Startup(ctx, s, errChan)
	}
}

// Shutdown stops services in the reverse order of their defined startup order.
func (s *Services) Shutdown(ctx context.Context) error {
	s.log.Info("stopping services")

	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(2*time.Second))
	defer cancel()

	var retErr error
	for order, svc := range reverseServiceOrder(serviceOrder) {
		s.log.Info(fmt.Sprintf("shutting down %s", svc.Name()), log.Attributes{
			"order": order + 1,
		})

		if err := svc.Shutdown(ctx, s); err != nil {
			retErr = err
			break
		}
	}

	return retErr
}
