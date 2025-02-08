package main

import "fmt"

func Map[T any, U any](x []T, f func(T) U) []U {
	if len(x) == 0 {
		return []U{}
	}

	return append([]U{f(x[0])}, Map(x[1:], f)...)
}

func main() {

	var x = []int{1, 4: 5, 6, 300, 10: 120, 125}

	for i := 0; i < len(x); i++ {
		fmt.Println(x[i])
	}

	figas := Map(x, func(x int) int { return x * x })

	fmt.Println(figas)

	slice := make([]int, 10, 10)
	fmt.Println(slice)
}
