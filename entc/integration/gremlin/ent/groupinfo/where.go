// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package groupinfo

import (
	"github.com/apocelipes/ent/dialect/gremlin/graph/dsl"
	"github.com/apocelipes/ent/dialect/gremlin/graph/dsl/__"
	"github.com/apocelipes/ent/dialect/gremlin/graph/dsl/p"
	"github.com/apocelipes/ent/entc/integration/gremlin/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id string) predicate.GroupInfo {
	return predicate.GroupInfo(func(t *dsl.Traversal) {
		t.HasID(id)
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.GroupInfo {
	return predicate.GroupInfo(func(t *dsl.Traversal) {
		t.HasID(p.EQ(id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.GroupInfo {
	return predicate.GroupInfo(func(t *dsl.Traversal) {
		t.HasID(p.NEQ(id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.GroupInfo {
	return predicate.GroupInfo(func(t *dsl.Traversal) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		t.HasID(p.Within(v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.GroupInfo {
	return predicate.GroupInfo(func(t *dsl.Traversal) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		t.HasID(p.Without(v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.GroupInfo {
	return predicate.GroupInfo(func(t *dsl.Traversal) {
		t.HasID(p.GT(id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.GroupInfo {
	return predicate.GroupInfo(func(t *dsl.Traversal) {
		t.HasID(p.GTE(id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.GroupInfo {
	return predicate.GroupInfo(func(t *dsl.Traversal) {
		t.HasID(p.LT(id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.GroupInfo {
	return predicate.GroupInfo(func(t *dsl.Traversal) {
		t.HasID(p.LTE(id))
	})
}

// Desc applies equality check predicate on the "desc" field. It's identical to DescEQ.
func Desc(v string) predicate.GroupInfo {
	return predicate.GroupInfo(func(t *dsl.Traversal) {
		t.Has(Label, FieldDesc, p.EQ(v))
	})
}

// MaxUsers applies equality check predicate on the "max_users" field. It's identical to MaxUsersEQ.
func MaxUsers(v int) predicate.GroupInfo {
	return predicate.GroupInfo(func(t *dsl.Traversal) {
		t.Has(Label, FieldMaxUsers, p.EQ(v))
	})
}

// DescEQ applies the EQ predicate on the "desc" field.
func DescEQ(v string) predicate.GroupInfo {
	return predicate.GroupInfo(func(t *dsl.Traversal) {
		t.Has(Label, FieldDesc, p.EQ(v))
	})
}

// DescNEQ applies the NEQ predicate on the "desc" field.
func DescNEQ(v string) predicate.GroupInfo {
	return predicate.GroupInfo(func(t *dsl.Traversal) {
		t.Has(Label, FieldDesc, p.NEQ(v))
	})
}

// DescIn applies the In predicate on the "desc" field.
func DescIn(vs ...string) predicate.GroupInfo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GroupInfo(func(t *dsl.Traversal) {
		t.Has(Label, FieldDesc, p.Within(v...))
	})
}

// DescNotIn applies the NotIn predicate on the "desc" field.
func DescNotIn(vs ...string) predicate.GroupInfo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GroupInfo(func(t *dsl.Traversal) {
		t.Has(Label, FieldDesc, p.Without(v...))
	})
}

// DescGT applies the GT predicate on the "desc" field.
func DescGT(v string) predicate.GroupInfo {
	return predicate.GroupInfo(func(t *dsl.Traversal) {
		t.Has(Label, FieldDesc, p.GT(v))
	})
}

// DescGTE applies the GTE predicate on the "desc" field.
func DescGTE(v string) predicate.GroupInfo {
	return predicate.GroupInfo(func(t *dsl.Traversal) {
		t.Has(Label, FieldDesc, p.GTE(v))
	})
}

// DescLT applies the LT predicate on the "desc" field.
func DescLT(v string) predicate.GroupInfo {
	return predicate.GroupInfo(func(t *dsl.Traversal) {
		t.Has(Label, FieldDesc, p.LT(v))
	})
}

// DescLTE applies the LTE predicate on the "desc" field.
func DescLTE(v string) predicate.GroupInfo {
	return predicate.GroupInfo(func(t *dsl.Traversal) {
		t.Has(Label, FieldDesc, p.LTE(v))
	})
}

// DescContains applies the Contains predicate on the "desc" field.
func DescContains(v string) predicate.GroupInfo {
	return predicate.GroupInfo(func(t *dsl.Traversal) {
		t.Has(Label, FieldDesc, p.Containing(v))
	})
}

// DescHasPrefix applies the HasPrefix predicate on the "desc" field.
func DescHasPrefix(v string) predicate.GroupInfo {
	return predicate.GroupInfo(func(t *dsl.Traversal) {
		t.Has(Label, FieldDesc, p.StartingWith(v))
	})
}

// DescHasSuffix applies the HasSuffix predicate on the "desc" field.
func DescHasSuffix(v string) predicate.GroupInfo {
	return predicate.GroupInfo(func(t *dsl.Traversal) {
		t.Has(Label, FieldDesc, p.EndingWith(v))
	})
}

// MaxUsersEQ applies the EQ predicate on the "max_users" field.
func MaxUsersEQ(v int) predicate.GroupInfo {
	return predicate.GroupInfo(func(t *dsl.Traversal) {
		t.Has(Label, FieldMaxUsers, p.EQ(v))
	})
}

// MaxUsersNEQ applies the NEQ predicate on the "max_users" field.
func MaxUsersNEQ(v int) predicate.GroupInfo {
	return predicate.GroupInfo(func(t *dsl.Traversal) {
		t.Has(Label, FieldMaxUsers, p.NEQ(v))
	})
}

// MaxUsersIn applies the In predicate on the "max_users" field.
func MaxUsersIn(vs ...int) predicate.GroupInfo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GroupInfo(func(t *dsl.Traversal) {
		t.Has(Label, FieldMaxUsers, p.Within(v...))
	})
}

// MaxUsersNotIn applies the NotIn predicate on the "max_users" field.
func MaxUsersNotIn(vs ...int) predicate.GroupInfo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GroupInfo(func(t *dsl.Traversal) {
		t.Has(Label, FieldMaxUsers, p.Without(v...))
	})
}

// MaxUsersGT applies the GT predicate on the "max_users" field.
func MaxUsersGT(v int) predicate.GroupInfo {
	return predicate.GroupInfo(func(t *dsl.Traversal) {
		t.Has(Label, FieldMaxUsers, p.GT(v))
	})
}

// MaxUsersGTE applies the GTE predicate on the "max_users" field.
func MaxUsersGTE(v int) predicate.GroupInfo {
	return predicate.GroupInfo(func(t *dsl.Traversal) {
		t.Has(Label, FieldMaxUsers, p.GTE(v))
	})
}

// MaxUsersLT applies the LT predicate on the "max_users" field.
func MaxUsersLT(v int) predicate.GroupInfo {
	return predicate.GroupInfo(func(t *dsl.Traversal) {
		t.Has(Label, FieldMaxUsers, p.LT(v))
	})
}

// MaxUsersLTE applies the LTE predicate on the "max_users" field.
func MaxUsersLTE(v int) predicate.GroupInfo {
	return predicate.GroupInfo(func(t *dsl.Traversal) {
		t.Has(Label, FieldMaxUsers, p.LTE(v))
	})
}

// HasGroups applies the HasEdge predicate on the "groups" edge.
func HasGroups() predicate.GroupInfo {
	return predicate.GroupInfo(func(t *dsl.Traversal) {
		t.InE(GroupsInverseLabel).InV()
	})
}

// HasGroupsWith applies the HasEdge predicate on the "groups" edge with a given conditions (other predicates).
func HasGroupsWith(preds ...predicate.Group) predicate.GroupInfo {
	return predicate.GroupInfo(func(t *dsl.Traversal) {
		tr := __.OutV()
		for _, p := range preds {
			p(tr)
		}
		t.InE(GroupsInverseLabel).Where(tr).InV()
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.GroupInfo) predicate.GroupInfo {
	return predicate.GroupInfo(func(tr *dsl.Traversal) {
		trs := make([]interface{}, 0, len(predicates))
		for _, p := range predicates {
			t := __.New()
			p(t)
			trs = append(trs, t)
		}
		tr.Where(__.And(trs...))
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.GroupInfo) predicate.GroupInfo {
	return predicate.GroupInfo(func(tr *dsl.Traversal) {
		trs := make([]interface{}, 0, len(predicates))
		for _, p := range predicates {
			t := __.New()
			p(t)
			trs = append(trs, t)
		}
		tr.Where(__.Or(trs...))
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.GroupInfo) predicate.GroupInfo {
	return predicate.GroupInfo(func(tr *dsl.Traversal) {
		t := __.New()
		p(t)
		tr.Where(__.Not(t))
	})
}
