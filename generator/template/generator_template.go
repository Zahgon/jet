package template

import (
	"github.com/go-jet/jet/v2/generator/metadata"
	"github.com/go-jet/jet/v2/internal/jet"
)

// Template is generator template used for file generation
type Template struct {
	Dialect jet.Dialect
	Schema  func(schemaMetaData metadata.Schema) Schema
}

// Default is default generator template implementation
func Default(dialect jet.Dialect) Template { _ = "STUB: not implemented"; return *new(Template) }

// UseSchema replaces current schema generate function with a new implementation and returns new generator template
func (t Template) UseSchema(schemaFunc func(schemaMetaData metadata.Schema) Schema) Template {
	_ = "STUB: not implemented"
	return *new(Template)
}

// Schema is schema generator template used to generate schema(model and sql builder) files
type Schema struct {
	Path       string
	Model      Model
	SQLBuilder SQLBuilder
}

// UsePath replaces path and returns new schema template
func (s Schema) UsePath(path string) Schema {
	_ = "STUB: not implemented"
	return *

	// UseModel returns new schema template with replaced template for model files generation
	new(Schema)
}

func (s Schema) UseModel(model Model) Schema {
	_ = "STUB: not implemented"
	return *

	// UseSQLBuilder returns new schema with replaced template for sql builder files generation
	new(Schema)
}

func (s Schema) UseSQLBuilder(sqlBuilder SQLBuilder) Schema {
	_ = "STUB: not implemented"
	return *new(Schema)
}

// DefaultSchema returns default schema template implementation
func DefaultSchema(schemaMetaData metadata.Schema) Schema {
	_ = "STUB: not implemented"
	return *new(Schema)
}
