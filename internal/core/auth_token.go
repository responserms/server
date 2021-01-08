package core

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/responserms/server/ent"
)

// AccessTokenType represents the type of object an AccessToken is issued to.
type AccessTokenType int

// CreateTokenOptions configures the CreateToken
type CreateTokenOptions struct {
	SessionID     int
	UserID        int
	TokenDuration time.Duration
	Claims        *TokenClaims
}

// ResponseTokenClaims contains the Response-specific claims stored in JWT's.
type ResponseTokenClaims struct {
	Name string
}

// TokenClaims are the standard and Respnose-specific claims stored in JWT's.
type TokenClaims struct {
	jwt.StandardClaims
	Response ResponseTokenClaims
}

const (
	// UserAccessTokenType represents an AccessToken issued to a User.
	UserAccessTokenType AccessTokenType = iota + 1

	// BotAccessTokenType represents an AccessToken issued to a Bot.
	BotAccessTokenType
)

func (a *authService) CreateToken(ctx context.Context, opts *CreateTokenOptions) (*ent.Token, error) {
	create := a.core.svcs.Database().
		Token.
		Create()

	claims, err := json.Marshal(opts.Claims)
	if err != nil {
		return nil, fmt.Errorf("Sessions.CreateSession: %w", err)
	}

	token, err := create.SetSessionID(opts.SessionID).
		SetExpiredAt(time.Now().Add(opts.TokenDuration)).
		SetClaims(string(claims)).
		Save(ctx)

	if err != nil {
		return nil, fmt.Errorf("Auth.CreateToken: %w", err)
	}

	return token, nil
}

// AccessToken is an access token for User and Bot access.
type AccessToken struct {
	Type AccessTokenType
	Data *ent.Token
	jwt  *jwt.Token
}

type createUserTokenOptions struct {
	Session       *ent.Session
	User          *ent.User
	TokenDuration time.Duration
}

func (a *authService) createUserToken(ctx context.Context, opts *createUserTokenOptions) (*AccessToken, error) {
	claims := &TokenClaims{
		StandardClaims: jwt.StandardClaims{
			Subject:   strconv.Itoa(opts.User.ID),
			ExpiresAt: time.Now().Add(opts.TokenDuration).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		Response: ResponseTokenClaims{
			Name: opts.User.Name,
		},
	}

	token, err := a.CreateToken(ctx, &CreateTokenOptions{
		SessionID:     opts.Session.ID,
		UserID:        opts.User.ID,
		TokenDuration: opts.TokenDuration,
		Claims:        claims,
	})

	if err != nil {
		return nil, fmt.Errorf("Auth.createUserToken: %w", err)
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	if err != nil {
		return nil, fmt.Errorf("Auth.createUserToken: %w", err)
	}

	t := &AccessToken{
		Type: UserAccessTokenType,
		Data: token,
		jwt:  jwtToken,
	}

	return t, nil
}

func (a *AccessToken) String() (string, error) {
	str, err := a.jwt.SignedString([]byte{})
	if err != nil {
		return "", fmt.Errorf("AccessToken.String: %w", err)
	}

	return str, nil
}
