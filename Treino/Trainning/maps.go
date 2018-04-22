package main

//O exemplo a seguir conta a frequência de palavras com as mesmas iniciais

import(
	"fmt"
	"os"
	"strings"
	"io"
)

func main(){
	palavras:=os.Args[1:]
	estatisticas:=colherEstatisticas(palavras)
	imprimir(estatisticas)
}

func colherEstatisticas(palavras []string)map[string]int{
	/*
	O tipo de retorno é um map com chaves do tipo string e valores do tipo int. Para cada item armazenado a chave é a letra
	inicial da palavra e o valor é a quantidade de palavras com esta inicial.
	*/
	estatisticas:=make(map[string]int) // inicializando o map
	for _,palavra:= range palavras{
		inicial :=strings.ToUpper(string(palavra[0]))
		contador,encontrado:=estatisticas[inicial]
		/*
		Em seguida, para cada palavra, extraímos sua letra inicial - palavra[0]- e, utilizando a função ToUpper(), convertemos
		a inicial para maiúscula e a armazenamos na variável inicial. Desta forma garantimos que não haja distinção de palavras
		minúsculas ou maiúsculas. Com a inicial em mãos, procuramos no mapa estatisticas uma encontrada cuja chave seja essa inicial.
		Para acessar um valor em um map, utilizamos a chave do valor entre os colchetes. Esta operação retorna dois valores: o valor
		armazenado sob aquela chave(contador) e um bool indicando se a chave existe ou não no map(encontrado)
		*/
		if encontrado{
			estatisticas[inicial]=contador+1
		}else {
			estatisticas[inicial] = 1
		}
	}
	return estatisticas
}

func imprimir(estatisticas map[string]int){
	fmt.Println("Contagem de palavras iniciadas em cada letra:")
	for inicial,contador:=range estatisticas{
		fmt.Printf("%s=%d\n",inicial,contador)
	}
}