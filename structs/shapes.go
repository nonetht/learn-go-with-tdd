package shapes

import (
	"math"
)

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width, Height float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Height float64
	Width  float64
}

func (r Rectangle) Perimeter() float64 {
	return 2*r.Width + 2*r.Height
}

// Area 接收器（receiver）选择是将该方法和接收器绑定一起。
// 不传结构体，而是传递结构体的地址。而且，Go语言之中，都是将接收器变量命名为类型首字母
func (r *Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Area 就是计算接收器为Circle类型的面积
func (c *Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (t *Triangle) Area() float64 {
	return t.Height * t.Width * 0.5
}
