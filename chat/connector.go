package chat

import (
	"Chatbox/config"
	"log"
	"os"
	"os/signal"
	"strings"

	soc "github.com/sacOO7/gowebsocket"
)

type Client struct {
	socket soc.Socket
	config config.Configuration
}

func InitBot() {
	conf := config.ConfigInit()

	interupt := make(chan os.Signal, 1)
	signal.Notify(interupt, os.Interrupt)

	client := Client{soc.New("wss://irc-ws.chat.twitch.tv"), conf}

	client.socket.OnConnectError = func(err error, socket soc.Socket) {
		log.Fatal("Received connect error -", err)
	}

	client.socket.OnConnected = func(socket soc.Socket) {
		client.auth()
		client.joinChannel()
	}

	client.socket.OnTextMessage = func(msg string, socket soc.Socket) {
		log.Println(msg)
		if strings.Contains(msg, "PRIVMSG") {
			fsplit := strings.Split(msg, "PRIVMSG")
			colonSplit := strings.Split(fsplit[1], ":")

			if strings.HasPrefix(colonSplit[1], "!donate") {
				client.sendMessage(getRandomMessage(client.config.Messages))
			}
		} else if strings.Contains(msg, "PING") {
			client.sendMessage("PONG :tmi.twitch.tv")
		} else if strings.Contains(msg, "RECONNECT") {
			client.socket.Close()
			client.socket.Connect()
			client.auth()
			client.joinChannel()
		}

	}

	go client.timedMsg()

	client.socket.Connect()

	for {
		select {
		case <-interupt:
			log.Println("Shutting down")
			client.socket.Close()
			return
		}
	}
}
