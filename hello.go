package main

import (
	"fmt"
	"net/http"
	"os"
	"reflect"
)

func main() {

	//exibeIntroducao()

	exibirNomes()

	for {

		//exibeMenu()
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

	return comandoLido
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	var sites [4]string
	sites[0] = "https://torp.ind.br/"
	sites[1] = "https://www.alura.com.br/"
	sites[2] = "https://ge.globo.com/"
	site := "https://torp.ind.br/"
	fmt.Println("Site: ", sites, " foi carregado com sucesso")

	resp, _ := http.Get(site)
	//fmt.Println(resp)
	if resp.StatusCode == 200 {
		fmt.Println("Site: ", site, " foi carregado com sucesso")
	} else {
		fmt.Println("Site: ", site, " Site com erro:", resp.StatusCode)
	}
}

func exibirNomes() {
	nomes := []string{"Wesley", "Daniel", "Maria"}
	fmt.Println(nomes)
	fmt.Println(reflect.TypeOf(nomes))
	fmt.Println(len(nomes))

}
