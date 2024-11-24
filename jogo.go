package main

import (
	"fmt"
	"math/rand"
)

func getNum() int {
	return rand.Intn(100)
}

func main() {

	fmt.Println("Adivinhe o número! Qual dificuldade voce escolhe? (1) fácil , (2) médio , (3) difícil : ")

	var diff int
	fmt.Scanf("%d", &diff)

	switch {
	case diff == 1:
		fmt.Println("Escolheu o fácil, você terá 10 vidas para acertar o número")
		diff = 10

	case diff == 2:
		fmt.Println("Escolheu o médio, você terá 7 vidas para acertar o número")
		diff = 7

	case diff == 3:
		fmt.Println("Escolheu o difícil, você terá 5 vidas para acertar o número")
		diff = 5

	default:
		fmt.Println("Dificuldade inválida, escolha entre 1, 2 ou 3")
		return
	}

	fmt.Println("\n\n Vamos começar o jogo! \n\n")
	//fmt.Println("Digite um número entre 0 e 100: ")
	var num int
	//fmt.Scanf("%d", &num)

	secret := getNum()

	for diff > 0 {

		fmt.Println("Digite um número entre 0 e 100: ")
		fmt.Scanf("%d", &num)

		if num == secret {
			fmt.Println("Parabéns, você acertou o número!")
			return
		} else if num < secret {
			fmt.Println("O número secreto é maior que o número digitado")
		} else {
			fmt.Println("O número secreto é menor que o número digitado")
		}

		diff--
		fmt.Println("Você ainda tem", diff, "vidas")

	}
	fmt.Println("Suas vidas acabaram, o número secreto era", secret)

}
