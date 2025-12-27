package main

import (
	"fmt"
	"math"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

// funções para serem aplicadas com as go routines

func testRoutines(arrA []int, n int) (int, []int) {
	j := 0
	for i := 0; i < len(arrA); i++ {
		arrA[i] = arrA[i] * n
		j = j + arrA[i]
	}

	return j, arrA
}

func mapFigas(arr []int, f func(int) int) []int {

	for i, v := range arr {
		arr[i] = f(v)
	}
	return arr
}

// go routines

func applyF1(arrA []int, n int) (int, []int) {
	var wg sync.WaitGroup

	numWorkers := runtime.NumCPU()

	fmt.Println("Fazendo isso com X nucleos: ", numWorkers)

	chunks := (len(arrA) + numWorkers - 1) / numWorkers

	ch := make(chan int, numWorkers)

	for i := 0; i < numWorkers; i++ {
		start := i * chunks
		end := start + chunks

		if start >= len(arrA) {
			break
		}
		if end > len(arrA) {
			end = len(arrA)
		}

		wg.Add(1)

		go func(slice []int) {
			defer wg.Done()

			par_j, _ := testRoutines(slice, n)

			ch <- par_j
		}(arrA[start:end])
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	totalJ := 0

	for val := range ch {
		totalJ += val
	}

	return totalJ, arrA
}

func applyF2(arrA []int, f func(int) int) []int {
	var wg sync.WaitGroup
	numWorkers := runtime.NumCPU()

	fmt.Println("Fazendo isso com X nucleos: ", numWorkers)

	chunks := (len(arrA) + numWorkers - 1) / numWorkers

	for i := 0; i < numWorkers; i++ {
		start := i * chunks
		end := start + chunks

		if start >= len(arrA) {
			break
		}
		if end > len(arrA) {
			end = len(arrA)
		}

		wg.Add(1)

		go func(slice []int) {
			defer wg.Done()

			mapFigas(slice, f)
		}(arrA[start:end])
	}

	wg.Wait()
	return arrA
}

// function to test benchmarking

func runTestSemGoR(arrA []int, n int) (int, int) {
	start1 := time.Now()

	j, arrB := testRoutines(arrA, n)

	duration := time.Since(start1)
	fmt.Println("\n\n O resultado do j foi: ", j, "\n\n O resultado do array em uma posição aleatória foi: ", arrB[250])
	fmt.Println("Tempo gasto para o normal foi de: ", duration, "com um array de: ", len(arrA))
	f := func(n int) int {
		return n * n
	}

	start2 := time.Now()
	i := mapFigas(arrB, f)
	duration2 := time.Since(start2)
	fmt.Println("O resultado do mapFigas normal foi: ", i[250])
	fmt.Println("Tempo gasto para o normal foi de: ", duration2, "com um array de: ", len(arrA))

	return int(duration), int(duration2)
}

func runTestComGoR(arrA []int, n int) (int, int) {

	start1 := time.Now()

	j, arrB := applyF1(arrA, n)

	duration1 := time.Since(start1)
	fmt.Println("\n\n O resultado do j foi: ", j, "\n\n O resultado do array em uma posição aleatória foi: ", arrB[250])
	fmt.Println("Tempo gasto para o normal foi de: ", duration1, "com um array de: ", len(arrA))

	f := func(n int) int {
		return n * n
	}
	start2 := time.Now()

	i := applyF2(arrA, f)
	duration2 := time.Since(start2)
	fmt.Println("O resultado do mapFigas normal foi: ", i[250])
	fmt.Println("Tempo gasto para o normal foi de: ", duration2, "com um array de: ", len(arrA))

	return int(duration1), int(duration2)
}

// auxiliares
func fill(arrA []int, arrB []int) {

	fmt.Println("Fazendo o preenchimento dos arrays")

	for i := 0; i < len(arrA); i++ {
		arrA[i] = rand.Intn(500)
	}

	fmt.Println("Feito o preenchimento do array A")

	j := len(arrB)
	jj := 0

	for jj < j {
		arrB[jj] = rand.Intn(100)
		jj += 1
	}

	fmt.Println("Feito o preenchimento do array B")
}

func showSome(arrA []int) {

	for i := 0; i < int(len(arrA)/10); i++ {
		fmt.Println("Printando -> ", arrA[i], "\n")
	}

}

func main() {

	fmt.Println("Iniciando o treinamento \n\n")

	arrA := make([]int, 1000)
	arrB := make([]int, 1000)

	fill(arrA, arrB)

	showSome(arrA)

	fmt.Println("FEito o procedimento e espero que seja sem erros")

	runTestSemGoR(arrA, 100)

	fmt.Println("Teste agora começa para valer ")

	valores1 := make([]int, 0, 20)
	valores2 := make([]int, 0, 20)

	for i := 7; i < 10; i++ {
		fmt.Println("Fazendo o teste com x 0: ", i)
		x := make([]int, int(math.Pow(10, float64(i))))

		for idx := range x {
			x[idx] = rand.Intn(500)
		}

		figas, figas2 := runTestSemGoR(x, 5)
		valores1 = append(valores1, figas)
		valores1 = append(valores1, figas2)
		fmt.Println("Adicionaod os valores", valores1)

		fmt.Println("\n")
	}

	fmt.Println("Teste agora com as go routines ")

	for i := 7; i < 10; i++ {
		fmt.Println("Fazendo o teste com x 0: ", i)
		x := make([]int, int(math.Pow(10, float64(i))))

		for idx := range x {
			x[idx] = rand.Intn(500)
		}

		runTestComGoR(x, 5)

		figas, figas2 := runTestSemGoR(x, 5)
		valores2 = append(valores2, figas)
		valores2 = append(valores2, figas2)

		fmt.Println("\n")
	}

	//for i := 0; i < len(valores1); i++ {
	//	fmt.Println("Tempo de demora F1 linear: ", valores1[i], "\n")
	//	fmt.Println("Tempo de demora F1 concur: ", valores2[i+1], "\n")
	//	fmt.Println("Tempo de demora F2 linear: ", valores1[i], "\n")
	//	fmt.Println("Tempo de demora F2 concur: ", valores2[i+1], "\n")
	//}

	fmt.Println(valores1)
	fmt.Println(valores2)
}
