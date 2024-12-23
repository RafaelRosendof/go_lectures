package main

import (
	"fmt"
)

func Map[T any, U any](x []T, f func(T) U) []U {
	if len(x) == 0 {
		return []U{}
	}

	return append([]U{f(x[0])}, Map(x[1:], f)...)
}

func main() {

	var x = []int{1, 5: 4, 6, 10: 100, 15}

	for i := 0; i < len(x); i++ {
		fmt.Println("Valor ", x[i])
	}

	x = append(x, 15)

	for i := 0; i < len(x); i++ {
		fmt.Println("Valor ", x[i])
	}

	figas := Map(x, func(x int) int { return x * x })

	fmt.Println(figas)
}
