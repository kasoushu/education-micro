package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"time"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Default(""),
		field.String("phone").Default(""),
		field.String("password"),
		field.Bool("is_admin").Default(false).Immutable(),
		field.Bool("is_student").Default(false).Immutable(),
		field.Bool("is_teacher").Default(false).Immutable(),
		field.Time("created").Default(time.Now().Local),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
