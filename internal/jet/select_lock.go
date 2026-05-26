package jet

// RowLock is interface for SELECT statement row lock types
type RowLock interface {
	Serializer

	OF(...Table) RowLock
	NOWAIT() RowLock
	SKIP_LOCKED() RowLock
}

type selectLockImpl struct {
	lockStrength       string
	of                 []Table
	noWait, skipLocked bool
}

// NewRowLock creates new RowLock
func NewRowLock(name string) func() RowLock { _ = "STUB: not implemented"; return nil }

func newSelectLock(lockStrength string) *selectLockImpl { _ = "STUB: not implemented"; return nil }

func (s *selectLockImpl) OF(tables ...Table) RowLock {
	_ = "STUB: not implemented"
	return *new(RowLock)
}

func (s *selectLockImpl) NOWAIT() RowLock { _ = "STUB: not implemented"; return *new(RowLock) }

func (s *selectLockImpl) SKIP_LOCKED() RowLock { _ = "STUB: not implemented"; return *new(RowLock) }

func (s *selectLockImpl) serialize(statement StatementType, out *SQLBuilder, options ...SerializeOption) {
	_ = "STUB: not implemented"
	return
}
