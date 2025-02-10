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
		array[i] = int(math.Sqrt(float64(array[i])) * math.Sqrt(float64(array[i])) * math.Sqrt(float64(array[i])))
		array[i] = array[i] * array[i]
		array[i] = int(math.Sqrt(float64(array[i])) * math.Sqrt(float64(array[i])) * math.Sqrt(float64(array[i])))
		array[i] = int(math.Sqrt(float64(array[i])) * math.Sqrt(float64(array[i])) * math.Sqrt(float64(array[i])))
	}
}

func ArrayMultiply2(array []int) {
	for i := 0; i < len(array); i++ {
		array[i] = int(math.Sqrt(float64(array[i])) * math.Sqrt(float64(array[i])) * math.Sqrt(float64(array[i])))
		array[i] = array[i] * array[i]
		array[i] = int(math.Sqrt(float64(array[i])) * math.Sqrt(float64(array[i])) * math.Sqrt(float64(array[i])))
		array[i] = int(math.Sqrt(float64(array[i])) * math.Sqrt(float64(array[i])) * math.Sqrt(float64(array[i])))

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
	wg.Add(8)

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

	fmt.Println(x1[:90])
	fmt.Println(x2[:90])
	fmt.Println(x3[:90])
	fmt.Println(x4[:90])

	x1 = makeArray(size)

	time33 := time.Now()
	ArrayMultiply2(x1)
	ArrayMultiply2(x2)
	ArrayMultiply2(x3)
	ArrayMultiply2(x4)
	ArrayMultiply2(x5)
	ArrayMultiply2(x6)
	ArrayMultiply2(x7)
	ArrayMultiply2(x8)

	time3 := time.Now()

	tempoSequencial := time3.Sub(time33)

	fmt.Println("Feito com o tempo de execução: ", tempoSequencial)
}
