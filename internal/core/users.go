package core

import (
	"context"
	"fmt"

	"github.com/responserms/server/ent"
	"github.com/responserms/server/ent/user"
	"github.com/responserms/server/internal/services/cluster"
)

var _ UsersService = (*usersService)(nil)

type (
	// Users represents many User instances.
	Users []*ent.User

	// UsersService is an instance of the service that contols users, updates to users, and user subscribers.
	UsersService interface {
		GetUserByID(ctx context.Context, id int) (*ent.User, error)
		GetUserByEmail(ctx context.Context, email string) (*ent.User, error)
	}

	usersService struct {
		core *Core
	}
)

func (s *usersService) userCacheKey(user int) cluster.Key {
	return cluster.Key(fmt.Sprintf("users.%d", user))
}

func (s *usersService) cacheStore() (cluster.Storer, error) {
	return s.core.svcs.Cluster().NewStore("users")
}

func (s *usersService) GetUserByID(ctx context.Context, id int) (*ent.User, error) {
	store, err := s.cacheStore()
	if err != nil {
		return nil, fmt.Errorf("GetUserByID: %w", err)
	}

	cached, err := s.userCacheKey(id).
		GetOrPut(store, func() (interface{}, error) {
			return s.core.svcs.Database().
				User.
				Get(ctx, id)
		})

	if ent.MaskNotFound(err) != nil {
		return nil, fmt.Errorf("core.GetUserByID: %w", err)
	}

	user, ok := cached.(*ent.User)
	if !ok {
		return nil, fmt.Errorf("core.GetUserByID: unable to convert user")
	}

	return user, nil
}

func (s *usersService) GetUserByEmail(ctx context.Context, email string) (*ent.User, error) {
	user, err := s.core.svcs.Database().
		User.
		Query().
		Where(user.EmailEQ(email)).
		First(ctx)

	if ent.MaskNotFound(err) != nil {
		return nil, fmt.Errorf("core.GetUserByEmail: %w", err)
	}

	return user, nil
}
