// Code generated by entc, DO NOT EDIT.

package loan

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/AndreasVikke-School/CPH-Bussiness-SI-Exam/applications/services/postgres/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Loan {
	return predicate.Loan(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Loan {
	return predicate.Loan(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Loan {
	return predicate.Loan(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Loan {
	return predicate.Loan(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Loan {
	return predicate.Loan(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Loan {
	return predicate.Loan(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Loan {
	return predicate.Loan(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Loan {
	return predicate.Loan(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Loan {
	return predicate.Loan(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// EntityId applies equality check predicate on the "entityId" field. It's identical to EntityIdEQ.
func EntityId(v int64) predicate.Loan {
	return predicate.Loan(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEntityId), v))
	})
}

// EntityIdEQ applies the EQ predicate on the "entityId" field.
func EntityIdEQ(v int64) predicate.Loan {
	return predicate.Loan(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEntityId), v))
	})
}

// EntityIdNEQ applies the NEQ predicate on the "entityId" field.
func EntityIdNEQ(v int64) predicate.Loan {
	return predicate.Loan(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEntityId), v))
	})
}

// EntityIdIn applies the In predicate on the "entityId" field.
func EntityIdIn(vs ...int64) predicate.Loan {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Loan(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldEntityId), v...))
	})
}

// EntityIdNotIn applies the NotIn predicate on the "entityId" field.
func EntityIdNotIn(vs ...int64) predicate.Loan {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Loan(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldEntityId), v...))
	})
}

// EntityIdGT applies the GT predicate on the "entityId" field.
func EntityIdGT(v int64) predicate.Loan {
	return predicate.Loan(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEntityId), v))
	})
}

// EntityIdGTE applies the GTE predicate on the "entityId" field.
func EntityIdGTE(v int64) predicate.Loan {
	return predicate.Loan(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEntityId), v))
	})
}

// EntityIdLT applies the LT predicate on the "entityId" field.
func EntityIdLT(v int64) predicate.Loan {
	return predicate.Loan(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEntityId), v))
	})
}

// EntityIdLTE applies the LTE predicate on the "entityId" field.
func EntityIdLTE(v int64) predicate.Loan {
	return predicate.Loan(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEntityId), v))
	})
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v Status) predicate.Loan {
	return predicate.Loan(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStatus), v))
	})
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v Status) predicate.Loan {
	return predicate.Loan(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStatus), v))
	})
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...Status) predicate.Loan {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Loan(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldStatus), v...))
	})
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...Status) predicate.Loan {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Loan(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldStatus), v...))
	})
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.Loan {
	return predicate.Loan(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.Loan {
	return predicate.Loan(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Loan) predicate.Loan {
	return predicate.Loan(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Loan) predicate.Loan {
	return predicate.Loan(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Loan) predicate.Loan {
	return predicate.Loan(func(s *sql.Selector) {
		p(s.Not())
	})
}
