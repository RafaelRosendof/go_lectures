package redblack

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

func RandomNumbers() []int {
	score := rand.Intn(901) + 100 // Random score between 100 and 1000

	return []int{score}
}

func Write_csv(csv_in string) {
	arq, err := os.Create(csv_in)

	if err != nil {
		fmt.Println("Error creating the file : ", err)
		return
	}

	defer arq.Close()
	writer := bufio.NewWriter(arq)
	writer.WriteString("Score\n")

	for i := 0; i < 10_000; i++ {
		numbers := RandomNumbers()
		writer.WriteString(fmt.Sprintf("%d\n", numbers[0]))
	}

	writer.Flush()
	fmt.Println("File created successfully")

}

func Read_csv(tree *Tree, csv_file string) *Tree {
	arq, err := os.Open(csv_file)

	if err != nil {
		fmt.Println("File not found: ", err)
		return nil
	}

	defer arq.Close()

	scan := bufio.NewScanner(arq)

	first := true

	for scan.Scan() {
		line := scan.Text()

		if first {
			first = false
			continue // Skip the header line

		}

		field := strings.Slplit(line, ",")

		if len(field) < 1 {
			fmt.Println("Invalid Line: ", line)
			continue
		}

		score, err := strconv.Atoi(field[0])
		if err != nil {
			fmt.Println("Invalid Score: ", field[0])
			continue
		}

		node := Node{
			score: score,
			cor:   Red,
			esq:   nil,
			dir:   nil,
			pai:   nil,
		}

		root := Arv_insereRB(tree.raiz, &node)

	}
	if err := scan.Err(); err != nil {
		fmt.Println("Error reading file: ", err)
		return nil
	}

	return tree
}
