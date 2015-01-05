package example

// A Shape is one of
//
//     - Circle
//     - Rectangle
type Shape interface {
	Circle() (Circle, bool)
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
