package beziercurve

import (
	"fmt"
	"testing"
)

func TestQuadraticBezierCurve(t *testing.T) {
	start := Point{X: 0, Y: 0}      //     起始位置
	control := Point{X: 128, Y: 54} //  最大较大弯曲程度
	end := Point{X: 256, Y: 108}    //   结束位置
	numPoints := 100                //   点的数量
	points := QuadraticBezierCurve(start, control, end, numPoints)
	fmt.Printf("坐标系:%v \n", points)
}
