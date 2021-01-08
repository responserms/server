package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/99designs/gqlgen/graphql"
	"github.com/responserms/server/graphql/server"
	"github.com/responserms/server/graphql/types"
	"github.com/responserms/server/internal/core"
	"github.com/responserms/server/internal/reqdata"
)

func (r *mutationResolver) LoginWithCredentials(ctx context.Context, input types.LoginWithCredentialsInput) (*types.LoginWithCredentialsPayload, error) {
	payload := &types.LoginWithCredentialsPayload{
		Errors: make([]types.LoginWithCredentialsError, 0),
	}

	token, err := r.core.Auth.TokenFromCredentials(ctx, &core.TokenFromCredentialsOptions{
		Email:                 input.Credentials.Email,
		Password:              input.Credentials.Password,
		BrowserName:           input.Browser.Name,
		BrowserVersion:        input.Browser.Version,
		DeviceType:            input.Device.Type,
		DeviceOperatingSystem: input.Device.OperatingSystem,
		IPAddress:             reqdata.IPAddressFromContext(ctx),
	})

	if err != nil {
		switch err {
		case core.ErrAuthInvalidCredentials:
			payload.Errors = append(payload.Errors, newInvalidCredentialsError())
		case core.ErrAuthPasswordNotSet:
			payload.Errors = append(payload.Errors, newPasswordIsNotAllowed())
		default:
			graphql.AddError(ctx, err)
		}
	}

	if token != nil {
		str, err := token.Token.String()
		if err != nil {
			return payload, fmt.Errorf("Auth.LoginWithCredentials: %w", err)
		}

		payload.Token = &types.LoginTokenPayload{
			AccessToken: &str,
			ExpiredAt:   &token.ExpiredAt,
		}
	}

	return payload, nil
}

func (r *mutationResolver) LoginWithProvider(ctx context.Context, input types.LoginWithProviderInput) (*types.LoginWithProviderPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns server.MutationResolver implementation.
func (r *Resolver) Mutation() server.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func newInvalidCredentialsError() types.InvalidCredentialsError {
	return types.InvalidCredentialsError{
		Message: "Your credentials are invalid",
		Path:    []string{"credentials"},
	}
}
func newPasswordIsNotAllowed() types.PasswordNotAllowedError {
	return types.PasswordNotAllowedError{
		Message: "You are not allowed to login with these credentials",
		Path:    []string{"credentials"},
	}
}
