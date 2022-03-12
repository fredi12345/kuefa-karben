// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/fredi12345/kuefa-karben/src/ent/predicate"
	"github.com/fredi12345/kuefa-karben/src/ent/titleimage"
	"github.com/google/uuid"
)

// TitleImageQuery is the builder for querying TitleImage entities.
type TitleImageQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.TitleImage
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the TitleImageQuery builder.
func (tiq *TitleImageQuery) Where(ps ...predicate.TitleImage) *TitleImageQuery {
	tiq.predicates = append(tiq.predicates, ps...)
	return tiq
}

// Limit adds a limit step to the query.
func (tiq *TitleImageQuery) Limit(limit int) *TitleImageQuery {
	tiq.limit = &limit
	return tiq
}

// Offset adds an offset step to the query.
func (tiq *TitleImageQuery) Offset(offset int) *TitleImageQuery {
	tiq.offset = &offset
	return tiq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (tiq *TitleImageQuery) Unique(unique bool) *TitleImageQuery {
	tiq.unique = &unique
	return tiq
}

// Order adds an order step to the query.
func (tiq *TitleImageQuery) Order(o ...OrderFunc) *TitleImageQuery {
	tiq.order = append(tiq.order, o...)
	return tiq
}

// First returns the first TitleImage entity from the query.
// Returns a *NotFoundError when no TitleImage was found.
func (tiq *TitleImageQuery) First(ctx context.Context) (*TitleImage, error) {
	nodes, err := tiq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{titleimage.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (tiq *TitleImageQuery) FirstX(ctx context.Context) *TitleImage {
	node, err := tiq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first TitleImage ID from the query.
// Returns a *NotFoundError when no TitleImage ID was found.
func (tiq *TitleImageQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = tiq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{titleimage.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (tiq *TitleImageQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := tiq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single TitleImage entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one TitleImage entity is not found.
// Returns a *NotFoundError when no TitleImage entities are found.
func (tiq *TitleImageQuery) Only(ctx context.Context) (*TitleImage, error) {
	nodes, err := tiq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{titleimage.Label}
	default:
		return nil, &NotSingularError{titleimage.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (tiq *TitleImageQuery) OnlyX(ctx context.Context) *TitleImage {
	node, err := tiq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only TitleImage ID in the query.
// Returns a *NotSingularError when exactly one TitleImage ID is not found.
// Returns a *NotFoundError when no entities are found.
func (tiq *TitleImageQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = tiq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{titleimage.Label}
	default:
		err = &NotSingularError{titleimage.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (tiq *TitleImageQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := tiq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of TitleImages.
func (tiq *TitleImageQuery) All(ctx context.Context) ([]*TitleImage, error) {
	if err := tiq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return tiq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (tiq *TitleImageQuery) AllX(ctx context.Context) []*TitleImage {
	nodes, err := tiq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of TitleImage IDs.
func (tiq *TitleImageQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := tiq.Select(titleimage.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (tiq *TitleImageQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := tiq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (tiq *TitleImageQuery) Count(ctx context.Context) (int, error) {
	if err := tiq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return tiq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (tiq *TitleImageQuery) CountX(ctx context.Context) int {
	count, err := tiq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (tiq *TitleImageQuery) Exist(ctx context.Context) (bool, error) {
	if err := tiq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return tiq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (tiq *TitleImageQuery) ExistX(ctx context.Context) bool {
	exist, err := tiq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the TitleImageQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (tiq *TitleImageQuery) Clone() *TitleImageQuery {
	if tiq == nil {
		return nil
	}
	return &TitleImageQuery{
		config:     tiq.config,
		limit:      tiq.limit,
		offset:     tiq.offset,
		order:      append([]OrderFunc{}, tiq.order...),
		predicates: append([]predicate.TitleImage{}, tiq.predicates...),
		// clone intermediate query.
		sql:  tiq.sql.Clone(),
		path: tiq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Created time.Time `json:"created,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.TitleImage.Query().
//		GroupBy(titleimage.FieldCreated).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (tiq *TitleImageQuery) GroupBy(field string, fields ...string) *TitleImageGroupBy {
	group := &TitleImageGroupBy{config: tiq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := tiq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return tiq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Created time.Time `json:"created,omitempty"`
//	}
//
//	client.TitleImage.Query().
//		Select(titleimage.FieldCreated).
//		Scan(ctx, &v)
//
func (tiq *TitleImageQuery) Select(fields ...string) *TitleImageSelect {
	tiq.fields = append(tiq.fields, fields...)
	return &TitleImageSelect{TitleImageQuery: tiq}
}

func (tiq *TitleImageQuery) prepareQuery(ctx context.Context) error {
	for _, f := range tiq.fields {
		if !titleimage.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if tiq.path != nil {
		prev, err := tiq.path(ctx)
		if err != nil {
			return err
		}
		tiq.sql = prev
	}
	return nil
}

func (tiq *TitleImageQuery) sqlAll(ctx context.Context) ([]*TitleImage, error) {
	var (
		nodes = []*TitleImage{}
		_spec = tiq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &TitleImage{config: tiq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, tiq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (tiq *TitleImageQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := tiq.querySpec()
	_spec.Node.Columns = tiq.fields
	if len(tiq.fields) > 0 {
		_spec.Unique = tiq.unique != nil && *tiq.unique
	}
	return sqlgraph.CountNodes(ctx, tiq.driver, _spec)
}

func (tiq *TitleImageQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := tiq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (tiq *TitleImageQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   titleimage.Table,
			Columns: titleimage.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: titleimage.FieldID,
			},
		},
		From:   tiq.sql,
		Unique: true,
	}
	if unique := tiq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := tiq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, titleimage.FieldID)
		for i := range fields {
			if fields[i] != titleimage.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := tiq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := tiq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := tiq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := tiq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (tiq *TitleImageQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(tiq.driver.Dialect())
	t1 := builder.Table(titleimage.Table)
	columns := tiq.fields
	if len(columns) == 0 {
		columns = titleimage.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if tiq.sql != nil {
		selector = tiq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if tiq.unique != nil && *tiq.unique {
		selector.Distinct()
	}
	for _, p := range tiq.predicates {
		p(selector)
	}
	for _, p := range tiq.order {
		p(selector)
	}
	if offset := tiq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := tiq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// TitleImageGroupBy is the group-by builder for TitleImage entities.
type TitleImageGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (tigb *TitleImageGroupBy) Aggregate(fns ...AggregateFunc) *TitleImageGroupBy {
	tigb.fns = append(tigb.fns, fns...)
	return tigb
}

// Scan applies the group-by query and scans the result into the given value.
func (tigb *TitleImageGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := tigb.path(ctx)
	if err != nil {
		return err
	}
	tigb.sql = query
	return tigb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (tigb *TitleImageGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := tigb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (tigb *TitleImageGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(tigb.fields) > 1 {
		return nil, errors.New("ent: TitleImageGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := tigb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (tigb *TitleImageGroupBy) StringsX(ctx context.Context) []string {
	v, err := tigb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (tigb *TitleImageGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = tigb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{titleimage.Label}
	default:
		err = fmt.Errorf("ent: TitleImageGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (tigb *TitleImageGroupBy) StringX(ctx context.Context) string {
	v, err := tigb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (tigb *TitleImageGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(tigb.fields) > 1 {
		return nil, errors.New("ent: TitleImageGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := tigb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (tigb *TitleImageGroupBy) IntsX(ctx context.Context) []int {
	v, err := tigb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (tigb *TitleImageGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = tigb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{titleimage.Label}
	default:
		err = fmt.Errorf("ent: TitleImageGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (tigb *TitleImageGroupBy) IntX(ctx context.Context) int {
	v, err := tigb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (tigb *TitleImageGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(tigb.fields) > 1 {
		return nil, errors.New("ent: TitleImageGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := tigb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (tigb *TitleImageGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := tigb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (tigb *TitleImageGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = tigb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{titleimage.Label}
	default:
		err = fmt.Errorf("ent: TitleImageGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (tigb *TitleImageGroupBy) Float64X(ctx context.Context) float64 {
	v, err := tigb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (tigb *TitleImageGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(tigb.fields) > 1 {
		return nil, errors.New("ent: TitleImageGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := tigb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (tigb *TitleImageGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := tigb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (tigb *TitleImageGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = tigb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{titleimage.Label}
	default:
		err = fmt.Errorf("ent: TitleImageGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (tigb *TitleImageGroupBy) BoolX(ctx context.Context) bool {
	v, err := tigb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (tigb *TitleImageGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range tigb.fields {
		if !titleimage.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := tigb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := tigb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (tigb *TitleImageGroupBy) sqlQuery() *sql.Selector {
	selector := tigb.sql.Select()
	aggregation := make([]string, 0, len(tigb.fns))
	for _, fn := range tigb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(tigb.fields)+len(tigb.fns))
		for _, f := range tigb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(tigb.fields...)...)
}

// TitleImageSelect is the builder for selecting fields of TitleImage entities.
type TitleImageSelect struct {
	*TitleImageQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (tis *TitleImageSelect) Scan(ctx context.Context, v interface{}) error {
	if err := tis.prepareQuery(ctx); err != nil {
		return err
	}
	tis.sql = tis.TitleImageQuery.sqlQuery(ctx)
	return tis.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (tis *TitleImageSelect) ScanX(ctx context.Context, v interface{}) {
	if err := tis.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (tis *TitleImageSelect) Strings(ctx context.Context) ([]string, error) {
	if len(tis.fields) > 1 {
		return nil, errors.New("ent: TitleImageSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := tis.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (tis *TitleImageSelect) StringsX(ctx context.Context) []string {
	v, err := tis.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (tis *TitleImageSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = tis.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{titleimage.Label}
	default:
		err = fmt.Errorf("ent: TitleImageSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (tis *TitleImageSelect) StringX(ctx context.Context) string {
	v, err := tis.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (tis *TitleImageSelect) Ints(ctx context.Context) ([]int, error) {
	if len(tis.fields) > 1 {
		return nil, errors.New("ent: TitleImageSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := tis.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (tis *TitleImageSelect) IntsX(ctx context.Context) []int {
	v, err := tis.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (tis *TitleImageSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = tis.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{titleimage.Label}
	default:
		err = fmt.Errorf("ent: TitleImageSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (tis *TitleImageSelect) IntX(ctx context.Context) int {
	v, err := tis.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (tis *TitleImageSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(tis.fields) > 1 {
		return nil, errors.New("ent: TitleImageSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := tis.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (tis *TitleImageSelect) Float64sX(ctx context.Context) []float64 {
	v, err := tis.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (tis *TitleImageSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = tis.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{titleimage.Label}
	default:
		err = fmt.Errorf("ent: TitleImageSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (tis *TitleImageSelect) Float64X(ctx context.Context) float64 {
	v, err := tis.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (tis *TitleImageSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(tis.fields) > 1 {
		return nil, errors.New("ent: TitleImageSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := tis.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (tis *TitleImageSelect) BoolsX(ctx context.Context) []bool {
	v, err := tis.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (tis *TitleImageSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = tis.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{titleimage.Label}
	default:
		err = fmt.Errorf("ent: TitleImageSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (tis *TitleImageSelect) BoolX(ctx context.Context) bool {
	v, err := tis.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (tis *TitleImageSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := tis.sql.Query()
	if err := tis.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}