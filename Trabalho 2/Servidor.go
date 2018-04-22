package main

import (
	"os"
	"fmt"
	"net"
)

func main(){

	prov:=":1200"
	address,err:=net.ResolveUDPAddr("udp",prov)
	if err !=nil{
		fmt.Println("Erro ao converter endereço! ",os.Args[0])
	}

	sock,err:=net.ListenUDP("udp",address)

	mens := make([]byte, 1024)
	for{
		n, addr, _:= sock.ReadFromUDP(mens)
		handle(mens[0:n],sock, addr)

	}

}

func handle(mens []byte,sock *net.UDPConn, addr *net.UDPAddr){
	fmt.Println("Mensagem chegou")
	_,er:=sock.WriteToUDP(mens,addr)
	if er != nil{
		fmt.Println("Erro ao ecoar mensagem ao cliente!")
		return
	}else{
		fmt.Println("Mensagem enviada")
	}
}

