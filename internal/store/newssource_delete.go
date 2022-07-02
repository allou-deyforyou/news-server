// Code generated by entc, DO NOT EDIT.

package store

import (
	"context"
	"fmt"
	"news/internal/store/newssource"
	"news/internal/store/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// NewsSourceDelete is the builder for deleting a NewsSource entity.
type NewsSourceDelete struct {
	config
	hooks    []Hook
	mutation *NewsSourceMutation
}

// Where appends a list predicates to the NewsSourceDelete builder.
func (nsd *NewsSourceDelete) Where(ps ...predicate.NewsSource) *NewsSourceDelete {
	nsd.mutation.Where(ps...)
	return nsd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (nsd *NewsSourceDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(nsd.hooks) == 0 {
		affected, err = nsd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*NewsSourceMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			nsd.mutation = mutation
			affected, err = nsd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(nsd.hooks) - 1; i >= 0; i-- {
			if nsd.hooks[i] == nil {
				return 0, fmt.Errorf("store: uninitialized hook (forgotten import store/runtime?)")
			}
			mut = nsd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, nsd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (nsd *NewsSourceDelete) ExecX(ctx context.Context) int {
	n, err := nsd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (nsd *NewsSourceDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: newssource.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: newssource.FieldID,
			},
		},
	}
	if ps := nsd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, nsd.driver, _spec)
}

// NewsSourceDeleteOne is the builder for deleting a single NewsSource entity.
type NewsSourceDeleteOne struct {
	nsd *NewsSourceDelete
}

// Exec executes the deletion query.
func (nsdo *NewsSourceDeleteOne) Exec(ctx context.Context) error {
	n, err := nsdo.nsd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{newssource.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (nsdo *NewsSourceDeleteOne) ExecX(ctx context.Context) {
	nsdo.nsd.ExecX(ctx)
}