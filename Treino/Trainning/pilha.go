package main

import(
	"errors"
	"fmt"
)

func main(){
	pilha:=Pilha{} // criamos uma instância da pilha e atribuímos o objeto retornado à variável 'pilha'.
	fmt.Println("Pilha criada com tamanho ",pilha.Tamanho())
	fmt.Println("Vazia? ",pilha.Vazia())

	pilha.Empilhar("Go")
	pilha.Empilhar(2009)
	pilha.Empilhar(3.14)
	pilha.Empilhar("Fim")

	for !pilha.Vazia(){
		v,_:=pilha.Desempilhar()
		fmt.Println("Desempilhando ",v)
		fmt.Println("Tamanho: ",pilha.Tamanho())
		fmt.Println("Vazia? ",pilha.Vazia())
	}
	_,err:=pilha.Desempilhar()
	if err !=nil{
		fmt.Println(err)
	}
}

type Pilha struct {
	valores []interface{} // armazena objetos do tipo interface
	/*
	Esse tipo é conhecido como interface vazia e descreve uma interface sem nenhum método. Na prática, isto faz com que a nossa
	implementação de pilha seja capaz de armazenar objetos de qualquer tipo válido. O slice valores foi intencionalmente declarado
	com a inicial minúscula para que ele não seja acessível a outro pacote.
	*/
}

func(pilha Pilha)Tamanho() int{
	return len(pilha.valores)
}
/*
A definição de um método é semelhante à da função, a diferença é que métodos definem um objeto receptor que deve ser especificado entre
parênteses antes do nome do método. Assim, o método Tamanho() acessa o slice pilha.valores e retorna seu tamanho utilizando a função len
*/

func(pilha Pilha)Vazia() bool{
	return pilha.Tamanho()==0
}

func(pilha *Pilha)Empilhar(valor interface{}){
	pilha.valores=append(pilha.valores,valor)
}
/*
No caso dos métodos Empilhar e Desempilhar, desejamos alterar a pilha na qual tais métodos foram chamados. Em Go, argumentos de funções
e métodos são sempre passados por cópia(exceto slices,maps e channels). Por isso, quando precisamos alterar qualquer argumento - incluindo
receptores de métodos - devemos declará-los como ponteiros.
*/

func(pilha *Pilha)Desempilhar()(interface{},error){
	if pilha.Vazia(){
		return nil,errors.New("Pilha vazia!")
	}
	valor:=pilha.valores[pilha.Tamanho()-1]
	pilha.valores=pilha.valores[:pilha.Tamanho()-1]
	return valor,nil
}
/*
este método possui dois valores de retorno: o objeto desempilhado e um valor do tipo error que é retornado quando a pilha está vazia.
Para isso, utilizamos o método pilha.Vazia() e, em caso de retorno positivo, criamos um novo erro e retornamos nil no lugar do objeto
desempilhado,junto como o erro recém-criado.Caso a pilha não esteja vazia, atribuímos o último objeto empilhado à variável valor. Em
seguida, atualizamos o slice pilha.valores com uma fatia do slice atual,incluindo todos os objetos empilhados com a exceção do último - que
acabou de ser desempilhado. Finalmente, retornamos o objeto removido e nil no lugar do erro, indicando que o objeto foi desempilhado com
sucesso.
*/