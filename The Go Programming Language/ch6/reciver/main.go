package main

import (
	"fmt"
	"math"
)

//Path是连接多个点的直线段
type Path []Point
type Point struct{ X, Y float64 }

//Distance方法返回路径的长度
func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

func main() {

	p := Point{1, 2}
	p.ScaleBy(2) //p隐式转换为&p
	fmt.Printf("p's Type is %T, value=%[1]v\n", p)

	p2 := Point{1, 2}
	q := Point{4, 6}
	pptr := &p2
	fmt.Printf("pptr's Type is %T, pptr.Distance(q)=%v\n", pptr, pptr.Distance(q)) //pptr隐式转换为*pptr

	list := IntList{Value: 5}
	fmt.Printf("List sum: %v, of list: %#v\n", list.Sum(), list)
}

//IntList是整型链表
//*IntList的类型nil代表空列表
type IntList struct {
	Value int
	Tail  *IntList
}

//Sum返回列表元素的总和
func (list *IntList) Sum() int {
	if list == nil {
		return 0
	}
	return list.Value + list.Tail.Sum()
}
