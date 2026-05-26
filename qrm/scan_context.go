package qrm

import (
	"database/sql"
	"reflect"
)

// ScanContext  contains information about current row processed, mapping from the row to the
// destination types and type grouping information.
type ScanContext struct {
	rowNum                   int64
	row                      []interface{}
	uniqueDestObjectsMap     map[string]int
	commonIdentToColumnIndex map[string]int
	groupKeyInfoCache        map[string]groupKeyInfo
	typeInfoMap              map[string]typeInfo

	typesVisited    typeStack // to prevent circular dependency scan
	columnAlias     []string
	columnIndexRead []bool

	unmappedFields []string
}

// NewScanContext creates new ScanContext from rows
func NewScanContext(rows *sql.Rows) (*ScanContext, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *ScanContext) ensureStrictness() {
	_ = "STUB: not implemented" // can panic
	return
}

// can panic

// can panic

func (s *ScanContext) ensureEveryColumnRead() { _ = "STUB: not implemented"; return }

func (s *ScanContext) recordUnmappedField(structType reflect.Type, parentField *reflect.StructField, field reflect.StructField) {
	_ = "STUB: not implemented"
	// skip private/unsettable fields (those are ignored by mapRowToStruct anyway)
	return
}

// NOTE: For unnamed/anonymous structs, Name() is empty, so String() is used for readability/uniqueness.

func (s *ScanContext) ensureEveryFieldMapped() { _ = "STUB: not implemented"; return }

func isOptionalQrmField(field *reflect.StructField) bool { _ = "STUB: not implemented"; return false }

func shouldRecordUnmappedField(parentField *reflect.StructField, field reflect.StructField, fieldMap fieldMapping) bool {
	_ = "STUB: not implemented"
	return false
}

func createScanSlice(columnCount int) []interface{} { _ = "STUB: not implemented"; return nil }

// if destination is pointer to interface sql.Scan will just forward driver value

type typeInfo struct {
	fieldMappings []fieldMapping
}

type fieldMappingType int

const (
	simpleType  fieldMappingType = iota
	complexType                  // slice and struct are complex types supported
	implementsScanner
	jsonUnmarshal
)

type fieldMapping struct {
	rowIndex int // index in ScanContext.row
	Type     fieldMappingType
}

func (s *ScanContext) getTypeInfo(structType reflect.Type, parentField *reflect.StructField) typeInfo {
	_ = "STUB: not implemented"
	return *new(typeInfo)
}

type groupKeyInfo struct {
	typeName  string
	pkIndexes []int
	subTypes  []groupKeyInfo
}

func (s *ScanContext) getGroupKey(structType reflect.Type, structField *reflect.StructField) string {
	_ = "STUB: not implemented"
	return ""
}

func (s *ScanContext) constructGroupKey(groupKeyInfo groupKeyInfo) string {
	_ = "STUB: not implemented"
	return ""
}

func (s *ScanContext) getGroupKeyInfo(
	structType reflect.Type,
	parentField *reflect.StructField,
	typeVisited *typeStack) groupKeyInfo {
	_ = "STUB: not implemented"
	return *new(groupKeyInfo)
}

func (s *ScanContext) typeToColumnIndex(typeName, fieldName string) int {
	_ = "STUB: not implemented"
	return 0
}

// rowElemValue always returns non-ptr value,
// invalid value is nil
func (s *ScanContext) rowElemValue(index int) reflect.Value {
	_ = "STUB: not implemented"
	return *new(reflect.Value)
}

// no need to check validity of Elem, because s.row[index] always contains interface in interface

func (s *ScanContext) rowElemToString(index int) string { _ = "STUB: not implemented"; return "" }

func (s *ScanContext) rowElemValueClonePtr(index int) reflect.Value {
	_ = "STUB: not implemented"
	return *new(reflect.Value)
}
