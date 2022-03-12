// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/fredi12345/kuefa-karben/src/ent/titleimage"
	"github.com/google/uuid"
)

// TitleImage is the model entity for the TitleImage schema.
type TitleImage struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Created holds the value of the "created" field.
	Created time.Time `json:"created,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*TitleImage) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case titleimage.FieldCreated:
			values[i] = new(sql.NullTime)
		case titleimage.FieldID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type TitleImage", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the TitleImage fields.
func (ti *TitleImage) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case titleimage.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				ti.ID = *value
			}
		case titleimage.FieldCreated:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created", values[i])
			} else if value.Valid {
				ti.Created = value.Time
			}
		}
	}
	return nil
}

// Update returns a builder for updating this TitleImage.
// Note that you need to call TitleImage.Unwrap() before calling this method if this TitleImage
// was returned from a transaction, and the transaction was committed or rolled back.
func (ti *TitleImage) Update() *TitleImageUpdateOne {
	return (&TitleImageClient{config: ti.config}).UpdateOne(ti)
}

// Unwrap unwraps the TitleImage entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ti *TitleImage) Unwrap() *TitleImage {
	tx, ok := ti.config.driver.(*txDriver)
	if !ok {
		panic("ent: TitleImage is not a transactional entity")
	}
	ti.config.driver = tx.drv
	return ti
}

// String implements the fmt.Stringer.
func (ti *TitleImage) String() string {
	var builder strings.Builder
	builder.WriteString("TitleImage(")
	builder.WriteString(fmt.Sprintf("id=%v", ti.ID))
	builder.WriteString(", created=")
	builder.WriteString(ti.Created.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// TitleImages is a parsable slice of TitleImage.
type TitleImages []*TitleImage

func (ti TitleImages) config(cfg config) {
	for _i := range ti {
		ti[_i].config = cfg
	}
}