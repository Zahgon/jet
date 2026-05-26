package jet

// BlobExpression interface
type BlobExpression interface {
	Expression

	isStringOrBlob()

	EQ(rhs BlobExpression) BoolExpression
	NOT_EQ(rhs BlobExpression) BoolExpression
	IS_DISTINCT_FROM(rhs BlobExpression) BoolExpression
	IS_NOT_DISTINCT_FROM(rhs BlobExpression) BoolExpression

	LT(rhs BlobExpression) BoolExpression
	LT_EQ(rhs BlobExpression) BoolExpression
	GT(rhs BlobExpression) BoolExpression
	GT_EQ(rhs BlobExpression) BoolExpression
	BETWEEN(min, max BlobExpression) BoolExpression
	NOT_BETWEEN(min, max BlobExpression) BoolExpression

	CONCAT(rhs BlobExpression) BlobExpression

	LIKE(pattern BlobExpression) BoolExpression
	NOT_LIKE(pattern BlobExpression) BoolExpression
}

type blobInterfaceImpl struct {
	root BlobExpression
}

func (b *blobInterfaceImpl) isStringOrBlob() { _ = "STUB: not implemented"; return }

func (b *blobInterfaceImpl) EQ(rhs BlobExpression) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

func (b *blobInterfaceImpl) NOT_EQ(rhs BlobExpression) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

func (b *blobInterfaceImpl) IS_DISTINCT_FROM(rhs BlobExpression) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

func (b *blobInterfaceImpl) IS_NOT_DISTINCT_FROM(rhs BlobExpression) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

func (b *blobInterfaceImpl) GT(rhs BlobExpression) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

func (b *blobInterfaceImpl) GT_EQ(rhs BlobExpression) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

func (b *blobInterfaceImpl) LT(rhs BlobExpression) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

func (b *blobInterfaceImpl) LT_EQ(rhs BlobExpression) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

func (b *blobInterfaceImpl) BETWEEN(min, max BlobExpression) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

func (b *blobInterfaceImpl) NOT_BETWEEN(min, max BlobExpression) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

func (b *blobInterfaceImpl) CONCAT(rhs BlobExpression) BlobExpression {
	_ = "STUB: not implemented"
	return *new(BlobExpression)
}

func (b *blobInterfaceImpl) LIKE(pattern BlobExpression) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

func (b *blobInterfaceImpl) NOT_LIKE(pattern BlobExpression) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

//---------------------------------------------------//

type blobExpressionWrapper struct {
	Expression
	blobInterfaceImpl
}

func newBlobExpressionWrap(expression Expression) BlobExpression {
	_ = "STUB: not implemented"
	return *new(BlobExpression)
}

// BlobExp is blob expression wrapper around arbitrary expression.
// Allows go compiler to see any expression as blob expression.
// Does not add sql cast to generated sql builder output.
func BlobExp(expression Expression) BlobExpression {
	_ = "STUB: not implemented"
	return *new(BlobExpression)
}
