package tree

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

/*
score
year
figas -> 0 or 1
*/

// give 3 values, one is a random beetween 100 - 1000 , second 1900 - 2025 , third is 0 or 1
func random_numbers() []int {

	seed := rand.Intn(100-1) + 1

	score := rand.Intn(1000-100) + 1000

	year := rand.Intn(2025-1900) + 1900

	figas := rand.Intn(10-5) + 5*seed

	return []int{score, year, figas}
}

func Write_csv(csv_in string) {
	arq, err := os.Create(csv_in)

	if err != nil {
		fmt.Println("Error creating file: ", err)
		return
	}

	defer arq.Close()

	writer := bufio.NewWriter(arq)
	writer.WriteString("score,year,figas\n")
	for i := 0; i < 20; i++ {
		numbers := random_numbers()
		writer.WriteString(fmt.Sprintf("%d,%d,%d\n", numbers[0], numbers[1], numbers[2]))
	}
	writer.Flush()
	fmt.Println("File created successfully")
}

func Read_csv(root *Tree, csv_file string) *Tree {
	arq, err := os.Open(csv_file)

	if err != nil {
		fmt.Println("File not found: ", err)
		return nil
	}

	defer arq.Close()

	scanner := bufio.NewScanner(arq)

	first := true

	for scanner.Scan() {

		line := scanner.Text()

		if first {
			first = false
			continue
		}

		field := strings.Split(line, ",")

		if len(field) < 3 {
			fmt.Println("Invalid Line: ", line)
			continue
		}

		score, _ := strconv.Atoi(field[0])
		year, _ := strconv.Atoi(field[1])
		figas, _ := strconv.Atoi(field[2])

		node := Node{
			score: score,
			year:  year,
			figas: figas,
		}

		root = Insert_avl(root, node)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error in reading the csv file: ", err)
	}

	return root
}
