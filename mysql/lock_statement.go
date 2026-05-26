package mysql

import "github.com/go-jet/jet/v2/internal/jet"

// LockStatement is interface for MySQL LOCK tables
type LockStatement interface {
	Statement
	READ() Statement
	WRITE() Statement
}

// LOCK creates LockStatement from list of tables
func LOCK(tables ...jet.SerializerTable) LockStatement {
	_ = "STUB: not implemented"
	return *new(LockStatement)
}

type lockStatementImpl struct {
	jet.SerializerStatement

	Lock  jet.ClauseStatementBegin
	Read  jet.ClauseOptional
	Write jet.ClauseOptional
}

func (l *lockStatementImpl) READ() Statement { _ = "STUB: not implemented"; return *new(Statement) }

func (l *lockStatementImpl) WRITE() Statement { _ = "STUB: not implemented"; return *new(Statement) }

// UNLOCK_TABLES explicitly releases any table locks held by the current session
func UNLOCK_TABLES() Statement { _ = "STUB: not implemented"; return *new(Statement) }

type unlockStatementImpl struct {
	jet.SerializerStatement
	Unlock jet.ClauseStatementBegin
}
