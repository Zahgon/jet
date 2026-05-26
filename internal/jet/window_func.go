package jet

// Window interface
type Window interface {
	Serializer
	ORDER_BY(expr ...OrderByClause) Window
	ROWS(start FrameExtent, end ...FrameExtent) Window
	RANGE(start FrameExtent, end ...FrameExtent) Window
	GROUPS(start FrameExtent, end ...FrameExtent) Window
}

type windowImpl struct {
	partitionBy []Expression
	orderBy     ClauseOrderBy
	frameUnits  string
	start, end  FrameExtent

	root Window
}

func newWindowImpl(root Window) *windowImpl { _ = "STUB: not implemented"; return nil }

func (w *windowImpl) serialize(statement StatementType, out *SQLBuilder, options ...SerializeOption) {
	_ = "STUB: not implemented"
	return
}

func (w *windowImpl) ORDER_BY(exprs ...OrderByClause) Window {
	_ = "STUB: not implemented"
	return *new(Window)
}

func (w *windowImpl) ROWS(start FrameExtent, end ...FrameExtent) Window {
	_ = "STUB: not implemented"
	return *new(Window)
}

func (w *windowImpl) RANGE(start FrameExtent, end ...FrameExtent) Window {
	_ = "STUB: not implemented"
	return *new(Window)
}

func (w *windowImpl) GROUPS(start FrameExtent, end ...FrameExtent) Window {
	_ = "STUB: not implemented"
	return *new(Window)
}

func (w *windowImpl) setFrameRange(start FrameExtent, end ...FrameExtent) {
	_ = "STUB: not implemented"
	return
}

// PARTITION_BY window function constructor
func PARTITION_BY(exp Expression, exprs ...Expression) Window {
	_ = "STUB: not implemented"
	return *new(Window)
}

// ORDER_BY window function constructor
func ORDER_BY(expr ...OrderByClause) Window { _ = "STUB: not implemented"; return *new(Window) }

// -----------------------------------------------

// FrameExtent interface
type FrameExtent interface {
	Serializer
	isFrameExtent()
}

// PRECEDING window frame clause
func PRECEDING(offset Serializer) FrameExtent { _ = "STUB: not implemented"; return *new(FrameExtent) }

// FOLLOWING window frame clause
func FOLLOWING(offset Serializer) FrameExtent { _ = "STUB: not implemented"; return *new(FrameExtent) }

type frameExtentImpl struct {
	preceding bool
	offset    Serializer
}

func (f *frameExtentImpl) isFrameExtent() { _ = "STUB: not implemented"; return }

func (f *frameExtentImpl) serialize(statement StatementType, out *SQLBuilder, options ...SerializeOption) {
	_ = "STUB: not implemented"
	return
}

// -----------------------------------------------

// Window function keywords
var (
	UNBOUNDED   = Keyword("UNBOUNDED")
	CURRENT_ROW = frameExtentKeyword{"CURRENT ROW"}
)

type frameExtentKeyword struct {
	Keyword
}

func (f frameExtentKeyword) isFrameExtent() {
	_ = "STUB: not implemented"

	// -----------------------------------------------
	return
}

// WindowName is used to specify window reference from WINDOW clause
func WindowName(name string) Window { _ = "STUB: not implemented"; return *new(Window) }

type windowName struct {
	windowImpl
	name string
}

func (w windowName) serialize(statement StatementType, out *SQLBuilder, options ...SerializeOption) {
	_ = "STUB: not implemented"
	return
}
