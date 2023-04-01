package main

import (
	"github.com/Kong/go-pdk"
	"github.com/Kong/go-pdk/server"
)

var Version = "0.2"
var Priority = 1

type Config struct {
	Message string
}

func New() interface{} {
	return &Config{}
}

func (conf Config) Access(kong *pdk.PDK) {

	message := conf.Message
	if message == "" {
		message = "hello"
	}

	kong.Log.Notice("Message: " + message)
}

func main() {
	server.StartServer(New, Version, Priority)
}
