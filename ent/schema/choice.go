package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Choice holds the schema definition for the Choice entity.
type Choice struct {
	ent.Schema
}

// Fields of the Choice.
func (Choice) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Unique().
			Default(uuid.New),
		field.Text("body").
			NotEmpty(),
		field.Int("order").
			NonNegative(),
	}
}

// Edges of the Choice.
func (Choice) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("quesion", Question.Type).
			Annotations(entsql.OnDelete(entsql.Cascade)).
			Ref("choices").
			Unique().
			Required(),
		edge.From("answer", Answer.Type).
			Annotations(entsql.OnDelete(entsql.Cascade)).
			Ref("chosen"),
	}
}
