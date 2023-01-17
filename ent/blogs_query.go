// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"entdemo/ent/blogs"
	"entdemo/ent/predicate"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// BlogsQuery is the builder for querying Blogs entities.
type BlogsQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	inters     []Interceptor
	predicates []predicate.Blogs
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the BlogsQuery builder.
func (bq *BlogsQuery) Where(ps ...predicate.Blogs) *BlogsQuery {
	bq.predicates = append(bq.predicates, ps...)
	return bq
}

// Limit the number of records to be returned by this query.
func (bq *BlogsQuery) Limit(limit int) *BlogsQuery {
	bq.limit = &limit
	return bq
}

// Offset to start from.
func (bq *BlogsQuery) Offset(offset int) *BlogsQuery {
	bq.offset = &offset
	return bq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (bq *BlogsQuery) Unique(unique bool) *BlogsQuery {
	bq.unique = &unique
	return bq
}

// Order specifies how the records should be ordered.
func (bq *BlogsQuery) Order(o ...OrderFunc) *BlogsQuery {
	bq.order = append(bq.order, o...)
	return bq
}

// First returns the first Blogs entity from the query.
// Returns a *NotFoundError when no Blogs was found.
func (bq *BlogsQuery) First(ctx context.Context) (*Blogs, error) {
	nodes, err := bq.Limit(1).All(newQueryContext(ctx, TypeBlogs, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{blogs.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (bq *BlogsQuery) FirstX(ctx context.Context) *Blogs {
	node, err := bq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Blogs ID from the query.
// Returns a *NotFoundError when no Blogs ID was found.
func (bq *BlogsQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = bq.Limit(1).IDs(newQueryContext(ctx, TypeBlogs, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{blogs.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (bq *BlogsQuery) FirstIDX(ctx context.Context) int {
	id, err := bq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Blogs entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Blogs entity is found.
// Returns a *NotFoundError when no Blogs entities are found.
func (bq *BlogsQuery) Only(ctx context.Context) (*Blogs, error) {
	nodes, err := bq.Limit(2).All(newQueryContext(ctx, TypeBlogs, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{blogs.Label}
	default:
		return nil, &NotSingularError{blogs.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (bq *BlogsQuery) OnlyX(ctx context.Context) *Blogs {
	node, err := bq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Blogs ID in the query.
// Returns a *NotSingularError when more than one Blogs ID is found.
// Returns a *NotFoundError when no entities are found.
func (bq *BlogsQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = bq.Limit(2).IDs(newQueryContext(ctx, TypeBlogs, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{blogs.Label}
	default:
		err = &NotSingularError{blogs.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (bq *BlogsQuery) OnlyIDX(ctx context.Context) int {
	id, err := bq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of BlogsSlice.
func (bq *BlogsQuery) All(ctx context.Context) ([]*Blogs, error) {
	ctx = newQueryContext(ctx, TypeBlogs, "All")
	if err := bq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Blogs, *BlogsQuery]()
	return withInterceptors[[]*Blogs](ctx, bq, qr, bq.inters)
}

// AllX is like All, but panics if an error occurs.
func (bq *BlogsQuery) AllX(ctx context.Context) []*Blogs {
	nodes, err := bq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Blogs IDs.
func (bq *BlogsQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	ctx = newQueryContext(ctx, TypeBlogs, "IDs")
	if err := bq.Select(blogs.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (bq *BlogsQuery) IDsX(ctx context.Context) []int {
	ids, err := bq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (bq *BlogsQuery) Count(ctx context.Context) (int, error) {
	ctx = newQueryContext(ctx, TypeBlogs, "Count")
	if err := bq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, bq, querierCount[*BlogsQuery](), bq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (bq *BlogsQuery) CountX(ctx context.Context) int {
	count, err := bq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (bq *BlogsQuery) Exist(ctx context.Context) (bool, error) {
	ctx = newQueryContext(ctx, TypeBlogs, "Exist")
	switch _, err := bq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (bq *BlogsQuery) ExistX(ctx context.Context) bool {
	exist, err := bq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the BlogsQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (bq *BlogsQuery) Clone() *BlogsQuery {
	if bq == nil {
		return nil
	}
	return &BlogsQuery{
		config:     bq.config,
		limit:      bq.limit,
		offset:     bq.offset,
		order:      append([]OrderFunc{}, bq.order...),
		inters:     append([]Interceptor{}, bq.inters...),
		predicates: append([]predicate.Blogs{}, bq.predicates...),
		// clone intermediate query.
		sql:    bq.sql.Clone(),
		path:   bq.path,
		unique: bq.unique,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		BlogTitle string `json:"blogTitle,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Blogs.Query().
//		GroupBy(blogs.FieldBlogTitle).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (bq *BlogsQuery) GroupBy(field string, fields ...string) *BlogsGroupBy {
	bq.fields = append([]string{field}, fields...)
	grbuild := &BlogsGroupBy{build: bq}
	grbuild.flds = &bq.fields
	grbuild.label = blogs.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		BlogTitle string `json:"blogTitle,omitempty"`
//	}
//
//	client.Blogs.Query().
//		Select(blogs.FieldBlogTitle).
//		Scan(ctx, &v)
func (bq *BlogsQuery) Select(fields ...string) *BlogsSelect {
	bq.fields = append(bq.fields, fields...)
	sbuild := &BlogsSelect{BlogsQuery: bq}
	sbuild.label = blogs.Label
	sbuild.flds, sbuild.scan = &bq.fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a BlogsSelect configured with the given aggregations.
func (bq *BlogsQuery) Aggregate(fns ...AggregateFunc) *BlogsSelect {
	return bq.Select().Aggregate(fns...)
}

func (bq *BlogsQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range bq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, bq); err != nil {
				return err
			}
		}
	}
	for _, f := range bq.fields {
		if !blogs.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if bq.path != nil {
		prev, err := bq.path(ctx)
		if err != nil {
			return err
		}
		bq.sql = prev
	}
	return nil
}

func (bq *BlogsQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Blogs, error) {
	var (
		nodes = []*Blogs{}
		_spec = bq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Blogs).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Blogs{config: bq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, bq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (bq *BlogsQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := bq.querySpec()
	_spec.Node.Columns = bq.fields
	if len(bq.fields) > 0 {
		_spec.Unique = bq.unique != nil && *bq.unique
	}
	return sqlgraph.CountNodes(ctx, bq.driver, _spec)
}

func (bq *BlogsQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   blogs.Table,
			Columns: blogs.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: blogs.FieldID,
			},
		},
		From:   bq.sql,
		Unique: true,
	}
	if unique := bq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := bq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, blogs.FieldID)
		for i := range fields {
			if fields[i] != blogs.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := bq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := bq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := bq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := bq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (bq *BlogsQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(bq.driver.Dialect())
	t1 := builder.Table(blogs.Table)
	columns := bq.fields
	if len(columns) == 0 {
		columns = blogs.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if bq.sql != nil {
		selector = bq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if bq.unique != nil && *bq.unique {
		selector.Distinct()
	}
	for _, p := range bq.predicates {
		p(selector)
	}
	for _, p := range bq.order {
		p(selector)
	}
	if offset := bq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := bq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// BlogsGroupBy is the group-by builder for Blogs entities.
type BlogsGroupBy struct {
	selector
	build *BlogsQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (bgb *BlogsGroupBy) Aggregate(fns ...AggregateFunc) *BlogsGroupBy {
	bgb.fns = append(bgb.fns, fns...)
	return bgb
}

// Scan applies the selector query and scans the result into the given value.
func (bgb *BlogsGroupBy) Scan(ctx context.Context, v any) error {
	ctx = newQueryContext(ctx, TypeBlogs, "GroupBy")
	if err := bgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*BlogsQuery, *BlogsGroupBy](ctx, bgb.build, bgb, bgb.build.inters, v)
}

func (bgb *BlogsGroupBy) sqlScan(ctx context.Context, root *BlogsQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(bgb.fns))
	for _, fn := range bgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*bgb.flds)+len(bgb.fns))
		for _, f := range *bgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*bgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := bgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// BlogsSelect is the builder for selecting fields of Blogs entities.
type BlogsSelect struct {
	*BlogsQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (bs *BlogsSelect) Aggregate(fns ...AggregateFunc) *BlogsSelect {
	bs.fns = append(bs.fns, fns...)
	return bs
}

// Scan applies the selector query and scans the result into the given value.
func (bs *BlogsSelect) Scan(ctx context.Context, v any) error {
	ctx = newQueryContext(ctx, TypeBlogs, "Select")
	if err := bs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*BlogsQuery, *BlogsSelect](ctx, bs.BlogsQuery, bs, bs.inters, v)
}

func (bs *BlogsSelect) sqlScan(ctx context.Context, root *BlogsQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(bs.fns))
	for _, fn := range bs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*bs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := bs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
