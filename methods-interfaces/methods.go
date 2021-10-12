package main

import "fmt"

type rect struct {
	width, height int
}

// methods can be defined for a pointer type
func (r *rect) area() int {
	return r.width * r.height
}

// methods can also just be for a value receiver
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 10, height: 5}

	fmt.Println("r")
	fmt.Println("area:", r.area())
	fmt.Println("perimiter:", r.perim())

	rp := &r
	fmt.Println("area:", rp.area())
	fmt.Println("perimeter:", rp.perim())
}
