package jet

// OrderByClause interface
type OrderByClause interface {
	// NULLS_FIRST specifies sort where null values appear before all non-null values.
	// For some dialects(mysql,mariadb), which do not support NULL_FIRST, NULL_FIRST is simulated
	// with additional IS_NOT_NULL expression.
	// For instance,
	//        Rental.ReturnDate.DESC().NULLS_FIRST()
	// would translate to,
	//        rental.return_date IS NOT NULL, rental.return_date DESC
	NULLS_FIRST() OrderByClause

	// NULLS_LAST specifies sort where null values appear after all non-null values.
	// For some dialects(mysql,mariadb), which do not support NULLS_LAST, NULLS_LAST is simulated
	// with additional IS_NULL expression.
	// For instance,
	//        Rental.ReturnDate.ASC().NULLS_LAST()
	// would translate to,
	//        rental.return_date IS NULL, rental.return_date ASC
	NULLS_LAST() OrderByClause

	serializeForOrderBy(statement StatementType, out *SQLBuilder)
}

type orderByClauseImpl struct {
	expression Expression
	ascending  *bool
	nullsFirst *bool
}

func (ord *orderByClauseImpl) NULLS_FIRST() OrderByClause {
	_ = "STUB: not implemented"
	return *new(OrderByClause)
}

func (ord *orderByClauseImpl) NULLS_LAST() OrderByClause {
	_ = "STUB: not implemented"
	return *new(OrderByClause)
}

func (ord *orderByClauseImpl) serializeForOrderBy(statement StatementType, out *SQLBuilder) {
	_ = "STUB: not implemented"
	return
}

func newOrderByAscending(expression Expression, ascending bool) OrderByClause {
	_ = "STUB: not implemented"
	return *new(OrderByClause)
}

func newOrderByNullsFirst(expression Expression, nullsFirst bool) OrderByClause {
	_ = "STUB: not implemented"
	return *new(OrderByClause)
}
