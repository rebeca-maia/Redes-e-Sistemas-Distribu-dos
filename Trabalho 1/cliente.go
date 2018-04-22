package main

import (
	"io"
	"log"
	"net"
	"os"
)
/* O programa cliente na seção 8.3 copia a entrada para o servidor na sua principal goroutine,
de modo que o programa cliente termina assim que o fluxo de entrada é fechado,
mesmo se a goroutine de background ainda estiver funcionando.
Para fazer com que o programa espere que a goroutine de fundo termine antes de sair,
usamos um canal para sincronizar as duas goroutines:*/

func main() {
	nome:=[]byte(os.Args[1])

	conn, err := net.Dial("tcp", "127.0.0.1:1234")
	if err != nil {
		log.Fatalln(err)
	}

	conn.Write(nome)
	done := make(chan struct{})
	go func() {
		_,err:=io.Copy(os.Stdout, conn) // NOTE: ignoring errors
		if err !=nil{
			log.Fatalln("Copy ",err)
			os.Exit(1)
		}
		log.Println("done")
		done <- struct{}{} // signal the main goroutine
		//tentar substituir esse struct por um booleano
	}()
	mustCopy(conn, os.Stdin)
	conn.Close()
	<-done // wait for background goroutine to finish
}

func mustCopy(conn io.Writer, src io.Reader) {
	_, err := io.Copy(conn, src)
	if  err != nil {
		log.Fatal(err)
	}
}