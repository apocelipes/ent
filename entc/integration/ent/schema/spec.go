// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"github.com/apocelipes/ent"
	"github.com/apocelipes/ent/schema/edge"
)

type Spec struct {
	ent.Schema
}

// Edges of the Spec.
func (Spec) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("card", Card.Type),
	}
}
