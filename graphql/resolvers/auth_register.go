package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/99designs/gqlgen/graphql"
	"github.com/responserms/server/graphql/types"
	"github.com/responserms/server/internal/core"
)

func (r *mutationResolver) RegisterWithCredentials(ctx context.Context, input types.RegisterWithCredentialsInput) (*types.RegisterWithCredentialsPayload, error) {
	payload := &types.RegisterWithCredentialsPayload{
		Errors: make([]types.RegisterWithCredentialsError, 0),
	}

	_, err := r.core.Auth.Register(ctx, &core.RegisterOptions{
		Name:     input.Name,
		Email:    input.Credentials.Email,
		Password: input.Credentials.Password,
	})

	if err != nil {
		switch err {
		case core.ErrAuthEmailAlreadyRegistered:
			payload.Errors = append(payload.Errors, newEmailAlreadyExistsError())
		default:
			graphql.AddError(ctx, err)
		}

		return payload, nil
	}

	payload.Success = true

	return payload, nil
}

func (r *mutationResolver) RegisterWithProvider(ctx context.Context, input types.RegisterWithProviderInput) (*types.RegisterWithProviderPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func newEmailAlreadyExistsError() types.EmailAlreadyExistsError {
	return types.EmailAlreadyExistsError{
		Message: "The email provided is already registered",
		Path:    []string{"credentials", "email"},
	}
}
