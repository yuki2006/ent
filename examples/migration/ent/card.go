// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/examples/migration/ent/card"
	"entgo.io/ent/examples/migration/ent/user"
)

// Card is the model entity for the Card schema.
type Card struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Type holds the value of the "type" field.
	Type string `json:"type,omitempty"`
	// NumberHash holds the value of the "number_hash" field.
	NumberHash string `json:"number_hash,omitempty"`
	// CvvHash holds the value of the "cvv_hash" field.
	CvvHash string `json:"cvv_hash,omitempty"`
	// ExpiresAt holds the value of the "expires_at" field.
	ExpiresAt time.Time `json:"expires_at,omitempty"`
	// OwnerID holds the value of the "owner_id" field.
	OwnerID int `json:"owner_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CardQuery when eager-loading is set.
	Edges        CardEdges `json:"edges"`
	selectValues sql.SelectValues
}

// CardEdges holds the relations/edges for other nodes in the graph.
type CardEdges struct {
	// Owner holds the value of the owner edge.
	Owner *User `json:"owner,omitempty"`
	// Payments holds the value of the payments edge.
	Payments []*Payment `json:"payments,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// OwnerOrErr returns the Owner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CardEdges) OwnerOrErr() (*User, error) {
	if e.Owner != nil {
		return e.Owner, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: user.Label}
	}
	return nil, &NotLoadedError{edge: "owner"}
}

// PaymentsOrErr returns the Payments value or an error if the edge
// was not loaded in eager-loading.
func (e CardEdges) PaymentsOrErr() ([]*Payment, error) {
	if e.loadedTypes[1] {
		return e.Payments, nil
	}
	return nil, &NotLoadedError{edge: "payments"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Card) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case card.FieldID, card.FieldOwnerID:
			values[i] = new(sql.NullInt64)
		case card.FieldType, card.FieldNumberHash, card.FieldCvvHash:
			values[i] = new(sql.NullString)
		case card.FieldExpiresAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Card fields.
func (c *Card) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case card.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			c.ID = int(value.Int64)
		case card.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				c.Type = value.String
			}
		case card.FieldNumberHash:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field number_hash", values[i])
			} else if value.Valid {
				c.NumberHash = value.String
			}
		case card.FieldCvvHash:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field cvv_hash", values[i])
			} else if value.Valid {
				c.CvvHash = value.String
			}
		case card.FieldExpiresAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field expires_at", values[i])
			} else if value.Valid {
				c.ExpiresAt = value.Time
			}
		case card.FieldOwnerID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field owner_id", values[i])
			} else if value.Valid {
				c.OwnerID = int(value.Int64)
			}
		default:
			c.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Card.
// This includes values selected through modifiers, order, etc.
func (c *Card) Value(name string) (ent.Value, error) {
	return c.selectValues.Get(name)
}

// QueryOwner queries the "owner" edge of the Card entity.
func (c *Card) QueryOwner() *UserQuery {
	return NewCardClient(c.config).QueryOwner(c)
}

// QueryPayments queries the "payments" edge of the Card entity.
func (c *Card) QueryPayments() *PaymentQuery {
	return NewCardClient(c.config).QueryPayments(c)
}

// Update returns a builder for updating this Card.
// Note that you need to call Card.Unwrap() before calling this method if this Card
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Card) Update() *CardUpdateOne {
	return NewCardClient(c.config).UpdateOne(c)
}

// Unwrap unwraps the Card entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Card) Unwrap() *Card {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Card is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Card) String() string {
	var builder strings.Builder
	builder.WriteString("Card(")
	builder.WriteString(fmt.Sprintf("id=%v, ", c.ID))
	builder.WriteString("type=")
	builder.WriteString(c.Type)
	builder.WriteString(", ")
	builder.WriteString("number_hash=")
	builder.WriteString(c.NumberHash)
	builder.WriteString(", ")
	builder.WriteString("cvv_hash=")
	builder.WriteString(c.CvvHash)
	builder.WriteString(", ")
	builder.WriteString("expires_at=")
	builder.WriteString(c.ExpiresAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("owner_id=")
	builder.WriteString(fmt.Sprintf("%v", c.OwnerID))
	builder.WriteByte(')')
	return builder.String()
}

// Cards is a parsable slice of Card.
type Cards []*Card
