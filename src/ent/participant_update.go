// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/fredi12345/kuefa-karben/src/ent/event"
	"github.com/fredi12345/kuefa-karben/src/ent/participant"
	"github.com/fredi12345/kuefa-karben/src/ent/predicate"
	"github.com/google/uuid"
)

// ParticipantUpdate is the builder for updating Participant entities.
type ParticipantUpdate struct {
	config
	hooks    []Hook
	mutation *ParticipantMutation
}

// Where appends a list predicates to the ParticipantUpdate builder.
func (pu *ParticipantUpdate) Where(ps ...predicate.Participant) *ParticipantUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetMenu sets the "menu" field.
func (pu *ParticipantUpdate) SetMenu(pa participant.Menu) *ParticipantUpdate {
	pu.mutation.SetMenu(pa)
	return pu
}

// SetEventID sets the "event" edge to the Event entity by ID.
func (pu *ParticipantUpdate) SetEventID(id uuid.UUID) *ParticipantUpdate {
	pu.mutation.SetEventID(id)
	return pu
}

// SetNillableEventID sets the "event" edge to the Event entity by ID if the given value is not nil.
func (pu *ParticipantUpdate) SetNillableEventID(id *uuid.UUID) *ParticipantUpdate {
	if id != nil {
		pu = pu.SetEventID(*id)
	}
	return pu
}

// SetEvent sets the "event" edge to the Event entity.
func (pu *ParticipantUpdate) SetEvent(e *Event) *ParticipantUpdate {
	return pu.SetEventID(e.ID)
}

// Mutation returns the ParticipantMutation object of the builder.
func (pu *ParticipantUpdate) Mutation() *ParticipantMutation {
	return pu.mutation
}

// ClearEvent clears the "event" edge to the Event entity.
func (pu *ParticipantUpdate) ClearEvent() *ParticipantUpdate {
	pu.mutation.ClearEvent()
	return pu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *ParticipantUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(pu.hooks) == 0 {
		if err = pu.check(); err != nil {
			return 0, err
		}
		affected, err = pu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ParticipantMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = pu.check(); err != nil {
				return 0, err
			}
			pu.mutation = mutation
			affected, err = pu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(pu.hooks) - 1; i >= 0; i-- {
			if pu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = pu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (pu *ParticipantUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *ParticipantUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *ParticipantUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pu *ParticipantUpdate) check() error {
	if v, ok := pu.mutation.Menu(); ok {
		if err := participant.MenuValidator(v); err != nil {
			return &ValidationError{Name: "menu", err: fmt.Errorf(`ent: validator failed for field "Participant.menu": %w`, err)}
		}
	}
	return nil
}

func (pu *ParticipantUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   participant.Table,
			Columns: participant.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: participant.FieldID,
			},
		},
	}
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.Menu(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: participant.FieldMenu,
		})
	}
	if pu.mutation.EventCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   participant.EventTable,
			Columns: []string{participant.EventColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: event.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.EventIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   participant.EventTable,
			Columns: []string{participant.EventColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: event.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{participant.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// ParticipantUpdateOne is the builder for updating a single Participant entity.
type ParticipantUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ParticipantMutation
}

// SetMenu sets the "menu" field.
func (puo *ParticipantUpdateOne) SetMenu(pa participant.Menu) *ParticipantUpdateOne {
	puo.mutation.SetMenu(pa)
	return puo
}

// SetEventID sets the "event" edge to the Event entity by ID.
func (puo *ParticipantUpdateOne) SetEventID(id uuid.UUID) *ParticipantUpdateOne {
	puo.mutation.SetEventID(id)
	return puo
}

// SetNillableEventID sets the "event" edge to the Event entity by ID if the given value is not nil.
func (puo *ParticipantUpdateOne) SetNillableEventID(id *uuid.UUID) *ParticipantUpdateOne {
	if id != nil {
		puo = puo.SetEventID(*id)
	}
	return puo
}

// SetEvent sets the "event" edge to the Event entity.
func (puo *ParticipantUpdateOne) SetEvent(e *Event) *ParticipantUpdateOne {
	return puo.SetEventID(e.ID)
}

// Mutation returns the ParticipantMutation object of the builder.
func (puo *ParticipantUpdateOne) Mutation() *ParticipantMutation {
	return puo.mutation
}

// ClearEvent clears the "event" edge to the Event entity.
func (puo *ParticipantUpdateOne) ClearEvent() *ParticipantUpdateOne {
	puo.mutation.ClearEvent()
	return puo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *ParticipantUpdateOne) Select(field string, fields ...string) *ParticipantUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Participant entity.
func (puo *ParticipantUpdateOne) Save(ctx context.Context) (*Participant, error) {
	var (
		err  error
		node *Participant
	)
	if len(puo.hooks) == 0 {
		if err = puo.check(); err != nil {
			return nil, err
		}
		node, err = puo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ParticipantMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = puo.check(); err != nil {
				return nil, err
			}
			puo.mutation = mutation
			node, err = puo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(puo.hooks) - 1; i >= 0; i-- {
			if puo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = puo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, puo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (puo *ParticipantUpdateOne) SaveX(ctx context.Context) *Participant {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *ParticipantUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *ParticipantUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (puo *ParticipantUpdateOne) check() error {
	if v, ok := puo.mutation.Menu(); ok {
		if err := participant.MenuValidator(v); err != nil {
			return &ValidationError{Name: "menu", err: fmt.Errorf(`ent: validator failed for field "Participant.menu": %w`, err)}
		}
	}
	return nil
}

func (puo *ParticipantUpdateOne) sqlSave(ctx context.Context) (_node *Participant, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   participant.Table,
			Columns: participant.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: participant.FieldID,
			},
		},
	}
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Participant.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, participant.FieldID)
		for _, f := range fields {
			if !participant.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != participant.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.Menu(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: participant.FieldMenu,
		})
	}
	if puo.mutation.EventCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   participant.EventTable,
			Columns: []string{participant.EventColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: event.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.EventIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   participant.EventTable,
			Columns: []string{participant.EventColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: event.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Participant{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{participant.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
