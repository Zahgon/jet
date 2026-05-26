package jet

// Operators
const (
	StringConcatOperator        = "||"
	StringRegexpLikeOperator    = "REGEXP"
	StringNotRegexpLikeOperator = "NOT REGEXP"
)

//----------- Logical operators ---------------//

// NOT returns negation of bool expression result
func NOT(exp BoolExpression) BoolExpression { _ = "STUB: not implemented"; return *new(BoolExpression) }

// BIT_NOT inverts every bit in integer expression result
func BIT_NOT(expr IntegerExpression) IntegerExpression {
	_ = "STUB: not implemented"
	return *new(IntegerExpression)
}

//----------- Comparison operators ---------------//

// EXISTS checks for existence of the rows in subQuery
func EXISTS(subQuery Expression) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

// Eq returns a representation of "a=b"
func Eq(lhs, rhs Expression) BoolExpression { _ = "STUB: not implemented"; return *new(BoolExpression) }

// NotEq returns a representation of "a!=b"
func NotEq(lhs, rhs Expression) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

// IsDistinctFrom returns a representation of "a IS DISTINCT FROM b"
func IsDistinctFrom(lhs, rhs Expression) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

// IsNotDistinctFrom returns a representation of "a IS NOT DISTINCT FROM b"
func IsNotDistinctFrom(lhs, rhs Expression) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

// Lt returns a representation of "a<b"
func Lt(lhs Expression, rhs Expression) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

// LtEq returns a representation of "a<=b"
func LtEq(lhs, rhs Expression) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

// Gt returns a representation of "a>b"
func Gt(lhs, rhs Expression) BoolExpression { _ = "STUB: not implemented"; return *new(BoolExpression) }

// GtEq returns a representation of "a>=b"
func GtEq(lhs, rhs Expression) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

// Contains returns a representation of "a @> b"
func Contains(lhs Expression, rhs Expression) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

// IsContainedBy returns a representation of "a <@ b"
func IsContainedBy(lhs Expression, rhs Expression) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

// Overlap returns a representation of "a && b"
func Overlap(lhs, rhs Expression) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

// Add notEq returns a representation of "a + b"
func Add(lhs, rhs Serializer) Expression { _ = "STUB: not implemented"; return *new(Expression) }

// Sub notEq returns a representation of "a - b"
func Sub(lhs, rhs Serializer) Expression { _ = "STUB: not implemented"; return *new(Expression) }

// Mul returns a representation of "a * b"
func Mul(lhs, rhs Serializer) Expression { _ = "STUB: not implemented"; return *new(Expression) }

// Div returns a representation of "a / b"
func Div(lhs, rhs Serializer) Expression { _ = "STUB: not implemented"; return *new(Expression) }

// Mod returns a representation of "a % b"
func Mod(lhs, rhs Serializer) Expression { _ = "STUB: not implemented"; return *new(Expression) }

// --------------- CASE operator -------------------//

// CaseOperator is interface for SQL case operator
type CaseOperator interface {
	Expression

	WHEN(condition Expression) CaseOperator
	THEN(then Expression) CaseOperator
	ELSE(els Expression) CaseOperator
}

type caseOperatorImpl struct {
	ExpressionInterfaceImpl

	expression Expression
	when       []Expression
	then       []Expression
	els        Expression
}

// CASE create CASE operator with optional list of expressions
func CASE(expression ...Expression) CaseOperator {
	_ = "STUB: not implemented"
	return *new(CaseOperator)
}

func (c *caseOperatorImpl) WHEN(when Expression) CaseOperator {
	_ = "STUB: not implemented"
	return *new(CaseOperator)
}

func (c *caseOperatorImpl) THEN(then Expression) CaseOperator {
	_ = "STUB: not implemented"
	return *new(CaseOperator)
}

func (c *caseOperatorImpl) ELSE(els Expression) CaseOperator {
	_ = "STUB: not implemented"
	return *new(CaseOperator)
}

func (c *caseOperatorImpl) serialize(statement StatementType, out *SQLBuilder, options ...SerializeOption) {
	_ = "STUB: not implemented"
	return
}

// DISTINCT operator can be used to return distinct values of expr
func DISTINCT(expr Expression) Expression { _ = "STUB: not implemented"; return *new(Expression) }

func BinaryOperator(lhs Expression, rhs Expression, operator string) Expression {
	_ = "STUB: not implemented"
	return *new(Expression)
}
