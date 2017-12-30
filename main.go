package main

import (
	"./duolingo"
	"encoding/json"
	"io/ioutil"
)

type (
	Config struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}
)

func main() {
	var config Config

	config_file, _ := ioutil.ReadFile("config.json")
	json.Unmarshal(config_file, &config)

	client := duolingo.Login(config.Login, config.Password)
	//client.GetActivity()
	client.GetCourses()
}
