// Code generated by entc, DO NOT EDIT.

package model

import (
	"fmt"
	"strings"
	"time"

	"mall-go/app/payment/service/internal/data/model/coupontype"

	"entgo.io/ent/dialect/sql"
)

// CouponType is the model entity for the CouponType schema.
type CouponType struct {
	config `json:"-"`
	// ID of the ent.
	ID int64 `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// DeleteTime holds the value of the "delete_time" field.
	DeleteTime time.Time `json:"delete_time,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Code holds the value of the "code" field.
	Code int `json:"code,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*CouponType) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case coupontype.FieldID, coupontype.FieldCode:
			values[i] = new(sql.NullInt64)
		case coupontype.FieldName, coupontype.FieldDescription:
			values[i] = new(sql.NullString)
		case coupontype.FieldCreateTime, coupontype.FieldUpdateTime, coupontype.FieldDeleteTime:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type CouponType", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the CouponType fields.
func (ct *CouponType) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case coupontype.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ct.ID = int64(value.Int64)
		case coupontype.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				ct.CreateTime = value.Time
			}
		case coupontype.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				ct.UpdateTime = value.Time
			}
		case coupontype.FieldDeleteTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field delete_time", values[i])
			} else if value.Valid {
				ct.DeleteTime = value.Time
			}
		case coupontype.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				ct.Name = value.String
			}
		case coupontype.FieldCode:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field code", values[i])
			} else if value.Valid {
				ct.Code = int(value.Int64)
			}
		case coupontype.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				ct.Description = value.String
			}
		}
	}
	return nil
}

// Update returns a builder for updating this CouponType.
// Note that you need to call CouponType.Unwrap() before calling this method if this CouponType
// was returned from a transaction, and the transaction was committed or rolled back.
func (ct *CouponType) Update() *CouponTypeUpdateOne {
	return (&CouponTypeClient{config: ct.config}).UpdateOne(ct)
}

// Unwrap unwraps the CouponType entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ct *CouponType) Unwrap() *CouponType {
	tx, ok := ct.config.driver.(*txDriver)
	if !ok {
		panic("model: CouponType is not a transactional entity")
	}
	ct.config.driver = tx.drv
	return ct
}

// String implements the fmt.Stringer.
func (ct *CouponType) String() string {
	var builder strings.Builder
	builder.WriteString("CouponType(")
	builder.WriteString(fmt.Sprintf("id=%v", ct.ID))
	builder.WriteString(", create_time=")
	builder.WriteString(ct.CreateTime.Format(time.ANSIC))
	builder.WriteString(", update_time=")
	builder.WriteString(ct.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", delete_time=")
	builder.WriteString(ct.DeleteTime.Format(time.ANSIC))
	builder.WriteString(", name=")
	builder.WriteString(ct.Name)
	builder.WriteString(", code=")
	builder.WriteString(fmt.Sprintf("%v", ct.Code))
	builder.WriteString(", description=")
	builder.WriteString(ct.Description)
	builder.WriteByte(')')
	return builder.String()
}

// CouponTypes is a parsable slice of CouponType.
type CouponTypes []*CouponType

func (ct CouponTypes) config(cfg config) {
	for _i := range ct {
		ct[_i].config = cfg
	}
}
