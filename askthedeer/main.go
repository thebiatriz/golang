package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("\nOlá! Bem-vindo ao Pergunte ao Cervo")
	fmt.Println(
		`Como funciona?
---------------------------------------
Você vai escrever duas opções das quais você quer que o Cervo decida uma para você.
Está em dúvida do que fazer entre duas coisas? O Cervo escolhe por você!
---------------------------------------`)

	fmt.Println("\nDito isso, digite a primeira opção de atividade:")
	scanner.Scan()
	option1 := scanner.Text()

	fmt.Println("\nAgora, digite a segunda opção de atividade:")
	scanner.Scan()
	option2 := scanner.Text()

	fmt.Println("\n--------- O Cervo está pensando ---------")

	result := rand.Intn(2)

	fmt.Printf("\nBatalha entre \"%s\" e \"%s\"\n", option1, option2)
	fmt.Printf("O Cervo escolheu: ")

	if result == 0 {
		fmt.Printf("%s\n", option1)
		return
	}

	fmt.Printf("%s\n", option2)
}
