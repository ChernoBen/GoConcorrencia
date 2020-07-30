package main

import (
	"fmt"
	"time"
)

func fale(pessoa, texto string, qtd int) {
	for i := 0; i < qtd; i++ {
		time.Sleep(time.Second)
		//uma goroutine so funcionara enquanto a tred principal estiver funcionando
		fmt.Printf("%s: %s (iteracao %d)\n", pessoa, texto, i+1)

	}
}

func main() {
	go fale("Robot", "Hello", 5)
	fale("RObot2", "Olá mundo", 5)
}
