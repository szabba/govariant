package example

//go:generate govariant Shape Circle Rectangle

type Circle struct {
	Center Point
	Radius float64
}

type Rectangle struct {
	LowerLeft     Point
	Width, Height float64
}

type Point struct {
	X, Y float64
}
