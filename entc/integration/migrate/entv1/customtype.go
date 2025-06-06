// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package entv1

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/entc/integration/migrate/entv1/customtype"
)

// CustomType is the model entity for the CustomType schema.
type CustomType struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Custom holds the value of the "custom" field.
	Custom string `json:"custom,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*CustomType) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case customtype.FieldID:
			values[i] = new(sql.NullInt64)
		case customtype.FieldCustom:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type CustomType", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the CustomType fields.
func (ct *CustomType) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case customtype.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ct.ID = int(value.Int64)
		case customtype.FieldCustom:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field custom", values[i])
			} else if value.Valid {
				ct.Custom = value.String
			}
		}
	}
	return nil
}

// Update returns a builder for updating this CustomType.
// Note that you need to call CustomType.Unwrap() before calling this method if this CustomType
// was returned from a transaction, and the transaction was committed or rolled back.
func (ct *CustomType) Update() *CustomTypeUpdateOne {
	return NewCustomTypeClient(ct.config).UpdateOne(ct)
}

// Unwrap unwraps the CustomType entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ct *CustomType) Unwrap() *CustomType {
	_tx, ok := ct.config.driver.(*txDriver)
	if !ok {
		panic("entv1: CustomType is not a transactional entity")
	}
	ct.config.driver = _tx.drv
	return ct
}

// String implements the fmt.Stringer.
func (ct *CustomType) String() string {
	var builder strings.Builder
	builder.WriteString("CustomType(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ct.ID))
	builder.WriteString("custom=")
	builder.WriteString(ct.Custom)
	builder.WriteByte(')')
	return builder.String()
}

// CustomTypes is a parsable slice of CustomType.
type CustomTypes []*CustomType

func (ct CustomTypes) config(cfg config) {
	for _i := range ct {
		ct[_i].config = cfg
	}
}
