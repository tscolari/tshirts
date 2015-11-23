package colors

import (
	"math"

	"github.com/lucasb-eyer/go-colorful"
)

func CalcDistanceHex(color1 string, color2 string) float64 {
	c1, _ := colorful.Hex(color1)
	c2, _ := colorful.Hex(color2)

	dist1 := calcDistanceDirectional(c1, c2)
	dist2 := calcDistanceDirectional(c2, c1)

	if dist1 < dist2 {
		return dist1
	}

	return dist2

}

func calcDistanceDirectional(color1 colorful.Color, color2 colorful.Color) float64 {
	l1, a1, b1 := color1.Lab()
	l2, a2, b2 := color2.Lab()

	l := math.Pow(l1-l2, 2)
	a := math.Pow(a1-a2, 2)
	b := math.Pow(b1-b2, 2)

	return math.Sqrt(l + a + b)
}
