package main

import (
	"net"
	"os"
	"fmt"
	//"log"
	//"strings"
	//"reflect"
	"bufio"
	"log"
)


/*func SetNames(culinaria,futebol,cinema,astronomia,viagens can){
	culinaria.name=[5]string("culinária")
	futebol.name=[5]string("futebol")
	astronomia.name=[5]string("astronomia")
	cinema.name=[5]string("cinema")
	viagens.name=[5]string("viagens")

}*/
//func deliver(conn net.TCPConn, name string, y []can,c string){
	/*
	//o cliente envia as mensagens digitadas pelo usuário para o servidor, o servidor faz o broadcast(usando sockets não bloqueantes ou canais bufferizados)
	defer func(){
		conn.Close()
		os.Exit(0)
	}()
	/*2. loop infinito
	3. Switch c
	3.1.  verificar se 'm' é uma mensagem regular([EU])
	3.1.1. Se for, extrair canais do array y,enviar ao seu canal respectivo(switch)
	3.2. verificar se c=="/end"
	3.2.1 se for, retira seu registro do map de can, retira seu registro do map global clientes, break(automaticamente o defer é executado)
	4. fim do loop infinito
	*/
	/*for _,n:=range y{
		for{

			if strings.Contains(c,"[EU]") && !strings.Contains(c,"/end") {
				//extraindo todos os objetos do tipo 'can' do array y
				switch n {
				//caso a conexão deste usuário seja chave do map clientes, enviar mensagem pelo seu canal objeto.message
				case n.name == [5]string("culinária") && reflect.DeepEqual(n.clientes, culinaria.clientes): //&& n.clientes: map[net.TCPConn]string == culinaria.clientes:conn:name:
					culinaria.message <- [MAX_CHAR]byte(c)

				case n.name == [5]string("futebol") && reflect.DeepEqual(n.clientes, futebol.clientes):
					futebol.message <- [MAX_CHAR]byte(c)

				case n.name == [5]string("cinema") && reflect.DeepEqual(n.clientes, cinema.clientes):
					cinema.message <- [MAX_CHAR]byte(c)

				case n.name == [5]string("astronomia") && reflect.DeepEqual(n.clientes, astronomia.clientes):
					astronomia.message <- [MAX_CHAR]byte(c)

				case n.name == [5]string("viagens") && reflect.DeepEqual(n.clientes, viagens.clientes):
					viagens.message <- [MAX_CHAR]byte(c)
				}
			}
			if strings.Contains(c,"[EU]") && strings.Contains(c,"/end"){
				val,ok:=n.clientes[conn]
				delete(clientes,cliente(name))
				switch conn {//select
				case reflect.DeepEqual(n.name,culinaria.name) && (val==name && ok)://fazer a mesma coisa para todos os cases contidos nessa função
					delete(culinaria.clientes, conn)

				case n.name == [5]string("futebol",nil,nil,nil,nil) && (val==name && ok):
					delete(futebol.clientes, conn)

				case n.name == [5]string("cinema") && (val==name && ok):
					delete(cinema.clientes, conn)

				case n.name == [5]string("astronomia") && (val==name && ok):
					delete(astronomia.clientes, conn)

				case n.name == [5]string("viagens") && (val==name && ok):
					delete(viagens.clientes, conn)
				}
				break
			}
		}
	}*/
//}

func listen(conn net.Conn,name string) {
	//mensagens <- []byte(name)
	t:=bufio.NewWriter(conn)
	t.Write([]byte(name))
	t.Flush()


	/*defer os.Exit(1)
	var c string
	j := 0
	for {
		fmt.Scanln(&c)
		mensagens<-[]byte(c)

		for i:= range []byte(c){
			if i > 200{
				log.Fatalln("Digite uma mensagem com menos de 200 caracteres!")
				os.Exit(1)
			}
		}
		switch c {
		case c == ("/list"):
			println("Escolha um dos canais a seguir:")
			for _, u := range a.name {
				fmt.Println(u)
			}

		case strings.Contains(c, "/join"):
			w:=string(<-mensagens)
			n := string(strings.Split(c, "/join"))
			print("[EU]")
			for _, u := range a.name {
				if n != u {
					j++
				}
			}
			if j == len(a.name) { //ou fazer essa verificação na parte servidor
				fmt.Println(SERVER_NO_CHANNEL_EXISTS)
			} else if j < len(a.name) {
				switch n {
				case n == "culinária":
					println(w)
					culinaria = can{name: [5]string{"culinária"},, clientes: map[net.TCPConn]string{conn: name}}
				case n == "futebol":
					println(w)
					futebol = can{name: [5]string{"futebol"}, ,clientes: map[net.TCPConn]string{conn: name}}
				case n == "cinema":
					println(w)
					cinema = can{name: [5]string{"cinema"}, ,clientes: map[net.TCPConn]string{conn: name}}
				case n == "astronomia":
					println(w)
					astronomia = can{name: [5]string{"astronomia"}, ,clientes: map[net.TCPConn]string{conn: name}}
				case n == "viagens":
					println(w)
					viagens = can{name: [5]string{"viagens"}, ,clientes: map[net.TCPConn]string{conn: name}}
				}
			}
			// mudar essa condição, pois todas as mensagens, sejam globais ou regulares começam com "[EU]"
		case  (strings.Contains(c,"/end")) || (!(strings.Contains(c,"/end") && (strings.Contains(c,"/list"))) && strings.Contains(c,"[EU]")): // case c=="/end" || (!(strings.Contains(c,"/join")) || (!(strings.Contains(c,"/list")) && a.map[conn]=name está registrado em algum map de 'can'
			//chama função deliver
			go deliver(conn,name,y,c)
		default:
			fmt.Println("Digite uma mensagem válida!")
		}

	}*/

}


func main() {
	if len(os.Args) != 4 {
		fmt.Fprintln(os.Stderr, "Digite seu nome, o ip e o número da porta!")
		os.Exit(1)
	}
	//name := os.Args[1]
	ip:=os.Args[2]
	port:=os.Args[3]
	//a:=string(<-mensagens) //pegando o numero da porta vindo do servidor
	//print(a)
	if ip != "127.0.0.1" || port != a{
		log.Fatalln("Digite um endereço válido!")
	}
	localhost:=ip+":"+port
	//tcpAddr, err := net.ResolveTCPAddr("tcp4", localhost)
	/*if err != nil{
		log.Fatalln("ResolveTCPAddr "+os.Args[0]+" ",err)
		os.Exit(1)
	}*/
	//for {
		conn, err := net.Dial("tcp", localhost)
		if err != nil {
			//log.Fatalln(CLIENT_CANNOT_CONNECT, ip, port)
			os.Exit(1)
		}

		go listen(conn,os.Args[1])
	//}
}

