package example

// A Shape is one of
//
//     - Circle
//     - Rectangle
//
// In each implementation exactly one of the methods should return a true
// boolean value.
type Shape interface {

	// Circle returns a Circle and a boolean. When the second return value is true,
	// the Shape is the returned Circle.
	Circle() (Circle, bool)

	// Rectangle returns a Rectangle and a boolean. When the second return value is true,
	// the Shape is the returned Rectangle.
	Rectangle() (Rectangle, bool)
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

// Circle implements the corresponding method of Shape on the Circle type.
func (sv Circle) Circle() (Circle, bool) {
	return sv, true
}

// Rectangle implements the corresponding method of Shape on the Circle type.
func (_ Circle) Rectangle() (Rectangle, bool) {
	var v Rectangle
	return v, false
}

// Circle implements the corresponding method of Shape on the Rectangle type.
func (_ Rectangle) Circle() (Circle, bool) {
	var v Circle
	return v, false
}

// Rectangle implements the corresponding method of Shape on the Rectangle type.
func (sv Rectangle) Rectangle() (Rectangle, bool) {
	return sv, true
}
