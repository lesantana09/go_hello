package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

const n_monitoramento = 3
const n_delay = 3

func main() {

	for {
		menuInicial()
		variaveisIniciais()

		comando := capturaComando()

		switch comando {
		case 1:
			monitoramento()
		case 2:
			fmt.Println("Exibindo Logs ...")
		case 0:
			fmt.Println("Saindo do Programa ...")
			os.Exit(0)
		default:
			fmt.Println("O comando informado não foi reconhecido ...")
			os.Exit(-1)
		}
	}

}

func menuInicial() {
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Programa")
}

func variaveisIniciais() {
	nome := "Leandro Santana"
	idade := 39
	version := 1.1

	fmt.Println("Hello World.", nome, "idade", idade)
	fmt.Println("Estamos na versão", version)
}

func capturaComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	return comandoLido
}

func monitoramento() {
	fmt.Println("Monitoramento ...")

	sites := lerdoArquivo()

	for i := 0; i < n_monitoramento; i++ {
		for _, site := range sites {
			fmt.Println("Testando o site", site)
			testaSite(site)
			fmt.Println("")
		}
		time.Sleep(n_delay * time.Second)
	}
}

func testaSite(site string) {
	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("O", site, "está em funcionamento.")
	} else {
		fmt.Println("O", site, "não está em funcionamento.")
	}
}

func lerdoArquivo() []string {
	var sites []string
	arquivo, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro ao abrir o arquivo", err)
	}

	leitor := bufio.NewReader(arquivo)

	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)
		sites = append(sites, linha)
		if err == io.EOF {
			break
		}
	}
	arquivo.Close()

	return sites
}
