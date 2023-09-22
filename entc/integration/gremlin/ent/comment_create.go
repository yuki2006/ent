// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"

	"entgo.io/ent/dialect/gremlin"
	"entgo.io/ent/dialect/gremlin/graph/dsl"
	"entgo.io/ent/dialect/gremlin/graph/dsl/__"
	"entgo.io/ent/dialect/gremlin/graph/dsl/g"
	"entgo.io/ent/dialect/gremlin/graph/dsl/p"
	schemadir "entgo.io/ent/entc/integration/ent/schema/dir"
	"entgo.io/ent/entc/integration/gremlin/ent/comment"
)

// CommentCreate is the builder for creating a Comment entity.
type CommentCreate struct {
	config
	mutation *CommentMutation
	hooks    []Hook
}

// SetUniqueInt sets the "unique_int" field.
func (cc *CommentCreate) SetUniqueInt(i int) *CommentCreate {
	cc.mutation.SetUniqueInt(i)
	return cc
}

// SetUniqueFloat sets the "unique_float" field.
func (cc *CommentCreate) SetUniqueFloat(f float64) *CommentCreate {
	cc.mutation.SetUniqueFloat(f)
	return cc
}

// SetNillableInt sets the "nillable_int" field.
func (cc *CommentCreate) SetNillableInt(i int) *CommentCreate {
	cc.mutation.SetNillableInt(i)
	return cc
}

// SetNillableNillableInt sets the "nillable_int" field if the given value is not nil.
func (cc *CommentCreate) SetNillableNillableInt(i *int) *CommentCreate {
	if i != nil {
		cc.SetNillableInt(*i)
	}
	return cc
}

// SetTable sets the "table" field.
func (cc *CommentCreate) SetTable(s string) *CommentCreate {
	cc.mutation.SetTable(s)
	return cc
}

// SetNillableTable sets the "table" field if the given value is not nil.
func (cc *CommentCreate) SetNillableTable(s *string) *CommentCreate {
	if s != nil {
		cc.SetTable(*s)
	}
	return cc
}

// SetDir sets the "dir" field.
func (cc *CommentCreate) SetDir(s schemadir.Dir) *CommentCreate {
	cc.mutation.SetDir(s)
	return cc
}

// SetNillableDir sets the "dir" field if the given value is not nil.
func (cc *CommentCreate) SetNillableDir(s *schemadir.Dir) *CommentCreate {
	if s != nil {
		cc.SetDir(*s)
	}
	return cc
}

// SetClient sets the "client" field.
func (cc *CommentCreate) SetClient(s string) *CommentCreate {
	cc.mutation.SetClient(s)
	return cc
}

// SetNillableClient sets the "client" field if the given value is not nil.
func (cc *CommentCreate) SetNillableClient(s *string) *CommentCreate {
	if s != nil {
		cc.SetClient(*s)
	}
	return cc
}

// Mutation returns the CommentMutation object of the builder.
func (cc *CommentCreate) Mutation() *CommentMutation {
	return cc.mutation
}

// Save creates the Comment in the database.
func (cc *CommentCreate) Save(ctx context.Context) (*Comment, error) {
	return withHooks(ctx, cc.gremlinSave, cc.mutation, cc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (cc *CommentCreate) SaveX(ctx context.Context) *Comment {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *CommentCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *CommentCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cc *CommentCreate) check() error {
	if _, ok := cc.mutation.UniqueInt(); !ok {
		return &ValidationError{Name: "unique_int", err: errors.New(`ent: missing required field "Comment.unique_int"`)}
	}
	if _, ok := cc.mutation.UniqueFloat(); !ok {
		return &ValidationError{Name: "unique_float", err: errors.New(`ent: missing required field "Comment.unique_float"`)}
	}
	return nil
}

func (cc *CommentCreate) gremlinSave(ctx context.Context) (*Comment, error) {
	if err := cc.check(); err != nil {
		return nil, err
	}
	res := &gremlin.Response{}
	query, bindings := cc.gremlin().Query()
	if err := cc.driver.Exec(ctx, query, bindings, res); err != nil {
		return nil, err
	}
	if err, ok := isConstantError(res); ok {
		return nil, err
	}
	rnode := &Comment{config: cc.config}
	if err := rnode.FromResponse(res); err != nil {
		return nil, err
	}
	cc.mutation.id = &rnode.ID
	cc.mutation.done = true
	return rnode, nil
}

func (cc *CommentCreate) gremlin() *dsl.Traversal {
	type constraint struct {
		pred *dsl.Traversal // constraint predicate.
		test *dsl.Traversal // test matches and its constant.
	}
	constraints := make([]*constraint, 0, 2)
	v := g.AddV(comment.Label)
	if value, ok := cc.mutation.UniqueInt(); ok {
		constraints = append(constraints, &constraint{
			pred: g.V().Has(comment.Label, comment.FieldUniqueInt, value).Count(),
			test: __.Is(p.NEQ(0)).Constant(NewErrUniqueField(comment.Label, comment.FieldUniqueInt, value)),
		})
		v.Property(dsl.Single, comment.FieldUniqueInt, value)
	}
	if value, ok := cc.mutation.UniqueFloat(); ok {
		constraints = append(constraints, &constraint{
			pred: g.V().Has(comment.Label, comment.FieldUniqueFloat, value).Count(),
			test: __.Is(p.NEQ(0)).Constant(NewErrUniqueField(comment.Label, comment.FieldUniqueFloat, value)),
		})
		v.Property(dsl.Single, comment.FieldUniqueFloat, value)
	}
	if value, ok := cc.mutation.NillableInt(); ok {
		v.Property(dsl.Single, comment.FieldNillableInt, value)
	}
	if value, ok := cc.mutation.Table(); ok {
		v.Property(dsl.Single, comment.FieldTable, value)
	}
	if value, ok := cc.mutation.Dir(); ok {
		v.Property(dsl.Single, comment.FieldDir, value)
	}
	if value, ok := cc.mutation.GetClient(); ok {
		v.Property(dsl.Single, comment.FieldClient, value)
	}
	if len(constraints) == 0 {
		return v.ValueMap(true)
	}
	tr := constraints[0].pred.Coalesce(constraints[0].test, v.ValueMap(true))
	for _, cr := range constraints[1:] {
		tr = cr.pred.Coalesce(cr.test, tr)
	}
	return tr
}

// CommentCreateBulk is the builder for creating many Comment entities in bulk.
type CommentCreateBulk struct {
	config
	err      error
	builders []*CommentCreate
}
