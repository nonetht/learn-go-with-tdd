package shapes

import "testing"

func TestPerimeter(t *testing.T) {

	rectangle := Rectangle{10.0, 10.0}
	got := rectangle.Perimeter()
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {

	//t.Run("rectangle", func(t *testing.T) {
	//	rectangle := Rectangle{12.0, 6.0}
	//	got := rectangle.Area()
	//	want := 72.0
	//
	//	if got != want {
	//		// 使用 %g取代了 %f，则是因为%g会在错误信息之中输出更加精确的小数
	//		t.Errorf("got %g want %g", got, want)
	//	}
	//})
	//
	//t.Run("circles", func(t *testing.T) {
	//	circle := Circle{10}
	//	got := circle.Area()
	//	want := 314.1592653589793
	//
	//	if got != want {
	//		t.Errorf("got %g want %g", got, want)
	//	}
	//})

	// 辅助函数无需关心形状（shape），辅助函数和具体类型解耦。
	//checkArea := func(t testing.TB, shape Shape, want float64) {
	//	t.Helper()
	//	got := shape.Area()
	//	if got != want {
	//		t.Errorf("got %g want %g", got, want)
	//	}
	//}
	//
	//t.Run("rectangle", func(t *testing.T) {
	//	rectangle := Rectangle{12.0, 6.0}
	//	checkArea(t, &rectangle, 72.0)
	//})
	//
	//t.Run("circle", func(t *testing.T) {
	//	circle := Circle{10.0}
	//	checkArea(t, &circle, 314.1592653589793)
	//})

	// table driven tests
	// 新的语法，就是创建了一个匿名结构体。如果想要引入一个新的形状，那也很简单。
	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "Rectangle", shape: &Rectangle{12, 6}, want: 72.0},
		{name: "Circle", shape: &Circle{10}, want: 314.1592653589793},
		{name: "Circle", shape: &Triangle{12, 6}, want: 36.0},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("got %g want %g", got, tt.want)
		}
	}
}
