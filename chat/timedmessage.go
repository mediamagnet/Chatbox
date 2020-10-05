package chat

import (
	"time"
)

func (c *Client) timedMsg() {
	for range time.Tick(time.Minute * 15) {
		c.sendMessage(getRandomMessage(c.config.Messages))
	}
}
