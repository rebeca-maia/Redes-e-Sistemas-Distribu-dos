package main

import(
	"fmt"
	"strings"
	"net"
	"strconv"
	"regexp"
)

var connectionCount int
var messagePool chan(string)

const (
	INPUT_BUFFER_LENGTH = 140
)
type User struct {
	Name string
	ID int
	Initiated bool
/*The initiated variable tells us that User is connected after a connection and announcement.
Let’s examine the following code to understand the way we’d listen on a channel for a
logged-in user:*/
	UChannel chan []byte
	Connection *net.Conn
}
/*The User struct contains all of the information we will maintain
for each connection. Keep in mind here we don't do any sanity
checking to make sure a user doesn't exist – this doesn't
necessarily pose a problem in an example, but a real chat client
would benefit from a response should a user name already be
in use.*/

func (u *User) Listen() {
	fmt.Println("Listening for",u.Name)
	for {
		select {
		case msg := <- u.UChannel:
			fmt.Println("Sending new message to",u.Name)
			fmt.Fprintln(*u.Connection,string(msg))
		}
	}
}

/*This is the core of our server: each User gets its own Listen() method, which maintains
the User struct’s channel and sends and receives messages across it. Put simply, each user
gets a concurrent channel of his or her own. Let’s take a look at the ConnectionManager
struct and the Initiate() function that creates our server in the following code:*/

type ConnectionManager struct {
	name string
	initiated bool
}

func Initiate() *ConnectionManager {
	cM := &ConnectionManager{ name: "Chat Server 1.0", initiated: false,}
	return cM
}

/*Our ConnectionManager struct is initiated just once. This sets some relatively ornamental
attributes, some of which could be returned on request or on chat login. We’ll examine the
evalMessageRecipient function that attempts to roughly identify the intended recipient of
any message sent as follows:*/

func evalMessageRecipient(msg []byte, uName string) bool {
	eval := true
	expression := "@"
	re, err := regexp.MatchString(expression, string(msg))
	if err != nil {
		fmt.Println("Error:", err)
	}
	
	if re == true {
		eval = false
		
		pmExpression := "@" + uName
		pmRe, pmErr := regexp.MatchString(pmExpression, string(msg))
		if pmErr != nil {
			fmt.Println("Regex error", err)
		}
		if pmRe == true {
			eval = true
		}
	}
	return eval
}
/*This is our router of sorts taking the @ part of the string and using it to detect an intended
recipient to hide from public consumption. We do not return an error if the user doesn’t
exist or has left the chat. Let’s take a look at the code for the Listen() method of the ConnectionManager struct:*/
func (cM *ConnectionManager) Listen(listener net.Listener) {
	fmt.Println(cM.name, "Started")
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Connection error", err)
		}
		
		connectionCount++
		
		fmt.Println(conn.RemoteAddr(), "connected")
		
		user := User{Name: "anonymous", ID: 0, Initiated: false}
		Users = append(Users, &user) // adicionando valores à variável Users do tipo []*User
		
		for _, u := range Users {
			fmt.Println("User online", u.Name)
		}
		
		fmt.Println(connectionCount, "connections active")
		go cM.messageReady(conn, &user)
	}
}

func (cM *ConnectionManager) messageReady(conn net.Conn, user *User) {
	uChan := make(chan []byte)
	for {
		buf := make([]byte, INPUT_BUFFER_LENGTH)
		
		n, err := conn.Read(buf)
		if err != nil || n==0{
			conn.Close()
			conn = nil
		}
		
		fmt.Println(n, "character message from user", user.Name)
		
		if user.Initiated == false {
			fmt.Println("New User is", string(buf))
			
			user.Initiated = true
			user.UChannel = uChan
			user.Name = string(buf[:n])
			user.Connection = &conn
			
			go user.Listen()
			
			minusYouCount := strconv.FormatInt(int64(connectionCount-1),10)
			
			conn.Write([]byte("Welcome to the chat, " + user.Name + ",there are " + minusYouCount + " other users"))
		}else {
			sendMessage := []byte(user.Name + ": " +
			strings.TrimRight(string(buf), " \t\r\n"))
			
			for _, u := range Users {
				if evalMessageRecipient(sendMessage, u.Name) == true {
					u.UChannel <- sendMessage
				}
			}
		}
	}	
}
/*geReady (per connectionManager) function instantiates new
connections into a User struct, utilizing first sent message as
the user's name.*/
var Users []*User
//This is our unbuffered array (or slice) of user structs.
var Users map[string] *User

func main() {
/*As expected, main() primarily handles the connection and error and keeps our server open
and nonblocked with the serverClosed channel.
There are a number of methods we could employ to improve the way we route messages.
The first method would be to invoke a map (or hash table) bound to a username. If the
map’s key exists, we could return some error functionality if a user already exists, as
shown in the following code snippet:*/

	Users := make(map[string] *User)
	connectionCount = 0
	serverClosed := make(chan bool)
	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		fmt.Println ("Could not start server!",err)
	}
	connManage := Initiate()
	go connManage.Listen(listener)
	<-serverClosed
}
