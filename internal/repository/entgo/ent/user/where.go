// Code generated by entc, DO NOT EDIT.

package user

import (
	"time"

	"github.com/DanielTitkov/teams-bot-example/internal/repository/entgo/ent/predicate"
	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.User {
	return predicate.User(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.User {
	return predicate.User(func(s *sql.Selector) {
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
func IDGT(id int) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateTime), v))
	})
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateTime), v))
	})
}

// Username applies equality check predicate on the "username" field. It's identical to UsernameEQ.
func Username(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUsername), v))
	})
}

// Email applies equality check predicate on the "email" field. It's identical to EmailEQ.
func Email(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEmail), v))
	})
}

// PasswordHash applies equality check predicate on the "password_hash" field. It's identical to PasswordHashEQ.
func PasswordHash(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPasswordHash), v))
	})
}

// Service applies equality check predicate on the "service" field. It's identical to ServiceEQ.
func Service(v bool) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldService), v))
	})
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateTime), v))
	})
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreateTime), v))
	})
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...time.Time) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreateTime), v...))
	})
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...time.Time) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreateTime), v...))
	})
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreateTime), v))
	})
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreateTime), v))
	})
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreateTime), v))
	})
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreateTime), v))
	})
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...time.Time) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUpdateTime), v...))
	})
}

// UpdateTimeNotIn applies the NotIn predicate on the "update_time" field.
func UpdateTimeNotIn(vs ...time.Time) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUpdateTime), v...))
	})
}

// UpdateTimeGT applies the GT predicate on the "update_time" field.
func UpdateTimeGT(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdateTime), v))
	})
}

// UsernameEQ applies the EQ predicate on the "username" field.
func UsernameEQ(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUsername), v))
	})
}

// UsernameNEQ applies the NEQ predicate on the "username" field.
func UsernameNEQ(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUsername), v))
	})
}

// UsernameIn applies the In predicate on the "username" field.
func UsernameIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUsername), v...))
	})
}

// UsernameNotIn applies the NotIn predicate on the "username" field.
func UsernameNotIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUsername), v...))
	})
}

// UsernameGT applies the GT predicate on the "username" field.
func UsernameGT(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUsername), v))
	})
}

// UsernameGTE applies the GTE predicate on the "username" field.
func UsernameGTE(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUsername), v))
	})
}

// UsernameLT applies the LT predicate on the "username" field.
func UsernameLT(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUsername), v))
	})
}

// UsernameLTE applies the LTE predicate on the "username" field.
func UsernameLTE(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUsername), v))
	})
}

// UsernameContains applies the Contains predicate on the "username" field.
func UsernameContains(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldUsername), v))
	})
}

// UsernameHasPrefix applies the HasPrefix predicate on the "username" field.
func UsernameHasPrefix(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldUsername), v))
	})
}

// UsernameHasSuffix applies the HasSuffix predicate on the "username" field.
func UsernameHasSuffix(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldUsername), v))
	})
}

// UsernameEqualFold applies the EqualFold predicate on the "username" field.
func UsernameEqualFold(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldUsername), v))
	})
}

// UsernameContainsFold applies the ContainsFold predicate on the "username" field.
func UsernameContainsFold(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldUsername), v))
	})
}

// EmailEQ applies the EQ predicate on the "email" field.
func EmailEQ(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEmail), v))
	})
}

// EmailNEQ applies the NEQ predicate on the "email" field.
func EmailNEQ(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEmail), v))
	})
}

// EmailIn applies the In predicate on the "email" field.
func EmailIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldEmail), v...))
	})
}

// EmailNotIn applies the NotIn predicate on the "email" field.
func EmailNotIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldEmail), v...))
	})
}

// EmailGT applies the GT predicate on the "email" field.
func EmailGT(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEmail), v))
	})
}

// EmailGTE applies the GTE predicate on the "email" field.
func EmailGTE(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEmail), v))
	})
}

// EmailLT applies the LT predicate on the "email" field.
func EmailLT(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEmail), v))
	})
}

// EmailLTE applies the LTE predicate on the "email" field.
func EmailLTE(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEmail), v))
	})
}

// EmailContains applies the Contains predicate on the "email" field.
func EmailContains(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldEmail), v))
	})
}

// EmailHasPrefix applies the HasPrefix predicate on the "email" field.
func EmailHasPrefix(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldEmail), v))
	})
}

// EmailHasSuffix applies the HasSuffix predicate on the "email" field.
func EmailHasSuffix(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldEmail), v))
	})
}

// EmailEqualFold applies the EqualFold predicate on the "email" field.
func EmailEqualFold(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldEmail), v))
	})
}

// EmailContainsFold applies the ContainsFold predicate on the "email" field.
func EmailContainsFold(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldEmail), v))
	})
}

// PasswordHashEQ applies the EQ predicate on the "password_hash" field.
func PasswordHashEQ(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPasswordHash), v))
	})
}

// PasswordHashNEQ applies the NEQ predicate on the "password_hash" field.
func PasswordHashNEQ(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPasswordHash), v))
	})
}

// PasswordHashIn applies the In predicate on the "password_hash" field.
func PasswordHashIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPasswordHash), v...))
	})
}

// PasswordHashNotIn applies the NotIn predicate on the "password_hash" field.
func PasswordHashNotIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPasswordHash), v...))
	})
}

// PasswordHashGT applies the GT predicate on the "password_hash" field.
func PasswordHashGT(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPasswordHash), v))
	})
}

// PasswordHashGTE applies the GTE predicate on the "password_hash" field.
func PasswordHashGTE(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPasswordHash), v))
	})
}

// PasswordHashLT applies the LT predicate on the "password_hash" field.
func PasswordHashLT(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPasswordHash), v))
	})
}

// PasswordHashLTE applies the LTE predicate on the "password_hash" field.
func PasswordHashLTE(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPasswordHash), v))
	})
}

// PasswordHashContains applies the Contains predicate on the "password_hash" field.
func PasswordHashContains(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldPasswordHash), v))
	})
}

// PasswordHashHasPrefix applies the HasPrefix predicate on the "password_hash" field.
func PasswordHashHasPrefix(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldPasswordHash), v))
	})
}

// PasswordHashHasSuffix applies the HasSuffix predicate on the "password_hash" field.
func PasswordHashHasSuffix(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldPasswordHash), v))
	})
}

// PasswordHashEqualFold applies the EqualFold predicate on the "password_hash" field.
func PasswordHashEqualFold(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldPasswordHash), v))
	})
}

// PasswordHashContainsFold applies the ContainsFold predicate on the "password_hash" field.
func PasswordHashContainsFold(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldPasswordHash), v))
	})
}

// ServiceEQ applies the EQ predicate on the "service" field.
func ServiceEQ(v bool) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldService), v))
	})
}

// ServiceNEQ applies the NEQ predicate on the "service" field.
func ServiceNEQ(v bool) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldService), v))
	})
}

// HasDialog applies the HasEdge predicate on the "dialog" edge.
func HasDialog() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(DialogTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, DialogTable, DialogColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDialogWith applies the HasEdge predicate on the "dialog" edge with a given conditions (other predicates).
func HasDialogWith(preds ...predicate.Dialog) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(DialogInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, DialogTable, DialogColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.User) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.User) predicate.User {
	return predicate.User(func(s *sql.Selector) {
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
func Not(p predicate.User) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		p(s.Not())
	})
}
