// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/apocelipes/ent/dialect/sql/sqlgraph"
	"github.com/apocelipes/ent/entc/integration/ent/card"
	"github.com/apocelipes/ent/entc/integration/ent/spec"
	"github.com/apocelipes/ent/schema/field"
)

// SpecCreate is the builder for creating a Spec entity.
type SpecCreate struct {
	config
	mutation *SpecMutation
	hooks    []Hook
}

// AddCardIDs adds the card edge to Card by ids.
func (sc *SpecCreate) AddCardIDs(ids ...int) *SpecCreate {
	sc.mutation.AddCardIDs(ids...)
	return sc
}

// AddCard adds the card edges to Card.
func (sc *SpecCreate) AddCard(c ...*Card) *SpecCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return sc.AddCardIDs(ids...)
}

// Mutation returns the SpecMutation object of the builder.
func (sc *SpecCreate) Mutation() *SpecMutation {
	return sc.mutation
}

// Save creates the Spec in the database.
func (sc *SpecCreate) Save(ctx context.Context) (*Spec, error) {
	if err := sc.preSave(); err != nil {
		return nil, err
	}
	var (
		err  error
		node *Spec
	)
	if len(sc.hooks) == 0 {
		node, err = sc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SpecMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
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
func (sc *SpecCreate) SaveX(ctx context.Context) *Spec {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (sc *SpecCreate) preSave() error {
	return nil
}

func (sc *SpecCreate) sqlSave(ctx context.Context) (*Spec, error) {
	s, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	s.ID = int(id)
	return s, nil
}

func (sc *SpecCreate) createSpec() (*Spec, *sqlgraph.CreateSpec) {
	var (
		s     = &Spec{config: sc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: spec.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: spec.FieldID,
			},
		}
	)
	if nodes := sc.mutation.CardIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   spec.CardTable,
			Columns: spec.CardPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: card.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return s, _spec
}

// SpecCreateBulk is the builder for creating a bulk of Spec entities.
type SpecCreateBulk struct {
	config
	builders []*SpecCreate
}

// Save creates the Spec entities in the database.
func (scb *SpecCreateBulk) Save(ctx context.Context) ([]*Spec, error) {
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Spec, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				if err := builder.preSave(); err != nil {
					return nil, err
				}
				mutation, ok := m.(*SpecMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
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
func (scb *SpecCreateBulk) SaveX(ctx context.Context) []*Spec {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
