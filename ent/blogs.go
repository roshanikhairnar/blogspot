// Code generated by ent, DO NOT EDIT.

package ent

import (
	"entdemo/ent/blogs"
	"entdemo/ent/user"
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
)

// Blogs is the model entity for the Blogs schema.
type Blogs struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// BlogTitle holds the value of the "blogTitle" field.
	BlogTitle *string `json:"blogTitle,omitempty"`
	// BlogType holds the value of the "blogType" field.
	BlogType *string `json:"blogType,omitempty"`
	// BlogContent holds the value of the "blogContent" field.
	BlogContent *string `json:"blogContent,omitempty"`
	// BlogAuthor holds the value of the "blogAuthor" field.
	BlogAuthor *string `json:"blogAuthor,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the BlogsQuery when eager-loading is set.
	Edges      BlogsEdges `json:"edges"`
	user_blogs *int
}

// BlogsEdges holds the relations/edges for other nodes in the graph.
type BlogsEdges struct {
	// Owner holds the value of the owner edge.
	Owner *User `json:"owner,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// OwnerOrErr returns the Owner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e BlogsEdges) OwnerOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.Owner == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.Owner, nil
	}
	return nil, &NotLoadedError{edge: "owner"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Blogs) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case blogs.FieldID:
			values[i] = new(sql.NullInt64)
		case blogs.FieldBlogTitle, blogs.FieldBlogType, blogs.FieldBlogContent, blogs.FieldBlogAuthor:
			values[i] = new(sql.NullString)
		case blogs.ForeignKeys[0]: // user_blogs
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Blogs", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Blogs fields.
func (b *Blogs) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case blogs.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			b.ID = int(value.Int64)
		case blogs.FieldBlogTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field blogTitle", values[i])
			} else if value.Valid {
				b.BlogTitle = new(string)
				*b.BlogTitle = value.String
			}
		case blogs.FieldBlogType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field blogType", values[i])
			} else if value.Valid {
				b.BlogType = new(string)
				*b.BlogType = value.String
			}
		case blogs.FieldBlogContent:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field blogContent", values[i])
			} else if value.Valid {
				b.BlogContent = new(string)
				*b.BlogContent = value.String
			}
		case blogs.FieldBlogAuthor:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field blogAuthor", values[i])
			} else if value.Valid {
				b.BlogAuthor = new(string)
				*b.BlogAuthor = value.String
			}
		case blogs.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field user_blogs", value)
			} else if value.Valid {
				b.user_blogs = new(int)
				*b.user_blogs = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryOwner queries the "owner" edge of the Blogs entity.
func (b *Blogs) QueryOwner() *UserQuery {
	return (&BlogsClient{config: b.config}).QueryOwner(b)
}

// Update returns a builder for updating this Blogs.
// Note that you need to call Blogs.Unwrap() before calling this method if this Blogs
// was returned from a transaction, and the transaction was committed or rolled back.
func (b *Blogs) Update() *BlogsUpdateOne {
	return (&BlogsClient{config: b.config}).UpdateOne(b)
}

// Unwrap unwraps the Blogs entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (b *Blogs) Unwrap() *Blogs {
	_tx, ok := b.config.driver.(*txDriver)
	if !ok {
		panic("ent: Blogs is not a transactional entity")
	}
	b.config.driver = _tx.drv
	return b
}

// String implements the fmt.Stringer.
func (b *Blogs) String() string {
	var builder strings.Builder
	builder.WriteString("Blogs(")
	builder.WriteString(fmt.Sprintf("id=%v, ", b.ID))
	if v := b.BlogTitle; v != nil {
		builder.WriteString("blogTitle=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := b.BlogType; v != nil {
		builder.WriteString("blogType=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := b.BlogContent; v != nil {
		builder.WriteString("blogContent=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := b.BlogAuthor; v != nil {
		builder.WriteString("blogAuthor=")
		builder.WriteString(*v)
	}
	builder.WriteByte(')')
	return builder.String()
}

// BlogsSlice is a parsable slice of Blogs.
type BlogsSlice []*Blogs

func (b BlogsSlice) config(cfg config) {
	for _i := range b {
		b[_i].config = cfg
	}
}
