package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Blogs holds the schema definition for the Blogs entity.
type Blogs struct {
	ent.Schema
}

// blogID int,blogType char(50),blogTitle char(50), blogContent text,blogAuthor int
// Fields of the Blogs.
func (Blogs) Fields() []ent.Field {
	return []ent.Field{
		field.String("blogTitle").Nillable().Optional(),
		field.String("blogType").Nillable().Optional(),
		field.String("blogContent").Nillable().Optional(),
		field.String("blogAuthor").Nillable().Optional(),
	}
}

// Edges of the Blogs.
func (Blogs) Edges() []ent.Edge {
	return nil
}
