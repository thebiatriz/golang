package calculos

import "errors"

func Somar(numero1 int, numero2 int) int {
	return numero1 + numero2
}

func Diminuir(numero1 int, numero2 int) int {
	return numero1 - numero2
}

func Multiplicar(numero1 int, numero2 int) int {
	return numero1 * numero2
}

//retorno de erro ou int. nil = nulo (operação sendo um sucesso)
func Dividir(numero1 int, numero2 int) (int, error) {
	if numero2 == 0 {
		return 0, errors.New("não é possivel realizar divisão por zero")
	}

	resultado := numero1 / numero2

	return resultado, nil
}