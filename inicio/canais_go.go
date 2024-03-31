package main

import (
  "fmt"
  "time"
)

func printaConta(c chan int){// função que recebe um canal 'c ' como argumento
  num :=0

  for num >= 0{ //loop que continua enquanto num for maior ou igual a 0 
    num = <-c  //Recebe um valor do canal 'c' e atribui a 'num'
    fmt.Print(num , " ") //imprime o valor recebido seguido de um espaço
  }
}

func main(){
  c:= make(chan int ) //Cria um canal de inteiros 
  a := []int{8, 6, 7, 5, 4, 3, 1, 0, -8} //definindo um slice de inteiros o ultimo digito tem que ser menor que 0 para parar a operação.

  go printaConta(c) //iniciando uma goroutine chamando a função printaConta com um canal c

  for _, v := range a{ //itera sobre os valores do slice 'a'

    c <- v //envia cada valor para o canal 'c' 


  }


  time.Sleep(time.Millisecond * 1) //aguarda esse tempo por operação 
  fmt.Println("End da operação")
}
