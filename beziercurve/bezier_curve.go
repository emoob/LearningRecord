package beziercurve

type Point struct {
	X, Y float64
}

/*
  - start := Point{X: 0, Y: 0}         起始位置
    control := Point{X: 128, Y: 54}    最大较大弯曲程度
    end := Point{X: 256, Y: 108}       结束位置
    numPoints := 100                   点的数量
    points := QuadraticBezierCurve(start, control, end, numPoints)
*/
func QuadraticBezierCurve(start, control, end Point, numPoints int) []Point {
	points := make([]Point, numPoints+1)
	for i := 0; i <= numPoints; i++ {
		t := float64(i) / float64(numPoints)
		x := (1-t)*(1-t)*start.X + 2*(1-t)*t*control.X + t*t*end.X
		y := (1-t)*(1-t)*start.Y + 2*(1-t)*t*control.Y + t*t*end.Y
		points[i] = Point{X: x, Y: y}
	}
	return points
}
