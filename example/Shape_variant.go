package example

// A Shape is one of
//
//     - Circle
//     - Rectangle
//
// In each implementation exactly one of the methods should have the second
// return value be true.
type Shape interface {

	// Circle returns a Circle and a boolean. When the second return value is true,
	// the Shape is the returned Circle.
	Circle() (Circle, bool)

	// Rectangle returns a Rectangle and a boolean. When the second return value is true,
	// the Shape is the returned Rectangle.
	Rectangle() (Rectangle, bool)
}

// A coreShape provides default implementations for the methods of a
// Shape. The wrapper types only redefine the one for the associated variant.
type coreShape struct{}

func (_ coreShape) Circle() (Circle, bool) {
	var v Circle
	return v, false
}

func (_ coreShape) Rectangle() (Rectangle, bool) {
	var v Rectangle
	return v, false
}

// Shape converts a Circle to an instance of the sum type Shape.
func (v Circle) Shape() Shape {
	return wrapCircleShape{
		wrappedCircle: v,
	}
}

type wrapCircleShape struct {
	coreShape
	wrappedCircle Circle
}

func (w wrapCircleShape) Circle() (Circle, bool) {
	return w.wrappedCircle, true
}

// Shape converts a Rectangle to an instance of the sum type Shape.
func (v Rectangle) Shape() Shape {
	return wrapRectangleShape{
		wrappedRectangle: v,
	}
}

type wrapRectangleShape struct {
	coreShape
	wrappedRectangle Rectangle
}

func (w wrapRectangleShape) Rectangle() (Rectangle, bool) {
	return w.wrappedRectangle, true
}

// A ShapeExhaustive is a Shape that can be used to check
// exhaustivity in tests
type ShapeExhaustive struct {
	Shape

	CircleCalled    bool
	RectangleCalled bool
}

// Checks whether all the variants of the Shape were considered.
func (se ShapeExhaustive) Exhaustive() bool {

	if !se.CircleCalled {
		return false
	}

	if !se.RectangleCalled {
		return false
	}

	return true
}
