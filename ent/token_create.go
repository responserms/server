// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/responserms/server/ent/session"
	"github.com/responserms/server/ent/token"
)

// TokenCreate is the builder for creating a Token entity.
type TokenCreate struct {
	config
	mutation *TokenMutation
	hooks    []Hook
}

// SetExpiredAt sets the expired_at field.
func (tc *TokenCreate) SetExpiredAt(t time.Time) *TokenCreate {
	tc.mutation.SetExpiredAt(t)
	return tc
}

// SetBlockedAt sets the blocked_at field.
func (tc *TokenCreate) SetBlockedAt(t time.Time) *TokenCreate {
	tc.mutation.SetBlockedAt(t)
	return tc
}

// SetNillableBlockedAt sets the blocked_at field if the given value is not nil.
func (tc *TokenCreate) SetNillableBlockedAt(t *time.Time) *TokenCreate {
	if t != nil {
		tc.SetBlockedAt(*t)
	}
	return tc
}

// SetClaims sets the claims field.
func (tc *TokenCreate) SetClaims(s string) *TokenCreate {
	tc.mutation.SetClaims(s)
	return tc
}

// SetSessionID sets the session edge to Session by id.
func (tc *TokenCreate) SetSessionID(id int) *TokenCreate {
	tc.mutation.SetSessionID(id)
	return tc
}

// SetNillableSessionID sets the session edge to Session by id if the given value is not nil.
func (tc *TokenCreate) SetNillableSessionID(id *int) *TokenCreate {
	if id != nil {
		tc = tc.SetSessionID(*id)
	}
	return tc
}

// SetSession sets the session edge to Session.
func (tc *TokenCreate) SetSession(s *Session) *TokenCreate {
	return tc.SetSessionID(s.ID)
}

// Mutation returns the TokenMutation object of the builder.
func (tc *TokenCreate) Mutation() *TokenMutation {
	return tc.mutation
}

// Save creates the Token in the database.
func (tc *TokenCreate) Save(ctx context.Context) (*Token, error) {
	var (
		err  error
		node *Token
	)
	if len(tc.hooks) == 0 {
		if err = tc.check(); err != nil {
			return nil, err
		}
		node, err = tc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TokenMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tc.check(); err != nil {
				return nil, err
			}
			tc.mutation = mutation
			node, err = tc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(tc.hooks) - 1; i >= 0; i-- {
			mut = tc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (tc *TokenCreate) SaveX(ctx context.Context) *Token {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// check runs all checks and user-defined validators on the builder.
func (tc *TokenCreate) check() error {
	if _, ok := tc.mutation.ExpiredAt(); !ok {
		return &ValidationError{Name: "expired_at", err: errors.New("ent: missing required field \"expired_at\"")}
	}
	if _, ok := tc.mutation.Claims(); !ok {
		return &ValidationError{Name: "claims", err: errors.New("ent: missing required field \"claims\"")}
	}
	return nil
}

func (tc *TokenCreate) sqlSave(ctx context.Context) (*Token, error) {
	_node, _spec := tc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (tc *TokenCreate) createSpec() (*Token, *sqlgraph.CreateSpec) {
	var (
		_node = &Token{config: tc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: token.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: token.FieldID,
			},
		}
	)
	if value, ok := tc.mutation.ExpiredAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: token.FieldExpiredAt,
		})
		_node.ExpiredAt = value
	}
	if value, ok := tc.mutation.BlockedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: token.FieldBlockedAt,
		})
		_node.BlockedAt = &value
	}
	if value, ok := tc.mutation.Claims(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: token.FieldClaims,
		})
		_node.Claims = value
	}
	if nodes := tc.mutation.SessionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   token.SessionTable,
			Columns: []string{token.SessionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: session.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// TokenCreateBulk is the builder for creating a bulk of Token entities.
type TokenCreateBulk struct {
	config
	builders []*TokenCreate
}

// Save creates the Token entities in the database.
func (tcb *TokenCreateBulk) Save(ctx context.Context) ([]*Token, error) {
	specs := make([]*sqlgraph.CreateSpec, len(tcb.builders))
	nodes := make([]*Token, len(tcb.builders))
	mutators := make([]Mutator, len(tcb.builders))
	for i := range tcb.builders {
		func(i int, root context.Context) {
			builder := tcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TokenMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, tcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				id := specs[i].ID.Value.(int64)
				nodes[i].ID = int(id)
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, tcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX calls Save and panics if Save returns an error.
func (tcb *TokenCreateBulk) SaveX(ctx context.Context) []*Token {
	v, err := tcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
