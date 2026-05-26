package must

import (
	"reflect"
)

// BeTrue panics when condition is false
func BeTrue(condition bool, errorStr string) { _ = "STUB: not implemented"; return }

// BeTypeKind panics with errorStr error, if v interface is not of reflect kind
func BeTypeKind(v interface{}, kind reflect.Kind, errorStr string) {
	_ = "STUB: not implemented"
	return
}

// ValueBeOfTypeKind panics with errorStr error, if v value is not of reflect kind
func ValueBeOfTypeKind(v reflect.Value, kind reflect.Kind, errorStr string) {
	_ = "STUB: not implemented"
	return
}

// TypeBeOfKind panics with errorStr error, if v type is not of reflect kind
func TypeBeOfKind(v reflect.Type, kind reflect.Kind, errorStr string) {
	_ = "STUB: not implemented"
	return
}

// BeInitializedPtr panics with errorStr if val interface is nil
func BeInitializedPtr(val interface{}, errorStr string) { _ = "STUB: not implemented"; return }
