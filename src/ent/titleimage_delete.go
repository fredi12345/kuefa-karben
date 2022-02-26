// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/fredi12345/kuefa-karben/src/ent/predicate"
	"github.com/fredi12345/kuefa-karben/src/ent/titleimage"
)

// TitleImageDelete is the builder for deleting a TitleImage entity.
type TitleImageDelete struct {
	config
	hooks    []Hook
	mutation *TitleImageMutation
}

// Where appends a list predicates to the TitleImageDelete builder.
func (tid *TitleImageDelete) Where(ps ...predicate.TitleImage) *TitleImageDelete {
	tid.mutation.Where(ps...)
	return tid
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (tid *TitleImageDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(tid.hooks) == 0 {
		affected, err = tid.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TitleImageMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			tid.mutation = mutation
			affected, err = tid.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(tid.hooks) - 1; i >= 0; i-- {
			if tid.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tid.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tid.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (tid *TitleImageDelete) ExecX(ctx context.Context) int {
	n, err := tid.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (tid *TitleImageDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: titleimage.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: titleimage.FieldID,
			},
		},
	}
	if ps := tid.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, tid.driver, _spec)
}

// TitleImageDeleteOne is the builder for deleting a single TitleImage entity.
type TitleImageDeleteOne struct {
	tid *TitleImageDelete
}

// Exec executes the deletion query.
func (tido *TitleImageDeleteOne) Exec(ctx context.Context) error {
	n, err := tido.tid.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{titleimage.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (tido *TitleImageDeleteOne) ExecX(ctx context.Context) {
	tido.tid.ExecX(ctx)
}
