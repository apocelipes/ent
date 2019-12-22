// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/entc/integration/template/ent/pet"
	"github.com/facebookincubator/ent/entc/integration/template/ent/predicate"
	"github.com/facebookincubator/ent/entc/integration/template/ent/user"
	"github.com/facebookincubator/ent/schema/field"
)

// PetUpdate is the builder for updating Pet entities.
type PetUpdate struct {
	config
	age              *int
	addage           *int
	licensed_at      *time.Time
	clearlicensed_at bool
	owner            map[int]struct{}
	clearedOwner     bool
	predicates       []predicate.Pet
}

// Where adds a new predicate for the builder.
func (pu *PetUpdate) Where(ps ...predicate.Pet) *PetUpdate {
	pu.predicates = append(pu.predicates, ps...)
	return pu
}

// SetAge sets the age field.
func (pu *PetUpdate) SetAge(i int) *PetUpdate {
	pu.age = &i
	pu.addage = nil
	return pu
}

// AddAge adds i to age.
func (pu *PetUpdate) AddAge(i int) *PetUpdate {
	if pu.addage == nil {
		pu.addage = &i
	} else {
		*pu.addage += i
	}
	return pu
}

// SetLicensedAt sets the licensed_at field.
func (pu *PetUpdate) SetLicensedAt(t time.Time) *PetUpdate {
	pu.licensed_at = &t
	return pu
}

// SetNillableLicensedAt sets the licensed_at field if the given value is not nil.
func (pu *PetUpdate) SetNillableLicensedAt(t *time.Time) *PetUpdate {
	if t != nil {
		pu.SetLicensedAt(*t)
	}
	return pu
}

// ClearLicensedAt clears the value of licensed_at.
func (pu *PetUpdate) ClearLicensedAt() *PetUpdate {
	pu.licensed_at = nil
	pu.clearlicensed_at = true
	return pu
}

// SetOwnerID sets the owner edge to User by id.
func (pu *PetUpdate) SetOwnerID(id int) *PetUpdate {
	if pu.owner == nil {
		pu.owner = make(map[int]struct{})
	}
	pu.owner[id] = struct{}{}
	return pu
}

// SetNillableOwnerID sets the owner edge to User by id if the given value is not nil.
func (pu *PetUpdate) SetNillableOwnerID(id *int) *PetUpdate {
	if id != nil {
		pu = pu.SetOwnerID(*id)
	}
	return pu
}

// SetOwner sets the owner edge to User.
func (pu *PetUpdate) SetOwner(u *User) *PetUpdate {
	return pu.SetOwnerID(u.ID)
}

// ClearOwner clears the owner edge to User.
func (pu *PetUpdate) ClearOwner() *PetUpdate {
	pu.clearedOwner = true
	return pu
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (pu *PetUpdate) Save(ctx context.Context) (int, error) {
	if len(pu.owner) > 1 {
		return 0, errors.New("ent: multiple assignments on a unique edge \"owner\"")
	}
	return pu.sqlSave(ctx)
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PetUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PetUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PetUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (pu *PetUpdate) sqlSave(ctx context.Context) (n int, err error) {
	spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   pet.Table,
			Columns: pet.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: pet.FieldID,
			},
		},
	}
	if ps := pu.predicates; len(ps) > 0 {
		spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value := pu.age; value != nil {
		spec.Fields.Set = append(spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  *value,
			Column: pet.FieldAge,
		})
	}
	if value := pu.addage; value != nil {
		spec.Fields.Add = append(spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  *value,
			Column: pet.FieldAge,
		})
	}
	if value := pu.licensed_at; value != nil {
		spec.Fields.Set = append(spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  *value,
			Column: pet.FieldLicensedAt,
		})
	}
	if pu.clearlicensed_at {
		spec.Fields.Clear = append(spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: pet.FieldLicensedAt,
		})
	}
	if pu.clearedOwner {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   pet.OwnerTable,
			Columns: []string{pet.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		spec.Edges.Clear = append(spec.Edges.Clear, edge)
	}
	if nodes := pu.owner; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   pet.OwnerTable,
			Columns: []string{pet.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		spec.Edges.Add = append(spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// PetUpdateOne is the builder for updating a single Pet entity.
type PetUpdateOne struct {
	config
	id               int
	age              *int
	addage           *int
	licensed_at      *time.Time
	clearlicensed_at bool
	owner            map[int]struct{}
	clearedOwner     bool
}

// SetAge sets the age field.
func (puo *PetUpdateOne) SetAge(i int) *PetUpdateOne {
	puo.age = &i
	puo.addage = nil
	return puo
}

// AddAge adds i to age.
func (puo *PetUpdateOne) AddAge(i int) *PetUpdateOne {
	if puo.addage == nil {
		puo.addage = &i
	} else {
		*puo.addage += i
	}
	return puo
}

// SetLicensedAt sets the licensed_at field.
func (puo *PetUpdateOne) SetLicensedAt(t time.Time) *PetUpdateOne {
	puo.licensed_at = &t
	return puo
}

// SetNillableLicensedAt sets the licensed_at field if the given value is not nil.
func (puo *PetUpdateOne) SetNillableLicensedAt(t *time.Time) *PetUpdateOne {
	if t != nil {
		puo.SetLicensedAt(*t)
	}
	return puo
}

// ClearLicensedAt clears the value of licensed_at.
func (puo *PetUpdateOne) ClearLicensedAt() *PetUpdateOne {
	puo.licensed_at = nil
	puo.clearlicensed_at = true
	return puo
}

// SetOwnerID sets the owner edge to User by id.
func (puo *PetUpdateOne) SetOwnerID(id int) *PetUpdateOne {
	if puo.owner == nil {
		puo.owner = make(map[int]struct{})
	}
	puo.owner[id] = struct{}{}
	return puo
}

// SetNillableOwnerID sets the owner edge to User by id if the given value is not nil.
func (puo *PetUpdateOne) SetNillableOwnerID(id *int) *PetUpdateOne {
	if id != nil {
		puo = puo.SetOwnerID(*id)
	}
	return puo
}

// SetOwner sets the owner edge to User.
func (puo *PetUpdateOne) SetOwner(u *User) *PetUpdateOne {
	return puo.SetOwnerID(u.ID)
}

// ClearOwner clears the owner edge to User.
func (puo *PetUpdateOne) ClearOwner() *PetUpdateOne {
	puo.clearedOwner = true
	return puo
}

// Save executes the query and returns the updated entity.
func (puo *PetUpdateOne) Save(ctx context.Context) (*Pet, error) {
	if len(puo.owner) > 1 {
		return nil, errors.New("ent: multiple assignments on a unique edge \"owner\"")
	}
	return puo.sqlSave(ctx)
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PetUpdateOne) SaveX(ctx context.Context) *Pet {
	pe, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return pe
}

// Exec executes the query on the entity.
func (puo *PetUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PetUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (puo *PetUpdateOne) sqlSave(ctx context.Context) (pe *Pet, err error) {
	spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   pet.Table,
			Columns: pet.Columns,
			ID: &sqlgraph.FieldSpec{
				Value:  puo.id,
				Type:   field.TypeInt,
				Column: pet.FieldID,
			},
		},
	}
	if value := puo.age; value != nil {
		spec.Fields.Set = append(spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  *value,
			Column: pet.FieldAge,
		})
	}
	if value := puo.addage; value != nil {
		spec.Fields.Add = append(spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  *value,
			Column: pet.FieldAge,
		})
	}
	if value := puo.licensed_at; value != nil {
		spec.Fields.Set = append(spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  *value,
			Column: pet.FieldLicensedAt,
		})
	}
	if puo.clearlicensed_at {
		spec.Fields.Clear = append(spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: pet.FieldLicensedAt,
		})
	}
	if puo.clearedOwner {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   pet.OwnerTable,
			Columns: []string{pet.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		spec.Edges.Clear = append(spec.Edges.Clear, edge)
	}
	if nodes := puo.owner; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   pet.OwnerTable,
			Columns: []string{pet.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		spec.Edges.Add = append(spec.Edges.Add, edge)
	}
	pe = &Pet{config: puo.config}
	spec.Assign = pe.assignValues
	spec.ScanTypes = pe.scanValues()
	if err = sqlgraph.UpdateNode(ctx, puo.driver, spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return pe, nil
}
