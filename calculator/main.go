package main

import (
	"fmt"
	"github.com/thebiatriz/golang/calculator/operations"
)

func main() {
	fmt.Println("Olá! Bem-vindo ao pacote principal.")

	var option, number1, number2, result int

	fmt.Println(`----MENU----
    1 - Soma
    2 - Subtração
    3 - Multiplicação
    4 - Divisão`)

	fmt.Print("Digite a opção escolhida: ")
	fmt.Scanln(&option)

	fmt.Print("Digite o primeiro número: ")
	fmt.Scanln(&number1)

	fmt.Print("Digite o segundo número: ")
	fmt.Scanln(&number2)

	switch option {
	case 1:
		result = operations.Add(number1, number2)
	case 2:
		result = operations.Subtract(number1, number2)
	case 3:
		result = operations.Multiply(number1, number2)
	case 4:
		var resultError error

		result, resultError = operations.Divide(number1, number2)

		if resultError != nil {
			fmt.Println("Erro:", resultError)
			return
		}
	default:
		fmt.Println("Opção escolhida inválida")
		return
	}

	fmt.Printf("O resultado da operação é: %d", result)
}
