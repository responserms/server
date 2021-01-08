package core

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/responserms/server/ent"
	"github.com/responserms/server/internal/log"
	"github.com/responserms/server/internal/services/cluster"
	"github.com/responserms/server/internal/services/events"
	"github.com/responserms/server/pkg/sdk/event"
)

var (
	// ErrSessionsSessionDoesNotExist is returned when a session being terminates does not
	// exist.
	ErrSessionsSessionDoesNotExist = errors.New("session does not exist")
)

type (
	// SessionsService allows managing sessions in Response.
	SessionsService interface {
		GetSession(ctx context.Context, id int) (*ent.Session, error)
		CreateSession(ctx context.Context, opts *CreateSessionOptions) (*ent.Session, error)
		TerminateSession(ctx context.Context, opts *TerminateSessionOptions) error
	}

	// CreateSessionOptions configures the CreateSession method.
	CreateSessionOptions struct {
		UserID                int
		BrowserName           string
		BrowserVersion        string
		DeviceType            string
		DeviceOperatingSystem string
		IPAddress             string
	}

	// TerminateSessionOptions configures the TerminateSession method.
	TerminateSessionOptions struct {
		SessionID int
		Reason    string
	}
)

type sessionsService struct {
	core *Core
}

func (s *sessionsService) sessionCacheKey(user int) cluster.Key {
	return cluster.Key(fmt.Sprintf("sessions.%d", user))
}

func (s *sessionsService) cacheStore() (cluster.Storer, error) {
	return s.core.svcs.Cluster().NewStore("sessions")
}

func (s *sessionsService) CreateSession(ctx context.Context, opts *CreateSessionOptions) (*ent.Session, error) {
	store, err := s.cacheStore()
	if err != nil {
		return nil, fmt.Errorf("Sessions.CreateSession: %w", err)
	}

	create := s.core.svcs.Database().
		Session.
		Create()

	session, err := create.SetBrowserName(opts.BrowserName).
		SetBrowserVersion(opts.BrowserVersion).
		SetDeviceType(opts.DeviceType).
		SetDeviceOs(opts.DeviceOperatingSystem).
		SetIPAddress(opts.IPAddress).
		SetUserID(opts.UserID).
		Save(ctx)

	if err != nil {
		return nil, fmt.Errorf("Sessions.CreateSession: %w", err)
	}

	s.sessionCacheKey(session.ID).
		Put(store, session)

	msgs, err := events.NewMessage(session, event.AnySessionStarted())
	if err != nil {
		s.core.log.Error("Failed to build event message.", log.Attributes{
			"error": err.Error(),
		})
	}

	if len(msgs) > 0 {
		err = s.core.svcs.Events().
			Publish(ctx, msgs...)

		if err != nil {
			s.core.log.Error("Failed to publish an event.", log.Attributes{
				"event": event.AnySessionStarted(),
				"error": err.Error(),
			})
		}
	}

	return session, nil
}

func (s *sessionsService) GetSession(ctx context.Context, id int) (*ent.Session, error) {
	store, err := s.cacheStore()
	if err != nil {
		return nil, fmt.Errorf("Sessions.GetSession: %w", err)
	}

	cached, err := s.sessionCacheKey(id).
		GetOrPut(store, func() (interface{}, error) {
			return s.core.svcs.Database().
				Session.
				Get(ctx, id)
		})

	if ent.MaskNotFound(err) != nil {
		return nil, fmt.Errorf("Sessions.GetSession: %w", err)
	}

	session, ok := cached.(*ent.Session)
	if !ok {
		return nil, fmt.Errorf("Sessions.GetSession: unable to convert session")
	}

	return session, nil
}

func (s *sessionsService) TerminateSession(ctx context.Context, opts *TerminateSessionOptions) error {
	session, err := s.GetSession(ctx, opts.SessionID)
	if err != nil {
		return fmt.Errorf("Sessions.TerminateSession: %w", err)
	}

	if session == nil {
		return ErrSessionsSessionDoesNotExist
	}

	session, err = session.Update().
		SetTerminatedAt(time.Now()).
		Save(ctx)

	msgs, err := events.NewMessage(
		session,
		event.SessionTerminated(session), event.AnySessionTerminated(),
	)

	if err != nil {
		s.core.log.Error("Failed to build event message.", log.Attributes{
			"error": err.Error(),
		})
	}

	if len(msgs) > 0 {
		err = s.core.svcs.Events().
			Publish(ctx, msgs...)

		if err != nil {
			s.core.log.Error("Failed to publish an event.", log.Attributes{
				"event": event.AnySessionStarted(),
				"error": err.Error(),
			})

			s.core.log.Error("Failed to publish an event.", log.Attributes{
				"event": event.SessionTerminated(session),
				"error": err.Error(),
			})
		}
	}

	if err != nil {
		return fmt.Errorf("Sessions.TerminateSession: %w", err)
	}

	return nil
}
