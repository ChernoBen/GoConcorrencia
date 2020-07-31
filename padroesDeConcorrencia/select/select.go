package main

import (
	"fmt"
	"time"

	"github.com/ChernoBen/area/html"
)

// Entendendo a estrutura de controle select p/ concorrencia GO

// funcao que testa qual url tem a resposta mais rapida usando estrutura select
func testaURL(url1, url2, url3 string) string {
	canal1 := html.Htmltitles(url1)
	canal2 := html.Htmltitles(url2)
	canal3 := html.Htmltitles(url3)

	//estrutura de controle select/ muito parecida com switch
	select {
	//retornara a url que obter a resposta mais rapida
	case title1 := <-canal1:
		return title1
	case title2 := <-canal2:
		return title2
	case title3 := <-canal3:
		return title3
	case <-time.After(1000 * time.Millisecond):
		return ("time out !")

	}
}

func main() {
	winner := testaURL(
		"https://www.youtube.com",
		"https://www.google.com",
		"https://www.amazon.com",
	)
	fmt.Println(winner)
}
