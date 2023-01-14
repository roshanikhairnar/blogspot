// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"entdemo/ent/blogs"
	"entdemo/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// BlogsDelete is the builder for deleting a Blogs entity.
type BlogsDelete struct {
	config
	hooks    []Hook
	mutation *BlogsMutation
}

// Where appends a list predicates to the BlogsDelete builder.
func (bd *BlogsDelete) Where(ps ...predicate.Blogs) *BlogsDelete {
	bd.mutation.Where(ps...)
	return bd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (bd *BlogsDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, BlogsMutation](ctx, bd.sqlExec, bd.mutation, bd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (bd *BlogsDelete) ExecX(ctx context.Context) int {
	n, err := bd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (bd *BlogsDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: blogs.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: blogs.FieldID,
			},
		},
	}
	if ps := bd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, bd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	bd.mutation.done = true
	return affected, err
}

// BlogsDeleteOne is the builder for deleting a single Blogs entity.
type BlogsDeleteOne struct {
	bd *BlogsDelete
}

// Exec executes the deletion query.
func (bdo *BlogsDeleteOne) Exec(ctx context.Context) error {
	n, err := bdo.bd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{blogs.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (bdo *BlogsDeleteOne) ExecX(ctx context.Context) {
	bdo.bd.ExecX(ctx)
}
