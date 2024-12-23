package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "Fala figas"
			time.Sleep(time.Nanosecond * 2000)
		}
	}()

	go func() {
		for {
			c2 <- "Sr. Firuza"
			time.Sleep(time.Microsecond * 500)
		}
	}()

	for i := 0; i < 5; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		}
	}

	c := make(chan string)
	go count(" Pigs ", c)

	for msg := range c {
		fmt.Println(msg)
	}

	rand.Seed(time.Now().UnixNano())

	size := 1_000_000
	x := make([]int, size)
	for i := 0; i < size; i++ {
		x[i] = rand.Intn(500)
	}
	//preenchido, agora vamos iterar sobre cada metade e colocar ao quadrado o array x
	c3 := make(chan []int)
	c4 := make(chan []int)

	go func(slice []int, c chan []int) {
		res := make([]int, len(slice))
		for i, val := range slice {
			res[i] = val * val
		}
		c <- res
	}(x[:size/2], c3)

	go func(slice []int, c chan []int) {
		res2 := make([]int, len(slice))
		for i, val := range slice {
			res2[i] = val * val
		}
		c <- res2
	}(x[:size/2], c4)

	fH1 := <-c3
	sH1 := <-c4

	x = append(fH1, sH1...)

	fmt.Println("Processamento concluido exemplo dos 100 primeiros", x[:100])

	for i := 0; i < size; i++ {
		fmt.Println(" ", x[i], " ")
	}
	fmt.Println("tamanho do array: ", len(x))

	size2 := 1_100_000
	y := make([]int, size2)
	for i := 0; i < size2; i++ {
		y[i] = rand.Intn(100)
	}

	temp1 := time.Now()
	num_workers := 500
	chunk_size := size2 / num_workers
	c5 := make([]chan []int, num_workers)

	for i := 0; i < num_workers; i++ {
		c5[i] = make(chan []int)
		go func(slice []int, ch chan []int) {
			res := make([]int, len(slice))
			for i, val := range slice {
				res[i] = val * val
			}
			ch <- res
		}(y[i*chunk_size:(i+1)*chunk_size], c5[i])
	}

	fim := []int{}

	for i := 0; i < num_workers; i++ {
		fim = append(fim, <-c5[i]...)
	}

	duration := time.Since(temp1)
	//9.861366ms -> 11 workers
	//10.602307ms -> 12
	//13.519712ms -> 110
	fmt.Println("Processado com 6 threads", fim[:1_100_000])
	fmt.Println("Processado em tempo x: ", duration)
}

func count(figas string, c chan string) {
	for i := 1; i <= 10; i++ {
		c <- figas
		time.Sleep(time.Millisecond * 200)
	}

	close(c)
}
