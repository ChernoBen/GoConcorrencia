package main

import (
	"fmt"
	"time"
)

//channel além de ser um tipo, também é a forma de comunicação entre as goroutines
func multiplicador(valor int, canal chan int) {
	time.Sleep(time.Second) //aguarde 1 segundo
	canal <- 2 * valor      // enviando resultado da multiplicão ao canal c
	time.Sleep(time.Second)

	canal <- 3 * valor
	time.Sleep(3 * time.Second) //aguardando 3 segundos
	canal <- 4 * valor

}

func main() {
	canal := make(chan int)
	// a partir daqui a função será executada de forma independente
	go multiplicador(2, canal)
	//atribuindo valor do canal a uma variavel
	a, b := <-canal, <-canal
	fmt.Println(a, b)
	// eu enviei 3 valores ao canal
	fmt.Println("----\n", <-canal)

}
