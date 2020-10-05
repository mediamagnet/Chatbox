package main

import (
	"Chatbox/chat"

	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetLevel(log.InfoLevel)
}

func main() {
	chat.InitBot()
}
