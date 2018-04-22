package main

const(
	// Client messages

	//Mensagem printada quando um cliente não pode se conectar ao IP e porta do servidor que foram passadas
	CLIENT_CANNOT_CONNECT = "Incapaz de conectar a %s:%d"

	//Mensagem printada antes de sair, se o servidor se desconectar.
	CLIENT_SERVER_DISCONNECTED = "Servidor  %s:%d está desconectado"

	//Mensagem printada antes do cliente digitar novas mensagens
	CLIENT_MESSAGE_PREFIX = "[EU] "

	//Mensagem que pode ser printada acima do prefixo para limpá-lo.
	//Essa string tem espaços brancos no fim para garantir que todos os
	//prefixos "[EU]" é substistuído por essa string. A próxima string escrita
	//para a saída irá precisar começar com "\r" para evitar desnecessários espaços brancos
	CLIENT_WIPE_ME = "\r"

	//Server messages

	//O cliente envia uma mensagem de controle (uma mensagem começando com "/") se ela não existir
	//(e.x., /foobar).
	SERVER_INVALID_CONTROL_MESSAGE = "não é uma válida mensagem de controle. Mensagens válidas são /create, /list, and /join."

	//Mensagem retornada quando um cliente tenta se juntar a um canal que não existe.
	SERVER_NO_CHANNEL_EXISTS = "Este canal não existe. Escolha um canal válido!"

	//Mensagem enviada à clientes que usam "/join" sem o nome do canal.
	SERVER_JOIN_REQUIRES_ARGUMENT = "O comando /join deve ser seguido do nome do canal a se juntar."

	//Mensagem enviada quando um novo cliente entra no canal
	SERVER_CLIENT_JOINED_CHANNEL = "entrou no chat"

	//Mensagem enviada quando um cliente sai do canal
	SERVER_CLIENT_LEFT_CHANNEL = " saiu do chat"

	//Mensagem enviada ao cliente quando ele tenta criar um canal que já existe.
	SERVER_CHANNEL_EXISTS = "Canal %s já existe, então não pode ser criado."

	//Mensagem enviada ao cliente quando ele digita "/create" sem o nome do canal.
	SERVER_CREATE_REQUIRES_ARGUMENT = "O comando /create  deve ser seguido do nome do canal a ser criado"

	//Mensagem enviada ao clinte que envia uma mensagem regular antes de entrar em algum canal.
	SERVER_CLIENT_NOT_IN_CHANNEL = "Você não está em nenhum canal. Deve se conectar num canal antes de enviar mensagens."
)