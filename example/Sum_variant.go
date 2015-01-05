
package example



// A Sum is one of 
// 
//     - Circle
//     - Rectangle
type Sum interface {
	 Circle() (Circle, bool)
	 Rectangle() (Rectangle, bool)
	
}

// A SumExhaustive is a Sum that can be use to check
// exhaustivity in tests
type SumExhaustive struct {
	Sum

	 CircleCalled bool
	 RectangleCalled bool
	
}

// Checks whether all the variants of the Sum were considered.
func (se SumExhaustive) Exhaustive() bool {
	
	if !se.CircleCalled {
		return false
	}
	
	if !se.RectangleCalled {
		return false
	}
	

	return true
}


	
		// Circle implements the corresponding method of Sum on the Circle type.
		 func (sv Circle) Circle() (Circle, bool) {
				return sv, true
			}
		
	
		// Rectangle implements the corresponding method of Sum on the Circle type.
		 func (_ Circle) Rectangle() (Rectangle, bool) {
				var v Rectangle
				return v, false
			}
		
	

	
		// Circle implements the corresponding method of Sum on the Rectangle type.
		 func (_ Rectangle) Circle() (Circle, bool) {
				var v Circle
				return v, false
			}
		
	
		// Rectangle implements the corresponding method of Sum on the Rectangle type.
		 func (sv Rectangle) Rectangle() (Rectangle, bool) {
				return sv, true
			}
		
	

