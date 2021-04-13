// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/go-kratos/beer-shop/apps/user/internal/data/ent/address"
	"github.com/go-kratos/beer-shop/apps/user/internal/data/ent/user"
)

// Address is the model entity for the Address schema.
type Address struct {
	config `json:"-"`
	// ID of the ent.
	ID int64 `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Mobile holds the value of the "mobile" field.
	Mobile string `json:"mobile,omitempty"`
	// Country holds the value of the "country" field.
	Country string `json:"country,omitempty"`
	// City holds the value of the "city" field.
	City string `json:"city,omitempty"`
	// Address holds the value of the "address" field.
	Address string `json:"address,omitempty"`
	// PostCode holds the value of the "post_code" field.
	PostCode string `json:"post_code,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the AddressQuery when eager-loading is set.
	Edges          AddressEdges `json:"edges"`
	user_addresses *int64
}

// AddressEdges holds the relations/edges for other nodes in the graph.
type AddressEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e AddressEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.User == nil {
			// The edge user was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Address) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case address.FieldID:
			values[i] = &sql.NullInt64{}
		case address.FieldName, address.FieldMobile, address.FieldCountry, address.FieldCity, address.FieldAddress, address.FieldPostCode:
			values[i] = &sql.NullString{}
		case address.FieldCreatedAt, address.FieldUpdatedAt:
			values[i] = &sql.NullTime{}
		case address.ForeignKeys[0]: // user_addresses
			values[i] = &sql.NullInt64{}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Address", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Address fields.
func (a *Address) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case address.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			a.ID = int64(value.Int64)
		case address.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				a.Name = value.String
			}
		case address.FieldMobile:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field mobile", values[i])
			} else if value.Valid {
				a.Mobile = value.String
			}
		case address.FieldCountry:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field country", values[i])
			} else if value.Valid {
				a.Country = value.String
			}
		case address.FieldCity:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field city", values[i])
			} else if value.Valid {
				a.City = value.String
			}
		case address.FieldAddress:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field address", values[i])
			} else if value.Valid {
				a.Address = value.String
			}
		case address.FieldPostCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field post_code", values[i])
			} else if value.Valid {
				a.PostCode = value.String
			}
		case address.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				a.CreatedAt = value.Time
			}
		case address.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				a.UpdatedAt = value.Time
			}
		case address.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field user_addresses", value)
			} else if value.Valid {
				a.user_addresses = new(int64)
				*a.user_addresses = int64(value.Int64)
			}
		}
	}
	return nil
}

// QueryUser queries the "user" edge of the Address entity.
func (a *Address) QueryUser() *UserQuery {
	return (&AddressClient{config: a.config}).QueryUser(a)
}

// Update returns a builder for updating this Address.
// Note that you need to call Address.Unwrap() before calling this method if this Address
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Address) Update() *AddressUpdateOne {
	return (&AddressClient{config: a.config}).UpdateOne(a)
}

// Unwrap unwraps the Address entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (a *Address) Unwrap() *Address {
	tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("ent: Address is not a transactional entity")
	}
	a.config.driver = tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Address) String() string {
	var builder strings.Builder
	builder.WriteString("Address(")
	builder.WriteString(fmt.Sprintf("id=%v", a.ID))
	builder.WriteString(", name=")
	builder.WriteString(a.Name)
	builder.WriteString(", mobile=")
	builder.WriteString(a.Mobile)
	builder.WriteString(", country=")
	builder.WriteString(a.Country)
	builder.WriteString(", city=")
	builder.WriteString(a.City)
	builder.WriteString(", address=")
	builder.WriteString(a.Address)
	builder.WriteString(", post_code=")
	builder.WriteString(a.PostCode)
	builder.WriteString(", created_at=")
	builder.WriteString(a.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(a.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Addresses is a parsable slice of Address.
type Addresses []*Address

func (a Addresses) config(cfg config) {
	for _i := range a {
		a[_i].config = cfg
	}
}
