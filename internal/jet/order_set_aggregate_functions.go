package jet

// MODE computes the most frequent value of the aggregated argument
func MODE() *OrderSetAggregateFunc { _ = "STUB: not implemented"; return nil }

// PERCENTILE_CONT computes a value corresponding to the specified fraction within the ordered set of
// aggregated argument values. This will interpolate between adjacent input items if needed.
func PERCENTILE_CONT(fraction FloatExpression) *OrderSetAggregateFunc {
	_ = "STUB: not implemented"
	return nil
}

// PERCENTILE_DISC computes  the first value within the ordered set of aggregated argument values whose position
// in the ordering equals or exceeds the specified fraction. The aggregated argument must be of a sortable type.
func PERCENTILE_DISC(fraction FloatExpression) *OrderSetAggregateFunc {
	_ = "STUB: not implemented"
	return nil
}

// OrderSetAggregateFunc implementation of order set aggregate function
type OrderSetAggregateFunc struct {
	name     string
	fraction FloatExpression
	orderBy  Window
}

func newOrderSetAggregateFunction(name string, fraction FloatExpression) *OrderSetAggregateFunc {
	_ = "STUB: not implemented"
	return nil
}

// WITHIN_GROUP_ORDER_BY specifies ordered set of aggregated argument values
func (p *OrderSetAggregateFunc) WITHIN_GROUP_ORDER_BY(orderBy OrderByClause) Expression {
	_ = "STUB: not implemented"
	return *new(Expression)
}

func newOrderSetAggregateFuncExpression(aggFunc *OrderSetAggregateFunc) Expression {
	_ = "STUB: not implemented"
	return *new(Expression)
}

type orderSetAggregateFuncSerializer struct {
	*OrderSetAggregateFunc
}

func (p *orderSetAggregateFuncSerializer) serialize(statement StatementType, out *SQLBuilder, options ...SerializeOption) {
	_ = "STUB: not implemented"
	return
}
