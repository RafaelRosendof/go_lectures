package main

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
	"time"
)

func makeArray(size int) []int {
	x := make([]int, size)
	for i := 0; i < size; i++ {
		x[i] = rand.Intn(100)
	}
	return x
}

func ArrayMultiply(array []int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < len(array); i++ {
		array[i] = array[i] * array[i]
		array[i] = int(math.Sqrt(float64(array[i])) * math.Sqrt(float64(array[i])) * math.Sqrt(float64(array[i])))
		array[i] = int(math.Sqrt(float64(array[i])) * math.Sqrt(float64(array[i])) * math.Sqrt(float64(array[i])))
		array[i] = int(math.Exp(float64(array[i])) * 42 * math.Sqrt(float64(array[i])))
		array[i] = int(math.Exp(float64(array[i])) * 42 * math.Sqrt(float64(array[i])))
		array[i] = int(math.Tanh(float64(array[i])) * math.Tanh(float64(array[i])))
	}
}

func ArrayMultiplyNormal(array []int) {
	for i := 0; i < len(array); i++ {
		array[i] = array[i] * array[i]
		array[i] = int(math.Sqrt(float64(array[i])) * math.Sqrt(float64(array[i])) * math.Sqrt(float64(array[i])))
		array[i] = int(math.Sqrt(float64(array[i])) * math.Sqrt(float64(array[i])) * math.Sqrt(float64(array[i])))
		array[i] = int(math.Exp(float64(array[i])) * 42 * math.Sqrt(float64(array[i])))
		array[i] = int(math.Exp(float64(array[i])) * 42 * math.Sqrt(float64(array[i])))
		array[i] = int(math.Tanh(float64(array[i])) * math.Tanh(float64(array[i])))
	}
}

func main() {

	size := 90_000_000

	//Criar 4 canais e preencher o array x_i

	//criar 4 canais e multiplicar o array x_i

	rand.Seed(time.Now().UnixNano())

	x1 := makeArray(size)
	x2 := makeArray(size)
	x3 := makeArray(size)
	x4 := makeArray(size)
	x5 := makeArray(size)
	x6 := makeArray(size)
	x7 := makeArray(size)
	x8 := makeArray(size)

	var wg sync.WaitGroup
	wg.Add(4)

	time1 := time.Now()
	go ArrayMultiply(x1, &wg)
	go ArrayMultiply(x2, &wg)
	go ArrayMultiply(x3, &wg)
	go ArrayMultiply(x4, &wg)
	go ArrayMultiply(x5, &wg)
	go ArrayMultiply(x6, &wg)
	go ArrayMultiply(x7, &wg)
	go ArrayMultiply(x8, &wg)

	wg.Wait()
	time2 := time.Now()

	tempoConcorrente := time2.Sub(time1)

	fmt.Println("Feito com o tempo de execução: Concorrente", tempoConcorrente)

	fmt.Println("Tempo total concorrente foi:  ", time2.Sub(time1))
	fmt.Println(x1[:20])
	fmt.Println(x2[:20])
	fmt.Println(x3[:20])
	fmt.Println(x4[:20])

	time3 := time.Now()

	ArrayMultiplyNormal(x1)
	ArrayMultiplyNormal(x2)
	ArrayMultiplyNormal(x3)
	ArrayMultiplyNormal(x4)
	time4 := time.Now()

	fmt.Println("tempo total normal foi : ", time4.Sub(time3))

}
