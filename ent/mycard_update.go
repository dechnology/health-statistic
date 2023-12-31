// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/eesoymilk/health-statistic-api/ent/mycard"
	"github.com/eesoymilk/health-statistic-api/ent/notification"
	"github.com/eesoymilk/health-statistic-api/ent/predicate"
	"github.com/eesoymilk/health-statistic-api/ent/user"
	"github.com/google/uuid"
)

// MyCardUpdate is the builder for updating MyCard entities.
type MyCardUpdate struct {
	config
	hooks    []Hook
	mutation *MyCardMutation
}

// Where appends a list predicates to the MyCardUpdate builder.
func (mcu *MyCardUpdate) Where(ps ...predicate.MyCard) *MyCardUpdate {
	mcu.mutation.Where(ps...)
	return mcu
}

// SetCreatedAt sets the "created_at" field.
func (mcu *MyCardUpdate) SetCreatedAt(t time.Time) *MyCardUpdate {
	mcu.mutation.SetCreatedAt(t)
	return mcu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (mcu *MyCardUpdate) SetNillableCreatedAt(t *time.Time) *MyCardUpdate {
	if t != nil {
		mcu.SetCreatedAt(*t)
	}
	return mcu
}

// SetTakenAt sets the "taken_at" field.
func (mcu *MyCardUpdate) SetTakenAt(t time.Time) *MyCardUpdate {
	mcu.mutation.SetTakenAt(t)
	return mcu
}

// SetNillableTakenAt sets the "taken_at" field if the given value is not nil.
func (mcu *MyCardUpdate) SetNillableTakenAt(t *time.Time) *MyCardUpdate {
	if t != nil {
		mcu.SetTakenAt(*t)
	}
	return mcu
}

// ClearTakenAt clears the value of the "taken_at" field.
func (mcu *MyCardUpdate) ClearTakenAt() *MyCardUpdate {
	mcu.mutation.ClearTakenAt()
	return mcu
}

// SetRecipientID sets the "recipient" edge to the User entity by ID.
func (mcu *MyCardUpdate) SetRecipientID(id string) *MyCardUpdate {
	mcu.mutation.SetRecipientID(id)
	return mcu
}

// SetNillableRecipientID sets the "recipient" edge to the User entity by ID if the given value is not nil.
func (mcu *MyCardUpdate) SetNillableRecipientID(id *string) *MyCardUpdate {
	if id != nil {
		mcu = mcu.SetRecipientID(*id)
	}
	return mcu
}

// SetRecipient sets the "recipient" edge to the User entity.
func (mcu *MyCardUpdate) SetRecipient(u *User) *MyCardUpdate {
	return mcu.SetRecipientID(u.ID)
}

// AddNotificationIDs adds the "notifications" edge to the Notification entity by IDs.
func (mcu *MyCardUpdate) AddNotificationIDs(ids ...uuid.UUID) *MyCardUpdate {
	mcu.mutation.AddNotificationIDs(ids...)
	return mcu
}

// AddNotifications adds the "notifications" edges to the Notification entity.
func (mcu *MyCardUpdate) AddNotifications(n ...*Notification) *MyCardUpdate {
	ids := make([]uuid.UUID, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return mcu.AddNotificationIDs(ids...)
}

// Mutation returns the MyCardMutation object of the builder.
func (mcu *MyCardUpdate) Mutation() *MyCardMutation {
	return mcu.mutation
}

// ClearRecipient clears the "recipient" edge to the User entity.
func (mcu *MyCardUpdate) ClearRecipient() *MyCardUpdate {
	mcu.mutation.ClearRecipient()
	return mcu
}

// ClearNotifications clears all "notifications" edges to the Notification entity.
func (mcu *MyCardUpdate) ClearNotifications() *MyCardUpdate {
	mcu.mutation.ClearNotifications()
	return mcu
}

// RemoveNotificationIDs removes the "notifications" edge to Notification entities by IDs.
func (mcu *MyCardUpdate) RemoveNotificationIDs(ids ...uuid.UUID) *MyCardUpdate {
	mcu.mutation.RemoveNotificationIDs(ids...)
	return mcu
}

// RemoveNotifications removes "notifications" edges to Notification entities.
func (mcu *MyCardUpdate) RemoveNotifications(n ...*Notification) *MyCardUpdate {
	ids := make([]uuid.UUID, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return mcu.RemoveNotificationIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (mcu *MyCardUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, mcu.sqlSave, mcu.mutation, mcu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (mcu *MyCardUpdate) SaveX(ctx context.Context) int {
	affected, err := mcu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mcu *MyCardUpdate) Exec(ctx context.Context) error {
	_, err := mcu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mcu *MyCardUpdate) ExecX(ctx context.Context) {
	if err := mcu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (mcu *MyCardUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(mycard.Table, mycard.Columns, sqlgraph.NewFieldSpec(mycard.FieldID, field.TypeString))
	if ps := mcu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mcu.mutation.CreatedAt(); ok {
		_spec.SetField(mycard.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := mcu.mutation.TakenAt(); ok {
		_spec.SetField(mycard.FieldTakenAt, field.TypeTime, value)
	}
	if mcu.mutation.TakenAtCleared() {
		_spec.ClearField(mycard.FieldTakenAt, field.TypeTime)
	}
	if mcu.mutation.RecipientCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   mycard.RecipientTable,
			Columns: []string{mycard.RecipientColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mcu.mutation.RecipientIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   mycard.RecipientTable,
			Columns: []string{mycard.RecipientColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if mcu.mutation.NotificationsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   mycard.NotificationsTable,
			Columns: []string{mycard.NotificationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(notification.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mcu.mutation.RemovedNotificationsIDs(); len(nodes) > 0 && !mcu.mutation.NotificationsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   mycard.NotificationsTable,
			Columns: []string{mycard.NotificationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(notification.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mcu.mutation.NotificationsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   mycard.NotificationsTable,
			Columns: []string{mycard.NotificationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(notification.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, mcu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{mycard.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	mcu.mutation.done = true
	return n, nil
}

// MyCardUpdateOne is the builder for updating a single MyCard entity.
type MyCardUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *MyCardMutation
}

// SetCreatedAt sets the "created_at" field.
func (mcuo *MyCardUpdateOne) SetCreatedAt(t time.Time) *MyCardUpdateOne {
	mcuo.mutation.SetCreatedAt(t)
	return mcuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (mcuo *MyCardUpdateOne) SetNillableCreatedAt(t *time.Time) *MyCardUpdateOne {
	if t != nil {
		mcuo.SetCreatedAt(*t)
	}
	return mcuo
}

// SetTakenAt sets the "taken_at" field.
func (mcuo *MyCardUpdateOne) SetTakenAt(t time.Time) *MyCardUpdateOne {
	mcuo.mutation.SetTakenAt(t)
	return mcuo
}

// SetNillableTakenAt sets the "taken_at" field if the given value is not nil.
func (mcuo *MyCardUpdateOne) SetNillableTakenAt(t *time.Time) *MyCardUpdateOne {
	if t != nil {
		mcuo.SetTakenAt(*t)
	}
	return mcuo
}

// ClearTakenAt clears the value of the "taken_at" field.
func (mcuo *MyCardUpdateOne) ClearTakenAt() *MyCardUpdateOne {
	mcuo.mutation.ClearTakenAt()
	return mcuo
}

// SetRecipientID sets the "recipient" edge to the User entity by ID.
func (mcuo *MyCardUpdateOne) SetRecipientID(id string) *MyCardUpdateOne {
	mcuo.mutation.SetRecipientID(id)
	return mcuo
}

// SetNillableRecipientID sets the "recipient" edge to the User entity by ID if the given value is not nil.
func (mcuo *MyCardUpdateOne) SetNillableRecipientID(id *string) *MyCardUpdateOne {
	if id != nil {
		mcuo = mcuo.SetRecipientID(*id)
	}
	return mcuo
}

// SetRecipient sets the "recipient" edge to the User entity.
func (mcuo *MyCardUpdateOne) SetRecipient(u *User) *MyCardUpdateOne {
	return mcuo.SetRecipientID(u.ID)
}

// AddNotificationIDs adds the "notifications" edge to the Notification entity by IDs.
func (mcuo *MyCardUpdateOne) AddNotificationIDs(ids ...uuid.UUID) *MyCardUpdateOne {
	mcuo.mutation.AddNotificationIDs(ids...)
	return mcuo
}

// AddNotifications adds the "notifications" edges to the Notification entity.
func (mcuo *MyCardUpdateOne) AddNotifications(n ...*Notification) *MyCardUpdateOne {
	ids := make([]uuid.UUID, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return mcuo.AddNotificationIDs(ids...)
}

// Mutation returns the MyCardMutation object of the builder.
func (mcuo *MyCardUpdateOne) Mutation() *MyCardMutation {
	return mcuo.mutation
}

// ClearRecipient clears the "recipient" edge to the User entity.
func (mcuo *MyCardUpdateOne) ClearRecipient() *MyCardUpdateOne {
	mcuo.mutation.ClearRecipient()
	return mcuo
}

// ClearNotifications clears all "notifications" edges to the Notification entity.
func (mcuo *MyCardUpdateOne) ClearNotifications() *MyCardUpdateOne {
	mcuo.mutation.ClearNotifications()
	return mcuo
}

// RemoveNotificationIDs removes the "notifications" edge to Notification entities by IDs.
func (mcuo *MyCardUpdateOne) RemoveNotificationIDs(ids ...uuid.UUID) *MyCardUpdateOne {
	mcuo.mutation.RemoveNotificationIDs(ids...)
	return mcuo
}

// RemoveNotifications removes "notifications" edges to Notification entities.
func (mcuo *MyCardUpdateOne) RemoveNotifications(n ...*Notification) *MyCardUpdateOne {
	ids := make([]uuid.UUID, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return mcuo.RemoveNotificationIDs(ids...)
}

// Where appends a list predicates to the MyCardUpdate builder.
func (mcuo *MyCardUpdateOne) Where(ps ...predicate.MyCard) *MyCardUpdateOne {
	mcuo.mutation.Where(ps...)
	return mcuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (mcuo *MyCardUpdateOne) Select(field string, fields ...string) *MyCardUpdateOne {
	mcuo.fields = append([]string{field}, fields...)
	return mcuo
}

// Save executes the query and returns the updated MyCard entity.
func (mcuo *MyCardUpdateOne) Save(ctx context.Context) (*MyCard, error) {
	return withHooks(ctx, mcuo.sqlSave, mcuo.mutation, mcuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (mcuo *MyCardUpdateOne) SaveX(ctx context.Context) *MyCard {
	node, err := mcuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (mcuo *MyCardUpdateOne) Exec(ctx context.Context) error {
	_, err := mcuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mcuo *MyCardUpdateOne) ExecX(ctx context.Context) {
	if err := mcuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (mcuo *MyCardUpdateOne) sqlSave(ctx context.Context) (_node *MyCard, err error) {
	_spec := sqlgraph.NewUpdateSpec(mycard.Table, mycard.Columns, sqlgraph.NewFieldSpec(mycard.FieldID, field.TypeString))
	id, ok := mcuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "MyCard.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := mcuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, mycard.FieldID)
		for _, f := range fields {
			if !mycard.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != mycard.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := mcuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mcuo.mutation.CreatedAt(); ok {
		_spec.SetField(mycard.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := mcuo.mutation.TakenAt(); ok {
		_spec.SetField(mycard.FieldTakenAt, field.TypeTime, value)
	}
	if mcuo.mutation.TakenAtCleared() {
		_spec.ClearField(mycard.FieldTakenAt, field.TypeTime)
	}
	if mcuo.mutation.RecipientCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   mycard.RecipientTable,
			Columns: []string{mycard.RecipientColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mcuo.mutation.RecipientIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   mycard.RecipientTable,
			Columns: []string{mycard.RecipientColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if mcuo.mutation.NotificationsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   mycard.NotificationsTable,
			Columns: []string{mycard.NotificationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(notification.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mcuo.mutation.RemovedNotificationsIDs(); len(nodes) > 0 && !mcuo.mutation.NotificationsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   mycard.NotificationsTable,
			Columns: []string{mycard.NotificationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(notification.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mcuo.mutation.NotificationsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   mycard.NotificationsTable,
			Columns: []string{mycard.NotificationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(notification.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &MyCard{config: mcuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, mcuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{mycard.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	mcuo.mutation.done = true
	return _node, nil
}
