// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package runtime

import (
	"context"

	"github.com/apocelipes/ent/entc/integration/privacy/ent/galaxy"
	"github.com/apocelipes/ent/entc/integration/privacy/ent/planet"
	"github.com/apocelipes/ent/entc/integration/privacy/ent/schema"

	"github.com/apocelipes/ent"
)

// The init function reads all schema descriptors with runtime
// code (default values, validators or hooks) and stitches it
// to their package variables.
func init() {
	galaxy.Policy = schema.Galaxy{}.Policy()
	galaxy.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := galaxy.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	galaxyFields := schema.Galaxy{}.Fields()
	_ = galaxyFields
	// galaxyDescName is the schema descriptor for name field.
	galaxyDescName := galaxyFields[0].Descriptor()
	// galaxy.NameValidator is a validator for the "name" field. It is called by the builders before save.
	galaxy.NameValidator = galaxyDescName.Validators[0].(func(string) error)
	planet.Policy = schema.Planet{}.Policy()
	planet.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := planet.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	planetHooks := schema.Planet{}.Hooks()

	planet.Hooks[1] = planetHooks[0]
	planetFields := schema.Planet{}.Fields()
	_ = planetFields
	// planetDescName is the schema descriptor for name field.
	planetDescName := planetFields[0].Descriptor()
	// planet.NameValidator is a validator for the "name" field. It is called by the builders before save.
	planet.NameValidator = planetDescName.Validators[0].(func(string) error)
}

const (
	Version = "(devel)" // Version of ent codegen.
)
