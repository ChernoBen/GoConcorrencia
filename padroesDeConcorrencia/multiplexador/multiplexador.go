package main

import (
	"fmt"

	"github.com/ChernoBen/area/html"
)

//multiplexar: pegar 2 fontes de dados diferentes e jogar em 1 unico canal/2 channels em 1

//funcao que encaminha dados do canal de origem p/ canal de destino
func encaminhador(origem <-chan string, destino chan string) {
	for {
		// sempre que o canal de origem receber um dado ele enviará p/ canal de destino
		destino <- <-origem
	}
}

// funcao multiplexadora
func multiplexar(canal1, canal2 <-chan string) <-chan string {
	receptor := make(chan string) //canal que recebe os dados dos demais canais
	go encaminhador(canal1, receptor)
	go encaminhador(canal2, receptor)
	return receptor
}

func main() {
	canalResult := multiplexar(
		html.Htmltitles("https://www.amazon.com", "https://www.google.com"),
		html.Htmltitles("https://youtube.com", "https://facebook.com"),
	)
	fmt.Println(<-canalResult, "|", <-canalResult)
	fmt.Println(<-canalResult, "|", <-canalResult)
}
