package main

import "fmt"

func main() {
	s := make([]string, 3)
	fmt.Println("emp:", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)
	// need to reassign s to be itself plus whatever is appended, because append is actually creating a copy IIRC
	fmt.Println("apd2:", append(s, "g"))
	fmt.Println(s)

	// making
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	//slice operators are the same as in Python
	l := s[2:5]
	fmt.Println("sl1:", l)

	// no -1 in slice indices in go
	l = s[:len(s)-1]
	fmt.Println("sl2:", l)

	l = s[2:]
	fmt.Println("sl3:", l)

	//declaration in single line
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	// Slices can be composed into multidimensional data structs, the length of inner slices can vary, unlike with arrays
	twoD := make([][]int, 3)
	for ii := 0; ii < 3; ii++ {
		innerLen := ii + 1
		twoD[ii] = make([]int, innerLen)
		for jj := 0; jj < innerLen; jj++ {
			twoD[ii][jj] = ii + jj
		}
	}
	fmt.Println("2d:", twoD)
}
