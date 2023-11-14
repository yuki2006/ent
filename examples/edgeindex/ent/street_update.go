// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/examples/edgeindex/ent/city"
	"entgo.io/ent/examples/edgeindex/ent/predicate"
	"entgo.io/ent/examples/edgeindex/ent/street"
	"entgo.io/ent/schema/field"
)

// StreetUpdate is the builder for updating Street entities.
type StreetUpdate struct {
	config
	hooks    []Hook
	mutation *StreetMutation
}

// Where appends a list predicates to the StreetUpdate builder.
func (su *StreetUpdate) Where(ps ...predicate.Street) *StreetUpdate {
	su.mutation.Where(ps...)
	return su
}

// SetName sets the "name" field.
func (su *StreetUpdate) SetName(s string) *StreetUpdate {
	su.mutation.SetName(s)
	return su
}

// SetNillableName sets the "name" field if the given value is not nil.
func (su *StreetUpdate) SetNillableName(s *string) *StreetUpdate {
	if s != nil {
		su.SetName(*s)
	}
	return su
}

// SetCityID sets the "city" edge to the City entity by ID.
func (su *StreetUpdate) SetCityID(id int) *StreetUpdate {
	su.mutation.SetCityID(id)
	return su
}

// SetNillableCityID sets the "city" edge to the City entity by ID if the given value is not nil.
func (su *StreetUpdate) SetNillableCityID(id *int) *StreetUpdate {
	if id != nil {
		su = su.SetCityID(*id)
	}
	return su
}

// SetCity sets the "city" edge to the City entity.
func (su *StreetUpdate) SetCity(c *City) *StreetUpdate {
	return su.SetCityID(c.ID)
}

// Mutation returns the StreetMutation object of the builder.
func (su *StreetUpdate) Mutation() *StreetMutation {
	return su.mutation
}

// ClearCity clears the "city" edge to the City entity.
func (su *StreetUpdate) ClearCity() *StreetUpdate {
	su.mutation.ClearCity()
	return su
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (su *StreetUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, su.sqlSave, su.mutation, su.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (su *StreetUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *StreetUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *StreetUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

func (su *StreetUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(street.Table, street.Columns, sqlgraph.NewFieldSpec(street.FieldID, field.TypeInt))
	if ps := su.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.Name(); ok {
		_spec.SetField(street.FieldName, field.TypeString, value)
	}
	if su.mutation.CityCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   street.CityTable,
			Columns: []string{street.CityColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(city.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.CityIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   street.CityTable,
			Columns: []string{street.CityColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(city.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{street.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	su.mutation.done = true
	return n, nil
}

// StreetUpdateOne is the builder for updating a single Street entity.
type StreetUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *StreetMutation
}

// SetName sets the "name" field.
func (suo *StreetUpdateOne) SetName(s string) *StreetUpdateOne {
	suo.mutation.SetName(s)
	return suo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (suo *StreetUpdateOne) SetNillableName(s *string) *StreetUpdateOne {
	if s != nil {
		suo.SetName(*s)
	}
	return suo
}

// SetCityID sets the "city" edge to the City entity by ID.
func (suo *StreetUpdateOne) SetCityID(id int) *StreetUpdateOne {
	suo.mutation.SetCityID(id)
	return suo
}

// SetNillableCityID sets the "city" edge to the City entity by ID if the given value is not nil.
func (suo *StreetUpdateOne) SetNillableCityID(id *int) *StreetUpdateOne {
	if id != nil {
		suo = suo.SetCityID(*id)
	}
	return suo
}

// SetCity sets the "city" edge to the City entity.
func (suo *StreetUpdateOne) SetCity(c *City) *StreetUpdateOne {
	return suo.SetCityID(c.ID)
}

// Mutation returns the StreetMutation object of the builder.
func (suo *StreetUpdateOne) Mutation() *StreetMutation {
	return suo.mutation
}

// ClearCity clears the "city" edge to the City entity.
func (suo *StreetUpdateOne) ClearCity() *StreetUpdateOne {
	suo.mutation.ClearCity()
	return suo
}

// Where appends a list predicates to the StreetUpdate builder.
func (suo *StreetUpdateOne) Where(ps ...predicate.Street) *StreetUpdateOne {
	suo.mutation.Where(ps...)
	return suo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (suo *StreetUpdateOne) Select(field string, fields ...string) *StreetUpdateOne {
	suo.fields = append([]string{field}, fields...)
	return suo
}

// Save executes the query and returns the updated Street entity.
func (suo *StreetUpdateOne) Save(ctx context.Context) (*Street, error) {
	return withHooks(ctx, suo.sqlSave, suo.mutation, suo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (suo *StreetUpdateOne) SaveX(ctx context.Context) *Street {
	node, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suo *StreetUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *StreetUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (suo *StreetUpdateOne) sqlSave(ctx context.Context) (_node *Street, err error) {
	_spec := sqlgraph.NewUpdateSpec(street.Table, street.Columns, sqlgraph.NewFieldSpec(street.FieldID, field.TypeInt))
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Street.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := suo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, street.FieldID)
		for _, f := range fields {
			if !street.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != street.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := suo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := suo.mutation.Name(); ok {
		_spec.SetField(street.FieldName, field.TypeString, value)
	}
	if suo.mutation.CityCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   street.CityTable,
			Columns: []string{street.CityColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(city.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.CityIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   street.CityTable,
			Columns: []string{street.CityColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(city.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Street{config: suo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{street.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	suo.mutation.done = true
	return _node, nil
}
