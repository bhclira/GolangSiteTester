package main

import (
	"fmt"
	"os"
	"net/http"
	// "reflect" // pesquisar os tipos de estruturas "fmt.println(reflect.TypeOf(nomes))"
	"time"
)

const monitoramentos = 3
const delay = 5

func main() {
	
	exibeIntro()

	for {
		exibeMenu()
	
		comando := leComando()
	
		switch comando {
		case 1:
			iniciarMonitoramento()
			break
		case 2:
			fmt.Println("Exibindo Logs...")
		case 0:
			fmt.Println("Saindo do programa.")
			os.Exit(0) // saiu como esperado
		default:
			fmt.Println("Não conheço este comando")
			os.Exit(-1) // saiu com algum problema
		}
	
		// switch é equivalente ao comando escolha do python
		// nao existe break no go, so executa uma vez
	
	}

}
	

func exibeIntro(){
	nome := "Breno"
	versao := 1.1
	fmt.Println("Olá, sr.", nome)
	fmt.Println("Este programa está na versâo", versao)
}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("o comando escolhido foi", comandoLido)
	fmt.Println("")

	return comandoLido

	// ou scanf("%d", &comando) -> (modificador + apontador)

}

func exibeMenu() {
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair")
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	sites := []string{"https://random-status-code.herokuapp.com/", "https://www.alura.com.br", "https://www.caelum.com.br"}

	// fmt.Println(sites)
	// site == sites[i] :: obtenha o indice e o valor da string
	
	for i := 0 ; i < monitoramentos ; i++ {
		for  i, site := range sites {
			fmt.Println("Testando site", i, ":", site)
			testaSite(site)
		}
		time.Sleep(delay * time.Second) // delay entre as verificações
		fmt.Println("")

	}
	

	fmt.Println("")

}

func testaSite(site string) {
	resp, _ := http.Get(site)
	
	if resp.StatusCode == 200 {
		fmt.Println("O site:", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("O site:", site, "Está com problemas. Status Code:", resp.StatusCode)
	}
}
	







	// saber se o site ta online
	// requisições web com Go
	// use o pacote "net/http"
	// função GHet retorna mais de uma coisa
	// para ignorar qualquer variavel num retorno duplo use underline



// if -> algo que retorne Boolean
// if nao usa parênteses

// if comando == 1 {
// 	fmt.Println("Monitorando...")
// else if comando == 2 {
// 	fmt.Println("Exibindo Logs...")
// else if comando == 0 {
// 	fmt.Println("Saindo do programa.")
// else {
// 	fmt.Println("Não conheço este comando")