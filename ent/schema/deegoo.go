package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Deegoo holds the schema definition for the Deegoo entity.
type Deegoo struct {
	ent.Schema
}

// Fields of the Deegoo.
func (Deegoo) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Immutable().
			Default(uuid.New).
			Unique(),
		field.Int8("perception").Min(0).Max(100).StructTag(`json:"perception"`),
		field.Int8("focus").Min(0).Max(100).StructTag(`json:"focus"`),
		field.Int8("execution").Min(0).Max(100).StructTag(`json:"execution"`),
		field.Int8("memory").Min(0).Max(100).StructTag(`json:"memory"`),
		field.Int8("language").Min(0).Max(100).StructTag(`json:"language"`),
		field.Time("created_at").
			Default(time.Now),
	}
}

// Edges of the Deegoo.
func (Deegoo) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("deegoo").
			Unique(),
	}
}
