package geometry

// #include <libgeometry-c/geometry.h>
import "C"

func CircleNew(x, y, radius float64) *C.Shape {
	pt := C.Point{0, 0}
	circle := C.geometry_circle_new(pt, C.double(radius))
	return circle
}

func ShapeDelete(shape *C.Shape) {
	C.geometry_shape_delete(shape)
}

func Area(shape *C.Shape) float64 {
	return float64(C.geometry_shape_area(shape))
}

func Radius(shape *C.Shape) float64 {
	return float64(C.geometry_circle_radius(shape))
}
