package template

import (
	"github.com/go-jet/jet/v2/generator/metadata"
)

// SQLBuilder is template for generating sql builder files
type SQLBuilder struct {
	Skip  bool
	Path  string
	Table func(table metadata.Table) TableSQLBuilder
	View  func(view metadata.Table) TableSQLBuilder
	Enum  func(enum metadata.Enum) EnumSQLBuilder
}

// DefaultSQLBuilder returns default SQLBuilder implementation
func DefaultSQLBuilder() SQLBuilder { _ = "STUB: not implemented"; return *new(SQLBuilder) }

// UsePath returns new SQLBuilder with new relative path set
func (sb SQLBuilder) UsePath(path string) SQLBuilder {
	_ = "STUB: not implemented"
	return *new(SQLBuilder)
}

// UseTable returns new SQLBuilder with new TableSQLBuilder template function set
func (sb SQLBuilder) UseTable(tableFunc func(table metadata.Table) TableSQLBuilder) SQLBuilder {
	_ = "STUB: not implemented"
	return *new(SQLBuilder)
}

// UseView returns new SQLBuilder with new ViewSQLBuilder template function set
func (sb SQLBuilder) UseView(viewFunc func(table metadata.Table) ViewSQLBuilder) SQLBuilder {
	_ = "STUB: not implemented"
	return *new(SQLBuilder)
}

// UseEnum returns new SQLBuilder with new EnumSQLBuilder template function set
func (sb SQLBuilder) UseEnum(enumFunc func(enum metadata.Enum) EnumSQLBuilder) SQLBuilder {
	_ = "STUB: not implemented"
	return *new(SQLBuilder)
}

// ShouldSkip returns new SQLBuilder with new skip flag set
func (sb SQLBuilder) ShouldSkip(skip bool) SQLBuilder {
	_ = "STUB: not implemented"
	return *new(SQLBuilder)
}

// TableSQLBuilder is template for generating table SQLBuilder files
type TableSQLBuilder struct {
	Skip         bool
	Path         string
	FileName     string
	InstanceName string
	TypeName     string
	DefaultAlias string
	Column       func(columnMetaData metadata.Column) TableSQLBuilderColumn
}

// ViewSQLBuilder is template for generating view SQLBuilder files
type ViewSQLBuilder = TableSQLBuilder

// DefaultTableSQLBuilder returns default implementation for TableSQLBuilder
func DefaultTableSQLBuilder(tableMetaData metadata.Table) TableSQLBuilder {
	_ = "STUB: not implemented"
	return *new(TableSQLBuilder)
}

// DefaultViewSQLBuilder returns default implementation for ViewSQLBuilder
func DefaultViewSQLBuilder(viewMetaData metadata.Table) ViewSQLBuilder {
	_ = "STUB: not implemented"
	return *new(ViewSQLBuilder)
}

// PackageName returns package name of table sql builder types
func (tb TableSQLBuilder) PackageName() string { _ = "STUB: not implemented"; return "" }

// UsePath returns new TableSQLBuilder with new relative path set
func (tb TableSQLBuilder) UsePath(path string) TableSQLBuilder {
	_ = "STUB: not implemented"
	return *new(TableSQLBuilder)
}

// UseFileName returns new TableSQLBuilder with new file name set
func (tb TableSQLBuilder) UseFileName(name string) TableSQLBuilder {
	_ = "STUB: not implemented"
	return *new(TableSQLBuilder)
}

// UseInstanceName returns new TableSQLBuilder with new instance name set
func (tb TableSQLBuilder) UseInstanceName(name string) TableSQLBuilder {
	_ = "STUB: not implemented"
	return *new(TableSQLBuilder)
}

// UseTypeName returns new TableSQLBuilder with new type name set
func (tb TableSQLBuilder) UseTypeName(name string) TableSQLBuilder {
	_ = "STUB: not implemented"
	return *new(TableSQLBuilder)
}

// UseDefaultAlias returns new TableSQLBuilder with new default alias set
func (tb TableSQLBuilder) UseDefaultAlias(defaultAlias string) TableSQLBuilder {
	_ = "STUB: not implemented"
	return *new(TableSQLBuilder)
}

// UseColumn returns new TableSQLBuilder with new column template function set
func (tb TableSQLBuilder) UseColumn(columnsFunc func(column metadata.Column) TableSQLBuilderColumn) TableSQLBuilder {
	_ = "STUB: not implemented"
	return *new(TableSQLBuilder)
}

// TableSQLBuilderColumn is template for table sql builder column
type TableSQLBuilderColumn struct {
	Skip bool
	Name string
	Type string
}

var reservedKeywords = []string{"TableName", "Table", "SchemaName", "Alias", "AllColumns", "MutableColumns", "DefaultColumns"}

func renameIfReserved(name string) string { _ = "STUB: not implemented"; return "" }

// DefaultTableSQLBuilderColumn returns default implementation of TableSQLBuilderColumn
func DefaultTableSQLBuilderColumn(columnMetaData metadata.Column) TableSQLBuilderColumn {
	_ = "STUB: not implemented"
	return *new(TableSQLBuilderColumn)
}

// getSqlBuilderColumnType returns type of jet sql builder column
func getSqlBuilderColumnType(columnMetaData metadata.Column) string {
	_ = "STUB: not implemented"
	return ""
}

// sqlToColumnType maps the type of a SQL column type to a go jet sql builder column. The second return value returns
// whether the given type is supported.
func sqlToColumnType(columnMetaData metadata.Column) string { _ = "STUB: not implemented"; return "" }

//MySQL

//MySQL:

//MySQL

// MySQL

// postgres

// mysql and sqlite

// MySQL

// EnumSQLBuilder is template for generating enum SQLBuilder files
type EnumSQLBuilder struct {
	Skip         bool
	Path         string
	FileName     string
	InstanceName string
	ValueName    func(enumValue string) string
}

// DefaultEnumSQLBuilder returns default implementation of EnumSQLBuilder
func DefaultEnumSQLBuilder(enumMetaData metadata.Enum) EnumSQLBuilder {
	_ = "STUB: not implemented"
	return *new(EnumSQLBuilder)
}

// PackageName returns enum sql builder package name
func (e EnumSQLBuilder) PackageName() string { _ = "STUB: not implemented"; return "" }

// UsePath returns new EnumSQLBuilder with new path set
func (e EnumSQLBuilder) UsePath(path string) EnumSQLBuilder {
	_ = "STUB: not implemented"
	return *

	// UseFileName returns new EnumSQLBuilder with new file name set
	new(EnumSQLBuilder)
}

func (e EnumSQLBuilder) UseFileName(name string) EnumSQLBuilder {
	_ = "STUB: not implemented"
	return *new(EnumSQLBuilder)
}

// UseInstanceName returns new EnumSQLBuilder with instance name set
func (e EnumSQLBuilder) UseInstanceName(name string) EnumSQLBuilder {
	_ = "STUB: not implemented"
	return *new(EnumSQLBuilder)
}

func defaultEnumValueName(enumName, enumValue string) string { _ = "STUB: not implemented"; return "" }
