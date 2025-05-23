// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/gremlin"
	"entgo.io/ent/entc/integration/gremlin/ent/filetype"
)

// FileType is the model entity for the FileType schema.
type FileType struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Type holds the value of the "type" field.
	Type filetype.Type `json:"type,omitempty"`
	// State holds the value of the "state" field.
	State filetype.State `json:"state,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the FileTypeQuery when eager-loading is set.
	Edges FileTypeEdges `json:"edges"`
}

// FileTypeEdges holds the relations/edges for other nodes in the graph.
type FileTypeEdges struct {
	// Files holds the value of the files edge.
	Files []*File `json:"files,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// FilesOrErr returns the Files value or an error if the edge
// was not loaded in eager-loading.
func (e FileTypeEdges) FilesOrErr() ([]*File, error) {
	if e.loadedTypes[0] {
		return e.Files, nil
	}
	return nil, &NotLoadedError{edge: "files"}
}

// FromResponse scans the gremlin response data into FileType.
func (ft *FileType) FromResponse(res *gremlin.Response) error {
	vmap, err := res.ReadValueMap()
	if err != nil {
		return err
	}
	var scanft struct {
		ID    string         `json:"id,omitempty"`
		Name  string         `json:"name,omitempty"`
		Type  filetype.Type  `json:"type,omitempty"`
		State filetype.State `json:"state,omitempty"`
	}
	if err := vmap.Decode(&scanft); err != nil {
		return err
	}
	ft.ID = scanft.ID
	ft.Name = scanft.Name
	ft.Type = scanft.Type
	ft.State = scanft.State
	return nil
}

// QueryFiles queries the "files" edge of the FileType entity.
func (ft *FileType) QueryFiles() *FileQuery {
	return NewFileTypeClient(ft.config).QueryFiles(ft)
}

// Update returns a builder for updating this FileType.
// Note that you need to call FileType.Unwrap() before calling this method if this FileType
// was returned from a transaction, and the transaction was committed or rolled back.
func (ft *FileType) Update() *FileTypeUpdateOne {
	return NewFileTypeClient(ft.config).UpdateOne(ft)
}

// Unwrap unwraps the FileType entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ft *FileType) Unwrap() *FileType {
	_tx, ok := ft.config.driver.(*txDriver)
	if !ok {
		panic("ent: FileType is not a transactional entity")
	}
	ft.config.driver = _tx.drv
	return ft
}

// String implements the fmt.Stringer.
func (ft *FileType) String() string {
	var builder strings.Builder
	builder.WriteString("FileType(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ft.ID))
	builder.WriteString("name=")
	builder.WriteString(ft.Name)
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(fmt.Sprintf("%v", ft.Type))
	builder.WriteString(", ")
	builder.WriteString("state=")
	builder.WriteString(fmt.Sprintf("%v", ft.State))
	builder.WriteByte(')')
	return builder.String()
}

// FileTypes is a parsable slice of FileType.
type FileTypes []*FileType

// FromResponse scans the gremlin response data into FileTypes.
func (ft *FileTypes) FromResponse(res *gremlin.Response) error {
	vmap, err := res.ReadValueMap()
	if err != nil {
		return err
	}
	var scanft []struct {
		ID    string         `json:"id,omitempty"`
		Name  string         `json:"name,omitempty"`
		Type  filetype.Type  `json:"type,omitempty"`
		State filetype.State `json:"state,omitempty"`
	}
	if err := vmap.Decode(&scanft); err != nil {
		return err
	}
	for _, v := range scanft {
		node := &FileType{ID: v.ID}
		node.Name = v.Name
		node.Type = v.Type
		node.State = v.State
		*ft = append(*ft, node)
	}
	return nil
}

func (ft FileTypes) config(cfg config) {
	for _i := range ft {
		ft[_i].config = cfg
	}
}
