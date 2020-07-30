package main

import (
	"fmt"
	"time"
)

// criando função p/ definir se numero é primo
func isPrimo(num int) bool {
	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}

	}
	return true
}

func obterPrimos(n int, canal chan int) {
	inicio := 2
	//laco para gerar quantidade N de primos
	for i := 0; i < n; i++ {
		for primo := inicio; ; primo++ {
			if isPrimo(primo) {
				canal <- primo
				inicio = primo + 1
				time.Sleep(time.Millisecond * 200)
				break // quebra o laco interior e retornando ao exterior
			}
		}
	}
	close(canal)
}

func main() {
	canal := make(chan int, 30)
	go obterPrimos(60, canal)
	//imprimindo valores primos do buffer do canal
	for primo := range canal {
		fmt.Printf("%d ", primo)
	}
	fmt.Println("FIM !")
}

/*
lembrar de aplicar esse script no script de função totiente de euler
*/
