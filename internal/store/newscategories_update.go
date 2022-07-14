// Code generated by entc, DO NOT EDIT.

package store

import (
	"context"
	"errors"
	"fmt"
	"news/internal/store/newscategories"
	"news/internal/store/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// NewsCategoriesUpdate is the builder for updating NewsCategories entities.
type NewsCategoriesUpdate struct {
	config
	hooks    []Hook
	mutation *NewsCategoriesMutation
}

// Where appends a list predicates to the NewsCategoriesUpdate builder.
func (ncu *NewsCategoriesUpdate) Where(ps ...predicate.NewsCategories) *NewsCategoriesUpdate {
	ncu.mutation.Where(ps...)
	return ncu
}

// SetStatus sets the "status" field.
func (ncu *NewsCategoriesUpdate) SetStatus(b bool) *NewsCategoriesUpdate {
	ncu.mutation.SetStatus(b)
	return ncu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (ncu *NewsCategoriesUpdate) SetNillableStatus(b *bool) *NewsCategoriesUpdate {
	if b != nil {
		ncu.SetStatus(*b)
	}
	return ncu
}

// SetLanguage sets the "language" field.
func (ncu *NewsCategoriesUpdate) SetLanguage(s string) *NewsCategoriesUpdate {
	ncu.mutation.SetLanguage(s)
	return ncu
}

// SetNillableLanguage sets the "language" field if the given value is not nil.
func (ncu *NewsCategoriesUpdate) SetNillableLanguage(s *string) *NewsCategoriesUpdate {
	if s != nil {
		ncu.SetLanguage(*s)
	}
	return ncu
}

// SetTvCategories sets the "tv_categories" field.
func (ncu *NewsCategoriesUpdate) SetTvCategories(m map[string]string) *NewsCategoriesUpdate {
	ncu.mutation.SetTvCategories(m)
	return ncu
}

// SetArticleCategories sets the "article_categories" field.
func (ncu *NewsCategoriesUpdate) SetArticleCategories(m map[string]string) *NewsCategoriesUpdate {
	ncu.mutation.SetArticleCategories(m)
	return ncu
}

// Mutation returns the NewsCategoriesMutation object of the builder.
func (ncu *NewsCategoriesUpdate) Mutation() *NewsCategoriesMutation {
	return ncu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ncu *NewsCategoriesUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ncu.hooks) == 0 {
		affected, err = ncu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*NewsCategoriesMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ncu.mutation = mutation
			affected, err = ncu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ncu.hooks) - 1; i >= 0; i-- {
			if ncu.hooks[i] == nil {
				return 0, fmt.Errorf("store: uninitialized hook (forgotten import store/runtime?)")
			}
			mut = ncu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ncu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ncu *NewsCategoriesUpdate) SaveX(ctx context.Context) int {
	affected, err := ncu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ncu *NewsCategoriesUpdate) Exec(ctx context.Context) error {
	_, err := ncu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ncu *NewsCategoriesUpdate) ExecX(ctx context.Context) {
	if err := ncu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ncu *NewsCategoriesUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   newscategories.Table,
			Columns: newscategories.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: newscategories.FieldID,
			},
		},
	}
	if ps := ncu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ncu.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: newscategories.FieldStatus,
		})
	}
	if value, ok := ncu.mutation.Language(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: newscategories.FieldLanguage,
		})
	}
	if value, ok := ncu.mutation.TvCategories(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: newscategories.FieldTvCategories,
		})
	}
	if value, ok := ncu.mutation.ArticleCategories(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: newscategories.FieldArticleCategories,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ncu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{newscategories.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// NewsCategoriesUpdateOne is the builder for updating a single NewsCategories entity.
type NewsCategoriesUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *NewsCategoriesMutation
}

// SetStatus sets the "status" field.
func (ncuo *NewsCategoriesUpdateOne) SetStatus(b bool) *NewsCategoriesUpdateOne {
	ncuo.mutation.SetStatus(b)
	return ncuo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (ncuo *NewsCategoriesUpdateOne) SetNillableStatus(b *bool) *NewsCategoriesUpdateOne {
	if b != nil {
		ncuo.SetStatus(*b)
	}
	return ncuo
}

// SetLanguage sets the "language" field.
func (ncuo *NewsCategoriesUpdateOne) SetLanguage(s string) *NewsCategoriesUpdateOne {
	ncuo.mutation.SetLanguage(s)
	return ncuo
}

// SetNillableLanguage sets the "language" field if the given value is not nil.
func (ncuo *NewsCategoriesUpdateOne) SetNillableLanguage(s *string) *NewsCategoriesUpdateOne {
	if s != nil {
		ncuo.SetLanguage(*s)
	}
	return ncuo
}

// SetTvCategories sets the "tv_categories" field.
func (ncuo *NewsCategoriesUpdateOne) SetTvCategories(m map[string]string) *NewsCategoriesUpdateOne {
	ncuo.mutation.SetTvCategories(m)
	return ncuo
}

// SetArticleCategories sets the "article_categories" field.
func (ncuo *NewsCategoriesUpdateOne) SetArticleCategories(m map[string]string) *NewsCategoriesUpdateOne {
	ncuo.mutation.SetArticleCategories(m)
	return ncuo
}

// Mutation returns the NewsCategoriesMutation object of the builder.
func (ncuo *NewsCategoriesUpdateOne) Mutation() *NewsCategoriesMutation {
	return ncuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ncuo *NewsCategoriesUpdateOne) Select(field string, fields ...string) *NewsCategoriesUpdateOne {
	ncuo.fields = append([]string{field}, fields...)
	return ncuo
}

// Save executes the query and returns the updated NewsCategories entity.
func (ncuo *NewsCategoriesUpdateOne) Save(ctx context.Context) (*NewsCategories, error) {
	var (
		err  error
		node *NewsCategories
	)
	if len(ncuo.hooks) == 0 {
		node, err = ncuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*NewsCategoriesMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ncuo.mutation = mutation
			node, err = ncuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ncuo.hooks) - 1; i >= 0; i-- {
			if ncuo.hooks[i] == nil {
				return nil, fmt.Errorf("store: uninitialized hook (forgotten import store/runtime?)")
			}
			mut = ncuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ncuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ncuo *NewsCategoriesUpdateOne) SaveX(ctx context.Context) *NewsCategories {
	node, err := ncuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ncuo *NewsCategoriesUpdateOne) Exec(ctx context.Context) error {
	_, err := ncuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ncuo *NewsCategoriesUpdateOne) ExecX(ctx context.Context) {
	if err := ncuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ncuo *NewsCategoriesUpdateOne) sqlSave(ctx context.Context) (_node *NewsCategories, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   newscategories.Table,
			Columns: newscategories.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: newscategories.FieldID,
			},
		},
	}
	id, ok := ncuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`store: missing "NewsCategories.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ncuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, newscategories.FieldID)
		for _, f := range fields {
			if !newscategories.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("store: invalid field %q for query", f)}
			}
			if f != newscategories.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ncuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ncuo.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: newscategories.FieldStatus,
		})
	}
	if value, ok := ncuo.mutation.Language(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: newscategories.FieldLanguage,
		})
	}
	if value, ok := ncuo.mutation.TvCategories(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: newscategories.FieldTvCategories,
		})
	}
	if value, ok := ncuo.mutation.ArticleCategories(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: newscategories.FieldArticleCategories,
		})
	}
	_node = &NewsCategories{config: ncuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ncuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{newscategories.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
