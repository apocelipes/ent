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
	"entgo.io/ent/entc/integration/ent/goods"
	"entgo.io/ent/entc/integration/ent/predicate"
	"entgo.io/ent/schema/field"
)

// GoodsUpdate is the builder for updating Goods entities.
type GoodsUpdate struct {
	config
	hooks     []Hook
	mutation  *GoodsMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the GoodsUpdate builder.
func (gu *GoodsUpdate) Where(ps ...predicate.Goods) *GoodsUpdate {
	gu.mutation.Where(ps...)
	return gu
}

// Mutation returns the GoodsMutation object of the builder.
func (gu *GoodsUpdate) Mutation() *GoodsMutation {
	return gu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (gu *GoodsUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, GoodsMutation](ctx, gu.sqlSave, gu.mutation, gu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (gu *GoodsUpdate) SaveX(ctx context.Context) int {
	affected, err := gu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (gu *GoodsUpdate) Exec(ctx context.Context) error {
	_, err := gu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gu *GoodsUpdate) ExecX(ctx context.Context) {
	if err := gu.Exec(ctx); err != nil {
		panic(err)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (gu *GoodsUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *GoodsUpdate {
	gu.modifiers = append(gu.modifiers, modifiers...)
	return gu
}

func (gu *GoodsUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   goods.Table,
			Columns: goods.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: goods.FieldID,
			},
		},
	}
	if ps := gu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	_spec.AddModifiers(gu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, gu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{goods.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	gu.mutation.done = true
	return n, nil
}

// GoodsUpdateOne is the builder for updating a single Goods entity.
type GoodsUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *GoodsMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Mutation returns the GoodsMutation object of the builder.
func (guo *GoodsUpdateOne) Mutation() *GoodsMutation {
	return guo.mutation
}

// Where appends a list predicates to the GoodsUpdate builder.
func (guo *GoodsUpdateOne) Where(ps ...predicate.Goods) *GoodsUpdateOne {
	guo.mutation.Where(ps...)
	return guo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (guo *GoodsUpdateOne) Select(field string, fields ...string) *GoodsUpdateOne {
	guo.fields = append([]string{field}, fields...)
	return guo
}

// Save executes the query and returns the updated Goods entity.
func (guo *GoodsUpdateOne) Save(ctx context.Context) (*Goods, error) {
	return withHooks[*Goods, GoodsMutation](ctx, guo.sqlSave, guo.mutation, guo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (guo *GoodsUpdateOne) SaveX(ctx context.Context) *Goods {
	node, err := guo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (guo *GoodsUpdateOne) Exec(ctx context.Context) error {
	_, err := guo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (guo *GoodsUpdateOne) ExecX(ctx context.Context) {
	if err := guo.Exec(ctx); err != nil {
		panic(err)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (guo *GoodsUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *GoodsUpdateOne {
	guo.modifiers = append(guo.modifiers, modifiers...)
	return guo
}

func (guo *GoodsUpdateOne) sqlSave(ctx context.Context) (_node *Goods, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   goods.Table,
			Columns: goods.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: goods.FieldID,
			},
		},
	}
	id, ok := guo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Goods.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := guo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, goods.FieldID)
		for _, f := range fields {
			if !goods.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != goods.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := guo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	_spec.AddModifiers(guo.modifiers...)
	_node = &Goods{config: guo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, guo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{goods.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	guo.mutation.done = true
	return _node, nil
}
