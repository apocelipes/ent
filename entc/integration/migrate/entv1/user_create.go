// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package entv1

import (
	"context"
	"errors"
	"fmt"

	"github.com/apocelipes/ent/dialect/sql/sqlgraph"
	"github.com/apocelipes/ent/entc/integration/migrate/entv1/car"
	"github.com/apocelipes/ent/entc/integration/migrate/entv1/user"
	"github.com/apocelipes/ent/schema/field"
)

// UserCreate is the builder for creating a User entity.
type UserCreate struct {
	config
	mutation *UserMutation
	hooks    []Hook
}

// SetAge sets the age field.
func (uc *UserCreate) SetAge(i int32) *UserCreate {
	uc.mutation.SetAge(i)
	return uc
}

// SetName sets the name field.
func (uc *UserCreate) SetName(s string) *UserCreate {
	uc.mutation.SetName(s)
	return uc
}

// SetNickname sets the nickname field.
func (uc *UserCreate) SetNickname(s string) *UserCreate {
	uc.mutation.SetNickname(s)
	return uc
}

// SetAddress sets the address field.
func (uc *UserCreate) SetAddress(s string) *UserCreate {
	uc.mutation.SetAddress(s)
	return uc
}

// SetNillableAddress sets the address field if the given value is not nil.
func (uc *UserCreate) SetNillableAddress(s *string) *UserCreate {
	if s != nil {
		uc.SetAddress(*s)
	}
	return uc
}

// SetRenamed sets the renamed field.
func (uc *UserCreate) SetRenamed(s string) *UserCreate {
	uc.mutation.SetRenamed(s)
	return uc
}

// SetNillableRenamed sets the renamed field if the given value is not nil.
func (uc *UserCreate) SetNillableRenamed(s *string) *UserCreate {
	if s != nil {
		uc.SetRenamed(*s)
	}
	return uc
}

// SetBlob sets the blob field.
func (uc *UserCreate) SetBlob(b []byte) *UserCreate {
	uc.mutation.SetBlob(b)
	return uc
}

// SetState sets the state field.
func (uc *UserCreate) SetState(u user.State) *UserCreate {
	uc.mutation.SetState(u)
	return uc
}

// SetNillableState sets the state field if the given value is not nil.
func (uc *UserCreate) SetNillableState(u *user.State) *UserCreate {
	if u != nil {
		uc.SetState(*u)
	}
	return uc
}

// SetStatus sets the status field.
func (uc *UserCreate) SetStatus(s string) *UserCreate {
	uc.mutation.SetStatus(s)
	return uc
}

// SetNillableStatus sets the status field if the given value is not nil.
func (uc *UserCreate) SetNillableStatus(s *string) *UserCreate {
	if s != nil {
		uc.SetStatus(*s)
	}
	return uc
}

// SetID sets the id field.
func (uc *UserCreate) SetID(i int) *UserCreate {
	uc.mutation.SetID(i)
	return uc
}

// SetParentID sets the parent edge to User by id.
func (uc *UserCreate) SetParentID(id int) *UserCreate {
	uc.mutation.SetParentID(id)
	return uc
}

// SetNillableParentID sets the parent edge to User by id if the given value is not nil.
func (uc *UserCreate) SetNillableParentID(id *int) *UserCreate {
	if id != nil {
		uc = uc.SetParentID(*id)
	}
	return uc
}

// SetParent sets the parent edge to User.
func (uc *UserCreate) SetParent(u *User) *UserCreate {
	return uc.SetParentID(u.ID)
}

// AddChildIDs adds the children edge to User by ids.
func (uc *UserCreate) AddChildIDs(ids ...int) *UserCreate {
	uc.mutation.AddChildIDs(ids...)
	return uc
}

// AddChildren adds the children edges to User.
func (uc *UserCreate) AddChildren(u ...*User) *UserCreate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uc.AddChildIDs(ids...)
}

// SetSpouseID sets the spouse edge to User by id.
func (uc *UserCreate) SetSpouseID(id int) *UserCreate {
	uc.mutation.SetSpouseID(id)
	return uc
}

// SetNillableSpouseID sets the spouse edge to User by id if the given value is not nil.
func (uc *UserCreate) SetNillableSpouseID(id *int) *UserCreate {
	if id != nil {
		uc = uc.SetSpouseID(*id)
	}
	return uc
}

// SetSpouse sets the spouse edge to User.
func (uc *UserCreate) SetSpouse(u *User) *UserCreate {
	return uc.SetSpouseID(u.ID)
}

// SetCarID sets the car edge to Car by id.
func (uc *UserCreate) SetCarID(id int) *UserCreate {
	uc.mutation.SetCarID(id)
	return uc
}

// SetNillableCarID sets the car edge to Car by id if the given value is not nil.
func (uc *UserCreate) SetNillableCarID(id *int) *UserCreate {
	if id != nil {
		uc = uc.SetCarID(*id)
	}
	return uc
}

// SetCar sets the car edge to Car.
func (uc *UserCreate) SetCar(c *Car) *UserCreate {
	return uc.SetCarID(c.ID)
}

// Mutation returns the UserMutation object of the builder.
func (uc *UserCreate) Mutation() *UserMutation {
	return uc.mutation
}

// Save creates the User in the database.
func (uc *UserCreate) Save(ctx context.Context) (*User, error) {
	if err := uc.preSave(); err != nil {
		return nil, err
	}
	var (
		err  error
		node *User
	)
	if len(uc.hooks) == 0 {
		node, err = uc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			uc.mutation = mutation
			node, err = uc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(uc.hooks) - 1; i >= 0; i-- {
			mut = uc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (uc *UserCreate) SaveX(ctx context.Context) *User {
	v, err := uc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (uc *UserCreate) preSave() error {
	if _, ok := uc.mutation.Age(); !ok {
		return &ValidationError{Name: "age", err: errors.New("entv1: missing required field \"age\"")}
	}
	if _, ok := uc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New("entv1: missing required field \"name\"")}
	}
	if v, ok := uc.mutation.Name(); ok {
		if err := user.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("entv1: validator failed for field \"name\": %w", err)}
		}
	}
	if _, ok := uc.mutation.Nickname(); !ok {
		return &ValidationError{Name: "nickname", err: errors.New("entv1: missing required field \"nickname\"")}
	}
	if v, ok := uc.mutation.State(); ok {
		if err := user.StateValidator(v); err != nil {
			return &ValidationError{Name: "state", err: fmt.Errorf("entv1: validator failed for field \"state\": %w", err)}
		}
	}
	return nil
}

func (uc *UserCreate) sqlSave(ctx context.Context) (*User, error) {
	u, _spec := uc.createSpec()
	if err := sqlgraph.CreateNode(ctx, uc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	if u.ID == 0 {
		id := _spec.ID.Value.(int64)
		u.ID = int(id)
	}
	return u, nil
}

func (uc *UserCreate) createSpec() (*User, *sqlgraph.CreateSpec) {
	var (
		u     = &User{config: uc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: user.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: user.FieldID,
			},
		}
	)
	if id, ok := uc.mutation.ID(); ok {
		u.ID = id
		_spec.ID.Value = id
	}
	if value, ok := uc.mutation.Age(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: user.FieldAge,
		})
		u.Age = value
	}
	if value, ok := uc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldName,
		})
		u.Name = value
	}
	if value, ok := uc.mutation.Nickname(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldNickname,
		})
		u.Nickname = value
	}
	if value, ok := uc.mutation.Address(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldAddress,
		})
		u.Address = value
	}
	if value, ok := uc.mutation.Renamed(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldRenamed,
		})
		u.Renamed = value
	}
	if value, ok := uc.mutation.Blob(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBytes,
			Value:  value,
			Column: user.FieldBlob,
		})
		u.Blob = value
	}
	if value, ok := uc.mutation.State(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: user.FieldState,
		})
		u.State = value
	}
	if value, ok := uc.mutation.Status(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldStatus,
		})
		u.Status = value
	}
	if nodes := uc.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   user.ParentTable,
			Columns: []string{user.ParentColumn},
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
	if nodes := uc.mutation.ChildrenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ChildrenTable,
			Columns: []string{user.ChildrenColumn},
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
	if nodes := uc.mutation.SpouseIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.SpouseTable,
			Columns: []string{user.SpouseColumn},
			Bidi:    true,
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
	if nodes := uc.mutation.CarIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.CarTable,
			Columns: []string{user.CarColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: car.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return u, _spec
}

// UserCreateBulk is the builder for creating a bulk of User entities.
type UserCreateBulk struct {
	config
	builders []*UserCreate
}

// Save creates the User entities in the database.
func (ucb *UserCreateBulk) Save(ctx context.Context) ([]*User, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ucb.builders))
	nodes := make([]*User, len(ucb.builders))
	mutators := make([]Mutator, len(ucb.builders))
	for i := range ucb.builders {
		func(i int, root context.Context) {
			builder := ucb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				if err := builder.preSave(); err != nil {
					return nil, err
				}
				mutation, ok := m.(*UserMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ucb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ucb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				if nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ucb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX calls Save and panics if Save returns an error.
func (ucb *UserCreateBulk) SaveX(ctx context.Context) []*User {
	v, err := ucb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
