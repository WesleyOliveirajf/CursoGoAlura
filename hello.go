package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

const monitoramento = 2
const delay = 5

func main() {

	exibeIntroducao()
	leSitesDoArquivos()

	for {

		exibeMenu()
		comando := leComando()

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo Logs...")
		case 0:
			fmt.Println("Saindo do programa")
			os.Exit(0)
		default:
			fmt.Println("Não conheço este comando")
			os.Exit(-1)
		}
	}

}

func exibeIntroducao() {
	nome := "Wesley"
	versao := 1.2
	fmt.Println("Olá, sr.", nome)
	fmt.Println("Este programa está na versão", versao)
}

func exibeMenu() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")
}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("O comando escolhido foi", comandoLido)
	fmt.Println("")

	return comandoLido
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	//sites := []string{"https://torp.ind.br/", "https://www.alura.com.br/", "https://ge.globo.com/"}

	sites := leSitesDoArquivos()

	for i := 0; i < monitoramento; i++ {
		for i, site := range sites {
			fmt.Println("Testando site", i, ": ", site)
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
		fmt.Println("Ocorreu um erro", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site: ", site, " foi carregado com sucesso")
	} else {
		fmt.Println("Site: ", site, " Site com erro:", resp.StatusCode)
	}
}

func leSitesDoArquivos() []string {

	var sites []string
	//arquivo, err := os.Open("site.txt")

	arquivo, err := os.ReadFile("site.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro", err)

	}
	fmt.Println(string(arquivo))
	return sites
}
