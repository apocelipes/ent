// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/entc/integration/customid/ent/other"
	"entgo.io/ent/entc/integration/customid/sid"
)

// Other is the model entity for the Other schema.
type Other struct {
	config
	// ID of the ent.
	ID sid.ID `json:"id,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Other) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case other.FieldID:
			values[i] = new(sid.ID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Other", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Other fields.
func (o *Other) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case other.FieldID:
			if value, ok := values[i].(*sid.ID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				o.ID = *value
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Other.
// Note that you need to call Other.Unwrap() before calling this method if this Other
// was returned from a transaction, and the transaction was committed or rolled back.
func (o *Other) Update() *OtherUpdateOne {
	return NewOtherClient(o.config).UpdateOne(o)
}

// Unwrap unwraps the Other entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (o *Other) Unwrap() *Other {
	_tx, ok := o.config.driver.(*txDriver)
	if !ok {
		panic("ent: Other is not a transactional entity")
	}
	o.config.driver = _tx.drv
	return o
}

// String implements the fmt.Stringer.
func (o *Other) String() string {
	var builder strings.Builder
	builder.WriteString("Other(")
	builder.WriteString(fmt.Sprintf("id=%v", o.ID))
	builder.WriteByte(')')
	return builder.String()
}

// Others is a parsable slice of Other.
type Others []*Other

func (o Others) config(cfg config) {
	for _i := range o {
		o[_i].config = cfg
	}
}
