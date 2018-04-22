package main

import (
	"fmt"
	"os"
	"strconv"
)

/*
Passos do quicksort:
1)Escolher um elemento da lista como pivô e removê-lo da lista;
2)Particionar a lista em duas listas distintas: uma contendo elementos menores que o pivô e outra os maiores;
3)Ordenar as duas listas recursivamente;
4)Retornar a combinação da lista ordenada de elementos menores, o próprio pivô e a lista ordenada dos elementos maiores.
*/
func main(){
	entrada:=os.Args[1:]
	numeros:=make([]int,len(entrada))
	/*
	Declaramos o slice numeros para armazenar os numeros inteiros: aqui usamos a função nativa make() para criar e inicializar um
	slice do tipo []int especificando também seu tamanho inicial como sendo o mesmo da lista recebida como argumento.
	*/
	for i,n:=range  entrada{
		numero,err:=strconv.Atoi(n)// convertendo cada numero para inteiro
		if err!=nil{
			fmt.Printf("%s não é um número válido!\n",n)
			os.Exit(1)
		}
		numeros[i]=numero
	}
	fmt.Println(quicksort(numeros))

}

func quicksort(numeros []int)[]int{
	/*
	O primeiro passo é verificar se a lista de entrada está vazia ou se contém apenas um número e,em caso positivo, retornar a
	própria lista,também chamada condição de parada.
	*/
	if len(numeros)<=1{
		return numeros
	}
	//O próximo passo é criar uma cópia do slice original para evitar que ele seja modificado
	n:=make([]int,len(numeros))
	copy(n,numeros)
	//Em seguida, fazemos a escolha do pivô(no meio da lista). Armazena o índice e o próprio pivô
	indicePivo:=len(n)/2
	pivo:=n[indicePivo]
	/*
	Com o pivô encontrado, precisamos removê-lo da lista original. Faremos isso através do uso da função append(). Ela adiciona um
	elemento ao final de um slice e sua forma geral é: novoSlice:=append(slice,elemento)
	Isso pode parecer um tanto estranho quando queremos de fato remover um elemento do slice. Entretanto, combinando o uso do append()
	com operações de slice,temos uma construção bastante poderosa e idiomática.
	*/
	n=append(n[:indicePivo],n[indicePivo+1:]...)
	/*
	Primeiro, fatiamos o slice n do primeiro elemento até o pivô - n[:indicePivo] - e utilizamos este novo slice com base para
	a operação append(). Depois fatiamos novamente n, partindo do elemento imediatamente posterior ao pivô até o último elemento
	dispoível - n[indicePivo+1:] - e utilizamos este slice como valor a ser adicionado ao slice-base
	*/
	menores,maiores:=particionar(n,pivo)
	return append(append(quicksort(menores),pivo),quicksort(maiores)...)
}

func particionar(numeros []int,pivo int)(menores []int,maiores []int){
	for _,n:= range numeros{
		if n<=pivo{
			menores=append(menores,n)
		}else{
			maiores = append(maiores,n)
		}
	}
	return menores,maiores
}