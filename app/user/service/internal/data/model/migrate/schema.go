// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// LinUserIdentiyColumns holds the columns for the "lin_user_identiy" table.
	LinUserIdentiyColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "delete_time", Type: field.TypeTime, Nullable: true},
		{Name: "user_id", Type: field.TypeInt64},
		{Name: "identity_type", Type: field.TypeString},
		{Name: "identifier", Type: field.TypeString},
		{Name: "credential", Type: field.TypeString},
	}
	// LinUserIdentiyTable holds the schema information for the "lin_user_identiy" table.
	LinUserIdentiyTable = &schema.Table{
		Name:       "lin_user_identiy",
		Columns:    LinUserIdentiyColumns,
		PrimaryKey: []*schema.Column{LinUserIdentiyColumns[0]},
	}
	// UserInfoColumns holds the columns for the "user_info" table.
	UserInfoColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "delete_time", Type: field.TypeTime, Nullable: true},
		{Name: "nickname", Type: field.TypeString},
		{Name: "phone", Type: field.TypeString},
		{Name: "avatar", Type: field.TypeString},
		{Name: "wx_profile", Type: field.TypeString},
	}
	// UserInfoTable holds the schema information for the "user_info" table.
	UserInfoTable = &schema.Table{
		Name:       "user_info",
		Columns:    UserInfoColumns,
		PrimaryKey: []*schema.Column{UserInfoColumns[0]},
	}
	// UserPointColumns holds the columns for the "user_point" table.
	UserPointColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "delete_time", Type: field.TypeTime, Nullable: true},
		{Name: "value", Type: field.TypeInt},
		{Name: "status", Type: field.TypeInt},
	}
	// UserPointTable holds the schema information for the "user_point" table.
	UserPointTable = &schema.Table{
		Name:       "user_point",
		Columns:    UserPointColumns,
		PrimaryKey: []*schema.Column{UserPointColumns[0]},
	}
	// UserPointDetailColumns holds the columns for the "user_point_detail" table.
	UserPointDetailColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "delete_time", Type: field.TypeTime, Nullable: true},
		{Name: "user_id", Type: field.TypeInt64},
		{Name: "value", Type: field.TypeInt},
	}
	// UserPointDetailTable holds the schema information for the "user_point_detail" table.
	UserPointDetailTable = &schema.Table{
		Name:       "user_point_detail",
		Columns:    UserPointDetailColumns,
		PrimaryKey: []*schema.Column{UserPointDetailColumns[0]},
	}
	// UserWalletColumns holds the columns for the "user_wallet" table.
	UserWalletColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "delete_time", Type: field.TypeTime, Nullable: true},
		{Name: "value", Type: field.TypeInt},
	}
	// UserWalletTable holds the schema information for the "user_wallet" table.
	UserWalletTable = &schema.Table{
		Name:       "user_wallet",
		Columns:    UserWalletColumns,
		PrimaryKey: []*schema.Column{UserWalletColumns[0]},
	}
	// UserWalletDetailColumns holds the columns for the "user_wallet_detail" table.
	UserWalletDetailColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "delete_time", Type: field.TypeTime, Nullable: true},
		{Name: "description", Type: field.TypeString},
		{Name: "op", Type: field.TypeInt, Default: 1},
		{Name: "current", Type: field.TypeInt},
		{Name: "value", Type: field.TypeInt},
		{Name: "type", Type: field.TypeInt, Default: 1},
	}
	// UserWalletDetailTable holds the schema information for the "user_wallet_detail" table.
	UserWalletDetailTable = &schema.Table{
		Name:       "user_wallet_detail",
		Columns:    UserWalletDetailColumns,
		PrimaryKey: []*schema.Column{UserWalletDetailColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		LinUserIdentiyTable,
		UserInfoTable,
		UserPointTable,
		UserPointDetailTable,
		UserWalletTable,
		UserWalletDetailTable,
	}
)

func init() {
	LinUserIdentiyTable.Annotation = &entsql.Annotation{
		Table: "lin_user_identiy",
	}
	UserInfoTable.Annotation = &entsql.Annotation{
		Table: "user_info",
	}
	UserPointDetailTable.Annotation = &entsql.Annotation{
		Table: "user_point_detail",
	}
}
