package main

import (
	"fmt"
	"start.com/inicial/calculos"
)

//toda funcao de pacote externo vai ser com a primeira letra maiuscula

func main() {
	fmt.Println("Olá! Bem-vindo ao pacote principal.")

	var opcao, numero1, numero2, resultado int

	fmt.Println(`----MENU----
    1 - Soma
    2 - Subtração
    3 - Multiplicação
    4 - Divisão`)

	fmt.Print("Digite a opção escolhida: ")
	fmt.Scanln(&opcao)

	fmt.Print("Digite o primeiro número: ")
	fmt.Scanln(&numero1)

	fmt.Print("Digite o segundo número: ")
	fmt.Scanln(&numero2)

	switch opcao {
	case 1:
		resultado = calculos.Somar(numero1, numero2)
	case 2:
		resultado = calculos.Diminuir(numero1, numero2)
	case 3:
		resultado = calculos.Multiplicar(numero1, numero2)
	case 4:
		var erro error

		resultado, erro = calculos.Dividir(numero1, numero2)

		if erro != nil {
			fmt.Println("Erro:", erro)
			return
		}
	default:
		fmt.Println("Opção escolhida inválida")
		return
	}

	fmt.Printf("O resultado da operação é: %d", resultado)
}
