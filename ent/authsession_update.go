// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/responserms/server/ent/authsession"
	"github.com/responserms/server/ent/predicate"
)

// AuthSessionUpdate is the builder for updating AuthSession entities.
type AuthSessionUpdate struct {
	config
	hooks    []Hook
	mutation *AuthSessionMutation
}

// Where adds a new predicate for the builder.
func (asu *AuthSessionUpdate) Where(ps ...predicate.AuthSession) *AuthSessionUpdate {
	asu.mutation.predicates = append(asu.mutation.predicates, ps...)
	return asu
}

// Mutation returns the AuthSessionMutation object of the builder.
func (asu *AuthSessionUpdate) Mutation() *AuthSessionMutation {
	return asu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (asu *AuthSessionUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(asu.hooks) == 0 {
		affected, err = asu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AuthSessionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			asu.mutation = mutation
			affected, err = asu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(asu.hooks) - 1; i >= 0; i-- {
			mut = asu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, asu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (asu *AuthSessionUpdate) SaveX(ctx context.Context) int {
	affected, err := asu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (asu *AuthSessionUpdate) Exec(ctx context.Context) error {
	_, err := asu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (asu *AuthSessionUpdate) ExecX(ctx context.Context) {
	if err := asu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (asu *AuthSessionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   authsession.Table,
			Columns: authsession.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: authsession.FieldID,
			},
		},
	}
	if ps := asu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if n, err = sqlgraph.UpdateNodes(ctx, asu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{authsession.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// AuthSessionUpdateOne is the builder for updating a single AuthSession entity.
type AuthSessionUpdateOne struct {
	config
	hooks    []Hook
	mutation *AuthSessionMutation
}

// Mutation returns the AuthSessionMutation object of the builder.
func (asuo *AuthSessionUpdateOne) Mutation() *AuthSessionMutation {
	return asuo.mutation
}

// Save executes the query and returns the updated entity.
func (asuo *AuthSessionUpdateOne) Save(ctx context.Context) (*AuthSession, error) {
	var (
		err  error
		node *AuthSession
	)
	if len(asuo.hooks) == 0 {
		node, err = asuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AuthSessionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			asuo.mutation = mutation
			node, err = asuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(asuo.hooks) - 1; i >= 0; i-- {
			mut = asuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, asuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (asuo *AuthSessionUpdateOne) SaveX(ctx context.Context) *AuthSession {
	node, err := asuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (asuo *AuthSessionUpdateOne) Exec(ctx context.Context) error {
	_, err := asuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (asuo *AuthSessionUpdateOne) ExecX(ctx context.Context) {
	if err := asuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (asuo *AuthSessionUpdateOne) sqlSave(ctx context.Context) (_node *AuthSession, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   authsession.Table,
			Columns: authsession.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: authsession.FieldID,
			},
		},
	}
	id, ok := asuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing AuthSession.ID for update")}
	}
	_spec.Node.ID.Value = id
	_node = &AuthSession{config: asuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, asuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{authsession.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
