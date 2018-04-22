package main

import(
	"fmt"
	"math/rand"
	"time"
)
func main(){
	rand.Seed(time.Now().UnixNano())
	n:=0
	for{
		n++
		i:=rand.Intn(4200)
		fmt.Println(i)

		if i%42==0 {
			break
		}
	}
	fmt.Printf("Saída após %d iterações.\n",n)
}
/*
Quando geramos numeros aleatorios, é sempre importante configurar o valor conhecido como seed do gerador. No exemplo, usamos o timestamp
atual no formato padrão do UNIX - o número de nanossegundos desde 1º de janeiro de 1970 - para garantir que, a cada execução do programa,
o gerador de números aleatórios produza números diferentes da vez anterior.
*/
