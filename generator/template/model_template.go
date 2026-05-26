package template

import (
	"github.com/go-jet/jet/v2/generator/metadata"
)

// Model is template for model files generation
type Model struct {
	Skip  bool
	Path  string
	Table func(table metadata.Table) TableModel
	View  func(table metadata.Table) ViewModel
	Enum  func(enum metadata.Enum) EnumModel
}

// PackageName returns package name of model types
func (m Model) PackageName() string { _ = "STUB: not implemented"; return "" }

// UsePath returns new Model template with replaced file path
func (m Model) UsePath(path string) Model {
	_ = "STUB: not implemented"
	return *

	// UseTable returns new Model template with replaced template for table model files generation
	new(Model)
}

func (m Model) UseTable(tableModelFunc func(table metadata.Table) TableModel) Model {
	_ = "STUB: not implemented"
	return *new(Model)
}

// UseView returns new Model template with replaced template for view model files generation
func (m Model) UseView(tableModelFunc func(table metadata.Table) TableModel) Model {
	_ = "STUB: not implemented"
	return *new(Model)
}

// UseEnum returns new Model template with replaced template for enum model files generation
func (m Model) UseEnum(enumFunc func(enumMetaData metadata.Enum) EnumModel) Model {
	_ = "STUB: not implemented"
	return *new(Model)
}

// ShouldSkip returns new Model template with new skip flag set
func (m Model) ShouldSkip(skip bool) Model {
	_ = "STUB: not implemented"
	return *

	// DefaultModel returns default Model template implementation
	new(Model)
}

func DefaultModel() Model { _ = "STUB: not implemented"; return *new(Model) }

// TableModel is template for table model files generation
type TableModel struct {
	Skip     bool
	FileName string
	TypeName string
	Field    func(columnMetaData metadata.Column) TableModelField
}

// ViewModel is template for view model files generation
type ViewModel = TableModel

// DefaultViewModel is default view template implementation
var DefaultViewModel = DefaultTableModel

// DefaultTableModel is default table template implementation
func DefaultTableModel(tableMetaData metadata.Table) TableModel {
	_ = "STUB: not implemented"
	return *new(TableModel)
}

// UseFileName returns new TableModel with new file name set
func (t TableModel) UseFileName(fileName string) TableModel {
	_ = "STUB: not implemented"
	return *new(TableModel)
}

// UseTypeName returns new TableModel with new type name set
func (t TableModel) UseTypeName(typeName string) TableModel {
	_ = "STUB: not implemented"
	return *new(TableModel)
}

// UseField returns new TableModel with new TableModelField template function
func (t TableModel) UseField(structFieldFunc func(columnMetaData metadata.Column) TableModelField) TableModel {
	_ = "STUB: not implemented"
	return *new(TableModel)
}

func getTableModelImports(modelType TableModel, tableMetaData metadata.Table) []string {
	_ = "STUB: not implemented"
	return nil
}

// EnumModel is template for enum model files generation
type EnumModel struct {
	Skip      bool
	FileName  string
	TypeName  string
	ValueName func(value string) string
}

// UseFileName returns new EnumModel with new file name set
func (em EnumModel) UseFileName(fileName string) EnumModel {
	_ = "STUB: not implemented"
	return *new(EnumModel)
}

// UseTypeName returns new EnumModel with new type name set
func (em EnumModel) UseTypeName(typeName string) EnumModel {
	_ = "STUB: not implemented"
	return *new(EnumModel)
}

// DefaultEnumModel returns default implementation for EnumModel
func DefaultEnumModel(enumMetaData metadata.Enum) EnumModel {
	_ = "STUB: not implemented"
	return *new(EnumModel)
}

// TableModelField is template for table model field generation
type TableModelField struct {
	Name string
	Type Type
	Tags []string
	Skip bool
}

// DefaultTableModelField returns default TableModelField implementation
func DefaultTableModelField(columnMetaData metadata.Column) TableModelField {
	_ = "STUB: not implemented"
	return *new(TableModelField)
}

// UseType returns new TypeModelField with a new field type set
func (f TableModelField) UseType(t Type) TableModelField {
	_ = "STUB: not implemented"
	return *

	// UseName returns new TableModelField implementation with new field name set
	new(TableModelField)
}

func (f TableModelField) UseName(name string) TableModelField {
	_ = "STUB: not implemented"
	return *

	// UseTags returns new TableModelField implementation with additional tags added.
	new(TableModelField)
}

func (f TableModelField) UseTags(tags ...string) TableModelField {
	_ = "STUB: not implemented"
	return *new(TableModelField)
}

// TagsString returns tags string representation
func (f TableModelField) TagsString() string { _ = "STUB: not implemented"; return "" }

// Type represents type of the struct field
type Type struct {
	ImportPath            string
	AdditionalImportPaths []string
	Name                  string
}

// NewType creates new type for dummy object
func NewType(dummyObject interface{}) Type { _ = "STUB: not implemented"; return *new(Type) }

func getTypeName(t interface{}) string { _ = "STUB: not implemented"; return "" }

func getImportPath(dummyData interface{}) string { _ = "STUB: not implemented"; return "" }

func getType(columnMetadata metadata.Column) Type { _ = "STUB: not implemented"; return *new(Type) }

func getUserDefinedType(column metadata.Column) string { _ = "STUB: not implemented"; return "" }

func getGoType(column metadata.Column) Type { _ = "STUB: not implemented"; return *new(Type) }

func toGoArrayType(elemType Type, column metadata.Column) Type {
	_ = "STUB: not implemented"
	return *new(Type)
}

// unsupported multidimensional arrays

// toGoType returns model type for column info.
func toGoType(column metadata.Column) Type { _ = "STUB: not implemented"; return *new(Type) }

//MySQL

// MySQL

//MySQL

// MySQL

// MySQL
