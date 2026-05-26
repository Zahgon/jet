package template

import (
	"text/template"

	"github.com/go-jet/jet/v2/generator/metadata"
	"github.com/go-jet/jet/v2/internal/jet"
)

// ProcessSchema will process schema metadata and constructs go files using generator Template
func ProcessSchema(dirPath string, schemaMetaData metadata.Schema, generatorTemplate Template) error {
	_ = "STUB: not implemented"
	return nil
}

func processModel(dirPath string, schemaMetaData metadata.Schema, schemaTemplate Schema) error {
	_ = "STUB: not implemented"
	return nil
}

func processSQLBuilder(dirPath string, dialect jet.Dialect, schemaMetaData metadata.Schema, schemaTemplate Schema) error {
	_ = "STUB: not implemented"
	return nil
}

func processEnumSQLBuilder(dirPath string, dialect jet.Dialect, enumsMetaData []metadata.Enum, sqlBuilder SQLBuilder) error {
	_ = "STUB: not implemented"
	return nil
}

func processTableSQLBuilder(fileTypes, dirPath string,
	dialect jet.Dialect,
	schemaMetaData metadata.Schema,
	tablesMetaData []metadata.Table,
	sqlBuilderTemplate SQLBuilder) error {
	_ = "STUB: not implemented"
	return nil
}

// postgres only

func generateUseSchemaFunc(dirPath, fileTypes string, builders []TableSQLBuilder) error {
	_ = "STUB: not implemented"
	return nil
}

func insertedRowAlias(dialect jet.Dialect) string { _ = "STUB: not implemented"; return "" }

func processTableModels(fileTypes, modelDirPath string, tablesMetaData []metadata.Table, modelTemplate Model) error {
	_ = "STUB: not implemented"
	return nil
}

func processEnumModels(modelDir string, enumsMetaData []metadata.Enum, modelTemplate Model) error {
	_ = "STUB: not implemented"
	return nil
}

func generateTemplate(templateText string, templateData interface{}, funcMap template.FuncMap) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
