package main

import (

	"os"
 	"fmt"
 	"net/http"
 	"time"
 	"bufio"
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

}

func exibeMenu() {
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair")
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	
	sites := leSitesDoArquivo()

	for i := 0 ; i < monitoramentos ; i++ {
		for  i, site := range sites {
			fmt.Println("Testando site", i, ":", site)
			testaSite(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")

	}
	fmt.Println("")

}

func testaSite(site string) {
	resp, err := http.Get(site)


	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}
	
	if resp.StatusCode == 200 {
		fmt.Println("O site:", site, "foi carregado com sucesso!")
		registraLog(site, true)
	} else {
		fmt.Println("O site:", site, "Está com problemas. Status Code:", resp.StatusCode)
		registraLog(site, false)
	}
}	

func leSitesDoArquivo() []string {

    var sites []string

    arquivo, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um Erro:", err)
		
	}

	leitor := bufio.NewReader(arquivo)

	linha, err := leitor.ReadString('\n')

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}
    
    return sites

}

func registraLog(site string, status bool) {
	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|O_CREATE, 0666)

a}
