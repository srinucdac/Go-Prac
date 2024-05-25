package main

import "fmt"

type shape interface {
	getArea() float64
}
type triangle struct {
	base   float64
	height float64
}
type square struct {
	len float64
}

func (t triangle) getArea() float64 {
	return (t.base * t.height)
}
func (s square) getArea() float64 {
	return (s.len * s.len)
}
func main() {
	tr := triangle{base: 1.0, height: 2.0}
	sq := square{len: 2.0}
	printArea(tr)
	printArea(sq)
	//fmt.Println(tr.getArea())
	//fmt.Println(sq.getArea())
}
func printArea(s shape) {
	fmt.Println(s.getArea())
}
