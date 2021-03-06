// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/responserms/server/ent/maplayer"
	"github.com/responserms/server/ent/maptype"
	"github.com/responserms/server/ent/metadata"
	"github.com/responserms/server/ent/predicate"
)

// MapLayerUpdate is the builder for updating MapLayer entities.
type MapLayerUpdate struct {
	config
	hooks    []Hook
	mutation *MapLayerMutation
}

// Where adds a new predicate for the builder.
func (mlu *MapLayerUpdate) Where(ps ...predicate.MapLayer) *MapLayerUpdate {
	mlu.mutation.predicates = append(mlu.mutation.predicates, ps...)
	return mlu
}

// SetUpdatedAt sets the updated_at field.
func (mlu *MapLayerUpdate) SetUpdatedAt(t time.Time) *MapLayerUpdate {
	mlu.mutation.SetUpdatedAt(t)
	return mlu
}

// SetName sets the name field.
func (mlu *MapLayerUpdate) SetName(s string) *MapLayerUpdate {
	mlu.mutation.SetName(s)
	return mlu
}

// SetURLTemplate sets the url_template field.
func (mlu *MapLayerUpdate) SetURLTemplate(s string) *MapLayerUpdate {
	mlu.mutation.SetURLTemplate(s)
	return mlu
}

// SetIsPublic sets the is_public field.
func (mlu *MapLayerUpdate) SetIsPublic(b bool) *MapLayerUpdate {
	mlu.mutation.SetIsPublic(b)
	return mlu
}

// SetMetadataID sets the metadata edge to Metadata by id.
func (mlu *MapLayerUpdate) SetMetadataID(id int) *MapLayerUpdate {
	mlu.mutation.SetMetadataID(id)
	return mlu
}

// SetNillableMetadataID sets the metadata edge to Metadata by id if the given value is not nil.
func (mlu *MapLayerUpdate) SetNillableMetadataID(id *int) *MapLayerUpdate {
	if id != nil {
		mlu = mlu.SetMetadataID(*id)
	}
	return mlu
}

// SetMetadata sets the metadata edge to Metadata.
func (mlu *MapLayerUpdate) SetMetadata(m *Metadata) *MapLayerUpdate {
	return mlu.SetMetadataID(m.ID)
}

// SetMapTypeID sets the map_type edge to MapType by id.
func (mlu *MapLayerUpdate) SetMapTypeID(id int) *MapLayerUpdate {
	mlu.mutation.SetMapTypeID(id)
	return mlu
}

// SetNillableMapTypeID sets the map_type edge to MapType by id if the given value is not nil.
func (mlu *MapLayerUpdate) SetNillableMapTypeID(id *int) *MapLayerUpdate {
	if id != nil {
		mlu = mlu.SetMapTypeID(*id)
	}
	return mlu
}

// SetMapType sets the map_type edge to MapType.
func (mlu *MapLayerUpdate) SetMapType(m *MapType) *MapLayerUpdate {
	return mlu.SetMapTypeID(m.ID)
}

// Mutation returns the MapLayerMutation object of the builder.
func (mlu *MapLayerUpdate) Mutation() *MapLayerMutation {
	return mlu.mutation
}

// ClearMetadata clears the "metadata" edge to type Metadata.
func (mlu *MapLayerUpdate) ClearMetadata() *MapLayerUpdate {
	mlu.mutation.ClearMetadata()
	return mlu
}

// ClearMapType clears the "map_type" edge to type MapType.
func (mlu *MapLayerUpdate) ClearMapType() *MapLayerUpdate {
	mlu.mutation.ClearMapType()
	return mlu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (mlu *MapLayerUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	mlu.defaults()
	if len(mlu.hooks) == 0 {
		affected, err = mlu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MapLayerMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			mlu.mutation = mutation
			affected, err = mlu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(mlu.hooks) - 1; i >= 0; i-- {
			mut = mlu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mlu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (mlu *MapLayerUpdate) SaveX(ctx context.Context) int {
	affected, err := mlu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mlu *MapLayerUpdate) Exec(ctx context.Context) error {
	_, err := mlu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mlu *MapLayerUpdate) ExecX(ctx context.Context) {
	if err := mlu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mlu *MapLayerUpdate) defaults() {
	if _, ok := mlu.mutation.UpdatedAt(); !ok {
		v := maplayer.UpdateDefaultUpdatedAt()
		mlu.mutation.SetUpdatedAt(v)
	}
}

func (mlu *MapLayerUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   maplayer.Table,
			Columns: maplayer.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: maplayer.FieldID,
			},
		},
	}
	if ps := mlu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mlu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: maplayer.FieldUpdatedAt,
		})
	}
	if value, ok := mlu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: maplayer.FieldName,
		})
	}
	if value, ok := mlu.mutation.URLTemplate(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: maplayer.FieldURLTemplate,
		})
	}
	if value, ok := mlu.mutation.IsPublic(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: maplayer.FieldIsPublic,
		})
	}
	if mlu.mutation.MetadataCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   maplayer.MetadataTable,
			Columns: []string{maplayer.MetadataColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: metadata.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mlu.mutation.MetadataIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   maplayer.MetadataTable,
			Columns: []string{maplayer.MetadataColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: metadata.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if mlu.mutation.MapTypeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   maplayer.MapTypeTable,
			Columns: []string{maplayer.MapTypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: maptype.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mlu.mutation.MapTypeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   maplayer.MapTypeTable,
			Columns: []string{maplayer.MapTypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: maptype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, mlu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{maplayer.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// MapLayerUpdateOne is the builder for updating a single MapLayer entity.
type MapLayerUpdateOne struct {
	config
	hooks    []Hook
	mutation *MapLayerMutation
}

// SetUpdatedAt sets the updated_at field.
func (mluo *MapLayerUpdateOne) SetUpdatedAt(t time.Time) *MapLayerUpdateOne {
	mluo.mutation.SetUpdatedAt(t)
	return mluo
}

// SetName sets the name field.
func (mluo *MapLayerUpdateOne) SetName(s string) *MapLayerUpdateOne {
	mluo.mutation.SetName(s)
	return mluo
}

// SetURLTemplate sets the url_template field.
func (mluo *MapLayerUpdateOne) SetURLTemplate(s string) *MapLayerUpdateOne {
	mluo.mutation.SetURLTemplate(s)
	return mluo
}

// SetIsPublic sets the is_public field.
func (mluo *MapLayerUpdateOne) SetIsPublic(b bool) *MapLayerUpdateOne {
	mluo.mutation.SetIsPublic(b)
	return mluo
}

// SetMetadataID sets the metadata edge to Metadata by id.
func (mluo *MapLayerUpdateOne) SetMetadataID(id int) *MapLayerUpdateOne {
	mluo.mutation.SetMetadataID(id)
	return mluo
}

// SetNillableMetadataID sets the metadata edge to Metadata by id if the given value is not nil.
func (mluo *MapLayerUpdateOne) SetNillableMetadataID(id *int) *MapLayerUpdateOne {
	if id != nil {
		mluo = mluo.SetMetadataID(*id)
	}
	return mluo
}

// SetMetadata sets the metadata edge to Metadata.
func (mluo *MapLayerUpdateOne) SetMetadata(m *Metadata) *MapLayerUpdateOne {
	return mluo.SetMetadataID(m.ID)
}

// SetMapTypeID sets the map_type edge to MapType by id.
func (mluo *MapLayerUpdateOne) SetMapTypeID(id int) *MapLayerUpdateOne {
	mluo.mutation.SetMapTypeID(id)
	return mluo
}

// SetNillableMapTypeID sets the map_type edge to MapType by id if the given value is not nil.
func (mluo *MapLayerUpdateOne) SetNillableMapTypeID(id *int) *MapLayerUpdateOne {
	if id != nil {
		mluo = mluo.SetMapTypeID(*id)
	}
	return mluo
}

// SetMapType sets the map_type edge to MapType.
func (mluo *MapLayerUpdateOne) SetMapType(m *MapType) *MapLayerUpdateOne {
	return mluo.SetMapTypeID(m.ID)
}

// Mutation returns the MapLayerMutation object of the builder.
func (mluo *MapLayerUpdateOne) Mutation() *MapLayerMutation {
	return mluo.mutation
}

// ClearMetadata clears the "metadata" edge to type Metadata.
func (mluo *MapLayerUpdateOne) ClearMetadata() *MapLayerUpdateOne {
	mluo.mutation.ClearMetadata()
	return mluo
}

// ClearMapType clears the "map_type" edge to type MapType.
func (mluo *MapLayerUpdateOne) ClearMapType() *MapLayerUpdateOne {
	mluo.mutation.ClearMapType()
	return mluo
}

// Save executes the query and returns the updated entity.
func (mluo *MapLayerUpdateOne) Save(ctx context.Context) (*MapLayer, error) {
	var (
		err  error
		node *MapLayer
	)
	mluo.defaults()
	if len(mluo.hooks) == 0 {
		node, err = mluo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MapLayerMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			mluo.mutation = mutation
			node, err = mluo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(mluo.hooks) - 1; i >= 0; i-- {
			mut = mluo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mluo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (mluo *MapLayerUpdateOne) SaveX(ctx context.Context) *MapLayer {
	node, err := mluo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (mluo *MapLayerUpdateOne) Exec(ctx context.Context) error {
	_, err := mluo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mluo *MapLayerUpdateOne) ExecX(ctx context.Context) {
	if err := mluo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mluo *MapLayerUpdateOne) defaults() {
	if _, ok := mluo.mutation.UpdatedAt(); !ok {
		v := maplayer.UpdateDefaultUpdatedAt()
		mluo.mutation.SetUpdatedAt(v)
	}
}

func (mluo *MapLayerUpdateOne) sqlSave(ctx context.Context) (_node *MapLayer, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   maplayer.Table,
			Columns: maplayer.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: maplayer.FieldID,
			},
		},
	}
	id, ok := mluo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing MapLayer.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := mluo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: maplayer.FieldUpdatedAt,
		})
	}
	if value, ok := mluo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: maplayer.FieldName,
		})
	}
	if value, ok := mluo.mutation.URLTemplate(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: maplayer.FieldURLTemplate,
		})
	}
	if value, ok := mluo.mutation.IsPublic(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: maplayer.FieldIsPublic,
		})
	}
	if mluo.mutation.MetadataCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   maplayer.MetadataTable,
			Columns: []string{maplayer.MetadataColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: metadata.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mluo.mutation.MetadataIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   maplayer.MetadataTable,
			Columns: []string{maplayer.MetadataColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: metadata.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if mluo.mutation.MapTypeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   maplayer.MapTypeTable,
			Columns: []string{maplayer.MapTypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: maptype.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mluo.mutation.MapTypeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   maplayer.MapTypeTable,
			Columns: []string{maplayer.MapTypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: maptype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &MapLayer{config: mluo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, mluo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{maplayer.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
