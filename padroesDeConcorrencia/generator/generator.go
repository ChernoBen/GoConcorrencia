package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

// <- chan -canal somente leitura

//funcao que recebe uma lista de strings que recebe urls e retorna um canal somente-leitura do tipo string
func htmltitles(urls ...string) <-chan string {
	//criando canal do tipo string
	canal := make(chan string)
	//
	for _, url := range urls {
		// chamando funcao anonima a partir de uma goroutine
		go func(url string) {
			//obtendo response do metodo http.Get e ignorando erro
			resp, _ := http.Get(url)
			// passando o corpo da resposta
			html, _ := ioutil.ReadAll(resp.Body)
			//aplicando regex p/ obter conteudo do title da url
			r, _ := regexp.Compile("<title>(.*?)<\\/title>")
			// convertendo var html p/string e aplicando expressao regular
			canal <- r.FindStringSubmatch(string(html))[1]

			//chamando funcao anonima
		}(url)
	}
	//retornando canal para leitura
	return canal
}

func main() {

	titulo1 := htmltitles("https://www.amazon.com", "https://www.youtube.com")
	titulo2 := htmltitles("https://www.google.com", "https://www.facebook.com")
	fmt.Println(<-titulo1)
	fmt.Println(<-titulo2)

}
