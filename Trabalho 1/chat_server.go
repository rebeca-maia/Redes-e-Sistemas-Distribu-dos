package main

import(
	"os"
	"log"
	"fmt"
	"net"
	//"strings"
	//"reflect"
	//"strings"
)

const MAX_CHAR=200 //bytes ou caracteres

type cliente struct{
	conexao net.TCPConn
	nome string
}

type can struct{
	name  [5] string //em vez de permitir o cliente criar canais, listar os canais previamente cadastrados e fazer com que o cliente escolha um (array, não slice)
	message chan [MAX_CHAR]byte //canal de arrays, não slices
	clientes map[net.TCPConn] string //todos os clientes conectados no canal
}

var (
	culinaria  can
	futebol    can
	cinema     can
	astronomia can
	viagens    can
	y =[]can{culinaria,futebol,cinema,astronomia,viagens}

	a = can{ name: [5]string{"cinema", "futebol", "culinária", "astronomia","viagens"}} //setando canais disponíveis
	mensagens = make(chan []byte,MAX_CHAR)
	clientes = make(map[cliente]bool) // todos os clientes conectados no servidor
	r chan string

)



/*func broadcast(conn net.TCPConn, name string, y []can,c string){
	defer func(){
		conn.Close()
		os.Exit(0)
	}()
	l:
	for{
		for _,n:=range y{
			if strings.Contains(c,"[EU]") && !strings.Contains(c,"/end") {
				//extraindo todos os objetos do tipo 'can' do array y
				// caso não dê certo por switch, usar if's
				//caso a conexão deste usuário seja chave do map clientes, enviar mensagem pelo seu canal objeto.message
				if n.name==[5]string(culinaria.name) { //&& n.clientes: map[net.TCPConn]string == culinaria.clientes:conn:name:
					r <- string(culinaria.message)
					<-r
				}
				if strings.Contains(string(n.name),string(futebol.name)) && reflect.DeepEqual(n.clientes, futebol.clientes){
					r<-string(futebol.message)
					<-r
				}
				if strings.Contains(string(n.name),string(cinema.name)) && reflect.DeepEqual(n.clientes, cinema.clientes) {
					r <- string(cinema.message)
					<-r
				}
				if strings.Contains(string(n.name),string(astronomia.name)) && reflect.DeepEqual(n.clientes, astronomia.clientes) {
					r <- string(astronomia.message)
					<-r
				}
				if strings.Contains(string(n.name),string(viagens.name)) && reflect.DeepEqual(n.clientes, viagens.clientes){
					r<-string(viagens.message)
					<-r
				}
			}
			if strings.Contains(c,"[EU]") && strings.Contains(c,"/end"){
				r <-"["+name+"] "+SERVER_CLIENT_LEFT_CHANNEL
				val,ok:=n.clientes[conn]
				// envia r pelo canal específico do usuário e deleta seu registro no map
				switch conn {
				case n.name == [5]string("culinária") && (val==name && ok):
					<-r

				case n.name == [5]string("futebol") && (val==name && ok):
					<-r

				case n.name == [5]string("cinema") && (val==name && ok):
					<-r

				case n.name == [5]string("astronomia") && (val==name && ok):
					<-r

				case n.name == [5]string("viagens") && (val==name && ok):
					<-r
				}
				break l
			}
			if strings.Contains(c,"/join"){
				r <-"["+name+"] "+SERVER_CLIENT_JOINED_CHANNEL
				val,ok:=n.clientes[conn]
				// envia r pelo canal específico do usuário e deleta seu registro no map
				switch conn {
				case n.name == [5]string("culinária") && (val==name && ok):
					<-r

				case n.name == [5]string("futebol") && (val==name && ok):
					<-r

				case n.name == [5]string("cinema") && (val==name && ok):
					<-r

				case n.name == [5]string("astronomia") && (val==name && ok):
					<-r

				case n.name == [5]string("viagens") && (val==name && ok):
					<-r
				}
			}
		}
	}
}

	/*for {
		_, err := conn.Read(buf[:]) // ou bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			log.Fatal("Erro ao tentar ler comando! ", err)
			os.Exit(1)
		}
		m := string(buf)
		switch m {
		case strings.Contains(m, "/create"):
			link := strings.Split(m, "/create")
			canais = link
			can1 := make(chan []byte, MAX_CHAR)
			mes := string(<-can1) //1
			if mes != " " { //2
				//broadcast(1 e 2 poderiam ficar dentro desta função)
			}
		case "/list":
			conn.Write([]byte(canais))
		case "/join":
			link := strings.Split(m, "/join")
			//inclui este cliente no canal informado e recebe as mensagens para fazer broadcast para os outros clientes inscritos

		}
	}
	*/
	/*i:=0
	for{
		f:
		select{
		case msg:=<-mensagens:
			//faz o broadcast das mensagens de entrada a todos os clientes no canal de saída
			for cli:= range clientes{
				if i>0 {
					cli <- string(msg)
					i++
				}
			}
		case cli:= <-entrada:
			clientes[cli]=true
		case cli:= <-saída:
			delete(clientes,cli)
			close(cli)
			break f
		}
	}*/



func handleClients(conn net.Conn){
	var name []byte=make([]byte,MAX_CHAR)
	conn.Read([]byte(name))
	//b:=cliente(name) //converte o nome do cliente para o tipo cliente( que um canal de string) e atribui a uma variável
	//for b= range clientes{ //faz essa variável percorrer o map[cliente]bool( que guarda todos os clientes que se conectam ao servidor clientes[b]=true //seta como true todos os clientes que se conectarem
	//}
	fmt.Println(string(name))
	/*
	err4:=conn.SetKeepAlivePeriod(0)
	if err4 != nil{
		log.Fatal("SetKeepAlivePeriod ",err4)
		os.Exit(1)
	}


	name,err:= bufio.NewReader(conn).ReadString('\n')
	if err != nil{
		log.Fatal("Erro ao tentar ler nome! ",err)
		os.Exit(1)
	}

	// ok:= broadcast(msg,canais)
	//if !(ok ==true){goto end}
	//apagando conteúdo do slice para entrada de novos dados
	clients[conn]=string(name)
	buf[len(buf)-1] = nil
	//end:
	defer func(){
		_,err:=conn.Read(buf[:]) // ou bufio.NewReader(conn).ReadString('\n')
		if err != nil{
			log.Fatal("Erro ao tentar fechar a conexão! ",err)
			os.Exit(1)
		}
		if string(buf) == "/end"{
			buf[len(buf)-1] = nil
			buf=[]byte(name+" saiu")
			conn.Write(buf)
			conn.Close()
		}

	}()
	broadcast(conn,buf,canais)
	*/
	//ch:=make(chan string) //A função handleClients cria um novo canal de mensagem de saída para o seu cliente
	/*go escritorCliente(conn,mensagens)
	//input:=bufio.NewScanner(conn)// usar ch para ler a mensagem ao invés de input *eliminar toda leitura de mensagem feita por bufio. As leituras serão feitas via canal

	defer os.Exit(1)

	//ch<-CLIENT_MESSAGE_PREFIX
	//mensagens<-who+SERVER_CLIENT_JOINED_CHANNEL
	//entrada<-ch
	for {
		c:=string(<-mensagens)
		if (strings.Contains(c,"/end"))||strings.Contains(c, "/join") || (!(strings.Contains(c,"/end") && (strings.Contains(c,"/list"))) && strings.Contains(c,"[EU]")){
			go broadcast(conn,name,y,c)

		}

		/*for input.Scan() {
			mensagens <- []byte("["+who+"]" + input.Text())
			// em vez de mensagens, enviar mensagem ao canal específico do cliente
			// verificar se este cliente está conectado a um canal, antes de enviar uma mensagem regular
		}
		err := input.Err()
		if err != nil {
			log.Fatalln("Erro na leitura de dados! ", err)
		}
		m:=string(bufio.NewScanner(conn))// é realmente necessário? deixar apenas a variável input ao invés de criar uma nova variável para ler mensagens (seja global ou regular)
		/*if strings.Contains(m,"/create"){
			//antes disso, conferir se o nome atribuído ao canal já existe

			if a.name == ""{
				mensagens<-SERVER_CREATE_REQUIRES_ARGUMENT
			}

			entrada<-ch

		}*/
		/*
		if m =="/end"{
			//excluir o cliente do map do objeto de 'can' e do map onde todos os clientes do servidor estão registrados(finalizar execução do programa)
			// enviar mensagem "Tal cliente saiu do canal"
			break

		}
		if strings.Contains(m,"/join"){
			a=can{clientes: map[net.TCPConn]string{conn:who}}
			a.message<-[MAX_CHAR]byte(who+SERVER_CLIENT_JOINED_CHANNEL)
			//enviar mensagem acima a todos os clientes conectados àquele específico canal

		}
		if m == "/list"{
			a.message<-[MAX_CHAR]byte(a.name)
		}
		if !((strings.Contains(m,"/join")) || (strings.Contains(m,"/list")) || (strings.Contains(m,"/end"))){
			// antes verificar se ele não está conectado a nenhum canal
				a.message <- [MAX_CHAR]byte(SERVER_INVALID_CONTROL_MESSAGE)

		}*/

	//}

}

func escritorCliente(conn net.TCPConn,ch <- chan []byte){
	for msg:= range ch{
		fmt.Println(conn,msg)
	}
}



func main(){
	if len(os.Args) != 2{
		fmt.Fprintln(os.Stderr," Informe o número da porta! ",os.Args[0])
		os.Exit(1)
	}

	mensagens<-[]byte(os.Args[1])
	localhost:="127.0.0.1:"+os.Args[1]

	tcpAddr,err1:= net.ResolveTCPAddr("tcp4",localhost)
	if err1 != nil{
		log.Fatalln("ResolveTCPAddrServer ",err1)
		os.Exit(1)
	}
	//print(tcpAddr)
	listener,err2:=net.Listen("tcp",localhost)
	//defer listener.Close()
	if err2 != nil {
		log.Fatalln("Servidor nao pode iniciar! ", err2)
		os.Exit(1)
	}
	//lembrar de fazer verificação caso o servidor caia enquanto clientes estão trocando informações (enviar uma mensagem global a todos os clientes que o servidor caiu)
	for {
		//err3:=listener.SetDeadline(time.Time{})
		//if err3 != nil{
		//log.Fatal("SetDeadline ",err3)
		//os.Exit(1)
		//}
		conn,err:=listener.Accept()
		if err != nil{
			log.Fatalln("Erro ao aceitar conexão ", err2)
			os.Exit(1)
		}

		//usar broadcast dentro de handleClients
		go handleClients(conn)
	}

}