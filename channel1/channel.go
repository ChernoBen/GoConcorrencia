package main

import "fmt"

func main() {
	//conceito de canal/declaração e envio da dados a um canal
	ch := make(chan int, 1) //canal do tipo inteiro e buffer 1
	ch <- 1                 //enviando dados
	<-ch                    // recebendo dados

	//exemplo de leitrura de valor atribuido a um canal
	ch <- 2
	fmt.Println(<-ch)
}
