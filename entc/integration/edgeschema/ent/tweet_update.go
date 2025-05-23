// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entc/integration/edgeschema/ent/predicate"
	"entgo.io/ent/entc/integration/edgeschema/ent/tag"
	"entgo.io/ent/entc/integration/edgeschema/ent/tweet"
	"entgo.io/ent/entc/integration/edgeschema/ent/tweettag"
	"entgo.io/ent/entc/integration/edgeschema/ent/user"
	"entgo.io/ent/entc/integration/edgeschema/ent/usertweet"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// TweetUpdate is the builder for updating Tweet entities.
type TweetUpdate struct {
	config
	hooks    []Hook
	mutation *TweetMutation
}

// Where appends a list predicates to the TweetUpdate builder.
func (tu *TweetUpdate) Where(ps ...predicate.Tweet) *TweetUpdate {
	tu.mutation.Where(ps...)
	return tu
}

// SetText sets the "text" field.
func (tu *TweetUpdate) SetText(s string) *TweetUpdate {
	tu.mutation.SetText(s)
	return tu
}

// AddLikedUserIDs adds the "liked_users" edge to the User entity by IDs.
func (tu *TweetUpdate) AddLikedUserIDs(ids ...int) *TweetUpdate {
	tu.mutation.AddLikedUserIDs(ids...)
	return tu
}

// AddLikedUsers adds the "liked_users" edges to the User entity.
func (tu *TweetUpdate) AddLikedUsers(u ...*User) *TweetUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return tu.AddLikedUserIDs(ids...)
}

// AddUserIDs adds the "user" edge to the User entity by IDs.
func (tu *TweetUpdate) AddUserIDs(ids ...int) *TweetUpdate {
	tu.mutation.AddUserIDs(ids...)
	return tu
}

// AddUser adds the "user" edges to the User entity.
func (tu *TweetUpdate) AddUser(u ...*User) *TweetUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return tu.AddUserIDs(ids...)
}

// AddTagIDs adds the "tags" edge to the Tag entity by IDs.
func (tu *TweetUpdate) AddTagIDs(ids ...int) *TweetUpdate {
	tu.mutation.AddTagIDs(ids...)
	return tu
}

// AddTags adds the "tags" edges to the Tag entity.
func (tu *TweetUpdate) AddTags(t ...*Tag) *TweetUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tu.AddTagIDs(ids...)
}

// AddTweetUserIDs adds the "tweet_user" edge to the UserTweet entity by IDs.
func (tu *TweetUpdate) AddTweetUserIDs(ids ...int) *TweetUpdate {
	tu.mutation.AddTweetUserIDs(ids...)
	return tu
}

// AddTweetUser adds the "tweet_user" edges to the UserTweet entity.
func (tu *TweetUpdate) AddTweetUser(u ...*UserTweet) *TweetUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return tu.AddTweetUserIDs(ids...)
}

// AddTweetTagIDs adds the "tweet_tags" edge to the TweetTag entity by IDs.
func (tu *TweetUpdate) AddTweetTagIDs(ids ...uuid.UUID) *TweetUpdate {
	tu.mutation.AddTweetTagIDs(ids...)
	return tu
}

// AddTweetTags adds the "tweet_tags" edges to the TweetTag entity.
func (tu *TweetUpdate) AddTweetTags(t ...*TweetTag) *TweetUpdate {
	ids := make([]uuid.UUID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tu.AddTweetTagIDs(ids...)
}

// Mutation returns the TweetMutation object of the builder.
func (tu *TweetUpdate) Mutation() *TweetMutation {
	return tu.mutation
}

// ClearLikedUsers clears all "liked_users" edges to the User entity.
func (tu *TweetUpdate) ClearLikedUsers() *TweetUpdate {
	tu.mutation.ClearLikedUsers()
	return tu
}

// RemoveLikedUserIDs removes the "liked_users" edge to User entities by IDs.
func (tu *TweetUpdate) RemoveLikedUserIDs(ids ...int) *TweetUpdate {
	tu.mutation.RemoveLikedUserIDs(ids...)
	return tu
}

// RemoveLikedUsers removes "liked_users" edges to User entities.
func (tu *TweetUpdate) RemoveLikedUsers(u ...*User) *TweetUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return tu.RemoveLikedUserIDs(ids...)
}

// ClearUser clears all "user" edges to the User entity.
func (tu *TweetUpdate) ClearUser() *TweetUpdate {
	tu.mutation.ClearUser()
	return tu
}

// RemoveUserIDs removes the "user" edge to User entities by IDs.
func (tu *TweetUpdate) RemoveUserIDs(ids ...int) *TweetUpdate {
	tu.mutation.RemoveUserIDs(ids...)
	return tu
}

// RemoveUser removes "user" edges to User entities.
func (tu *TweetUpdate) RemoveUser(u ...*User) *TweetUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return tu.RemoveUserIDs(ids...)
}

// ClearTags clears all "tags" edges to the Tag entity.
func (tu *TweetUpdate) ClearTags() *TweetUpdate {
	tu.mutation.ClearTags()
	return tu
}

// RemoveTagIDs removes the "tags" edge to Tag entities by IDs.
func (tu *TweetUpdate) RemoveTagIDs(ids ...int) *TweetUpdate {
	tu.mutation.RemoveTagIDs(ids...)
	return tu
}

// RemoveTags removes "tags" edges to Tag entities.
func (tu *TweetUpdate) RemoveTags(t ...*Tag) *TweetUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tu.RemoveTagIDs(ids...)
}

// ClearTweetUser clears all "tweet_user" edges to the UserTweet entity.
func (tu *TweetUpdate) ClearTweetUser() *TweetUpdate {
	tu.mutation.ClearTweetUser()
	return tu
}

// RemoveTweetUserIDs removes the "tweet_user" edge to UserTweet entities by IDs.
func (tu *TweetUpdate) RemoveTweetUserIDs(ids ...int) *TweetUpdate {
	tu.mutation.RemoveTweetUserIDs(ids...)
	return tu
}

// RemoveTweetUser removes "tweet_user" edges to UserTweet entities.
func (tu *TweetUpdate) RemoveTweetUser(u ...*UserTweet) *TweetUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return tu.RemoveTweetUserIDs(ids...)
}

// ClearTweetTags clears all "tweet_tags" edges to the TweetTag entity.
func (tu *TweetUpdate) ClearTweetTags() *TweetUpdate {
	tu.mutation.ClearTweetTags()
	return tu
}

// RemoveTweetTagIDs removes the "tweet_tags" edge to TweetTag entities by IDs.
func (tu *TweetUpdate) RemoveTweetTagIDs(ids ...uuid.UUID) *TweetUpdate {
	tu.mutation.RemoveTweetTagIDs(ids...)
	return tu
}

// RemoveTweetTags removes "tweet_tags" edges to TweetTag entities.
func (tu *TweetUpdate) RemoveTweetTags(t ...*TweetTag) *TweetUpdate {
	ids := make([]uuid.UUID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tu.RemoveTweetTagIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tu *TweetUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, TweetMutation](ctx, tu.sqlSave, tu.mutation, tu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tu *TweetUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *TweetUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *TweetUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (tu *TweetUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   tweet.Table,
			Columns: tweet.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: tweet.FieldID,
			},
		},
	}
	if ps := tu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tu.mutation.Text(); ok {
		_spec.SetField(tweet.FieldText, field.TypeString, value)
	}
	if tu.mutation.LikedUsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   tweet.LikedUsersTable,
			Columns: tweet.LikedUsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		createE := &TweetLikeCreate{config: tu.config, mutation: newTweetLikeMutation(tu.config, OpCreate)}
		_ = createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.RemovedLikedUsersIDs(); len(nodes) > 0 && !tu.mutation.LikedUsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   tweet.LikedUsersTable,
			Columns: tweet.LikedUsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &TweetLikeCreate{config: tu.config, mutation: newTweetLikeMutation(tu.config, OpCreate)}
		_ = createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.LikedUsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   tweet.LikedUsersTable,
			Columns: tweet.LikedUsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &TweetLikeCreate{config: tu.config, mutation: newTweetLikeMutation(tu.config, OpCreate)}
		_ = createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   tweet.UserTable,
			Columns: tweet.UserPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		createE := &UserTweetCreate{config: tu.config, mutation: newUserTweetMutation(tu.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.RemovedUserIDs(); len(nodes) > 0 && !tu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   tweet.UserTable,
			Columns: tweet.UserPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &UserTweetCreate{config: tu.config, mutation: newUserTweetMutation(tu.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   tweet.UserTable,
			Columns: tweet.UserPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &UserTweetCreate{config: tu.config, mutation: newUserTweetMutation(tu.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tu.mutation.TagsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   tweet.TagsTable,
			Columns: tweet.TagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tag.FieldID,
				},
			},
		}
		createE := &TweetTagCreate{config: tu.config, mutation: newTweetTagMutation(tu.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		if specE.ID.Value != nil {
			edge.Target.Fields = append(edge.Target.Fields, specE.ID)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.RemovedTagsIDs(); len(nodes) > 0 && !tu.mutation.TagsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   tweet.TagsTable,
			Columns: tweet.TagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tag.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &TweetTagCreate{config: tu.config, mutation: newTweetTagMutation(tu.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		if specE.ID.Value != nil {
			edge.Target.Fields = append(edge.Target.Fields, specE.ID)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.TagsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   tweet.TagsTable,
			Columns: tweet.TagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tag.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &TweetTagCreate{config: tu.config, mutation: newTweetTagMutation(tu.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		if specE.ID.Value != nil {
			edge.Target.Fields = append(edge.Target.Fields, specE.ID)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tu.mutation.TweetUserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   tweet.TweetUserTable,
			Columns: []string{tweet.TweetUserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: usertweet.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.RemovedTweetUserIDs(); len(nodes) > 0 && !tu.mutation.TweetUserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   tweet.TweetUserTable,
			Columns: []string{tweet.TweetUserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: usertweet.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.TweetUserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   tweet.TweetUserTable,
			Columns: []string{tweet.TweetUserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: usertweet.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tu.mutation.TweetTagsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   tweet.TweetTagsTable,
			Columns: []string{tweet.TweetTagsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: tweettag.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.RemovedTweetTagsIDs(); len(nodes) > 0 && !tu.mutation.TweetTagsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   tweet.TweetTagsTable,
			Columns: []string{tweet.TweetTagsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: tweettag.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.TweetTagsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   tweet.TweetTagsTable,
			Columns: []string{tweet.TweetTagsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: tweettag.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{tweet.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	tu.mutation.done = true
	return n, nil
}

// TweetUpdateOne is the builder for updating a single Tweet entity.
type TweetUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TweetMutation
}

// SetText sets the "text" field.
func (tuo *TweetUpdateOne) SetText(s string) *TweetUpdateOne {
	tuo.mutation.SetText(s)
	return tuo
}

// AddLikedUserIDs adds the "liked_users" edge to the User entity by IDs.
func (tuo *TweetUpdateOne) AddLikedUserIDs(ids ...int) *TweetUpdateOne {
	tuo.mutation.AddLikedUserIDs(ids...)
	return tuo
}

// AddLikedUsers adds the "liked_users" edges to the User entity.
func (tuo *TweetUpdateOne) AddLikedUsers(u ...*User) *TweetUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return tuo.AddLikedUserIDs(ids...)
}

// AddUserIDs adds the "user" edge to the User entity by IDs.
func (tuo *TweetUpdateOne) AddUserIDs(ids ...int) *TweetUpdateOne {
	tuo.mutation.AddUserIDs(ids...)
	return tuo
}

// AddUser adds the "user" edges to the User entity.
func (tuo *TweetUpdateOne) AddUser(u ...*User) *TweetUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return tuo.AddUserIDs(ids...)
}

// AddTagIDs adds the "tags" edge to the Tag entity by IDs.
func (tuo *TweetUpdateOne) AddTagIDs(ids ...int) *TweetUpdateOne {
	tuo.mutation.AddTagIDs(ids...)
	return tuo
}

// AddTags adds the "tags" edges to the Tag entity.
func (tuo *TweetUpdateOne) AddTags(t ...*Tag) *TweetUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tuo.AddTagIDs(ids...)
}

// AddTweetUserIDs adds the "tweet_user" edge to the UserTweet entity by IDs.
func (tuo *TweetUpdateOne) AddTweetUserIDs(ids ...int) *TweetUpdateOne {
	tuo.mutation.AddTweetUserIDs(ids...)
	return tuo
}

// AddTweetUser adds the "tweet_user" edges to the UserTweet entity.
func (tuo *TweetUpdateOne) AddTweetUser(u ...*UserTweet) *TweetUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return tuo.AddTweetUserIDs(ids...)
}

// AddTweetTagIDs adds the "tweet_tags" edge to the TweetTag entity by IDs.
func (tuo *TweetUpdateOne) AddTweetTagIDs(ids ...uuid.UUID) *TweetUpdateOne {
	tuo.mutation.AddTweetTagIDs(ids...)
	return tuo
}

// AddTweetTags adds the "tweet_tags" edges to the TweetTag entity.
func (tuo *TweetUpdateOne) AddTweetTags(t ...*TweetTag) *TweetUpdateOne {
	ids := make([]uuid.UUID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tuo.AddTweetTagIDs(ids...)
}

// Mutation returns the TweetMutation object of the builder.
func (tuo *TweetUpdateOne) Mutation() *TweetMutation {
	return tuo.mutation
}

// ClearLikedUsers clears all "liked_users" edges to the User entity.
func (tuo *TweetUpdateOne) ClearLikedUsers() *TweetUpdateOne {
	tuo.mutation.ClearLikedUsers()
	return tuo
}

// RemoveLikedUserIDs removes the "liked_users" edge to User entities by IDs.
func (tuo *TweetUpdateOne) RemoveLikedUserIDs(ids ...int) *TweetUpdateOne {
	tuo.mutation.RemoveLikedUserIDs(ids...)
	return tuo
}

// RemoveLikedUsers removes "liked_users" edges to User entities.
func (tuo *TweetUpdateOne) RemoveLikedUsers(u ...*User) *TweetUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return tuo.RemoveLikedUserIDs(ids...)
}

// ClearUser clears all "user" edges to the User entity.
func (tuo *TweetUpdateOne) ClearUser() *TweetUpdateOne {
	tuo.mutation.ClearUser()
	return tuo
}

// RemoveUserIDs removes the "user" edge to User entities by IDs.
func (tuo *TweetUpdateOne) RemoveUserIDs(ids ...int) *TweetUpdateOne {
	tuo.mutation.RemoveUserIDs(ids...)
	return tuo
}

// RemoveUser removes "user" edges to User entities.
func (tuo *TweetUpdateOne) RemoveUser(u ...*User) *TweetUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return tuo.RemoveUserIDs(ids...)
}

// ClearTags clears all "tags" edges to the Tag entity.
func (tuo *TweetUpdateOne) ClearTags() *TweetUpdateOne {
	tuo.mutation.ClearTags()
	return tuo
}

// RemoveTagIDs removes the "tags" edge to Tag entities by IDs.
func (tuo *TweetUpdateOne) RemoveTagIDs(ids ...int) *TweetUpdateOne {
	tuo.mutation.RemoveTagIDs(ids...)
	return tuo
}

// RemoveTags removes "tags" edges to Tag entities.
func (tuo *TweetUpdateOne) RemoveTags(t ...*Tag) *TweetUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tuo.RemoveTagIDs(ids...)
}

// ClearTweetUser clears all "tweet_user" edges to the UserTweet entity.
func (tuo *TweetUpdateOne) ClearTweetUser() *TweetUpdateOne {
	tuo.mutation.ClearTweetUser()
	return tuo
}

// RemoveTweetUserIDs removes the "tweet_user" edge to UserTweet entities by IDs.
func (tuo *TweetUpdateOne) RemoveTweetUserIDs(ids ...int) *TweetUpdateOne {
	tuo.mutation.RemoveTweetUserIDs(ids...)
	return tuo
}

// RemoveTweetUser removes "tweet_user" edges to UserTweet entities.
func (tuo *TweetUpdateOne) RemoveTweetUser(u ...*UserTweet) *TweetUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return tuo.RemoveTweetUserIDs(ids...)
}

// ClearTweetTags clears all "tweet_tags" edges to the TweetTag entity.
func (tuo *TweetUpdateOne) ClearTweetTags() *TweetUpdateOne {
	tuo.mutation.ClearTweetTags()
	return tuo
}

// RemoveTweetTagIDs removes the "tweet_tags" edge to TweetTag entities by IDs.
func (tuo *TweetUpdateOne) RemoveTweetTagIDs(ids ...uuid.UUID) *TweetUpdateOne {
	tuo.mutation.RemoveTweetTagIDs(ids...)
	return tuo
}

// RemoveTweetTags removes "tweet_tags" edges to TweetTag entities.
func (tuo *TweetUpdateOne) RemoveTweetTags(t ...*TweetTag) *TweetUpdateOne {
	ids := make([]uuid.UUID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tuo.RemoveTweetTagIDs(ids...)
}

// Where appends a list predicates to the TweetUpdate builder.
func (tuo *TweetUpdateOne) Where(ps ...predicate.Tweet) *TweetUpdateOne {
	tuo.mutation.Where(ps...)
	return tuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tuo *TweetUpdateOne) Select(field string, fields ...string) *TweetUpdateOne {
	tuo.fields = append([]string{field}, fields...)
	return tuo
}

// Save executes the query and returns the updated Tweet entity.
func (tuo *TweetUpdateOne) Save(ctx context.Context) (*Tweet, error) {
	return withHooks[*Tweet, TweetMutation](ctx, tuo.sqlSave, tuo.mutation, tuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *TweetUpdateOne) SaveX(ctx context.Context) *Tweet {
	node, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tuo *TweetUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *TweetUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (tuo *TweetUpdateOne) sqlSave(ctx context.Context) (_node *Tweet, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   tweet.Table,
			Columns: tweet.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: tweet.FieldID,
			},
		},
	}
	id, ok := tuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Tweet.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, tweet.FieldID)
		for _, f := range fields {
			if !tweet.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != tweet.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tuo.mutation.Text(); ok {
		_spec.SetField(tweet.FieldText, field.TypeString, value)
	}
	if tuo.mutation.LikedUsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   tweet.LikedUsersTable,
			Columns: tweet.LikedUsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		createE := &TweetLikeCreate{config: tuo.config, mutation: newTweetLikeMutation(tuo.config, OpCreate)}
		_ = createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.RemovedLikedUsersIDs(); len(nodes) > 0 && !tuo.mutation.LikedUsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   tweet.LikedUsersTable,
			Columns: tweet.LikedUsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &TweetLikeCreate{config: tuo.config, mutation: newTweetLikeMutation(tuo.config, OpCreate)}
		_ = createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.LikedUsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   tweet.LikedUsersTable,
			Columns: tweet.LikedUsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &TweetLikeCreate{config: tuo.config, mutation: newTweetLikeMutation(tuo.config, OpCreate)}
		_ = createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tuo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   tweet.UserTable,
			Columns: tweet.UserPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		createE := &UserTweetCreate{config: tuo.config, mutation: newUserTweetMutation(tuo.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.RemovedUserIDs(); len(nodes) > 0 && !tuo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   tweet.UserTable,
			Columns: tweet.UserPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &UserTweetCreate{config: tuo.config, mutation: newUserTweetMutation(tuo.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   tweet.UserTable,
			Columns: tweet.UserPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &UserTweetCreate{config: tuo.config, mutation: newUserTweetMutation(tuo.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tuo.mutation.TagsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   tweet.TagsTable,
			Columns: tweet.TagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tag.FieldID,
				},
			},
		}
		createE := &TweetTagCreate{config: tuo.config, mutation: newTweetTagMutation(tuo.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		if specE.ID.Value != nil {
			edge.Target.Fields = append(edge.Target.Fields, specE.ID)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.RemovedTagsIDs(); len(nodes) > 0 && !tuo.mutation.TagsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   tweet.TagsTable,
			Columns: tweet.TagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tag.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &TweetTagCreate{config: tuo.config, mutation: newTweetTagMutation(tuo.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		if specE.ID.Value != nil {
			edge.Target.Fields = append(edge.Target.Fields, specE.ID)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.TagsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   tweet.TagsTable,
			Columns: tweet.TagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tag.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &TweetTagCreate{config: tuo.config, mutation: newTweetTagMutation(tuo.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		if specE.ID.Value != nil {
			edge.Target.Fields = append(edge.Target.Fields, specE.ID)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tuo.mutation.TweetUserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   tweet.TweetUserTable,
			Columns: []string{tweet.TweetUserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: usertweet.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.RemovedTweetUserIDs(); len(nodes) > 0 && !tuo.mutation.TweetUserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   tweet.TweetUserTable,
			Columns: []string{tweet.TweetUserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: usertweet.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.TweetUserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   tweet.TweetUserTable,
			Columns: []string{tweet.TweetUserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: usertweet.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tuo.mutation.TweetTagsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   tweet.TweetTagsTable,
			Columns: []string{tweet.TweetTagsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: tweettag.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.RemovedTweetTagsIDs(); len(nodes) > 0 && !tuo.mutation.TweetTagsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   tweet.TweetTagsTable,
			Columns: []string{tweet.TweetTagsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: tweettag.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.TweetTagsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   tweet.TweetTagsTable,
			Columns: []string{tweet.TweetTagsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: tweettag.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Tweet{config: tuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{tweet.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	tuo.mutation.done = true
	return _node, nil
}
