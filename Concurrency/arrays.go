package main

import (
	"fmt"
	"math/rand"
	"time"
	"sync"
)

func makeArray(size int) []int{
	x := make([]int , size)
	for i := 0; i < size; i++ {
		x[i] = rand.Intn(100)
	}
	return x
}


func ArrayMultiply(array [] int , wg *sync.WaitGroup){
	defer wg.Done()
	for i := 0 ; i < len(array); i++{
		array[i] = array[i] * array[i]
	}
}


func main(){

	size := 9_000_000

	//Criar 4 canais e preencher o array x_i

	//criar 4 canais e multiplicar o array x_i

	rand.Seed(time.Now().UnixNano())

	x1 := makeArray(size)
	x2 := makeArray(size)
	x3 := makeArray(size)
	x4 := makeArray(size)

	var wg sync.WaitGroup	
	wg.Add(4)

	go ArrayMultiply(x1, &wg)
	go ArrayMultiply(x2, &wg)
	go ArrayMultiply(x3, &wg)
	go ArrayMultiply(x4, &wg)

	wg.Wait()

	fmt.Println("Feito")

	fmt.Println(x1[:90])
	fmt.Println(x2[:90])
	fmt.Println(x3[:90])
	fmt.Println(x4[:90])
}