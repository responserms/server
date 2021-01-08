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
	"github.com/responserms/server/ent/user"
)

// SessionCreate is the builder for creating a Session entity.
type SessionCreate struct {
	config
	mutation *SessionMutation
	hooks    []Hook
}

// SetCreateTime sets the create_time field.
func (sc *SessionCreate) SetCreateTime(t time.Time) *SessionCreate {
	sc.mutation.SetCreateTime(t)
	return sc
}

// SetNillableCreateTime sets the create_time field if the given value is not nil.
func (sc *SessionCreate) SetNillableCreateTime(t *time.Time) *SessionCreate {
	if t != nil {
		sc.SetCreateTime(*t)
	}
	return sc
}

// SetUpdateTime sets the update_time field.
func (sc *SessionCreate) SetUpdateTime(t time.Time) *SessionCreate {
	sc.mutation.SetUpdateTime(t)
	return sc
}

// SetNillableUpdateTime sets the update_time field if the given value is not nil.
func (sc *SessionCreate) SetNillableUpdateTime(t *time.Time) *SessionCreate {
	if t != nil {
		sc.SetUpdateTime(*t)
	}
	return sc
}

// SetIPAddress sets the ip_address field.
func (sc *SessionCreate) SetIPAddress(s string) *SessionCreate {
	sc.mutation.SetIPAddress(s)
	return sc
}

// SetBrowserName sets the browser_name field.
func (sc *SessionCreate) SetBrowserName(s string) *SessionCreate {
	sc.mutation.SetBrowserName(s)
	return sc
}

// SetBrowserVersion sets the browser_version field.
func (sc *SessionCreate) SetBrowserVersion(s string) *SessionCreate {
	sc.mutation.SetBrowserVersion(s)
	return sc
}

// SetDeviceOs sets the device_os field.
func (sc *SessionCreate) SetDeviceOs(s string) *SessionCreate {
	sc.mutation.SetDeviceOs(s)
	return sc
}

// SetDeviceType sets the device_type field.
func (sc *SessionCreate) SetDeviceType(s string) *SessionCreate {
	sc.mutation.SetDeviceType(s)
	return sc
}

// SetClaims sets the claims field.
func (sc *SessionCreate) SetClaims(s string) *SessionCreate {
	sc.mutation.SetClaims(s)
	return sc
}

// SetTerminatedAt sets the terminated_at field.
func (sc *SessionCreate) SetTerminatedAt(t time.Time) *SessionCreate {
	sc.mutation.SetTerminatedAt(t)
	return sc
}

// SetTokenID sets the token edge to Token by id.
func (sc *SessionCreate) SetTokenID(id int) *SessionCreate {
	sc.mutation.SetTokenID(id)
	return sc
}

// SetNillableTokenID sets the token edge to Token by id if the given value is not nil.
func (sc *SessionCreate) SetNillableTokenID(id *int) *SessionCreate {
	if id != nil {
		sc = sc.SetTokenID(*id)
	}
	return sc
}

// SetToken sets the token edge to Token.
func (sc *SessionCreate) SetToken(t *Token) *SessionCreate {
	return sc.SetTokenID(t.ID)
}

// SetUserID sets the user edge to User by id.
func (sc *SessionCreate) SetUserID(id int) *SessionCreate {
	sc.mutation.SetUserID(id)
	return sc
}

// SetNillableUserID sets the user edge to User by id if the given value is not nil.
func (sc *SessionCreate) SetNillableUserID(id *int) *SessionCreate {
	if id != nil {
		sc = sc.SetUserID(*id)
	}
	return sc
}

// SetUser sets the user edge to User.
func (sc *SessionCreate) SetUser(u *User) *SessionCreate {
	return sc.SetUserID(u.ID)
}

// Mutation returns the SessionMutation object of the builder.
func (sc *SessionCreate) Mutation() *SessionMutation {
	return sc.mutation
}

// Save creates the Session in the database.
func (sc *SessionCreate) Save(ctx context.Context) (*Session, error) {
	var (
		err  error
		node *Session
	)
	sc.defaults()
	if len(sc.hooks) == 0 {
		if err = sc.check(); err != nil {
			return nil, err
		}
		node, err = sc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SessionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = sc.check(); err != nil {
				return nil, err
			}
			sc.mutation = mutation
			node, err = sc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(sc.hooks) - 1; i >= 0; i-- {
			mut = sc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, sc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (sc *SessionCreate) SaveX(ctx context.Context) *Session {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (sc *SessionCreate) defaults() {
	if _, ok := sc.mutation.CreateTime(); !ok {
		v := session.DefaultCreateTime()
		sc.mutation.SetCreateTime(v)
	}
	if _, ok := sc.mutation.UpdateTime(); !ok {
		v := session.DefaultUpdateTime()
		sc.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sc *SessionCreate) check() error {
	if _, ok := sc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New("ent: missing required field \"create_time\"")}
	}
	if _, ok := sc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New("ent: missing required field \"update_time\"")}
	}
	if _, ok := sc.mutation.IPAddress(); !ok {
		return &ValidationError{Name: "ip_address", err: errors.New("ent: missing required field \"ip_address\"")}
	}
	if _, ok := sc.mutation.BrowserName(); !ok {
		return &ValidationError{Name: "browser_name", err: errors.New("ent: missing required field \"browser_name\"")}
	}
	if _, ok := sc.mutation.BrowserVersion(); !ok {
		return &ValidationError{Name: "browser_version", err: errors.New("ent: missing required field \"browser_version\"")}
	}
	if _, ok := sc.mutation.DeviceOs(); !ok {
		return &ValidationError{Name: "device_os", err: errors.New("ent: missing required field \"device_os\"")}
	}
	if _, ok := sc.mutation.DeviceType(); !ok {
		return &ValidationError{Name: "device_type", err: errors.New("ent: missing required field \"device_type\"")}
	}
	if _, ok := sc.mutation.Claims(); !ok {
		return &ValidationError{Name: "claims", err: errors.New("ent: missing required field \"claims\"")}
	}
	if _, ok := sc.mutation.TerminatedAt(); !ok {
		return &ValidationError{Name: "terminated_at", err: errors.New("ent: missing required field \"terminated_at\"")}
	}
	return nil
}

func (sc *SessionCreate) sqlSave(ctx context.Context) (*Session, error) {
	_node, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (sc *SessionCreate) createSpec() (*Session, *sqlgraph.CreateSpec) {
	var (
		_node = &Session{config: sc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: session.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: session.FieldID,
			},
		}
	)
	if value, ok := sc.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: session.FieldCreateTime,
		})
		_node.CreateTime = value
	}
	if value, ok := sc.mutation.UpdateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: session.FieldUpdateTime,
		})
		_node.UpdateTime = value
	}
	if value, ok := sc.mutation.IPAddress(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: session.FieldIPAddress,
		})
		_node.IPAddress = value
	}
	if value, ok := sc.mutation.BrowserName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: session.FieldBrowserName,
		})
		_node.BrowserName = value
	}
	if value, ok := sc.mutation.BrowserVersion(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: session.FieldBrowserVersion,
		})
		_node.BrowserVersion = value
	}
	if value, ok := sc.mutation.DeviceOs(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: session.FieldDeviceOs,
		})
		_node.DeviceOs = value
	}
	if value, ok := sc.mutation.DeviceType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: session.FieldDeviceType,
		})
		_node.DeviceType = value
	}
	if value, ok := sc.mutation.Claims(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: session.FieldClaims,
		})
		_node.Claims = value
	}
	if value, ok := sc.mutation.TerminatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: session.FieldTerminatedAt,
		})
		_node.TerminatedAt = value
	}
	if nodes := sc.mutation.TokenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   session.TokenTable,
			Columns: []string{session.TokenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: token.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   session.UserTable,
			Columns: []string{session.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
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

// SessionCreateBulk is the builder for creating a bulk of Session entities.
type SessionCreateBulk struct {
	config
	builders []*SessionCreate
}

// Save creates the Session entities in the database.
func (scb *SessionCreateBulk) Save(ctx context.Context) ([]*Session, error) {
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Session, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SessionMutation)
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
					_, err = mutators[i+1].Mutate(root, scb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, scb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, scb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX calls Save and panics if Save returns an error.
func (scb *SessionCreateBulk) SaveX(ctx context.Context) []*Session {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
