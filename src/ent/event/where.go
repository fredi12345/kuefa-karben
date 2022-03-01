// Code generated by entc, DO NOT EDIT.

package event

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/fredi12345/kuefa-karben/src/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
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
func IDNotIn(ids ...uuid.UUID) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
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
func IDGT(id uuid.UUID) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Created applies equality check predicate on the "created" field. It's identical to CreatedEQ.
func Created(v time.Time) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreated), v))
	})
}

// LastModified applies equality check predicate on the "last_modified" field. It's identical to LastModifiedEQ.
func LastModified(v time.Time) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLastModified), v))
	})
}

// Theme applies equality check predicate on the "theme" field. It's identical to ThemeEQ.
func Theme(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTheme), v))
	})
}

// StartingTime applies equality check predicate on the "starting_time" field. It's identical to StartingTimeEQ.
func StartingTime(v time.Time) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStartingTime), v))
	})
}

// ClosingTime applies equality check predicate on the "closing_time" field. It's identical to ClosingTimeEQ.
func ClosingTime(v time.Time) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldClosingTime), v))
	})
}

// Starter applies equality check predicate on the "starter" field. It's identical to StarterEQ.
func Starter(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStarter), v))
	})
}

// MainDish applies equality check predicate on the "main_dish" field. It's identical to MainDishEQ.
func MainDish(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMainDish), v))
	})
}

// Dessert applies equality check predicate on the "dessert" field. It's identical to DessertEQ.
func Dessert(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDessert), v))
	})
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDescription), v))
	})
}

// CreatedEQ applies the EQ predicate on the "created" field.
func CreatedEQ(v time.Time) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreated), v))
	})
}

// CreatedNEQ applies the NEQ predicate on the "created" field.
func CreatedNEQ(v time.Time) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreated), v))
	})
}

// CreatedIn applies the In predicate on the "created" field.
func CreatedIn(vs ...time.Time) predicate.Event {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Event(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreated), v...))
	})
}

// CreatedNotIn applies the NotIn predicate on the "created" field.
func CreatedNotIn(vs ...time.Time) predicate.Event {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Event(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreated), v...))
	})
}

// CreatedGT applies the GT predicate on the "created" field.
func CreatedGT(v time.Time) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreated), v))
	})
}

// CreatedGTE applies the GTE predicate on the "created" field.
func CreatedGTE(v time.Time) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreated), v))
	})
}

// CreatedLT applies the LT predicate on the "created" field.
func CreatedLT(v time.Time) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreated), v))
	})
}

// CreatedLTE applies the LTE predicate on the "created" field.
func CreatedLTE(v time.Time) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreated), v))
	})
}

// LastModifiedEQ applies the EQ predicate on the "last_modified" field.
func LastModifiedEQ(v time.Time) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLastModified), v))
	})
}

// LastModifiedNEQ applies the NEQ predicate on the "last_modified" field.
func LastModifiedNEQ(v time.Time) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLastModified), v))
	})
}

// LastModifiedIn applies the In predicate on the "last_modified" field.
func LastModifiedIn(vs ...time.Time) predicate.Event {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Event(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldLastModified), v...))
	})
}

// LastModifiedNotIn applies the NotIn predicate on the "last_modified" field.
func LastModifiedNotIn(vs ...time.Time) predicate.Event {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Event(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldLastModified), v...))
	})
}

// LastModifiedGT applies the GT predicate on the "last_modified" field.
func LastModifiedGT(v time.Time) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLastModified), v))
	})
}

// LastModifiedGTE applies the GTE predicate on the "last_modified" field.
func LastModifiedGTE(v time.Time) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLastModified), v))
	})
}

// LastModifiedLT applies the LT predicate on the "last_modified" field.
func LastModifiedLT(v time.Time) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLastModified), v))
	})
}

// LastModifiedLTE applies the LTE predicate on the "last_modified" field.
func LastModifiedLTE(v time.Time) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLastModified), v))
	})
}

// ThemeEQ applies the EQ predicate on the "theme" field.
func ThemeEQ(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTheme), v))
	})
}

// ThemeNEQ applies the NEQ predicate on the "theme" field.
func ThemeNEQ(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTheme), v))
	})
}

// ThemeIn applies the In predicate on the "theme" field.
func ThemeIn(vs ...string) predicate.Event {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Event(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldTheme), v...))
	})
}

// ThemeNotIn applies the NotIn predicate on the "theme" field.
func ThemeNotIn(vs ...string) predicate.Event {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Event(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldTheme), v...))
	})
}

// ThemeGT applies the GT predicate on the "theme" field.
func ThemeGT(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTheme), v))
	})
}

// ThemeGTE applies the GTE predicate on the "theme" field.
func ThemeGTE(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTheme), v))
	})
}

// ThemeLT applies the LT predicate on the "theme" field.
func ThemeLT(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTheme), v))
	})
}

// ThemeLTE applies the LTE predicate on the "theme" field.
func ThemeLTE(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTheme), v))
	})
}

// ThemeContains applies the Contains predicate on the "theme" field.
func ThemeContains(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTheme), v))
	})
}

// ThemeHasPrefix applies the HasPrefix predicate on the "theme" field.
func ThemeHasPrefix(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTheme), v))
	})
}

// ThemeHasSuffix applies the HasSuffix predicate on the "theme" field.
func ThemeHasSuffix(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTheme), v))
	})
}

// ThemeEqualFold applies the EqualFold predicate on the "theme" field.
func ThemeEqualFold(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTheme), v))
	})
}

// ThemeContainsFold applies the ContainsFold predicate on the "theme" field.
func ThemeContainsFold(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTheme), v))
	})
}

// StartingTimeEQ applies the EQ predicate on the "starting_time" field.
func StartingTimeEQ(v time.Time) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStartingTime), v))
	})
}

// StartingTimeNEQ applies the NEQ predicate on the "starting_time" field.
func StartingTimeNEQ(v time.Time) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStartingTime), v))
	})
}

// StartingTimeIn applies the In predicate on the "starting_time" field.
func StartingTimeIn(vs ...time.Time) predicate.Event {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Event(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldStartingTime), v...))
	})
}

// StartingTimeNotIn applies the NotIn predicate on the "starting_time" field.
func StartingTimeNotIn(vs ...time.Time) predicate.Event {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Event(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldStartingTime), v...))
	})
}

// StartingTimeGT applies the GT predicate on the "starting_time" field.
func StartingTimeGT(v time.Time) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldStartingTime), v))
	})
}

// StartingTimeGTE applies the GTE predicate on the "starting_time" field.
func StartingTimeGTE(v time.Time) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldStartingTime), v))
	})
}

// StartingTimeLT applies the LT predicate on the "starting_time" field.
func StartingTimeLT(v time.Time) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldStartingTime), v))
	})
}

// StartingTimeLTE applies the LTE predicate on the "starting_time" field.
func StartingTimeLTE(v time.Time) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldStartingTime), v))
	})
}

// ClosingTimeEQ applies the EQ predicate on the "closing_time" field.
func ClosingTimeEQ(v time.Time) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldClosingTime), v))
	})
}

// ClosingTimeNEQ applies the NEQ predicate on the "closing_time" field.
func ClosingTimeNEQ(v time.Time) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldClosingTime), v))
	})
}

// ClosingTimeIn applies the In predicate on the "closing_time" field.
func ClosingTimeIn(vs ...time.Time) predicate.Event {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Event(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldClosingTime), v...))
	})
}

// ClosingTimeNotIn applies the NotIn predicate on the "closing_time" field.
func ClosingTimeNotIn(vs ...time.Time) predicate.Event {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Event(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldClosingTime), v...))
	})
}

// ClosingTimeGT applies the GT predicate on the "closing_time" field.
func ClosingTimeGT(v time.Time) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldClosingTime), v))
	})
}

// ClosingTimeGTE applies the GTE predicate on the "closing_time" field.
func ClosingTimeGTE(v time.Time) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldClosingTime), v))
	})
}

// ClosingTimeLT applies the LT predicate on the "closing_time" field.
func ClosingTimeLT(v time.Time) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldClosingTime), v))
	})
}

// ClosingTimeLTE applies the LTE predicate on the "closing_time" field.
func ClosingTimeLTE(v time.Time) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldClosingTime), v))
	})
}

// StarterEQ applies the EQ predicate on the "starter" field.
func StarterEQ(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStarter), v))
	})
}

// StarterNEQ applies the NEQ predicate on the "starter" field.
func StarterNEQ(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStarter), v))
	})
}

// StarterIn applies the In predicate on the "starter" field.
func StarterIn(vs ...string) predicate.Event {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Event(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldStarter), v...))
	})
}

// StarterNotIn applies the NotIn predicate on the "starter" field.
func StarterNotIn(vs ...string) predicate.Event {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Event(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldStarter), v...))
	})
}

// StarterGT applies the GT predicate on the "starter" field.
func StarterGT(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldStarter), v))
	})
}

// StarterGTE applies the GTE predicate on the "starter" field.
func StarterGTE(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldStarter), v))
	})
}

// StarterLT applies the LT predicate on the "starter" field.
func StarterLT(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldStarter), v))
	})
}

// StarterLTE applies the LTE predicate on the "starter" field.
func StarterLTE(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldStarter), v))
	})
}

// StarterContains applies the Contains predicate on the "starter" field.
func StarterContains(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldStarter), v))
	})
}

// StarterHasPrefix applies the HasPrefix predicate on the "starter" field.
func StarterHasPrefix(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldStarter), v))
	})
}

// StarterHasSuffix applies the HasSuffix predicate on the "starter" field.
func StarterHasSuffix(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldStarter), v))
	})
}

// StarterEqualFold applies the EqualFold predicate on the "starter" field.
func StarterEqualFold(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldStarter), v))
	})
}

// StarterContainsFold applies the ContainsFold predicate on the "starter" field.
func StarterContainsFold(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldStarter), v))
	})
}

// MainDishEQ applies the EQ predicate on the "main_dish" field.
func MainDishEQ(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMainDish), v))
	})
}

// MainDishNEQ applies the NEQ predicate on the "main_dish" field.
func MainDishNEQ(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldMainDish), v))
	})
}

// MainDishIn applies the In predicate on the "main_dish" field.
func MainDishIn(vs ...string) predicate.Event {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Event(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldMainDish), v...))
	})
}

// MainDishNotIn applies the NotIn predicate on the "main_dish" field.
func MainDishNotIn(vs ...string) predicate.Event {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Event(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldMainDish), v...))
	})
}

// MainDishGT applies the GT predicate on the "main_dish" field.
func MainDishGT(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldMainDish), v))
	})
}

// MainDishGTE applies the GTE predicate on the "main_dish" field.
func MainDishGTE(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldMainDish), v))
	})
}

// MainDishLT applies the LT predicate on the "main_dish" field.
func MainDishLT(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldMainDish), v))
	})
}

// MainDishLTE applies the LTE predicate on the "main_dish" field.
func MainDishLTE(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldMainDish), v))
	})
}

// MainDishContains applies the Contains predicate on the "main_dish" field.
func MainDishContains(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldMainDish), v))
	})
}

// MainDishHasPrefix applies the HasPrefix predicate on the "main_dish" field.
func MainDishHasPrefix(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldMainDish), v))
	})
}

// MainDishHasSuffix applies the HasSuffix predicate on the "main_dish" field.
func MainDishHasSuffix(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldMainDish), v))
	})
}

// MainDishEqualFold applies the EqualFold predicate on the "main_dish" field.
func MainDishEqualFold(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldMainDish), v))
	})
}

// MainDishContainsFold applies the ContainsFold predicate on the "main_dish" field.
func MainDishContainsFold(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldMainDish), v))
	})
}

// DessertEQ applies the EQ predicate on the "dessert" field.
func DessertEQ(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDessert), v))
	})
}

// DessertNEQ applies the NEQ predicate on the "dessert" field.
func DessertNEQ(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDessert), v))
	})
}

// DessertIn applies the In predicate on the "dessert" field.
func DessertIn(vs ...string) predicate.Event {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Event(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDessert), v...))
	})
}

// DessertNotIn applies the NotIn predicate on the "dessert" field.
func DessertNotIn(vs ...string) predicate.Event {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Event(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDessert), v...))
	})
}

// DessertGT applies the GT predicate on the "dessert" field.
func DessertGT(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDessert), v))
	})
}

// DessertGTE applies the GTE predicate on the "dessert" field.
func DessertGTE(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDessert), v))
	})
}

// DessertLT applies the LT predicate on the "dessert" field.
func DessertLT(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDessert), v))
	})
}

// DessertLTE applies the LTE predicate on the "dessert" field.
func DessertLTE(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDessert), v))
	})
}

// DessertContains applies the Contains predicate on the "dessert" field.
func DessertContains(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldDessert), v))
	})
}

// DessertHasPrefix applies the HasPrefix predicate on the "dessert" field.
func DessertHasPrefix(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldDessert), v))
	})
}

// DessertHasSuffix applies the HasSuffix predicate on the "dessert" field.
func DessertHasSuffix(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldDessert), v))
	})
}

// DessertEqualFold applies the EqualFold predicate on the "dessert" field.
func DessertEqualFold(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldDessert), v))
	})
}

// DessertContainsFold applies the ContainsFold predicate on the "dessert" field.
func DessertContainsFold(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldDessert), v))
	})
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDescription), v))
	})
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDescription), v))
	})
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.Event {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Event(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDescription), v...))
	})
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.Event {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Event(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDescription), v...))
	})
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDescription), v))
	})
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDescription), v))
	})
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDescription), v))
	})
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDescription), v))
	})
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldDescription), v))
	})
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldDescription), v))
	})
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldDescription), v))
	})
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldDescription), v))
	})
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldDescription), v))
	})
}

// HasParticipants applies the HasEdge predicate on the "participants" edge.
func HasParticipants() predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ParticipantsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ParticipantsTable, ParticipantsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasParticipantsWith applies the HasEdge predicate on the "participants" edge with a given conditions (other predicates).
func HasParticipantsWith(preds ...predicate.Participant) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ParticipantsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ParticipantsTable, ParticipantsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasComments applies the HasEdge predicate on the "comments" edge.
func HasComments() predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CommentsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, CommentsTable, CommentsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCommentsWith applies the HasEdge predicate on the "comments" edge with a given conditions (other predicates).
func HasCommentsWith(preds ...predicate.Comment) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CommentsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, CommentsTable, CommentsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasImages applies the HasEdge predicate on the "images" edge.
func HasImages() predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ImagesTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ImagesTable, ImagesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasImagesWith applies the HasEdge predicate on the "images" edge with a given conditions (other predicates).
func HasImagesWith(preds ...predicate.Image) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ImagesInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ImagesTable, ImagesColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasTitleImage applies the HasEdge predicate on the "title_image" edge.
func HasTitleImage() predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TitleImageTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, TitleImageTable, TitleImageColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTitleImageWith applies the HasEdge predicate on the "title_image" edge with a given conditions (other predicates).
func HasTitleImageWith(preds ...predicate.TitleImage) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TitleImageInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, TitleImageTable, TitleImageColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Event) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Event) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
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
func Not(p predicate.Event) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		p(s.Not())
	})
}
