// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entc/integration/edgeschema/ent/role"
	"entgo.io/ent/entc/integration/edgeschema/ent/roleuser"
	"entgo.io/ent/entc/integration/edgeschema/ent/user"
	"entgo.io/ent/schema/field"
)

// RoleUserCreate is the builder for creating a RoleUser entity.
type RoleUserCreate struct {
	config
	mutation *RoleUserMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (ruc *RoleUserCreate) SetCreatedAt(t time.Time) *RoleUserCreate {
	ruc.mutation.SetCreatedAt(t)
	return ruc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ruc *RoleUserCreate) SetNillableCreatedAt(t *time.Time) *RoleUserCreate {
	if t != nil {
		ruc.SetCreatedAt(*t)
	}
	return ruc
}

// SetRoleID sets the "role_id" field.
func (ruc *RoleUserCreate) SetRoleID(i int) *RoleUserCreate {
	ruc.mutation.SetRoleID(i)
	return ruc
}

// SetUserID sets the "user_id" field.
func (ruc *RoleUserCreate) SetUserID(i int) *RoleUserCreate {
	ruc.mutation.SetUserID(i)
	return ruc
}

// SetRole sets the "role" edge to the Role entity.
func (ruc *RoleUserCreate) SetRole(r *Role) *RoleUserCreate {
	return ruc.SetRoleID(r.ID)
}

// SetUser sets the "user" edge to the User entity.
func (ruc *RoleUserCreate) SetUser(u *User) *RoleUserCreate {
	return ruc.SetUserID(u.ID)
}

// Mutation returns the RoleUserMutation object of the builder.
func (ruc *RoleUserCreate) Mutation() *RoleUserMutation {
	return ruc.mutation
}

// Save creates the RoleUser in the database.
func (ruc *RoleUserCreate) Save(ctx context.Context) (*RoleUser, error) {
	ruc.defaults()
	return withHooks(ctx, ruc.sqlSave, ruc.mutation, ruc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ruc *RoleUserCreate) SaveX(ctx context.Context) *RoleUser {
	v, err := ruc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ruc *RoleUserCreate) Exec(ctx context.Context) error {
	_, err := ruc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ruc *RoleUserCreate) ExecX(ctx context.Context) {
	if err := ruc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ruc *RoleUserCreate) defaults() {
	if _, ok := ruc.mutation.CreatedAt(); !ok {
		v := roleuser.DefaultCreatedAt()
		ruc.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ruc *RoleUserCreate) check() error {
	if _, ok := ruc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "RoleUser.created_at"`)}
	}
	if _, ok := ruc.mutation.RoleID(); !ok {
		return &ValidationError{Name: "role_id", err: errors.New(`ent: missing required field "RoleUser.role_id"`)}
	}
	if _, ok := ruc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "RoleUser.user_id"`)}
	}
	if _, ok := ruc.mutation.RoleID(); !ok {
		return &ValidationError{Name: "role", err: errors.New(`ent: missing required edge "RoleUser.role"`)}
	}
	if _, ok := ruc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user", err: errors.New(`ent: missing required edge "RoleUser.user"`)}
	}
	return nil
}

func (ruc *RoleUserCreate) sqlSave(ctx context.Context) (*RoleUser, error) {
	if err := ruc.check(); err != nil {
		return nil, err
	}
	_node, _spec := ruc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ruc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}

func (ruc *RoleUserCreate) createSpec() (*RoleUser, *sqlgraph.CreateSpec) {
	var (
		_node = &RoleUser{config: ruc.config}
		_spec = sqlgraph.NewCreateSpec(roleuser.Table, nil)
	)
	_spec.OnConflict = ruc.conflict
	if value, ok := ruc.mutation.CreatedAt(); ok {
		_spec.SetField(roleuser.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if nodes := ruc.mutation.RoleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   roleuser.RoleTable,
			Columns: []string{roleuser.RoleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(role.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.RoleID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ruc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   roleuser.UserTable,
			Columns: []string{roleuser.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.UserID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.RoleUser.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.RoleUserUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (ruc *RoleUserCreate) OnConflict(opts ...sql.ConflictOption) *RoleUserUpsertOne {
	ruc.conflict = opts
	return &RoleUserUpsertOne{
		create: ruc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.RoleUser.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ruc *RoleUserCreate) OnConflictColumns(columns ...string) *RoleUserUpsertOne {
	ruc.conflict = append(ruc.conflict, sql.ConflictColumns(columns...))
	return &RoleUserUpsertOne{
		create: ruc,
	}
}

type (
	// RoleUserUpsertOne is the builder for "upsert"-ing
	//  one RoleUser node.
	RoleUserUpsertOne struct {
		create *RoleUserCreate
	}

	// RoleUserUpsert is the "OnConflict" setter.
	RoleUserUpsert struct {
		*sql.UpdateSet
	}
)

// SetCreatedAt sets the "created_at" field.
func (u *RoleUserUpsert) SetCreatedAt(v time.Time) *RoleUserUpsert {
	u.Set(roleuser.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *RoleUserUpsert) UpdateCreatedAt() *RoleUserUpsert {
	u.SetExcluded(roleuser.FieldCreatedAt)
	return u
}

// SetRoleID sets the "role_id" field.
func (u *RoleUserUpsert) SetRoleID(v int) *RoleUserUpsert {
	u.Set(roleuser.FieldRoleID, v)
	return u
}

// UpdateRoleID sets the "role_id" field to the value that was provided on create.
func (u *RoleUserUpsert) UpdateRoleID() *RoleUserUpsert {
	u.SetExcluded(roleuser.FieldRoleID)
	return u
}

// SetUserID sets the "user_id" field.
func (u *RoleUserUpsert) SetUserID(v int) *RoleUserUpsert {
	u.Set(roleuser.FieldUserID, v)
	return u
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *RoleUserUpsert) UpdateUserID() *RoleUserUpsert {
	u.SetExcluded(roleuser.FieldUserID)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.RoleUser.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *RoleUserUpsertOne) UpdateNewValues() *RoleUserUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.RoleUser.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *RoleUserUpsertOne) Ignore() *RoleUserUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *RoleUserUpsertOne) DoNothing() *RoleUserUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the RoleUserCreate.OnConflict
// documentation for more info.
func (u *RoleUserUpsertOne) Update(set func(*RoleUserUpsert)) *RoleUserUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&RoleUserUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *RoleUserUpsertOne) SetCreatedAt(v time.Time) *RoleUserUpsertOne {
	return u.Update(func(s *RoleUserUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *RoleUserUpsertOne) UpdateCreatedAt() *RoleUserUpsertOne {
	return u.Update(func(s *RoleUserUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetRoleID sets the "role_id" field.
func (u *RoleUserUpsertOne) SetRoleID(v int) *RoleUserUpsertOne {
	return u.Update(func(s *RoleUserUpsert) {
		s.SetRoleID(v)
	})
}

// UpdateRoleID sets the "role_id" field to the value that was provided on create.
func (u *RoleUserUpsertOne) UpdateRoleID() *RoleUserUpsertOne {
	return u.Update(func(s *RoleUserUpsert) {
		s.UpdateRoleID()
	})
}

// SetUserID sets the "user_id" field.
func (u *RoleUserUpsertOne) SetUserID(v int) *RoleUserUpsertOne {
	return u.Update(func(s *RoleUserUpsert) {
		s.SetUserID(v)
	})
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *RoleUserUpsertOne) UpdateUserID() *RoleUserUpsertOne {
	return u.Update(func(s *RoleUserUpsert) {
		s.UpdateUserID()
	})
}

// Exec executes the query.
func (u *RoleUserUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for RoleUserCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *RoleUserUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// RoleUserCreateBulk is the builder for creating many RoleUser entities in bulk.
type RoleUserCreateBulk struct {
	config
	err      error
	builders []*RoleUserCreate
	conflict []sql.ConflictOption
}

// Save creates the RoleUser entities in the database.
func (rucb *RoleUserCreateBulk) Save(ctx context.Context) ([]*RoleUser, error) {
	if rucb.err != nil {
		return nil, rucb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(rucb.builders))
	nodes := make([]*RoleUser, len(rucb.builders))
	mutators := make([]Mutator, len(rucb.builders))
	for i := range rucb.builders {
		func(i int, root context.Context) {
			builder := rucb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*RoleUserMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, rucb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = rucb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, rucb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, rucb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (rucb *RoleUserCreateBulk) SaveX(ctx context.Context) []*RoleUser {
	v, err := rucb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rucb *RoleUserCreateBulk) Exec(ctx context.Context) error {
	_, err := rucb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rucb *RoleUserCreateBulk) ExecX(ctx context.Context) {
	if err := rucb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.RoleUser.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.RoleUserUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (rucb *RoleUserCreateBulk) OnConflict(opts ...sql.ConflictOption) *RoleUserUpsertBulk {
	rucb.conflict = opts
	return &RoleUserUpsertBulk{
		create: rucb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.RoleUser.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (rucb *RoleUserCreateBulk) OnConflictColumns(columns ...string) *RoleUserUpsertBulk {
	rucb.conflict = append(rucb.conflict, sql.ConflictColumns(columns...))
	return &RoleUserUpsertBulk{
		create: rucb,
	}
}

// RoleUserUpsertBulk is the builder for "upsert"-ing
// a bulk of RoleUser nodes.
type RoleUserUpsertBulk struct {
	create *RoleUserCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.RoleUser.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *RoleUserUpsertBulk) UpdateNewValues() *RoleUserUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.RoleUser.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *RoleUserUpsertBulk) Ignore() *RoleUserUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *RoleUserUpsertBulk) DoNothing() *RoleUserUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the RoleUserCreateBulk.OnConflict
// documentation for more info.
func (u *RoleUserUpsertBulk) Update(set func(*RoleUserUpsert)) *RoleUserUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&RoleUserUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *RoleUserUpsertBulk) SetCreatedAt(v time.Time) *RoleUserUpsertBulk {
	return u.Update(func(s *RoleUserUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *RoleUserUpsertBulk) UpdateCreatedAt() *RoleUserUpsertBulk {
	return u.Update(func(s *RoleUserUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetRoleID sets the "role_id" field.
func (u *RoleUserUpsertBulk) SetRoleID(v int) *RoleUserUpsertBulk {
	return u.Update(func(s *RoleUserUpsert) {
		s.SetRoleID(v)
	})
}

// UpdateRoleID sets the "role_id" field to the value that was provided on create.
func (u *RoleUserUpsertBulk) UpdateRoleID() *RoleUserUpsertBulk {
	return u.Update(func(s *RoleUserUpsert) {
		s.UpdateRoleID()
	})
}

// SetUserID sets the "user_id" field.
func (u *RoleUserUpsertBulk) SetUserID(v int) *RoleUserUpsertBulk {
	return u.Update(func(s *RoleUserUpsert) {
		s.SetUserID(v)
	})
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *RoleUserUpsertBulk) UpdateUserID() *RoleUserUpsertBulk {
	return u.Update(func(s *RoleUserUpsert) {
		s.UpdateUserID()
	})
}

// Exec executes the query.
func (u *RoleUserUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the RoleUserCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for RoleUserCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *RoleUserUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
