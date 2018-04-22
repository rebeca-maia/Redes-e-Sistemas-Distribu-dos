package main
import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

type client chan<- string // an outgoing message channel

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string) // all incoming client messages
)

func broadcast() {
	clients := make(map[client]bool) // all connected clients
	for {
		select {
		case msg := <-messages:
			// Broadcast entrada de mensagens a todos os canais de saída de mensagens dos clientes.
			for cli := range clients {
				cli <- msg
			}

		case cli := <-entering:
			clients[cli] = true

		case cli := <-leaving:
			delete(clients, cli)
			close(cli)
		}
	}
}


func handleConn(conn net.Conn) {
	defer conn.Close()
	who:=make([]byte,200)

	ch := make(chan string) // outgoing client messages
	go clientWriter(conn, ch)

	_,err := conn.Read(who)
	if err !=nil{
		log.Fatal("Read ",err)
	}
	ch <- "You are " + string(who)
	messages <- string(who) + " entrou no chat"
	//entering <- ch

	input := bufio.NewScanner(conn)

	err1:=input.Err()
	if err1 !=nil{
		fmt.Println("Erro ao ler mensagem! ")
		os.Exit(1)
	}

	for input.Scan() {
		messages <- string(who) + ": " + input.Text()
	}


	leaving <- ch
	messages <- string(who) + " saiu do chat"

}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg) // NOTE: ignoring network errors
	}
}



func main() {
	b:=make([]byte,200)
	listener, err := net.Listen("tcp", "127.0.0.1:1234")
	if err != nil {
		log.Fatalln("Listen ",err)
		os.Exit(1)
	}

	go broadcast()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Accept ",err)
			os.Exit(1)
		}

		_,err1 := conn.Read(b)
		if err1!=nil{
			log.Println("Erro ao ler mensagem")
			os.Exit(1)
		}
		go handleConn(conn)
	}
}

