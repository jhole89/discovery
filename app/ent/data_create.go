// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebook/ent/dialect/gremlin"
	"github.com/facebook/ent/dialect/gremlin/graph/dsl"
	"github.com/facebook/ent/dialect/gremlin/graph/dsl/g"
	"github.com/jhole89/orbital/ent/data"
)

// DataCreate is the builder for creating a Data entity.
type DataCreate struct {
	config
	mutation *DataMutation
	hooks    []Hook
}

// SetName sets the name field.
func (dc *DataCreate) SetName(s string) *DataCreate {
	dc.mutation.SetName(s)
	return dc
}

// SetNillableName sets the name field if the given value is not nil.
func (dc *DataCreate) SetNillableName(s *string) *DataCreate {
	if s != nil {
		dc.SetName(*s)
	}
	return dc
}

// SetContext sets the context field.
func (dc *DataCreate) SetContext(s string) *DataCreate {
	dc.mutation.SetContext(s)
	return dc
}

// SetNillableContext sets the context field if the given value is not nil.
func (dc *DataCreate) SetNillableContext(s *string) *DataCreate {
	if s != nil {
		dc.SetContext(*s)
	}
	return dc
}

// AddHasTableIDs adds the has_table edge to Data by ids.
func (dc *DataCreate) AddHasTableIDs(ids ...int) *DataCreate {
	dc.mutation.AddHasTableIDs(ids...)
	return dc
}

// AddHasTable adds the has_table edges to Data.
func (dc *DataCreate) AddHasTable(d ...*Data) *DataCreate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return dc.AddHasTableIDs(ids...)
}

// AddHasFieldIDs adds the has_field edge to Data by ids.
func (dc *DataCreate) AddHasFieldIDs(ids ...int) *DataCreate {
	dc.mutation.AddHasFieldIDs(ids...)
	return dc
}

// AddHasField adds the has_field edges to Data.
func (dc *DataCreate) AddHasField(d ...*Data) *DataCreate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return dc.AddHasFieldIDs(ids...)
}

// Mutation returns the DataMutation object of the builder.
func (dc *DataCreate) Mutation() *DataMutation {
	return dc.mutation
}

// Save creates the Data in the database.
func (dc *DataCreate) Save(ctx context.Context) (*Data, error) {
	var (
		err  error
		node *Data
	)
	dc.defaults()
	if len(dc.hooks) == 0 {
		if err = dc.check(); err != nil {
			return nil, err
		}
		node, err = dc.gremlinSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DataMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = dc.check(); err != nil {
				return nil, err
			}
			dc.mutation = mutation
			node, err = dc.gremlinSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(dc.hooks) - 1; i >= 0; i-- {
			mut = dc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, dc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (dc *DataCreate) SaveX(ctx context.Context) *Data {
	v, err := dc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (dc *DataCreate) defaults() {
	if _, ok := dc.mutation.Name(); !ok {
		v := data.DefaultName
		dc.mutation.SetName(v)
	}
	if _, ok := dc.mutation.Context(); !ok {
		v := data.DefaultContext
		dc.mutation.SetContext(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (dc *DataCreate) check() error {
	if _, ok := dc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New("ent: missing required field \"name\"")}
	}
	if _, ok := dc.mutation.Context(); !ok {
		return &ValidationError{Name: "context", err: errors.New("ent: missing required field \"context\"")}
	}
	return nil
}

func (dc *DataCreate) gremlinSave(ctx context.Context) (*Data, error) {
	res := &gremlin.Response{}
	query, bindings := dc.gremlin().Query()
	if err := dc.driver.Exec(ctx, query, bindings, res); err != nil {
		return nil, err
	}
	if err, ok := isConstantError(res); ok {
		return nil, err
	}
	d := &Data{config: dc.config}
	if err := d.FromResponse(res); err != nil {
		return nil, err
	}
	return d, nil
}

func (dc *DataCreate) gremlin() *dsl.Traversal {
	v := g.AddV(data.Label)
	if value, ok := dc.mutation.Name(); ok {
		v.Property(dsl.Single, data.FieldName, value)
	}
	if value, ok := dc.mutation.Context(); ok {
		v.Property(dsl.Single, data.FieldContext, value)
	}
	for _, id := range dc.mutation.HasTableIDs() {
		v.AddE(data.HasTableLabel).To(g.V(id)).OutV()
	}
	for _, id := range dc.mutation.HasFieldIDs() {
		v.AddE(data.HasFieldLabel).To(g.V(id)).OutV()
	}
	return v.ValueMap(true)
}

// DataCreateBulk is the builder for creating a bulk of Data entities.
type DataCreateBulk struct {
	config
	builders []*DataCreate
}