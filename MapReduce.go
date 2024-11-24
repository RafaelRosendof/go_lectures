package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func Map(arr []int, f func(int) int) []int { //em go tem expressão lambdas?
	var result []int

	for _, v := range arr {
		result = append(result, f(v))
	}

	return result
}

func reduce(arr []int, f func(int, int) int) int {

	if len(arr) == 0 {
		return 0
	}

	res := arr[0]

	for _, v := range arr[1:] {
		res = f(res, v)
	}
	return res
}

func mapW(arr []int, f func(int) int) []int {
	res := make([]int, len(arr))
	for i, v := range arr {
		res[i] = f(v)
	}
	return res
}

func reduceW(arr []int, f func(int, int) int) int {
	if len(arr) == 0 {
		return 0
	}

	res := arr[0]

	for _, v := range arr[1:] {
		res = f(res, v)
	}
	return res

}

func mapWorker(arr []int, f func(int) int, results chan<- []int) {
	results <- Map(arr, f) // que danado é isso?
}

func reduceWorker(arr []int, f func(int, int) int, results chan<- int) {
	results <- reduce(arr, f)
}

func main() {

	rand.Seed(time.Now().UnixNano())

	size := 4_000_000
	arr := make([]int, size)
	for i := range arr {
		arr[i] = rand.Intn(100) // o que é o Intn
	}

	const workers = 4

	partSize := size / workers

	results := make(chan []int, workers)

	reduceRes := make(chan int, workers)

	var wg sync.WaitGroup

	mapFun := func(x int) int { return x * x }
	reduceFun := func(x, y int) int { return x + y }

	for i := 0; i < workers; i++ {
		start := i * partSize
		end := start + partSize

		wg.Add(1) // Adiciona uma goroutine ao wait group
		go func(start, end int) {
			defer wg.Done()                            // Marca a goroutine como concluída ao final
			mapWorker(arr[start:end], mapFun, results) // Executa Map na parte do array
		}(start, end)
	}

	//espera as goroutines do map terminarem
	wg.Wait()
	close(results)

	//junta os dados do map
	var mappedData []int
	for part := range results {
		mappedData = append(mappedData, part...)
	}

	//reduce em paralelo
	wg.Add(workers)
	chunkSize := len(mappedData) / workers
	for i := 0; i < workers; i++ {
		start := i * chunkSize
		end := start + chunkSize

		if i == workers-1 {
			end = len(mappedData)
		}

		go func(start, end int) {
			defer wg.Done()
			reduceWorker(mappedData[start:end], reduceFun, reduceRes)
		}(start, end)
	}

	wg.Wait()
	close(reduceRes)

	//junta os dados do reduce
	final := 0
	for part := range reduceRes {
		final += part
	}

	fmt.Println("A soma dos quadrados do array é ", final)

	/*
		for _, v := range arr {
			fmt.Println(v)
		}
	*/

	//res := Map(arr, func(x int) int { return x * x })
	/*
		for _, v := range res {
			fmt.Println(v)
		}

	*/
	//	red := reduce(res, func(x, y int) int { return x + y })

	//fmt.Println("A soma dos quadrados do array é ", red)

}
