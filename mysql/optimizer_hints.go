package mysql

import (
	"github.com/go-jet/jet/v2/internal/jet"
)

// OptimizerHint provides a way to optimize query execution per-statement basis
type OptimizerHint = jet.OptimizerHint

// MAX_EXECUTION_TIME limits statement execution time
func MAX_EXECUTION_TIME(miliseconds int) OptimizerHint {
	_ = "STUB: not implemented"
	return *new(OptimizerHint)
}

// QB_NAME assigns name to query block
func QB_NAME(name string) OptimizerHint { _ = "STUB: not implemented"; return *new(OptimizerHint) }
