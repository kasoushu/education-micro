package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"time"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("uid", uuid.UUID{}).Default(uuid.New).Unique(),
		field.String("name").Default(""),
		field.String("password"),
		field.Bool("is_admin"),
		field.Bool("is_student"),
		field.Bool("is_teacher"),
		field.Time("created").Default(time.Now().Local),
		field.String("phone").Default(""),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
