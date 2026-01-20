package types

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}
type rect2 struct {
	width, height float64
}
type circle struct {
	radius float64
}

// rectangel / persegi panjang
func (r rect2) area() float64 {
	return r.width * r.height
}
func (r rect2) perim() float64 {
	return r.height*2 + r.width*2
}

// lingkaran
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}
func detectCircle(g geometry) {
	if c, ok := g.(circle); ok {
		fmt.Println("circle with radius", c.radius)
	} else {
		fmt.Println("this is not a circle")
	}
}
func RunInterface() {
	// rectangel
	r := rect2{width: 10, height: 5}
	measure(r)
	// circle
	c := circle{radius: 5}
	measure(c)

	//circle detect
	detectCircle(r)
	detectCircle(c)
}
