
package example


type Sum interface {
	 Circle() (Circle, bool)
	 Rectangle() (Rectangle, bool)
	
}

type SumExhaustive struct {
	Sum

	 CircleCalled bool
	 RectangleCalled bool
	
}

func (se SumExhaustive) Exhaustive() bool {
	
	if !se.CircleCalled {
		return false
	}
	
	if !se.RectangleCalled {
		return false
	}
	

	return true
}


	
		
			func (sv Circle) Circle() (Circle, bool) {
				return sv, true
			}
		
	
		
			func (_ Circle) Rectangle() (Rectangle, bool) {
				var v Rectangle
				return v, false
			}
		
	

	
		
			func (_ Rectangle) Circle() (Circle, bool) {
				var v Circle
				return v, false
			}
		
	
		
			func (sv Rectangle) Rectangle() (Rectangle, bool) {
				return sv, true
			}
		
	

