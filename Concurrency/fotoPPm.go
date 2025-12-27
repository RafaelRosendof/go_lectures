package main

import (
	"bufio"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func gerarTripla() (int, int, int) {
	x_coef := rand.Intn(10) + 1
	y_coef := rand.Intn(10) + 1
	z_coef := rand.Intn(10) + 1

	x := rand.Intn(255) / x_coef
	y := rand.Intn(255) / y_coef
	z := rand.Intn(255) / z_coef

	return x, y, z
}

func gerar_tripla_cond(n int) (int, int, int) {
	cores := [20][3]int{
		{255, 0, 0},     // Vermelho
		{0, 255, 0},     // Verde
		{0, 0, 255},     // Azul
		{255, 255, 0},   // Amarelo
		{255, 0, 255},   // Magenta
		{0, 255, 255},   // Ciano
		{255, 165, 0},   // Laranja
		{128, 0, 128},   // Roxo
		{0, 128, 128},   // Verde-azulado
		{255, 192, 203}, // Rosa
		{150, 150, 150},
		{100, 250, 202},
		{190, 10, 92},
		{250, 100, 12},
		{200, 50, 192},
		{0, 90, 162},
		{10, 150, 152},
		{190, 110, 142},
		{180, 200, 132},
		{150, 99, 200},
	}

	index := n
	return cores[index][0], cores[index][1], cores[index][2]
}

func makeBatchMatrix(matrix [][][3]int, startLine, startCol, sizeLine, sizeCol, color int) {

	for i := startLine; i < startLine+sizeLine; i++ {
		for j := startCol; j < startCol+sizeCol; j++ {
			x, y, z := gerar_tripla_cond(color)
			matrix[i][j] = [3]int{x, y, z}
		}
	}
}

// vai ser do foramto em blocos
func gerar_matix_comp(totalLines int, totalCols, divisoes int) [][][3]int {

	matrix := make([][][3]int, totalLines)
	for i := range matrix {
		matrix[i] = make([][3]int, totalCols)
	}

	blockH := totalLines / divisoes
	blockW := totalCols / divisoes

	for bL := 0; bL < divisoes; bL++ {
		for bC := 0; bC < divisoes; bC++ {
			startL := bL * blockH
			startC := bC * blockW

			cor := rand.Intn(20)
			if cor == 0 {
				cor = cor + 1
			}

			go makeBatchMatrix(matrix, startL, startC, blockH, blockW, cor)
		}
	}
	return matrix
}

func calcularFractal(px, py, largura, altura int) (int, int, int) {
	// Mapeia as coordenadas do pixel (0 a 1000) para o plano complexo (-2.0 a 1.0)
	x0 := float64(px)/float64(largura)*3.5 - 2.5
	y0 := float64(py)/float64(altura)*2.0 - 1.0

	x, y := 0.0, 0.0
	iteracao := 0
	maxIteracoes := 500

	for x*x+y*y <= 4 && iteracao < maxIteracoes {
		xtemp := x*x - y*y + x0
		y = 2*x*y + y0
		x = xtemp
		iteracao++
	}

	cor := uint8(iteracao * 255 / maxIteracoes)
	return int(cor), int(cor / 2), int(255 - cor)
}

// gerar a matrix de forma linear
func generateMatrix(n int, m int) [][][3]int {
	matrix := make([][][3]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([][3]int, m)
		for j := 0; j < m; j++ {
			x, y, z := gerarTripla()
			matrix[i][j] = [3]int{x, y, z}
		}
	}
	return matrix
}

func generateMatrix2(n int, m int) [][][3]int {
	matrix := make([][][3]int, n)

	for i := 0; i < n; i++ {
		matrix[i] = make([][3]int, m)
		for j := 0; j < m; j++ {
			r, g, b := calcularFractal(j, i, m, n)

			if r > 200 {
				matrix[i][j] = [3]int{255, 255, 255}
			} else {
				matrix[i][j] = [3]int{r, g, b}
			}
		}
	}

	return matrix
}

func receiveAndWriteBatch(filename string, lines int, cols, divisoes int) error {
	matrix := gerar_matix_comp(lines, cols, divisoes)

	file, err := os.Create(filename)

	if err != nil {
		return err
	}

	defer file.Close()

	writer := bufio.NewWriter(file)
	fmt.Fprint(writer, "P3\n")
	//fmt.Fprint(writer, "%d %d\n", n, m)
	fmt.Fprintf(writer, "%d %d\n", cols, lines)
	fmt.Fprint(writer, "255\n")

	for i := 0; i < lines; i++ {
		for j := 0; j < cols; j++ {
			pixel := matrix[i][j]
			fmt.Fprintf(writer, "%d %d %d ", pixel[0], pixel[1], pixel[2])
		}
		//fmt.Fprint(writer, "\n")
		writer.WriteString("\n")
	}

	return writer.Flush()

}

func receiveAndWrite(fileName string, n int, m int) error {

	matrix := generateMatrix(n, m)

	file, err := os.Create(fileName)

	if err != nil {
		return err
	}

	defer file.Close()

	writer := bufio.NewWriter(file)

	fmt.Fprint(writer, "P3\n")
	//fmt.Fprint(writer, "%d %d\n", n, m)
	fmt.Fprintf(writer, "%d %d\n", m, n)
	fmt.Fprint(writer, "255\n")

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			pixel := matrix[i][j]
			fmt.Fprintf(writer, "%d %d %d ", pixel[0], pixel[1], pixel[2])
		}
		//fmt.Fprint(writer, "\n")
		writer.WriteString("\n")
	}

	return writer.Flush()
}

func receiveAndWriteFrac(fileName string, n int, m int) error {

	matrix := generateMatrix2(n, m)

	file, err := os.Create(fileName)

	if err != nil {
		return err
	}

	defer file.Close()

	writer := bufio.NewWriter(file)

	fmt.Fprint(writer, "P3\n")
	//fmt.Fprint(writer, "%d %d\n", n, m)
	fmt.Fprintf(writer, "%d %d\n", m, n)
	fmt.Fprint(writer, "255\n")

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			pixel := matrix[i][j]
			fmt.Fprintf(writer, "%d %d %d ", pixel[0], pixel[1], pixel[2])
		}
		//fmt.Fprint(writer, "\n")
		writer.WriteString("\n")
	}

	return writer.Flush()
}

// montar function para fazer o teste de tempo

func runTestLiner(fileName string, n int, m int) {

	start1 := time.Now()

	receiveAndWrite(fileName, n, m)

	duration1 := time.Since(start1)

	fmt.Println("Para essa quantidade de NxM: ", n, m)
	fmt.Println("Retornou algo como: ", duration1)

}

func main() {

	fmt.Println("Começando os experimentos: ")
	receiveAndWriteFrac("ppmFile.ppm", 4000, 4000)
	lines := flag.Int("lines", 1080, "an int")
	cols := flag.Int("cols", 1920, "an int")
	quads := flag.Int("quads", 4, "an int")
	flag.Parse()

	receiveAndWriteBatch("ppmFFile.ppm", *lines, *cols, *quads)

	//runTestLiner("ppmFile.ppm", 1080, 1920)

	fmt.Println("Teste liner realizado com sucesso")
}
