package main

import "fmt"

var total float64
var vezesRegistrado int

func titulo() {
	fmt.Println("=== CAIXA REGISTRADORA v1 ===") // espero adicionar um belo de um slice nisso...
}

func menu() {
	fmt.Println("1 - registrar valor de compra")
	fmt.Println("2 - mostrar total acumulado")
	fmt.Println("0 - sair")
	fmt.Print("escolha: ")
}

func registrar() {
	fmt.Println("digite o valor que quer adicionar: ")
	var valor float64
	fmt.Scanln(&valor)

	if valor < 0 {
		fmt.Println("erro: valor menor que 0")
		return
	}

	total += valor
	vezesRegistrado += 1

	fmt.Println("valor adicionado!")
}

func mostrar() {
	fmt.Printf("valor total da compra: %.2f\n", total)
	fmt.Printf("quantidade de registros: %d\n", vezesRegistrado)
}

func separador() {
	fmt.Println("==============")
}

func main() {
	var opcao int

	for {
		titulo()
		menu()
		fmt.Scanln(&opcao)

		switch opcao {
		case 0:
			separador()
			return
		case 1:
			separador()
			registrar()

		case 2:
			separador()
			mostrar()
		default:
			fmt.Println("opcao nao existe")
		}
	}
}
