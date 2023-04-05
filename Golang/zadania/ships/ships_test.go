package ships

import (
	"fmt"
)

func ExampleShip_MoveTo() {
	testCases := [4]struct {
		ship  Ship
		point Point
	}{
		{ship: []Point{{0, 0}, {0, 1}, {0, 2}}, point: Point{3, 5}},
		{ship: []Point{{3, 5}, {4, 5}}, point: Point{0, 0}},
		{ship: []Point{{7, 5}, {8, 5}, {9, 5}}, point: Point{-4, 0}},
		{ship: []Point{{-2, 7}}, point: Point{5, -2}},
	}
	for _, c := range testCases {
		ship := c.ship.MoveTo(c.point)
		fmt.Println(ship)
		// Output: [{3 5} {3 6} {3 7}]
		//[{0 0} {1 0}]
		//[{-4 0} {-3 0} {-2 0}]
		//[{5 -2}]
	}

}

func ExamplePoint_Add_1() {
	oldPoint := Point{X: 2, Y: 0}
	point := oldPoint.Add(Point{X: 3, Y: 3})
	fmt.Println(point.X, point.Y)
	// Output: 5 3
}

func ExamplePoint_Add_2() {
	oldPoint := Point{X: 5, Y: 3}
	point := oldPoint.Add(Point{X: -3, Y: -3})
	fmt.Println(point.X, point.Y)
	// Output: 2 0
}
