package main

import (
	"fmt"

	"github.com/ghazan/golang-to-cpp-opaque/client/geometry"
)

func main() {
	circle := geometry.CircleNew(0, 0, 2)
	area := geometry.Area(circle)
	fmt.Println("circle area =", area)
	radius := geometry.Radius(circle)
	fmt.Println("circle radius =", radius)
	geometry.ShapeDelete(circle)
}
