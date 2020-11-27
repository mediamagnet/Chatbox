package chat

import (
	"Chatbox/config"
	"fmt"
	"math/rand"
)

func (c *Client) auth() {
	c.socket.SendText("CAP REQ :twitch.tv/tags twitch.tv/commands twitch.tv/membership\r\n")
	c.socket.SendText("PASS " + c.config.Twitch.OAuth + "\r\n")
	c.socket.SendText("NICK " + c.config.Twitch.User + "\r\n")
}

func (c *Client) joinChannel() {
	c.socket.SendText("JOIN #" + c.config.Twitch.Channel + "\r\n")
}

func (c *Client) sendMessage(msg string) {
	message := fmt.Sprintf("PRIVMSG #%s :%s\r\n", c.config.Twitch.Channel, msg)
	c.socket.SendText(message)
}
func (c *Client) sendPing(msg string) {
	c.socket.SendText(msg)
}

func getRandomMessage(conf config.MessageConfiguration) string {
	msgs := conf.Messages
	randIndex := rand.Intn(len(msgs))
	msg := msgs[randIndex]

	return msg
}
