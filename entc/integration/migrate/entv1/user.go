// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package entv1

import (
	"fmt"
	"strings"

	"github.com/apocelipes/ent/dialect/sql"
	"github.com/apocelipes/ent/entc/integration/migrate/entv1/car"
	"github.com/apocelipes/ent/entc/integration/migrate/entv1/user"
)

// User is the model entity for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Age holds the value of the "age" field.
	Age int32 `json:"age,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Nickname holds the value of the "nickname" field.
	Nickname string `json:"nickname,omitempty"`
	// Address holds the value of the "address" field.
	Address string `json:"address,omitempty"`
	// Renamed holds the value of the "renamed" field.
	Renamed string `json:"renamed,omitempty"`
	// Blob holds the value of the "blob" field.
	Blob []byte `json:"blob,omitempty"`
	// State holds the value of the "state" field.
	State user.State `json:"state,omitempty"`
	// Status holds the value of the "status" field.
	Status string `json:"status,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserQuery when eager-loading is set.
	Edges         UserEdges `json:"edges"`
	user_children *int
	user_spouse   *int
}

// UserEdges holds the relations/edges for other nodes in the graph.
type UserEdges struct {
	// Parent holds the value of the parent edge.
	Parent *User
	// Children holds the value of the children edge.
	Children []*User
	// Spouse holds the value of the spouse edge.
	Spouse *User
	// Car holds the value of the car edge.
	Car *Car
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// ParentOrErr returns the Parent value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserEdges) ParentOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.Parent == nil {
			// The edge parent was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.Parent, nil
	}
	return nil, &NotLoadedError{edge: "parent"}
}

// ChildrenOrErr returns the Children value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) ChildrenOrErr() ([]*User, error) {
	if e.loadedTypes[1] {
		return e.Children, nil
	}
	return nil, &NotLoadedError{edge: "children"}
}

// SpouseOrErr returns the Spouse value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserEdges) SpouseOrErr() (*User, error) {
	if e.loadedTypes[2] {
		if e.Spouse == nil {
			// The edge spouse was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.Spouse, nil
	}
	return nil, &NotLoadedError{edge: "spouse"}
}

// CarOrErr returns the Car value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserEdges) CarOrErr() (*Car, error) {
	if e.loadedTypes[3] {
		if e.Car == nil {
			// The edge car was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: car.Label}
		}
		return e.Car, nil
	}
	return nil, &NotLoadedError{edge: "car"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullInt64{},  // age
		&sql.NullString{}, // name
		&sql.NullString{}, // nickname
		&sql.NullString{}, // address
		&sql.NullString{}, // renamed
		&[]byte{},         // blob
		&sql.NullString{}, // state
		&sql.NullString{}, // status
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*User) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // user_children
		&sql.NullInt64{}, // user_spouse
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(values ...interface{}) error {
	if m, n := len(values), len(user.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	u.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field age", values[0])
	} else if value.Valid {
		u.Age = int32(value.Int64)
	}
	if value, ok := values[1].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field name", values[1])
	} else if value.Valid {
		u.Name = value.String
	}
	if value, ok := values[2].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field nickname", values[2])
	} else if value.Valid {
		u.Nickname = value.String
	}
	if value, ok := values[3].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field address", values[3])
	} else if value.Valid {
		u.Address = value.String
	}
	if value, ok := values[4].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field renamed", values[4])
	} else if value.Valid {
		u.Renamed = value.String
	}
	if value, ok := values[5].(*[]byte); !ok {
		return fmt.Errorf("unexpected type %T for field blob", values[5])
	} else if value != nil {
		u.Blob = *value
	}
	if value, ok := values[6].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field state", values[6])
	} else if value.Valid {
		u.State = user.State(value.String)
	}
	if value, ok := values[7].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field status", values[7])
	} else if value.Valid {
		u.Status = value.String
	}
	values = values[8:]
	if len(values) == len(user.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field user_children", value)
		} else if value.Valid {
			u.user_children = new(int)
			*u.user_children = int(value.Int64)
		}
		if value, ok := values[1].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field user_spouse", value)
		} else if value.Valid {
			u.user_spouse = new(int)
			*u.user_spouse = int(value.Int64)
		}
	}
	return nil
}

// QueryParent queries the parent edge of the User.
func (u *User) QueryParent() *UserQuery {
	return (&UserClient{config: u.config}).QueryParent(u)
}

// QueryChildren queries the children edge of the User.
func (u *User) QueryChildren() *UserQuery {
	return (&UserClient{config: u.config}).QueryChildren(u)
}

// QuerySpouse queries the spouse edge of the User.
func (u *User) QuerySpouse() *UserQuery {
	return (&UserClient{config: u.config}).QuerySpouse(u)
}

// QueryCar queries the car edge of the User.
func (u *User) QueryCar() *CarQuery {
	return (&UserClient{config: u.config}).QueryCar(u)
}

// Update returns a builder for updating this User.
// Note that, you need to call User.Unwrap() before calling this method, if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return (&UserClient{config: u.config}).UpdateOne(u)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("entv1: User is not a transactional entity")
	}
	u.config.driver = tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v", u.ID))
	builder.WriteString(", age=")
	builder.WriteString(fmt.Sprintf("%v", u.Age))
	builder.WriteString(", name=")
	builder.WriteString(u.Name)
	builder.WriteString(", nickname=")
	builder.WriteString(u.Nickname)
	builder.WriteString(", address=")
	builder.WriteString(u.Address)
	builder.WriteString(", renamed=")
	builder.WriteString(u.Renamed)
	builder.WriteString(", blob=")
	builder.WriteString(fmt.Sprintf("%v", u.Blob))
	builder.WriteString(", state=")
	builder.WriteString(fmt.Sprintf("%v", u.State))
	builder.WriteString(", status=")
	builder.WriteString(u.Status)
	builder.WriteByte(')')
	return builder.String()
}

// Users is a parsable slice of User.
type Users []*User

func (u Users) config(cfg config) {
	for _i := range u {
		u[_i].config = cfg
	}
}
