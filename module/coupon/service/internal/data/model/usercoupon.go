// Code generated by ent, DO NOT EDIT.

package model

import (
	"fmt"
	"mall-go/module/coupon/service/internal/data/model/usercoupon"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// UserCoupon is the model entity for the UserCoupon schema.
type UserCoupon struct {
	config `json:"-"`
	// ID of the ent.
	ID int64 `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// DeleteTime holds the value of the "delete_time" field.
	DeleteTime time.Time `json:"delete_time,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID int64 `json:"user_id,omitempty"`
	// CouponID holds the value of the "coupon_id" field.
	CouponID int64 `json:"coupon_id,omitempty"`
	// 1:未使用，2：已使用， 3：已过期
	Status int `json:"status,omitempty"`
	// OrderID holds the value of the "order_id" field.
	OrderID int `json:"order_id,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*UserCoupon) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case usercoupon.FieldID, usercoupon.FieldUserID, usercoupon.FieldCouponID, usercoupon.FieldStatus, usercoupon.FieldOrderID:
			values[i] = new(sql.NullInt64)
		case usercoupon.FieldCreateTime, usercoupon.FieldUpdateTime, usercoupon.FieldDeleteTime:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type UserCoupon", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the UserCoupon fields.
func (uc *UserCoupon) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case usercoupon.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			uc.ID = int64(value.Int64)
		case usercoupon.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				uc.CreateTime = value.Time
			}
		case usercoupon.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				uc.UpdateTime = value.Time
			}
		case usercoupon.FieldDeleteTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field delete_time", values[i])
			} else if value.Valid {
				uc.DeleteTime = value.Time
			}
		case usercoupon.FieldUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				uc.UserID = value.Int64
			}
		case usercoupon.FieldCouponID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field coupon_id", values[i])
			} else if value.Valid {
				uc.CouponID = value.Int64
			}
		case usercoupon.FieldStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				uc.Status = int(value.Int64)
			}
		case usercoupon.FieldOrderID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field order_id", values[i])
			} else if value.Valid {
				uc.OrderID = int(value.Int64)
			}
		}
	}
	return nil
}

// Update returns a builder for updating this UserCoupon.
// Note that you need to call UserCoupon.Unwrap() before calling this method if this UserCoupon
// was returned from a transaction, and the transaction was committed or rolled back.
func (uc *UserCoupon) Update() *UserCouponUpdateOne {
	return (&UserCouponClient{config: uc.config}).UpdateOne(uc)
}

// Unwrap unwraps the UserCoupon entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (uc *UserCoupon) Unwrap() *UserCoupon {
	_tx, ok := uc.config.driver.(*txDriver)
	if !ok {
		panic("model: UserCoupon is not a transactional entity")
	}
	uc.config.driver = _tx.drv
	return uc
}

// String implements the fmt.Stringer.
func (uc *UserCoupon) String() string {
	var builder strings.Builder
	builder.WriteString("UserCoupon(")
	builder.WriteString(fmt.Sprintf("id=%v, ", uc.ID))
	builder.WriteString("create_time=")
	builder.WriteString(uc.CreateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("update_time=")
	builder.WriteString(uc.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("delete_time=")
	builder.WriteString(uc.DeleteTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", uc.UserID))
	builder.WriteString(", ")
	builder.WriteString("coupon_id=")
	builder.WriteString(fmt.Sprintf("%v", uc.CouponID))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", uc.Status))
	builder.WriteString(", ")
	builder.WriteString("order_id=")
	builder.WriteString(fmt.Sprintf("%v", uc.OrderID))
	builder.WriteByte(')')
	return builder.String()
}

// UserCoupons is a parsable slice of UserCoupon.
type UserCoupons []*UserCoupon

func (uc UserCoupons) config(cfg config) {
	for _i := range uc {
		uc[_i].config = cfg
	}
}
