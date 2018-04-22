package main

import(
	"fmt"
	"os" //pacote para lidar com o sistema operacional do host
	"strconv"//conversão de string para outros formatos e vice-versa
)
/*
No exemplo a seguir, usaremos esse pacote para ter acesso aos argumentos passados ao nosso programa via linha de comando, e também para
instruir o programa a interromper sua execução e retornar um código de erro adequado em casos especiais.
*/
func main() {
	if len(os.Args) < 3 {
		/*
		Por padrão, o primeiro elemento sempre será o próprio nome do programa executado, depois os dois argumentos necessários
		para a computação dos dados no programa.
		*/
		fmt.Println("Uso: conversor <valores> <unidade>")
		os.Exit(1)
	}

	unidadeOrigem := os.Args[len(os.Args)-1]
	// acessa o último elemento do slice
	valoresOrigem := os.Args[1:len(os.Args)-1]
	// a variável recebe uma sublista dos argumentos, descartando o primeiro(nome do programa) e o último(a unidade já atribuída
	// na outra variável

	var unidadeDestino string

	if unidadeOrigem == "celsius" {
		unidadeDestino = "fahrenheit"
	} else if unidadeOrigem == "quilometros" {
		unidadeDestino = "milhas"
	} else {
		fmt.Printf("%s nao eh uma unidade conhecida!", unidadeDestino)
		os.Exit(1)
	}

	for i, v := range valoresOrigem {
		/*
		Para converter todos os valores informados, precisamos percorrer o slice valoresOrigem e transformar cada valor em
		decimal. Go só possui uma estrutura de repetição,for; neste caso, usamos o for junto com o operador range para obter
		acesso a cada elemento do slice. Range, quando aplicado a um slice, retorna dois valores para cada elemento: o índice
		do elemento e o próprio elemento.
		*/
		valorOrigem, err := strconv.ParseFloat(v, 64)
		if err != nil {
			fmt.Printf("O valor %s na posicao %d nao eh um numero valido\n", v, i)
			os.Exit(1)
			/*
			Usamos a função ParseFloat() para converter a string em ponto flutuante. Essa função recebe dois argumentos -
			o valor a ser convertido e a precisão do valor (32 ou 64 bits) - e retorna dois valores: o valor convertido
			e um erro de conversão(que é nil quando o valor é convertido com sucesso)
			*/
		}
		var valorDestino float64
		if unidadeOrigem == "celsius" {
			valorDestino = valorOrigem*1.8 + 32
		} else {
			valorDestino = valorOrigem / 1.60934
		}
		fmt.Printf("%.2f %s = %.2f %s\n", valorOrigem, unidadeOrigem, valorDestino, unidadeDestino)
	}
}
