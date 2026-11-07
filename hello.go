package main

import "fmt"

func main() {
	nome := "Wesley"
	versao := 1.1
	fmt.Println("Ola, sr.", nome)
	fmt.Println("Este Programa esta na versao = ", versao)

	fmt.Println("==== MENU =======")
	fmt.Println("1- Iniciar Monitoramento ")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Sistema")

	var comando int
	fmt.Scan(&comando) //  capitura o que digitamos
	fmt.Println("O comando escolhido foi: ", comando)

}
