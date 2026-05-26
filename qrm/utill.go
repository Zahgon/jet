package qrm

import (
	"database/sql"
	"encoding/json"
	"reflect"
	"strings"
	"time"

	"github.com/google/uuid"
)

var scannerInterfaceType = reflect.TypeOf((*sql.Scanner)(nil)).Elem()

func implementsScannerType(fieldType reflect.Type) bool { _ = "STUB: not implemented"; return false }

func getScanner(value reflect.Value) sql.Scanner {
	_ = "STUB: not implemented"
	return *new(sql.Scanner)
}

func getSliceElemType(slicePtrValue reflect.Value) reflect.Type {
	_ = "STUB: not implemented"
	return *new(reflect.Type)
}

func getSliceElemPtrAt(slicePtrValue reflect.Value, index int) reflect.Value {
	_ = "STUB: not implemented"
	return *new(reflect.Value)
}

func appendElemToSlice(slicePtrValue reflect.Value, objPtrValue reflect.Value) error {
	_ = "STUB: not implemented"
	return nil
}

func newElemPtrValueForSlice(slicePtrValue reflect.Value) reflect.Value {
	_ = "STUB: not implemented"
	return *new(reflect.Value)
}

func getTypeName(structType reflect.Type, parentField *reflect.StructField) string {
	_ = "STUB: not implemented"
	return ""
}

func getTypeAndFieldName(structType string, field reflect.StructField) (string, string, bool) {
	_ = "STUB: not implemented"
	return "", "", false
}

var replacer = strings.NewReplacer(" ", "", "-", "", "_", "")

func toCommonIdentifier(name string) string { _ = "STUB: not implemented"; return "" }

func initializeValueIfNilPtr(value reflect.Value) { _ = "STUB: not implemented"; return }

var timeType = reflect.TypeOf(time.Now())
var uuidType = reflect.TypeOf(uuid.New())
var byteArrayType = reflect.TypeOf([]byte(""))
var jsonRawMessageType = reflect.TypeOf(json.RawMessage{})

func isSimpleModelType(objType reflect.Type) bool { _ = "STUB: not implemented"; return false }

// source can't be pointer
// destination can be pointer
func assign(source, destination reflect.Value) error { _ = "STUB: not implemented"; return nil }

// needs for the type conversions are rare, so we leave conversion as a last assign step if everything else fails

func assignIfAssignable(source, destination reflect.Value) bool {
	_ = "STUB: not implemented"
	return false
}

// source and destination are non-ptr values
func tryAssign(source, destination reflect.Value) error { _ = "STUB: not implemented"; return nil }

func tryConvert(source, destination reflect.Value) bool { _ = "STUB: not implemented"; return false }

func setZeroValue(value reflect.Value) { _ = "STUB: not implemented"; return }

func isPrimaryKey(field reflect.StructField, primaryKeyOverwrites []string) bool {
	_ = "STUB: not implemented"
	return false
}

func parentFieldPrimaryKeyOverwrite(parentField *reflect.StructField) []string {
	_ = "STUB: not implemented"
	return nil
}

func indirectType(reflectType reflect.Type) reflect.Type {
	_ = "STUB: not implemented"
	return *new(reflect.Type)
}

func fieldToString(field *reflect.StructField) string { _ = "STUB: not implemented"; return "" }

func cloneBytes(b []byte) []byte { _ = "STUB: not implemented"; return nil }

func concat(stringList ...string) string { _ = "STUB: not implemented"; return "" }
