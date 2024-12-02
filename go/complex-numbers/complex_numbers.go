package complexnumbers

import (
	"math"
)

// Define the Number type here.
type Number struct {
	real float64
	img  float64
}

func (n Number) Real() float64 {
	return n.real
}

func (n Number) Imaginary() float64 {
	return n.img
}

func (n1 Number) Add(n2 Number) Number {
	return Number{n1.real + n2.real, n1.img + n2.img}
}

func (n1 Number) Subtract(n2 Number) Number {
	return Number{n1.real - n2.real, n1.img - n2.img}
}

func (n1 Number) Multiply(n2 Number) Number {
	return Number{n1.real*n2.real - n1.img*n2.img, n1.img*n2.real + n1.real*n2.img}
}

func (n Number) Times(factor float64) Number {
	return Number{n.real * factor, n.img * factor}
}

func (n1 Number) Divide(n2 Number) Number {
	squared := 2.0
	return Number{
		(n1.real*n2.real + n1.img*n2.img) / (math.Pow(n2.real, squared) + math.Pow(n2.img, squared)),
		(n1.img*n2.real - n1.real*n2.img) / (math.Pow(n2.real, squared) + math.Pow(n2.img, squared)),
	}
}

func (n Number) Conjugate() Number {
	return Number{n.real, -n.img}
}

func (n Number) Abs() float64 {
	squared := 2.0
	return math.Sqrt(math.Pow(n.real, squared) + math.Pow(n.img, squared))
}

func (n Number) Exp() Number {
	factor := math.Exp(n.real)
	return Number{math.Cos(n.img) * factor, math.Sin(n.img) * factor}
}
