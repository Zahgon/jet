package jet

// ColumnList is a helper type to support list of columns as single projection
type ColumnList []ColumnExpression

func (cl ColumnList) isExpressionOrColumnList() {
	_ = "STUB: not implemented"

	// SET creates a column assignment from the current ColumnList using the provided expression.
	// This assignment can be used in INSERT queries (e.g., to set columns on conflict) or in UPDATE queries
	// (e.g., to assign new values to columns).
	//
	// The expression can be:
	//   - Another ColumnList: It must have the same length as the current ColumnList and each column must match by name
	//   - A ROW expression containing values.
	//   - A SELECT statement that returns a matching column list structure.
	//
	// Examples:
	//
	//	Link.AllColumns.SET(ROW(String("github.com"), Bool(false)))
	//
	//	Link.MutableColumns.SET(Link.EXCLUDED.MutableColumns)
	//
	//	Link.MutableColumns.SET(
	//	  SELECT(Link.MutableColumns).
	//	    FROM(Link).
	//	    WHERE(Link.ID.EQ(Int(200))),
	//	)
	return
}

func (cl ColumnList) SET(toAssignExp expressionOrColumnList) ColumnAssigment {
	_ = "STUB: not implemented"
	return *new(ColumnAssigment)
}

// Except will create new column list in which columns contained in list of excluded column names are removed
//
//	Address.AllColumns.Except(Address.PostalCode, Address.Phone)
func (cl ColumnList) Except(excludedColumns ...Column) ColumnList {
	_ = "STUB: not implemented"
	return *new(ColumnList)
}

// As will create new projection list where each column is wrapped with a new table alias.
// tableAlias should be in the form 'name' or 'name.*', or it can also be an empty string.
// For instance: If projection list has a column 'Artist.Name', and tableAlias is 'Musician.*', returned projection list will
// have a column wrapped in alias 'Musician.Name'. If tableAlias is empty string, it removes existing table alias ('Artist.Name' becomes 'Name').
func (cl ColumnList) As(tableAlias string) ProjectionList {
	_ = "STUB: not implemented"
	return *new(ProjectionList)
}

// From creates a new ColumnList that references the specified subquery.
// This method is typically used to project columns from a subquery into the surrounding query.
func (cl ColumnList) From(subQuery SelectTable) ColumnList {
	_ = "STUB: not implemented"
	return *new(ColumnList)
}

func (cl ColumnList) fromImpl(subQuery SelectTable) Projection {
	_ = "STUB: not implemented"
	return *new(Projection)
}

func (cl ColumnList) serialize(statement StatementType, out *SQLBuilder, options ...SerializeOption) {
	_ = "STUB: not implemented"
	return
}

func (cl ColumnList) serializeForProjection(statement StatementType, out *SQLBuilder) {
	_ = "STUB: not implemented"
	return
}

func (cl ColumnList) serializeForJsonObjEntry(statement StatementType, out *SQLBuilder) {
	_ = "STUB: not implemented"
	return
}

func (cl ColumnList) serializeForRowToJsonProjection(statement StatementType, out *SQLBuilder) {
	_ = "STUB: not implemented"
	return
}

// dummy column interface implementation

// Name is placeholder for ColumnList to implement Column interface
func (cl ColumnList) Name() string {
	_ = "STUB: not implemented"

	// TableName is placeholder for ColumnList to implement Column interface
	return ""
}

func (cl ColumnList) TableName() string                { _ = "STUB: not implemented"; return "" }
func (cl ColumnList) setTableName(name string)         { _ = "STUB: not implemented"; return }
func (cl ColumnList) setSubQuery(subQuery SelectTable) { _ = "STUB: not implemented"; return }
func (cl ColumnList) defaultAlias() string             { _ = "STUB: not implemented"; return "" }

// SetTableName is utility function to set table name from outside of jet package to avoid making public setTableName
func SetTableName(columnExpression ColumnExpression, tableName string) {
	_ = "STUB: not implemented"
	return
}

// SetSubQuery is utility function to set table name from outside of jet package to avoid making public setSubQuery
func SetSubQuery(columnExpression ColumnExpression, subQuery SelectTable) {
	_ = "STUB: not implemented"
	return
}
