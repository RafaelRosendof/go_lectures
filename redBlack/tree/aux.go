package tree

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func RandomNumbers() []int {
	score := rand.Intn(1201) + 100 // Random score between 100 and 1300

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

	for i := 0; i < 1_000; i++ {
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

	root := Arv_criaArv()

	defer arq.Close()

	scan := bufio.NewScanner(arq)

	first := true

	for scan.Scan() {
		line := scan.Text()

		if first {
			first = false
			continue // Skip the header line

		}

		field := strings.Split(line, ",")

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
		}

		Arv_insereRB(root, &node)

	}
	if err := scan.Err(); err != nil {
		fmt.Println("Error reading file: ", err)
		return nil
	}

	return root
}

func Remove_nodes(num_nodes int, csv_in string) []int {

	rand.Seed(time.Now().UnixNano())

	arq, err := os.Open(csv_in)

	if err != nil {
		fmt.Println("File not found: ", err)
		return nil
	}

	defer arq.Close()

	scan := bufio.NewScanner(arq)

	var lines []string

	for scan.Scan() {
		lines = append(lines, scan.Text())
	}

	if len(lines) > 0 && strings.EqualFold(strings.TrimSpace(lines[0]), "score") {
		lines = lines[1:]
	}

	if num_nodes > len(lines) {
		num_nodes = len(lines)
	}

	perm := rand.Perm(len(lines))[:num_nodes]

	var scores []int

	for _, idx := range perm {
		field := strings.Split(lines[idx], ",")
		score, err := strconv.Atoi(strings.TrimSpace(field[0]))

		if err != nil {
			fmt.Println("Invalid Score: ", field[0])
			continue
		}
		scores = append(scores, score)
	}
	return scores
}
