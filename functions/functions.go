package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}

// you can omit type name for multiple consecutive params of the same time
func plusPlus(a, b, c int) int {
	return a + b + c
}

//variadic functions are like Python's kwargs/args
func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
	return total
}

// with closures we can create generator-type functions in Go?
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func fact(n int) int {
	if n == 0 {
		return 1
	}
	fmt.Println("fact:", n)
	return n * fact(n-1)
}

func main() {

	res := plus(10, 20)
	fmt.Println("10 + 20 =", res)

	res = plusPlus(10, 20, 30)
	fmt.Println("10 + 20 + 30 =", res)

	nums := []int{1, 2, 3, 4, 5, 6}
	res = sum(nums...)
	fmt.Println(res)

	fmt.Println(intSeq())
	fmt.Println(intSeq())
	fmt.Println(intSeq())

	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInt := intSeq()
	fmt.Println(newInt())

	fmt.Println(fact(7))

	// closures can be recursive, but must be declared with a typed var explicitly before it's defined
	var fib func(n int) int

	fib = func(n int) int {
		if n < 2 {
			return n
		}
		return fib(n-1) + fib(n-2)
	}

	fmt.Println(fib(7))
}
