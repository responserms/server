package core

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/responserms/server/ent"
	"github.com/responserms/server/ent/user"
	"github.com/responserms/server/internal/log"
	"github.com/responserms/server/internal/services/events"
	"github.com/responserms/server/pkg/sdk/event"
	"golang.org/x/crypto/bcrypt"
)

var (
	// ErrAuthInvalidCredentials is returned when the credentials provided are invalid.
	ErrAuthInvalidCredentials = errors.New("credentials provided are not valid")

	// ErrAuthPasswordNotSet is returned when a User attempts to authenticate with a password but
	// has not yet set a password (such as when creating their account with a provider).
	ErrAuthPasswordNotSet = errors.New("password is not set")

	// ErrAuthEmailAlreadyRegistered is returned when a User attempts to register with an email address
	// that is already in-use.
	ErrAuthEmailAlreadyRegistered = errors.New("email already registered")
)

type (
	// AuthService handles user authentication.
	AuthService interface {

		// TokenFromCredentials returns a token given valid credentials or an error if the credetnails are invalid.
		TokenFromCredentials(ctx context.Context, opts *TokenFromCredentialsOptions) (*TokenResult, error)

		// TokenFromProvider returns a token given valid provider credentials or an error if the credetnails are invalid.
		TokenFromProvider(ctx context.Context, opts *TokenFromProviderOptions) (*TokenResult, error)

		// Register is called when registering a User.
		Register(ctx context.Context, opts *RegisterOptions) (*ent.User, error)
	}

	// RegisterOptions configures the registration of a User.
	RegisterOptions struct {
		Name     string
		Email    string
		Password string
	}

	// TokenFromCredentialsOptions configures the TokenFromCredentials method.
	TokenFromCredentialsOptions struct {
		Email                 string
		Password              string
		BrowserName           string
		BrowserVersion        string
		DeviceType            string
		DeviceOperatingSystem string
		IPAddress             string
	}

	// TokenFromProviderOptions configures the TokenFromProvider method.
	TokenFromProviderOptions struct {
		Provider    int
		AccessToken string
	}

	// TokenResult provides a token when the token has been succcessfully created from an authentication
	// attempt.
	TokenResult struct {
		Token     *AccessToken
		ExpiredAt time.Time
	}

	authService struct {
		core *Core
	}
)

// Login handles logs in a user using their credentials.
func (s *authService) TokenFromCredentials(ctx context.Context, opts *TokenFromCredentialsOptions) (*TokenResult, error) {
	user, err := s.core.svcs.Database().User.
		Query().
		Where(user.EmailEQ(opts.Email)).
		First(ctx)

	if ent.MaskNotFound(err) != nil {
		return nil, fmt.Errorf("Auth.TokenFromCredentials: %w", err)
	}

	if user == nil {
		return nil, ErrAuthInvalidCredentials
	}

	if user.Password == nil {
		return nil, ErrAuthPasswordNotSet
	}

	if err := bcrypt.CompareHashAndPassword([]byte(*user.Password), []byte(opts.Password)); err != nil {
		return nil, ErrAuthInvalidCredentials
	}

	session, err := s.core.Sessions.CreateSession(ctx, &CreateSessionOptions{
		UserID:                user.ID,
		BrowserName:           opts.BrowserName,
		BrowserVersion:        opts.BrowserVersion,
		DeviceType:            opts.DeviceType,
		DeviceOperatingSystem: opts.DeviceOperatingSystem,
		IPAddress:             opts.IPAddress,
	})
	if err != nil {
		return nil, fmt.Errorf("Auth.TokenFromCredentials: %w", err)
	}

	token, err := s.createUserToken(ctx, &createUserTokenOptions{
		Session:       session,
		User:          user,
		TokenDuration: time.Duration(8) * time.Hour,
	})
	if err != nil {
		return nil, fmt.Errorf("Auth.TokenFromCredentials: %w", err)
	}

	return &TokenResult{
		Token:     token,
		ExpiredAt: time.Now().Add(8 * time.Hour),
	}, nil
}

func (s *authService) TokenFromProvider(ctx context.Context, opts *TokenFromProviderOptions) (*TokenResult, error) {
	panic("not implemented") // TODO: Implement
}

// Register is called when registering a User.
func (s *authService) Register(ctx context.Context, opts *RegisterOptions) (*ent.User, error) {
	existingUser, err := s.core.Users.GetUserByEmail(ctx, opts.Email)
	if err != nil {
		return nil, fmt.Errorf("Auth.Register: %w", err)
	}

	if existingUser != nil {
		return nil, ErrAuthEmailAlreadyRegistered
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(opts.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("Auth.Register: %w", err)
	}

	create := s.core.svcs.Database().
		User.
		Create()

	user, err := create.SetName(opts.Name).
		SetEmail(opts.Email).
		SetPassword(string(hashedPassword)).
		Save(ctx)

	if err != nil {
		return nil, fmt.Errorf("Auth.Register: %w", err)
	}

	msgs, err := events.NewMessage(user, event.AnyUserRegistered())
	if err != nil {
		s.core.log.Error("Failed to build event message.", log.Attributes{
			"error": err.Error(),
		})
	}

	if len(msgs) > 0 {
		err := s.core.svcs.Events().
			Publish(ctx, msgs...)

		if err != nil {
			s.core.log.Error("Failed to publish an event.", log.Attributes{
				"event": event.AnyUserRegistered(),
				"error": err.Error(),
			})
		}
	}

	return user, nil
}
