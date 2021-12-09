// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/AndreasVikke-School/CPH-Bussines-SI-Exam/applications/services/postgres/ent/loan"
	"github.com/AndreasVikke-School/CPH-Bussines-SI-Exam/applications/services/postgres/ent/predicate"
	"github.com/AndreasVikke-School/CPH-Bussines-SI-Exam/applications/services/postgres/ent/user"
)

// LoanUpdate is the builder for updating Loan entities.
type LoanUpdate struct {
	config
	hooks    []Hook
	mutation *LoanMutation
}

// Where appends a list predicates to the LoanUpdate builder.
func (lu *LoanUpdate) Where(ps ...predicate.Loan) *LoanUpdate {
	lu.mutation.Where(ps...)
	return lu
}

// SetEntityId sets the "entityId" field.
func (lu *LoanUpdate) SetEntityId(i int64) *LoanUpdate {
	lu.mutation.ResetEntityId()
	lu.mutation.SetEntityId(i)
	return lu
}

// AddEntityId adds i to the "entityId" field.
func (lu *LoanUpdate) AddEntityId(i int64) *LoanUpdate {
	lu.mutation.AddEntityId(i)
	return lu
}

// SetStatus sets the "status" field.
func (lu *LoanUpdate) SetStatus(l loan.Status) *LoanUpdate {
	lu.mutation.SetStatus(l)
	return lu
}

// SetUserID sets the "user" edge to the User entity by ID.
func (lu *LoanUpdate) SetUserID(id int) *LoanUpdate {
	lu.mutation.SetUserID(id)
	return lu
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (lu *LoanUpdate) SetNillableUserID(id *int) *LoanUpdate {
	if id != nil {
		lu = lu.SetUserID(*id)
	}
	return lu
}

// SetUser sets the "user" edge to the User entity.
func (lu *LoanUpdate) SetUser(u *User) *LoanUpdate {
	return lu.SetUserID(u.ID)
}

// Mutation returns the LoanMutation object of the builder.
func (lu *LoanUpdate) Mutation() *LoanMutation {
	return lu.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (lu *LoanUpdate) ClearUser() *LoanUpdate {
	lu.mutation.ClearUser()
	return lu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (lu *LoanUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(lu.hooks) == 0 {
		if err = lu.check(); err != nil {
			return 0, err
		}
		affected, err = lu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*LoanMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = lu.check(); err != nil {
				return 0, err
			}
			lu.mutation = mutation
			affected, err = lu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(lu.hooks) - 1; i >= 0; i-- {
			if lu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = lu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, lu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (lu *LoanUpdate) SaveX(ctx context.Context) int {
	affected, err := lu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (lu *LoanUpdate) Exec(ctx context.Context) error {
	_, err := lu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lu *LoanUpdate) ExecX(ctx context.Context) {
	if err := lu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (lu *LoanUpdate) check() error {
	if v, ok := lu.mutation.Status(); ok {
		if err := loan.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf("ent: validator failed for field \"status\": %w", err)}
		}
	}
	return nil
}

func (lu *LoanUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   loan.Table,
			Columns: loan.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: loan.FieldID,
			},
		},
	}
	if ps := lu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := lu.mutation.EntityId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: loan.FieldEntityId,
		})
	}
	if value, ok := lu.mutation.AddedEntityId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: loan.FieldEntityId,
		})
	}
	if value, ok := lu.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: loan.FieldStatus,
		})
	}
	if lu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   loan.UserTable,
			Columns: []string{loan.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   loan.UserTable,
			Columns: []string{loan.UserColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, lu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{loan.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// LoanUpdateOne is the builder for updating a single Loan entity.
type LoanUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *LoanMutation
}

// SetEntityId sets the "entityId" field.
func (luo *LoanUpdateOne) SetEntityId(i int64) *LoanUpdateOne {
	luo.mutation.ResetEntityId()
	luo.mutation.SetEntityId(i)
	return luo
}

// AddEntityId adds i to the "entityId" field.
func (luo *LoanUpdateOne) AddEntityId(i int64) *LoanUpdateOne {
	luo.mutation.AddEntityId(i)
	return luo
}

// SetStatus sets the "status" field.
func (luo *LoanUpdateOne) SetStatus(l loan.Status) *LoanUpdateOne {
	luo.mutation.SetStatus(l)
	return luo
}

// SetUserID sets the "user" edge to the User entity by ID.
func (luo *LoanUpdateOne) SetUserID(id int) *LoanUpdateOne {
	luo.mutation.SetUserID(id)
	return luo
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (luo *LoanUpdateOne) SetNillableUserID(id *int) *LoanUpdateOne {
	if id != nil {
		luo = luo.SetUserID(*id)
	}
	return luo
}

// SetUser sets the "user" edge to the User entity.
func (luo *LoanUpdateOne) SetUser(u *User) *LoanUpdateOne {
	return luo.SetUserID(u.ID)
}

// Mutation returns the LoanMutation object of the builder.
func (luo *LoanUpdateOne) Mutation() *LoanMutation {
	return luo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (luo *LoanUpdateOne) ClearUser() *LoanUpdateOne {
	luo.mutation.ClearUser()
	return luo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (luo *LoanUpdateOne) Select(field string, fields ...string) *LoanUpdateOne {
	luo.fields = append([]string{field}, fields...)
	return luo
}

// Save executes the query and returns the updated Loan entity.
func (luo *LoanUpdateOne) Save(ctx context.Context) (*Loan, error) {
	var (
		err  error
		node *Loan
	)
	if len(luo.hooks) == 0 {
		if err = luo.check(); err != nil {
			return nil, err
		}
		node, err = luo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*LoanMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = luo.check(); err != nil {
				return nil, err
			}
			luo.mutation = mutation
			node, err = luo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(luo.hooks) - 1; i >= 0; i-- {
			if luo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = luo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, luo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (luo *LoanUpdateOne) SaveX(ctx context.Context) *Loan {
	node, err := luo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (luo *LoanUpdateOne) Exec(ctx context.Context) error {
	_, err := luo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (luo *LoanUpdateOne) ExecX(ctx context.Context) {
	if err := luo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (luo *LoanUpdateOne) check() error {
	if v, ok := luo.mutation.Status(); ok {
		if err := loan.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf("ent: validator failed for field \"status\": %w", err)}
		}
	}
	return nil
}

func (luo *LoanUpdateOne) sqlSave(ctx context.Context) (_node *Loan, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   loan.Table,
			Columns: loan.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: loan.FieldID,
			},
		},
	}
	id, ok := luo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Loan.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := luo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, loan.FieldID)
		for _, f := range fields {
			if !loan.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != loan.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := luo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := luo.mutation.EntityId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: loan.FieldEntityId,
		})
	}
	if value, ok := luo.mutation.AddedEntityId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: loan.FieldEntityId,
		})
	}
	if value, ok := luo.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: loan.FieldStatus,
		})
	}
	if luo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   loan.UserTable,
			Columns: []string{loan.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := luo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   loan.UserTable,
			Columns: []string{loan.UserColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Loan{config: luo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, luo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{loan.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}