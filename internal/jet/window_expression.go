package jet

type commonWindowImpl struct {
	expression Expression
	window     Window
}

func (w *commonWindowImpl) over(window ...Window) { _ = "STUB: not implemented"; return }

func (w *commonWindowImpl) serialize(statement StatementType, out *SQLBuilder, options ...SerializeOption) {
	_ = "STUB: not implemented"
	return
}

// --------------------------------------

type windowExpression interface {
	Expression
	OVER(window ...Window) Expression
}

func newWindowExpression(exp Expression) windowExpression {
	_ = "STUB: not implemented"
	return *new(windowExpression)
}

type windowExpressionImpl struct {
	Expression
	commonWindowImpl
}

func (f *windowExpressionImpl) OVER(window ...Window) Expression {
	_ = "STUB: not implemented"
	return *new(Expression)
}

func (f *windowExpressionImpl) serialize(statement StatementType, out *SQLBuilder, options ...SerializeOption) {
	_ = "STUB: not implemented"
	return
}

// -----------------------------------------------------

type floatWindowExpression interface {
	FloatExpression
	OVER(window ...Window) FloatExpression
}

func newFloatWindowExpression(floatExp FloatExpression) floatWindowExpression {
	_ = "STUB: not implemented"
	return *new(floatWindowExpression)
}

type floatWindowExpressionImpl struct {
	FloatExpression
	commonWindowImpl
}

func (f *floatWindowExpressionImpl) OVER(window ...Window) FloatExpression {
	_ = "STUB: not implemented"
	return *new(FloatExpression)
}

func (f *floatWindowExpressionImpl) serialize(statement StatementType, out *SQLBuilder, options ...SerializeOption) {
	_ = "STUB: not implemented"
	return
}

// ------------------------------------------------

type integerWindowExpression interface {
	IntegerExpression
	OVER(window ...Window) IntegerExpression
}

func newIntegerWindowExpression(intExp IntegerExpression) integerWindowExpression {
	_ = "STUB: not implemented"
	return *new(integerWindowExpression)
}

type integerWindowExpressionImpl struct {
	IntegerExpression
	commonWindowImpl
}

func (f *integerWindowExpressionImpl) OVER(window ...Window) IntegerExpression {
	_ = "STUB: not implemented"
	return *new(IntegerExpression)
}

func (f *integerWindowExpressionImpl) serialize(statement StatementType, out *SQLBuilder, options ...SerializeOption) {
	_ = "STUB: not implemented"
	return
}

// ------------------------------------------------

type boolWindowExpression interface {
	BoolExpression
	OVER(window ...Window) BoolExpression
}

func newBoolWindowExpression(boolExp BoolExpression) boolWindowExpression {
	_ = "STUB: not implemented"
	return *new(boolWindowExpression)
}

type boolWindowExpressionImpl struct {
	BoolExpression
	commonWindowImpl
}

func (f *boolWindowExpressionImpl) OVER(window ...Window) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

func (f *boolWindowExpressionImpl) serialize(statement StatementType, out *SQLBuilder, options ...SerializeOption) {
	_ = "STUB: not implemented"
	return
}
