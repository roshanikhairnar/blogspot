// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// BlogsColumns holds the columns for the "blogs" table.
	BlogsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "blog_title", Type: field.TypeString, Nullable: true},
		{Name: "blog_type", Type: field.TypeString, Nullable: true},
		{Name: "blog_content", Type: field.TypeString, Nullable: true},
		{Name: "blog_author", Type: field.TypeString, Nullable: true},
		{Name: "user_blogs", Type: field.TypeInt, Nullable: true},
	}
	// BlogsTable holds the schema information for the "blogs" table.
	BlogsTable = &schema.Table{
		Name:       "blogs",
		Columns:    BlogsColumns,
		PrimaryKey: []*schema.Column{BlogsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "blogs_users_blogs",
				Columns:    []*schema.Column{BlogsColumns[5]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "type", Type: field.TypeString},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		BlogsTable,
		UsersTable,
	}
)

func init() {
	BlogsTable.ForeignKeys[0].RefTable = UsersTable
}
