// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package card

import (
	"time"

	"github.com/apocelipes/ent"
)

const (
	// Label holds the string label denoting the card type in the database.
	Label = "card"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldNumber holds the string denoting the number field in the database.
	FieldNumber = "number"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"

	// EdgeOwner holds the string denoting the owner edge name in mutations.
	EdgeOwner = "owner"

	// Table holds the table name of the card in the database.
	Table = "cards"
	// OwnerTable is the table the holds the owner relation/edge.
	OwnerTable = "cards"
	// OwnerInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	OwnerInverseTable = "users"
	// OwnerColumn is the table column denoting the owner relation/edge.
	OwnerColumn = "user_cards"
)

// Columns holds all SQL columns for card fields.
var Columns = []string{
	FieldID,
	FieldNumber,
	FieldName,
	FieldCreatedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Card type.
var ForeignKeys = []string{
	"user_cards",
}

// Note that the variables below are initialized by the runtime
// package on the initialization of the application. Therefore,
// it should be imported in the main as follows:
//
//	import _ "github.com/apocelipes/ent/entc/integration/hooks/ent/runtime"
//
var (
	Hooks [3]ent.Hook
	// DefaultNumber holds the default value on creation for the number field.
	DefaultNumber string
	// NumberValidator is a validator for the "number" field. It is called by the builders before save.
	NumberValidator func(string) error
	// DefaultCreatedAt holds the default value on creation for the created_at field.
	DefaultCreatedAt func() time.Time
)
