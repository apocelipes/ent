// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/apocelipes/ent/dialect/sql/sqlgraph"
	"github.com/apocelipes/ent/examples/edgeindex/ent/city"
	"github.com/apocelipes/ent/examples/edgeindex/ent/street"
	"github.com/apocelipes/ent/schema/field"
)

// CityCreate is the builder for creating a City entity.
type CityCreate struct {
	config
	mutation *CityMutation
	hooks    []Hook
}

// SetName sets the name field.
func (cc *CityCreate) SetName(s string) *CityCreate {
	cc.mutation.SetName(s)
	return cc
}

// AddStreetIDs adds the streets edge to Street by ids.
func (cc *CityCreate) AddStreetIDs(ids ...int) *CityCreate {
	cc.mutation.AddStreetIDs(ids...)
	return cc
}

// AddStreets adds the streets edges to Street.
func (cc *CityCreate) AddStreets(s ...*Street) *CityCreate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return cc.AddStreetIDs(ids...)
}

// Mutation returns the CityMutation object of the builder.
func (cc *CityCreate) Mutation() *CityMutation {
	return cc.mutation
}

// Save creates the City in the database.
func (cc *CityCreate) Save(ctx context.Context) (*City, error) {
	if err := cc.preSave(); err != nil {
		return nil, err
	}
	var (
		err  error
		node *City
	)
	if len(cc.hooks) == 0 {
		node, err = cc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CityMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cc.mutation = mutation
			node, err = cc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(cc.hooks) - 1; i >= 0; i-- {
			mut = cc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (cc *CityCreate) SaveX(ctx context.Context) *City {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (cc *CityCreate) preSave() error {
	if _, ok := cc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New("ent: missing required field \"name\"")}
	}
	return nil
}

func (cc *CityCreate) sqlSave(ctx context.Context) (*City, error) {
	c, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	c.ID = int(id)
	return c, nil
}

func (cc *CityCreate) createSpec() (*City, *sqlgraph.CreateSpec) {
	var (
		c     = &City{config: cc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: city.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: city.FieldID,
			},
		}
	)
	if value, ok := cc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: city.FieldName,
		})
		c.Name = value
	}
	if nodes := cc.mutation.StreetsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   city.StreetsTable,
			Columns: []string{city.StreetsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: street.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return c, _spec
}

// CityCreateBulk is the builder for creating a bulk of City entities.
type CityCreateBulk struct {
	config
	builders []*CityCreate
}

// Save creates the City entities in the database.
func (ccb *CityCreateBulk) Save(ctx context.Context) ([]*City, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*City, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				if err := builder.preSave(); err != nil {
					return nil, err
				}
				mutation, ok := m.(*CityMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX calls Save and panics if Save returns an error.
func (ccb *CityCreateBulk) SaveX(ctx context.Context) []*City {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
