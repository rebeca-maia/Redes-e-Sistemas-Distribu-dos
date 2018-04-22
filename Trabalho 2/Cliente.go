package main

import(
	"os"
	"fmt"
	"net"
	"time"
	"strconv"
)

func main(){
	ad:="127.0.0.1:1200"
	addr,er:=net.ResolveUDPAddr("udp",ad)
	if er !=nil{
		fmt.Println("Erro ao converter endereço! ",os.Args[0])
		os.Exit(1)
	}
	LocalAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:20076")
	if err !=nil{
		fmt.Println("Erro ao converter endereço local!")
		os.Exit(1)
	}
	sock,err:=net.DialUDP("udp",LocalAddr,addr)
	if err !=nil{
		fmt.Println("Erro ao estabelecer conexão com o servidor!")
		os.Exit(1)
	}
	defer sock.Close()
	t := time.After(time.Second * 1)
	message := make([]byte,7)

	L:
	for i:=0;i<10;i++{
		select {
		case <-t:
			fmt.Println("Estouro do timeout!")
			break L
		default:
			//Ping 10 vezes
			beg := time.Now()
			mens:= "PING "+strconv.Itoa(i)
			_,erro:=sock.Write([]byte(mens))
			if erro !=nil{
				fmt.Println("Erro ao enviar mensagem ao servidor!", erro.Error())
			}else{
				fmt.Println("Mensagem enviada")
			}
			_, _, err:=sock.ReadFromUDP(message)
			if err !=nil{
				fmt.Println("Erro ao ler mensagem vinda do servidor!")
			}else{
				fmt.Println("Mensagem chegou")
			}
			end:=time.Now()
			rtt:= end.Sub(beg) // rtt = end-beg
			fmt.Println( "Message Received ", string(message))
			fmt.Println("Round Trip Time ", rtt)

		}
	}
}


