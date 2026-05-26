package jet

// Projection is interface for all projection types. Types that can be part of, for instance SELECT clause.
type Projection interface {
	serializeForProjection(statement StatementType, out *SQLBuilder)
	serializeForJsonObjEntry(statement StatementType, out *SQLBuilder)
	serializeForRowToJsonProjection(statement StatementType, out *SQLBuilder)
	fromImpl(subQuery SelectTable) Projection
}

// SerializeForProjection is helper function for serializing projection outside of jet package
func SerializeForProjection(projection Projection, statementType StatementType, out *SQLBuilder) {
	_ = "STUB: not implemented"
	return
}

// ProjectionList is a redefined type, so that ProjectionList can be used as a Projection.
type ProjectionList []Projection

func (pl ProjectionList) fromImpl(subQuery SelectTable) Projection {
	_ = "STUB: not implemented"
	return *new(Projection)
}

func (pl ProjectionList) serializeForProjection(statement StatementType, out *SQLBuilder) {
	_ = "STUB: not implemented"
	return
}

func (pl ProjectionList) serializeForJsonObjEntry(statement StatementType, out *SQLBuilder) {
	_ = "STUB: not implemented"
	return
}

// As will create new projection list where each column is wrapped with a new table alias.
// tableAlias should be in the form 'name' or 'name.*', or it can be an empty string, which will remove existing table alias.
// For instance: If projection list has a column 'Artist.Name', and tableAlias is 'Musician.*', returned projection list will
// have a column wrapped in alias 'Musician.Name'. If tableAlias is empty string, it removes existing table alias ('Artist.Name' becomes 'Name').
func (pl ProjectionList) As(tableAlias string) ProjectionList {
	_ = "STUB: not implemented"
	return *new(ProjectionList)
}

// Except will create new projection list in which columns contained in excluded column names are removed
func (pl ProjectionList) Except(toExclude ...Column) ProjectionList {
	_ = "STUB: not implemented"
	return *new(ProjectionList)
}

func (pl ProjectionList) serializeForRowToJsonProjection(statement StatementType, out *SQLBuilder) {
	_ = "STUB: not implemented"
	return
}

// JsonObjProjectionList redefines []Projection so projections can be serialized as json object key/values
type JsonObjProjectionList []Projection

func (j JsonObjProjectionList) serialize(statement StatementType, out *SQLBuilder, options ...SerializeOption) {
	_ = "STUB: not implemented"
	return
}
