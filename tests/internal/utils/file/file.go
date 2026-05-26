package file

import (
	"testing"
)

// Exists expects file to exist on path constructed from pathElems and returns content of the file
func Exists(t *testing.T, pathElems ...string) (fileContent string) {
	_ = "STUB: not implemented"
	return ""
}

// #nosec G304

// NotExists expects file not to exist on path constructed from pathElems
func NotExists(t *testing.T, pathElems ...string) { _ = "STUB: not implemented"; return }

// #nosec G304
