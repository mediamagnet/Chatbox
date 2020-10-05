package config

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/pelletier/go-toml"
)

func ConfigInit() Configuration {

	var conf Configuration

	if _, err := os.Stat("./config.toml"); os.IsNotExist(err) {
		conf = Configuration{
			TwitchConfiguration{
				Channel: "bees",
				OAuth:   "oauth:a;lkjsdafuhp;oasdjhf;aosiuhdf;kiusdhgf;aiudhgf;jkldhglkjasdhlk",
				User:    "Jeff",
			},
			MessageConfiguration{
				Messages: []string{"No the butts aren't the problem",
					"Yellow is the problem needs less cheese",
					"Maybe, grease the wheel but add a bit of sand in the axel",
					"Why is bacon so good? is it because the cows will it or that the sheep curse them to prevent us from eating them more"},
			},
		}
		file, err := toml.Marshal(conf)
		if err != nil {
			log.Fatal(err)
		}

		err = ioutil.WriteFile("./config.toml", file, 0664)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Please fill in config file")
		os.Exit(1)
	}

	file, err := ioutil.ReadFile("./config.toml")

	if err != nil {
		log.Fatalln(err)
	}

	err = toml.Unmarshal([]byte(file), &conf)

	if err != nil {
		log.Fatalln(err)
	}

	return conf
}
