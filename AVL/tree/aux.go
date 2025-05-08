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

func random_line() int {
	return rand.Intn(1000-100) + 1000
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

	if len(lines) > 0 && strings.HasPrefix(lines[0], "score") {
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

		if err == nil {
			scores = append(scores, score)
		}
	}

	return scores
}
