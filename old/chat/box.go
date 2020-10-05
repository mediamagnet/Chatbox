package chat

import (
	"Chatbox/config"
	"fmt"
	"github.com/gempir/go-twitch-irc/v2"

	//	"time"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)
// Box does the thing that makes the messages happen
func Box() {

	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")

	var cfg config.Configuration

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	err := viper.Unmarshal(&cfg)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}

	if err != nil {
		log.Error(err)
	}
	client := twitch.NewClient(cfg.Twitch.User, cfg.Twitch.OAuth)
	client.OnPrivateMessage(func(message twitch.PrivateMessage) {
		log.Infof("[%v] %v \n", message.User.DisplayName, message.Message)
	})
	client.OnConnect(func() {
		client.Join(cfg.IRC.Channel)
		log.Infoln("Chatbox Connected")
		client.Say(cfg.IRC.Channel, "Chatbox Connected")
		client.Say(cfg.IRC.Channel, fmt.Sprintf("%s", cfg.Messages.Messages[3]))
	})
	err = client.Connect()
	if err != nil {
		log.Panic(err)
	}
}
