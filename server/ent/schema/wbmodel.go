package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/dialect/entsql"
	"github.com/facebook/ent/schema"
	"github.com/facebook/ent/schema/field"
)

// WBModel holds the schema definition for the WBModel entity.
type WBModel struct {
	ent.Schema
}

// Annotations of the WBModel.
func (WBModel) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "wb"},
	}
}

// Fields of the WBModel.
func (WBModel) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("uid"),
		field.Int64("phone_number"),
	}
}

// Edges of the WBModel.
func (WBModel) Edges() []ent.Edge {
	return nil
}
