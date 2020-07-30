package main

import "fmt"

//entendendo funcionamento do buffer de um canal

func rotina(canal chan int) {
	canal <- 1
	canal <- 2
	fmt.Println("Executou")
	canal <- 3
	canal <- 4
}

func main() {
	// declarando um canal com buffer de 3 posicoes
	canal := make(chan int, 3)
	go rotina(canal)
	//lendo o buffer
	fmt.Println(<-canal)
	//nesse caso ele conseguirá ler até o quarto elemento por que funcao print irá liberar um espaco após a leitura do primeiro elemento
}

//como o buffer possui 3 posicoes ele ficará bloqueado até que se cosuma o dado de uma das posicoes p/ assim um novo dado ser inserido
