package main

import (
	"github.com/zserge/lorca"
	"github/shawncos/file-sync/config"
	"github/shawncos/file-sync/server"
	"log"
	"os"
	"os/signal"
)

func main() {

	go server.Run()

	ui, err := lorca.New("http://localhost:"+config.GetPort()+"/static/index.html", "", 800, 600)
	if err != nil {
		return
	}
	sign := make(chan os.Signal)
	signal.Notify(sign, os.Interrupt)
	select {
	case <-sign:
	case <-ui.Done():
	}

	log.Println("exiting...")
}
